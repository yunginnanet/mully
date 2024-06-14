//go:build !windows

package socket

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/user"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"git.tcp.direct/kayos/common/hash"
	"git.tcp.direct/kayos/common/pool"
	"git.tcp.direct/kayos/zwrap"
	"github.com/davecgh/go-spew/spew"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/sys/unix"
	"golang.org/x/text/encoding/unicode"
)

type AllowedMode uint8

const (
	AllowAll AllowedMode = iota
	AllowUser
	AllowGroup
	AllowPID
)

type Allowed[T comparable] interface {
	Value() T
	Type() AllowedMode
}

type AllowedUser uint32

func (u AllowedUser) Value() uint32 {
	return uint32(u)
}

func CurrentUser() AllowedUser {
	return AllowedUser(syscall.Getuid())
}

func (u AllowedUser) String() string {
	uu, e := user.LookupId(fmt.Sprintf("%d", u))
	if e != nil {
		return fmt.Sprintf("%d", u)
	}
	return uu.Name
}

func (u AllowedUser) Type() AllowedMode {
	return AllowUser
}

type AllowedGroup uint32

func (g AllowedGroup) String() string {
	return strconv.Itoa(int(g))
}

func (g AllowedGroup) Type() AllowedMode {
	return AllowGroup
}

func (g AllowedGroup) CheckUser(u user.User) bool {
	gids, err := u.GroupIds()
	if err != nil {
		return false
	}
	for _, gid := range gids {
		if gid == strconv.Itoa(int(g)) {
			return true
		}
	}
	return false
}

type UnixSocket struct {
	path              string
	conns             map[Allowed[uint32]][]net.Conn
	allowed           map[Allowed[uint32]]bool
	commands          map[hash.Type]func([]byte) []byte
	maxConnsPerClient int
	l                 net.Listener
	connMu            sync.RWMutex
	allowedMu         sync.RWMutex
	allowAll          bool
	log               zwrap.StdCompatLogger
	logMu             sync.RWMutex
	workers           *ants.Pool
	commandHandler    func([]byte) []byte
}

func NewUnixSocket(path string, logger zwrap.StdCompatLogger) (*UnixSocket, error) {
	unixAddr, err := net.ResolveUnixAddr("unix", path)
	if err != nil {
		return nil, fmt.Errorf("invalid unix socket address: %w", err)
	}

	_ = syscall.Unlink(path)
	oldMask := syscall.Umask(0o077)

	var unixListener net.Listener
	if unixListener, err = net.ListenUnix("unix", unixAddr); err != nil {
		syscall.Umask(oldMask)
		return nil, fmt.Errorf("failed to listen on unix socket: %w", err)
	}
	syscall.Umask(oldMask)
	if err = os.Chmod(path, 0o666); err != nil {
		return nil, fmt.Errorf("failed to chmod unix socket: %w", err)
	}

	w, _ := ants.NewPool(10, ants.WithNonblocking(true))
	us := &UnixSocket{
		conns:             make(map[Allowed[uint32]][]net.Conn),
		allowed:           make(map[Allowed[uint32]]bool),
		allowAll:          false,
		maxConnsPerClient: 1,
		l:                 unixListener,
		log:               logger,
		path:              unixAddr.String(),
		workers:           w,
	}

	us.allowed[CurrentUser()] = true

	go func() {
		for {
			c, e := us.Accept()
			if e != nil {
				us.Log().Println("failed to accept connection: " + e.Error())
				continue
			}
			if c == nil {
				continue
			}
			if acceptErr := w.Submit(func() {
				if err = us.HandleConn(c); err != nil {
					if strings.Contains(err.Error(), "use of closed network connection") {
						return
					}
					if c != nil {
						_ = c.Close()
					}
					us.Log().Printf("failed to handle connection: %s\n", err.Error())
				}
			}); acceptErr != nil {
				_, _ = c.Write([]byte("TOO_MANY_CONNECTIONS\n"))
				_ = c.Close()
			}
		}

	}()

	time.Sleep(250 * time.Millisecond)
	return us, nil
}

func (u *UnixSocket) SetMaxConnsPerClient(limit int) {
	u.allowedMu.Lock()
	u.maxConnsPerClient = limit
	u.allowedMu.Unlock()
}

func (u *UnixSocket) SetLogger(logger zwrap.StdCompatLogger) {
	u.logMu.Lock()
	u.log = logger
	u.logMu.Unlock()
}

