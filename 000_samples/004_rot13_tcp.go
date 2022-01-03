package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		encryptedText := scanner.Text()
		encryptedText = strings.ToLower(encryptedText)
		decryptedText := rotate(encryptedText, 13)
		fmt.Fprintf(conn, "%v - %v\n", encryptedText, decryptedText)
	}
	fmt.Println("** Closing Connection **")
}

func rotate(str string, rot int) string {
	ebs := []rune(str)
	dbs := make([]rune, len(ebs))
	for ind, val := range ebs {
		if val < 97 || val > 122 {
			dbs[ind] = val
		} else if val+13 > 122 {
			dbs[ind] = 97 + (val + 13 - 122 - 1)
		} else {
			dbs[ind] = val + 13
		}
	}
	fmt.Print(ebs, dbs, "\n")

	return string(dbs)
}
