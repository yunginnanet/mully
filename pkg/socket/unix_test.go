package socket

import (
	"net"
	"path/filepath"
	"testing"

	"git.tcp.direct/kayos/common/entropy"
	"git.tcp.direct/kayos/zwrap"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
)

func TestNewUnixSocket(t *testing.T) {
	logger := zwrap.Wrap(zerolog.New(zerolog.ConsoleWriter{Out: &testWriter{t: t}}))
	listenOn := filepath.Join(t.TempDir(), entropy.RandStrWithUpper(5))
	socket, err := NewUnixSocket(listenOn, logger)
	if err != nil {
		t.Error(err)
	}
	if socket == nil {
		t.Fatal("socket is nil")
	}
	t.Logf("socket: %s", spew.Sdump(socket))
	if err = socket.Close(); err != nil {
		t.Error(err)
	}
}

func TestUnixSocket_HandleConn(t *testing.T) {
	logger := zwrap.Wrap(zerolog.New(zerolog.ConsoleWriter{Out: &testWriter{t: t}}))
	listenOn := filepath.Join(t.TempDir(), entropy.RandStrWithUpper(5))
	socket, err := NewUnixSocket(listenOn, logger)

	conn, err := net.Dial("unix", listenOn)
	if err != nil {
		t.Fatal(err)
	}
	if conn == nil {
		t.Fatal("conn is nil")
	}

	if err = socket.HandleConn(conn); err != nil {
		t.Error(err)
	}
	if err = socket.Close(); err != nil {
		t.Error(err)
	}

}

func TestUnixSocket_Allowed(t *testing.T) {
	logger := zwrap.Wrap(zerolog.New(zerolog.ConsoleWriter{Out: &testWriter{t: t}}))
	listenOn := filepath.Join(t.TempDir(), entropy.RandStrWithUpper(5))
	socket, err := NewUnixSocket(listenOn, logger)
	if err != nil {
		t.Fatal(err)
	}
	if !socket.Allowed(CurrentUser()) {
		t.Fatal("expected true")
	}

	socket.SetCommandHandler(func(b []byte) []byte {
		return []byte("test")
	})

	if c, cErr := net.Dial("unix", listenOn); cErr != nil {
		t.Fatal(cErr)
	} else {
		if _, wErr := c.Write([]byte("test\n")); wErr != nil {
			t.Fatal(wErr)
		}
		b := make([]byte, len("test"))
		if _, rErr := c.Read(b); rErr != nil {
			t.Fatal(rErr)
		}
		if string(b) != "test" {
			t.Fatalf("expected to read '%s', but read '%s'", "test", string(b))
		}
		b = b[:1]
		if _, rErr := c.Read(b); rErr != nil {
			t.Fatal(rErr)
		}
		if string(b) != "\n" {
			t.Fatalf("expected to read '%s', but read '%s'", "\n", string(b))
		}
	}

	if err = socket.Close(); err != nil {
		t.Error(err)
	}
}

type testWriter struct{ t *testing.T }

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.t.Helper()
	w.t.Logf("%s", p)
	return len(p), nil
}

func TestUnixSocket_RemoteAddrs(t *testing.T) {
	logger := zwrap.Wrap(zerolog.New(zerolog.ConsoleWriter{Out: &testWriter{t: t}}))
	listenOn := filepath.Join(t.TempDir(), entropy.RandStrWithUpper(5))
	socket, err := NewUnixSocket(listenOn, logger)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = net.Dial("unix", listenOn); err != nil {
		t.Fatal(err)
	}
	addrs := socket.RemoteAddrs()
	t.Logf("addrs: %s", spew.Sdump(addrs))
	if addrs == nil {
		t.Fatal("addrs is nil")
	}
	if err = socket.Close(); err != nil {
		t.Error(err)

	}
}

func TestUnixSocket_LocalAddr(t *testing.T) {
	logger := zwrap.Wrap(zerolog.New(zerolog.ConsoleWriter{Out: &testWriter{t: t}}))
	listenOn := filepath.Join(t.TempDir(), entropy.RandStrWithUpper(5))
	socket, err := NewUnixSocket(listenOn, logger)
	if err != nil {
		t.Fatal(err)
	}
	addr := socket.LocalAddr()
	if addr == nil {
		t.Fatal("addr is nil")
	}
	t.Logf("addr: %s", addr)

	if err = socket.Close(); err != nil {
		t.Error(err)

	}
}
