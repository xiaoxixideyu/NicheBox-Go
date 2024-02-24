package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Packet struct {
	Length int32
	Token  string
	Type   int8

	Data []byte
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10031")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 0; i++ {
		token := "token"
		lenToken := len([]byte(token))

		pkgType := int8(2)
		data := []byte(`Hello, Hello. How are you?`)

		len := int32(lenToken + 1 + len(data))

		fmt.Println("len:", len)

		err := binary.Write(conn, binary.LittleEndian, len)
		if err != nil {
			fmt.Println("err:", err)
		}
		err = binary.Write(conn, binary.LittleEndian, []byte(token))
		if err != nil {
			fmt.Println("err:", err)
		}
		err = binary.Write(conn, binary.LittleEndian, pkgType)
		if err != nil {
			fmt.Println("err:", err)
		}
		err = binary.Write(conn, binary.LittleEndian, data)
		if err != nil {
			fmt.Println("err:", err)
		}

		//for err != io.EOF {
		//	b := make([]byte, 5)
		//	_, err = conn.Read(b)
		//	if err != nil {
		//		fmt.Println("err:", err)
		//	}
		//	fmt.Printf("%v", b)
		//
		//	time.Sleep(time.Second)
		//}

		time.Sleep(8 * time.Second)
	}

	b := make([]byte, 128)
	_, err = conn.Read(b)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%v", b)

	time.Sleep(time.Hour)
	conn.Close()

}
