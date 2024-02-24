package connection

import "nichebox/service/long-connection/rpc/internal/access/protocol"

type LongConn interface {
	GetUid() int64
	GetUserAgent() string
	RemoteAddress() string
	ReadPacket() (*protocol.Packet, error)
	WritePacket(packet *protocol.Packet) error
	Close() error
}

type Listener interface {
	Listen() (<-chan LongConn, error)
}
