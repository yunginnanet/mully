package socket

import (
	"io"
	"net"

	"git.tcp.direct/kayos/zwrap"
)

type Socket[T comparable] interface {
	net.Listener
	io.Closer
	Log() zwrap.StdCompatLogger
	SetLogger(zwrap.StdCompatLogger)
	SetCommandHandler(func([]byte) []byte)
	Allowed(Allowed[T]) bool
	RemoteAddrs() []net.Addr
	LocalAddr() net.Addr
	HandleConn(net.Conn)
}
