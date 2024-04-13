package main

import (
	"fmt"
	"strings"
)

type Response struct {
	StatusCode  int
	Content     string
	ContentType string
}

func (r Response) ToString() string {
	var response strings.Builder
	var status_code_message string

	status_code_str := fmt.Sprint(r.StatusCode)
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		status_code_message = status_code_str + " OK"
	} else if r.StatusCode == 404 {
		status_code_message = status_code_str + " Not Found"
	} else if r.StatusCode == 400 {
		status_code_message = status_code_str + " Bad Request"
	}

	first_line := fmt.Sprintf("HTTP/1.1 %v\r\n", status_code_message)
	response.WriteString(first_line)

	content_length := len(r.Content)
	if content_length > 0 {
		content_length := fmt.Sprint(content_length)
		content_type := fmt.Sprintf("Content-Type: %s\r\n", r.ContentType)
		content_length = fmt.Sprintf("Content-Length: %s\r\n", content_length)
		response.WriteString(content_type)
		response.WriteString(content_length)
		response.WriteString("\r\n")
		response.WriteString(r.Content)
	} else {
		response.WriteString("\r\n")
	}
	return response.String()
}

func handleGet(request *Request) string {
	split_path := strings.SplitN(request.Path[1:], "/", 2)

	if request.Path == "/" {
		return Response{200, "", ""}.ToString()
	}
	prefix := split_path[0]
	fmt.Println("prefix: " + prefix)
	switch prefix {
	case "echo":
		// If has echo prefix
		remainder := split_path[1]
		return Response{200, remainder, "text/plain"}.ToString()
	case "user-agent":
		// If user-agent path
		user_agent := request.Headers["User-Agent"]
		return Response{200, user_agent, "text/plain"}.ToString()
	case "files":
		// look for files in <directory>
		remainder := split_path[1]
		file_contents, err := findFile(remainder)
		if err != nil {
			return Response{404, "", ""}.ToString()
		}
		return Response{200, file_contents, "application/octet-stream"}.ToString()
	default:
		return Response{404, "", ""}.ToString()
	}
}

func handlePost(request *Request) string {
	split_path := strings.SplitN(request.Path[1:], "/", 2)

	if split_path[0] == "files" {
		file_path := split_path[1]
		file_contents := request.Body
		saveFile(file_path, file_contents)
		return Response{201, file_contents, "application/octet-stream"}.ToString()
	} else {

		return Response{404, "", ""}.ToString()
	}
}

func getResponse(request *Request) string {
	// If basic GET / return 200
	if request.Method == "GET" {
		return handleGet(request)
	} else if request.Method == "POST" {
		return handlePost(request)
	} else {
		return Response{404, "", ""}.ToString()
	}
}
