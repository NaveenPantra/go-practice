package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		var incomingMessage []byte
		conn.Read(incomingMessage)
		fmt.Println(string(incomingMessage))
		_, err = fmt.Fprintln(conn, "Hello from the server")
		if err != nil {
			fmt.Println(err)
			continue
		}
		conn.Close()
	}
}
