package sip

import (
	"context"
	"crypto/tls"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

// The hook overrides the dialled host, which is the whole point: an IP
// destination has no name to send, and the application knows the domain the
// certificate was issued to.
func TestTransportTLSServerNameHookNamesTheConnection(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer func() { _ = l.Close() }()
	go func() {
		for {
			conn, acceptErr := l.Accept()
			if acceptErr != nil {
				return
			}
			_ = conn.Close()
		}
	}()

	addr := l.Addr().(*net.TCPAddr)
	for _, c := range []struct {
		name string
		hook func(Addr) string
		want string
	}{
		{"hook names it", func(Addr) string { return "carrier.test" }, "carrier.test"},
		{"empty keeps the dialled host", func(Addr) string { return "" }, "127.0.0.1"},
		{"nil keeps the dialled host", nil, "127.0.0.1"},
	} {
		t.Run(c.name, func(t *testing.T) {
			tp := &TransportTLS{TransportTCP: &TransportTCP{}, ServerName: c.hook}
			tp.init(NewParser(), &tls.Config{MinVersion: tls.VersionTLS12})
			seen := make(chan string, 1)
			tp.tlsClient = func(conn net.Conn, hostname string) *tls.Conn {
				seen <- hostname
				return tls.Client(conn, &tls.Config{MinVersion: tls.VersionTLS12})
			}
			// The handshake fails against a listener that speaks no TLS; what is
			// under test happened before it.
			_, _ = tp.CreateConnection(context.TODO(), Addr{}, Addr{IP: addr.IP, Port: addr.Port}, nil)
			require.Equal(t, c.want, <-seen)
		})
	}
}
