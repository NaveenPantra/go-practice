package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var store = map[string]string{}
var commands = map[string]string{
	"GET": "GET",
	"SET": "SET",
	"DEL": "DEL",
}

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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		processCommand(text, conn)
	}
	fmt.Println("** Connection Closed **")
}

func processCommand(req string, conn net.Conn) {
	reqs := strings.Split(req, " ")
	cmd := reqs[0]
	key := reqs[1]
	switch cmd {
	case commands["GET"]:
		{
			value := getKey(key)
			fmt.Fprintln(conn, value)
		}
	case commands["SET"]:
		{
			if len(reqs) < 3 {
				fmt.Fprintln(conn, "Expecting value")
				return
			}
			value := reqs[2]
			setKey(key, value)
			fmt.Fprintln(conn, value)
		}
	case commands["DEL"]:
		delKey(key)
		fmt.Fprintln(conn, "Deleted ", key)
	}
}

func setKey(key, value string) {
	store[key] = value
}

func getKey(key string) string {
	return store[key]
}

func delKey(key string) {
	delete(store, key)
}
