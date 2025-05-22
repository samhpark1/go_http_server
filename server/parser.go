package server

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

func ParseRequest(raw []byte) /*(*Request, error)*/ {
	reqString := string(raw)

	lines := strings.Split(reqString, "\r\n")

	fmt.Println(lines)
}
