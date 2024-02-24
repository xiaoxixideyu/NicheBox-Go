package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	Length = 4

	TokenLen = 5
	TypeLen  = 1
)

const (
	PacketTypeNormal = 1
	PacketTypePing   = 2
	PacketTypePong   = 3
)

type Packet struct {
	Token string
	Type  int8

	Data []byte
}

func Encode(packet *Packet) ([]byte, error) {
	// todo: 加密
	length := int32(Length + TypeLen + len(packet.Data))

	pkg := new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息体
	err = binary.Write(pkg, binary.LittleEndian, packet.Type)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, packet.Data)
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(pack []byte) (*Packet, error) {

	start := 0
	stop := TokenLen
	token := string(pack[start:stop])

	fmt.Println("token:", token)

	start = stop
	stop = start + TypeLen
	typeBytes := pack[start:stop]

	start = stop
	data := make([]byte, len(pack)-start)
	copy(data, pack[start:])

	var typeInt8 int8

	bytesBuffer := bytes.NewBuffer(typeBytes)
	err := binary.Read(bytesBuffer, binary.LittleEndian, &typeInt8)
	if err != nil {
		return nil, err
	}

	p := Packet{
		Token: token,
		Type:  typeInt8,
		Data:  data,
	}

	return &p, nil
}