func (u *UnixSocket) HandleConn(c net.Conn) error {
	uc, ok := c.(*net.UnixConn)
	if !ok {
		return fmt.Errorf("invalid connection type: %T", c)
	}
	sysCallConn, err := uc.SyscallConn()
	if err != nil {
		return fmt.Errorf("failed to get syscall conn: %w", err)
	}
	var ucred *unix.Ucred
	if err = sysCallConn.Control(func(fd uintptr) {
		ucred, err = unix.GetsockoptUcred(int(fd), unix.SOL_SOCKET, unix.SO_PEERCRED)
	}); err != nil {
		return fmt.Errorf("failed to get SO_PEERCRED for incoming connection: %w", err)
	}

	permitted := false

	u.allowedMu.RLock()
	if u.allowAll {
		permitted = true
	}
	if !permitted {
		allowed, aok := u.allowed[AllowedUser(ucred.Uid)]
		if aok && allowed {
			permitted = true
		}
	}
	u.allowedMu.RUnlock()

	u.connMu.RLock()
	if _, exists := u.conns[AllowedUser(ucred.Uid)]; exists {
		if len(u.conns[AllowedUser(ucred.Uid)])+1 > u.maxConnsPerClient {
			permitted = false
		}
	}
	u.connMu.RUnlock()

	if !permitted {
		_, _ = c.Write([]byte("DENIED\n"))
		_ = c.Close()
		return fmt.Errorf("connection from %s not permitted", AllowedUser(ucred.Uid))
	}

	u.log.(*zwrap.Logger).Logger.Info().Str("caller", spew.Sdump(ucred)).Msg("accepting connection")
	u.connMu.Lock()
	u.conns[AllowedUser(ucred.Uid)] = append(u.conns[AllowedUser(ucred.Uid)], c)
	u.connMu.Unlock()
	go u.handleConn(c)
	return nil
}

var bufs = pool.NewBufferFactory()
var dec = unicode.UTF8.NewDecoder()

func (u *UnixSocket) handleConn(c net.Conn) {
	log := u.log.(*zwrap.Logger).Logger.With().Str("caller", c.RemoteAddr().String()).Logger()
	log.Trace().Msg("handling connection")
	defer func() {
		if err := c.Close(); err != nil {
			log.Warn().Err(err).Msg("failed to close connection")
		}
	}()
	errCt := 0

	checkErr := func(err error) (ok bool, bail bool) {
		if err == nil {
			return true, false
		}
		if errors.Is(err, io.EOF) {
			log.Trace().Msg("connection EOF")
			return true, true
		}
		errCt++
		log.Warn().Err(err).Msg("failed to read from connection")
		if errCt > 5 {
			log.Warn().Msg("too many errors, closing connection")
			return false, true
		}
		return true, false
	}

	buf := bufs.Get()
	buf.MustReset()
	_ = buf.Grow(1024)

	for {
		var ok, bail bool
		n, err := c.Read(buf.Bytes()[:1024])
		if ok, bail = checkErr(err); !ok || bail {
			break
		}

		if n == 0 {
			continue
		}

		u.log.(*zwrap.Logger).Logger.Trace().Msg("received data: " + string(buf.Bytes()[:n]))

		if !bytes.Contains(buf.Bytes(), []byte{0x0a}) {
			continue
		}
		incoming, err := buf.ReadBytes(byte(0x0a))
		if err != nil {
			continue
		}

		incoming = bytes.TrimSpace(incoming)

		if len(incoming) == 0 {
			continue
		}

		var utf8Err error

		if incoming, utf8Err = dec.Bytes(incoming); utf8Err != nil {
			log.Warn().Err(utf8Err).Msg("failed to decode incoming data to utf8")
			break
		}
		log.Trace().Msg("received data: " + string(incoming))
		if u.commandHandler != nil {
			resp := u.commandHandler(incoming)
			if len(resp) == 0 {
				_, _ = c.Write([]byte{0x0a})
				continue
			}
			if _, err = c.Write(resp); err != nil {
				log.Warn().Err(err).Msg("failed to write response")
				break
			}
		}
	}

}

/*

 */

func (u *UnixSocket) Log() zwrap.StdCompatLogger {
	u.logMu.RLock()
	lg := u.log
	u.logMu.RUnlock()
	return lg
}

func (u *UnixSocket) SetCommandHandler(handler func([]byte) []byte) {
	u.commandHandler = handler
}

func (u *UnixSocket) Accept() (net.Conn, error) {
	c, e := u.l.Accept()
	if e != nil {
		if strings.Contains(e.Error(), "use of closed network connection") {
			return nil, nil
		}
		u.log.(*zwrap.Logger).Logger.Error().Err(e).Msg("failed to accept connection")
	} else {
		// fixme
		u.log.(*zwrap.Logger).Logger.Info().Msg("accepted connection")
	}
	return c, e
}

func (u *UnixSocket) Close() error {
	u.Log().Println("closing unix socket: " + u.path)
	return u.l.Close()
}

func (u *UnixSocket) Addr() net.Addr {
	return u.l.Addr()
}

func (u *UnixSocket) Allowed(val Allowed[uint32]) bool {
	u.allowedMu.RLock()
	if u.allowAll {
		u.allowedMu.RUnlock()
		return true
	}
	var (
		allowed bool
		ok      bool
	)

	allowed, ok = u.allowed[val]
	u.allowedMu.RUnlock()
	return ok && allowed
}

func (u *UnixSocket) RemoteAddrs() []Allowed[uint32] {
	u.connMu.RLock()
	var addrs = make([]Allowed[uint32], 0, len(u.conns))
	for ip := range u.conns {
		addrs = append(addrs, ip)
	}
	u.connMu.RUnlock()
	return addrs
}

func (u *UnixSocket) LocalAddr() net.Addr {
	return u.l.Addr()
}
