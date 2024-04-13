package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// http server
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	// Keep connection open for multiple requests
	for {
		// Open connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// go subroutine to enable concurrency
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) error {
	var response string
	// ensure connection tcp
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
    response = Response{404, "", ""}.ToString()
    return fmt.Errorf("Connection is not TCP")
	}

	// get parsed request
	request, err := parseRequest(tcpConn)
	if err != nil {
		fmt.Printf("Error parsing request: %v\n", err)

    response = Response{400, "", ""}.ToString()
    writeResponse(tcpConn, response)
    return err
	}

	// get response
	response = getResponse(request)

	// Write response to the connection
	writeResponse(tcpConn, response)
  return nil
}

func writeResponse(conn *net.TCPConn, response string) {
  _, err := conn.Write([]byte(response))
  if err != nil {
    fmt.Println("Failed to write response!")
  }
	conn.CloseWrite()
}
