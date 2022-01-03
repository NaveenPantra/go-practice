package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	err := conn.SetDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintln(conn, "You said: ", ln)
	}

	fmt.Println("*** End of connection ***")

}
