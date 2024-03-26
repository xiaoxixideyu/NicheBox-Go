package tcp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"nichebox/common/biz"
	"nichebox/service/long-connection/rpc/internal/access/connection"
	"nichebox/service/long-connection/rpc/internal/access/protocol"
	"nichebox/service/long-connection/rpc/internal/routes"
	"os"
	"sync"
	"time"
)

const readTimeout = 10

type TCPLongConn struct {
	conn *net.TCPConn
	rw   *bufio.ReadWriter

	sync.Mutex

	uid int64
	ua  string
}

func (lc *TCPLongConn) RemoteAddress() string {
	return lc.conn.RemoteAddr().String()
}

func (lc *TCPLongConn) GetUid() int64 {
	return lc.uid
}

func (lc *TCPLongConn) GetUserAgent() string {
	return lc.ua
}

func (lc *TCPLongConn) ReadPacket() (*protocol.Packet, error) {

	//lengthByte := make([]byte, protocol.Length)
	//n, err := lc.rw.Read(lengthByte)
	//if err != nil {
	//	return nil, err
	//}
	//if n < protocol.Length {
	//	return nil, biz.ErrConnectionNotEnoughBytes
	//}

	err := lc.conn.SetReadDeadline(time.Now().Add(readTimeout * time.Second))
	if err != nil {
		return nil, err
	}

	reader := lc.rw.Reader
	// 读取消息的长度
	// block here
	lengthByte, err := reader.Peek(protocol.Length) // 读取前4个字节的数据
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, biz.ErrConnectionReadTimeout
		}
		return nil, err
	}

	lengthBuff := bytes.NewBuffer(lengthByte)
	var packetBytesLen int32
	err = binary.Read(lengthBuff, binary.LittleEndian, &packetBytesLen)
	if err != nil {
		return nil, err
	}

	//if int32(reader.Buffered()) < protocol.Length+packetBytesLen {
	//	return nil, biz.ErrConnectionNotEnoughBytes
	//}

	fmt.Printf("buff len:%v real len:%v\n", int32(reader.Buffered()), packetBytesLen)

	// read packet bytes
	rawLen := int(protocol.Length + packetBytesLen)
	raw := make([]byte, 0, rawLen)
	curr := 0
	for curr != rawLen {
		// cant use slice as a parameter here, must array
		b := make([]byte, rawLen-curr)
		n, err := reader.Read(b)
		if err != nil {
			return nil, err
		}

		raw = append(raw, b[:n]...)
		curr += n
	}

	pack := make([]byte, rawLen)
	err = binary.Read(bytes.NewBuffer(raw), binary.LittleEndian, pack)
	if err != nil {
		return nil, err
	}

	// decode
	packet, err := protocol.Decode(raw[protocol.Length:])
	if err != nil {
		return nil, err
	}
	return packet, nil
}

func (lc *TCPLongConn) WritePacket(packet *protocol.Packet) error {
	b, err := protocol.Encode(packet)
	if err != nil {
		return err
	}

	lc.Mutex.Lock()
	defer lc.Mutex.Unlock()

	n, err := lc.rw.Write(b)
	fmt.Println("n=", n)
	if err != nil {
		return err
	}
	err = lc.rw.Flush()
	if err != nil {
		return err
	}
	return nil
}

//func (lc *TCPLongConn) Peek(n int) (b []byte, err error) {
//	return lc.rw.Peek(n)
//}
//
//func (lc *TCPLongConn) Write(b []byte) (n int, err error) {
//	return lc.rw.Write(b)
//}
//
//func (lc *TCPLongConn) Read(b []byte) (n int, err error) {
//	return lc.rw.Read(b)
//}

func (lc *TCPLongConn) Close() error {
	var c connection.LongConn = lc
	routes.GetRouter().UnregisterConn(c)
	return lc.conn.Close()
}

func newTCPLongConn(c *net.TCPConn) connection.LongConn {
	return &TCPLongConn{
		conn: c,
		rw:   bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c)),
	}
}
