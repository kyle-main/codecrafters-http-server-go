package main

import (
	"fmt"
	"net"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

func parseRequest(conn *net.TCPConn) (*Request, error) {
	input := make([]byte, 1024) // create slice for input
	length, err := conn.Read(input)
	if err != nil {
		fmt.Println("Error reading from connection")
	}
	// get input
	input = input[:length]

	// split input into []lines
	lines := strings.Split(string(input), "\r\n")

	// Check we have enough lines to parse
	if len(lines) < 3 { // Aww <3
		return nil, fmt.Errorf("Invalid Request")
	}

	// Check the first line is shaped right
	firstLine := strings.Split(lines[0], " ")
	if len(firstLine) != 3 {
		return nil, fmt.Errorf("Invalid Request")
	}

	// Parse start line
	method := firstLine[0]
	path := firstLine[1]
	version := firstLine[2]
  fmt.Println(method)
  fmt.Println(path)
  fmt.Println(version)
	// Alloc header map
	headers := make(map[string]string)

	// Validate header lines up to penulimate line
	for i := 1; i < len(lines)-2; i++ {
		line := strings.Split(lines[i], ": ")
    fmt.Println(line)
		if len(line) < 2 {
			return nil, fmt.Errorf("Invalid Request, line %s isn't a valid header.", line)
		}
		// Add validated header to map
		headers[line[0]] = strings.Join(line[1:], ": ")
	}

	// Check penulimate line empty
	if lines[len(lines)-2] != "" {
		return nil, fmt.Errorf("Invalid Request, penulimate line is nonempty")
	}

	// Build & return Request object
	body := lines[len(lines)-1] // body on final line
	request := Request{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
		Body:    body,
	}
	return &request, nil
}
