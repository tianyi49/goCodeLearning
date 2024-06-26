package main

import (
	"fmt"
	"goCodeLearning/socket_stick/proto"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":30000")
	if err != nil {
		fmt.Println("dial failed err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
