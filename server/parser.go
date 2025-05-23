package server

import (
	"errors"
	"strings"
)

type Request struct {
	Method  string
	Path    []string
	Version string
	Headers map[string]string
	Body    []byte
}

func ParseRequest(raw []byte) (*Request, error) {
	reqString := string(raw)

	lines := strings.Split(reqString, "\r\n")

	// get request line and extract path dir
	request_line := strings.Split(lines[0], " ")
	path_dir := strings.Split(request_line[1], "/")

	// map for headers
	headers := make(map[string]string)

	// extract headers and add to map
	i := 1
	for {
		if lines[i] == "" {
			break
		}
		key_val := strings.Split(lines[i], ": ")
		headers[key_val[0]] = key_val[1]
		i++
	}

	// get optional body in bytes
	var body []byte

	if i+1 < len(lines) {
		body_str := strings.Join(lines[i+1:], "\r\n")
		body = []byte(body_str)
	}

	if (request_line[0] == "GET" || request_line[1] == "HEAD") && len(body) > 0 {
		return nil, errors.New("invalid request format")
	}

	req := Request{
		Method:  request_line[0],
		Path:    path_dir,
		Version: request_line[2],
		Headers: headers,
		Body:    body,
	}

	return &req, nil

}
