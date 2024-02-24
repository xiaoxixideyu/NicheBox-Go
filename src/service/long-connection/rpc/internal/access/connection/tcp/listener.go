package tcp

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"nichebox/service/long-connection/rpc/internal/access/connection"
)

const BufSize = 100

type TCPListener struct {
	ch   chan connection.LongConn
	port int

	logx.Logger
}

func (l *TCPListener) Listen() (<-chan connection.LongConn, error) {
	addr := fmt.Sprintf("127.0.0.1:%d", l.port)
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		l.Logger.Errorf("[TCP] Server listen failed, err:", err)
		return nil, err
	}

	// listen
	go func(listen net.Listener) {

		for {
			conn, err := listen.Accept()
			if err != nil {
				l.Logger.Errorf("[TCP] Server accept failed, err:", err)
				continue
			}
			lc := newTCPLongConn(conn.(*net.TCPConn))

			if len(l.ch) == cap(l.ch) {
				// prevent blocking
				go func() {
					l.ch <- lc
				}()
			} else {
				l.ch <- lc
			}

		}
	}(listen)

	fmt.Println("Tcp server starts listening at", addr)

	return l.ch, nil
}

func NewTCPListener(port int) connection.Listener {
	ch := make(chan connection.LongConn, BufSize)
	return &TCPListener{
		ch:     ch,
		port:   port,
		Logger: logx.WithContext(context.Background()),
	}
}
