package session

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/common/biz"
	"nichebox/service/long-connection/rpc/internal/access/connection"
	"nichebox/service/long-connection/rpc/internal/access/connection/tcp"
	Pool "nichebox/service/long-connection/rpc/internal/access/pool"
	"nichebox/service/long-connection/rpc/internal/access/protocol"
	"nichebox/service/long-connection/rpc/internal/config"
	"nichebox/service/long-connection/rpc/internal/routes"
	"strings"
	"time"
)

type Manager struct {
	ctx context.Context

	pingInterval time.Duration
	pool         *Pool.Pool

	logger logx.Logger
}

func NewSessionManager(ctx context.Context, pingInterval int) Manager {
	return Manager{
		ctx:          ctx,
		pingInterval: time.Duration(pingInterval) * time.Second,
		pool:         Pool.NewPool(),
		logger:       logx.WithContext(ctx),
	}
}

func (m *Manager) Start(c config.Config) error {
	// tcp
	tcpListener := tcp.NewTCPListener(c.ServerConf.TCPPort)
	tcpCh, err := tcpListener.Listen()
	if err != nil {
		m.logger.Errorf("[SessionManager] Start tcp listening failed, err:", err)
		return err
	}
	fmt.Println("TCP Starting listen===================")
	go m.process(tcpCh)

	// todo: websocket

	return nil
}

func (m *Manager) PushToUser(uid int64, data []byte) error {
	r := routes.GetRouter()
	conns, ok := r.GetConnsByUid(uid)
	if !ok {
		return biz.ErrConnectionNotFound
	}
	packet := protocol.Packet{
		Token: "",
		Type:  protocol.PacketTypeNormal,
		Data:  data,
	}
	for _, conn := range conns {
		m.pool.Go(push, &packet, conn, m.logger)
	}
	return nil
}

func (m *Manager) process(ch <-chan connection.LongConn) {
	for conn := range ch {
		addr := conn.RemoteAddress()
		ip := strings.Split(addr, ":")[0]
		r := routes.GetRouter()
		uid, ua, ok := r.GetUidAndUAByAddr(ip)
		if !ok {
			m.logger.Errorf("[SessionManager] Addr did not register, addr:", ip)
			continue
		}

		r.RegisterConn(uid, ua, conn)
		m.logger.Infof("[SessionManager] Successfully register conn:", uid, ua)
		go m.reader(conn)
	}
}

func (m *Manager) reader(conn connection.LongConn) {
	defer func() {
		if p := recover(); p != nil {
			m.logger.Errorf("[Session] Reader goroutine panic, err:", p)
		}
	}()

	timer := time.NewTimer(m.pingInterval)

	for {
		// todo: timer receive ping and send pong
		select {
		case <-timer.C:
			fmt.Println("timer timeout")
			conn.Close()
			return
		default:
		}

		packet, err := conn.ReadPacket()
		if err != nil {
			if err == biz.ErrConnectionReadTimeout {
				continue
			}
			m.logger.Errorf("[Session] Read packet failed, err:", err)
			return
		}
		fmt.Printf("packet:%v\n", packet)

		m.pool.Go(forward, packet, conn, m.logger)

		timer.Reset(m.pingInterval)
	}
}

func forward(packet *protocol.Packet, conn connection.LongConn, logger logx.Logger) {
	// todo: check token

	if int(packet.Type) == protocol.PacketTypePing {
		pong := protocol.Packet{
			Token: "",
			Type:  protocol.PacketTypePong,
			Data:  nil,
		}
		err := conn.WritePacket(&pong)
		if err != nil {
			logger.Errorf("[Session] Write pong packet failed, err:", err)
			panic(err)
		}
		fmt.Println("send!!!!!")

	} else {

		switch int(packet.Type) {
		case protocol.PacketTypeNormal:

		default:

		}
	}

}

func push(packet *protocol.Packet, conn connection.LongConn, logger logx.Logger) {
	err := conn.WritePacket(packet)
	if err != nil {
		logger.Errorf("[Session] Push failed, err:", err)
	}
}
