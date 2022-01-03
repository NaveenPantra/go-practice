package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type responsePathHandler interface {
	get(conn net.Conn)
	post(conn net.Conn)
}

type homeHandler struct{}

func (homeHandler homeHandler) get(conn net.Conn) {
	body := `<html><body><p>Hello World!!</p></body></html>`
	fmt.Fprintf(conn, "HTTP/1.1 %v %v\r\n", http.StatusOK, http.StatusText(http.StatusOK))
	fmt.Fprintf(conn, "Content-Length: %v\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func (homeHandler homeHandler) post(conn net.Conn) {
	fmt.Fprintf(conn, "HTTP/1.1 %v %v\r\n", http.StatusTemporaryRedirect, http.StatusText(http.StatusTemporaryRedirect))
	fmt.Fprintf(conn, "Location: locahost:8080/e40\r\n")
	fmt.Fprintf(conn, "\r\n")
}

type e404Handler struct{}

func (e404Handler e404Handler) get(conn net.Conn) {
	e404(conn)
}

func (e404Handler e404Handler) post(conn net.Conn) {
	e404(conn)
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
		go handleIncomingRequest(conn)
	}
}

func handleIncomingRequest(conn net.Conn) {
	defer conn.Close()
	method, path := processRequest(conn)
	processResponse(conn, method, path)
}

func processRequest(conn net.Conn) (method, path string) {
	scanner := bufio.NewScanner(conn)
	lineCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineCounter == 0 {
			header := strings.Fields(line)
			method = header[0]
			path = header[1]
			fmt.Printf("\n\n\nMethod: %v, Path: %v\n", method, path)
		}
		fmt.Println(line)
		if line == "" {
			break
		}
		lineCounter++
	}
	return
}

func processResponse(conn net.Conn, method, path string) {
	switch path {
	case "/":
		{
			handler := homeHandler{}
			pathHandler(handler, conn, method)
		}
	case "e404":
	default:
		{
			handler := e404Handler{}
			pathHandler(handler, conn, method)
		}
	}
}

func pathHandler(handler responsePathHandler, conn net.Conn, method string) {
	switch method {
	case http.MethodGet:
		handler.get(conn)
	case http.MethodPost:
		handler.post(conn)
	default:
		e404 := e404Handler{}
		e404.get(conn)
	}

}

func e404(conn net.Conn) {
	body := `<html><body><p>E404 - Not Found</p></body></html>`
	fmt.Fprintf(conn, "HTTP/1.1 %v %v\r\n", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	fmt.Fprintf(conn, "Content-Length: %v\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
