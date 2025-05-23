package server

import (
	"errors"
	"fmt"
	"net"
)

// struct to store port to listen (exported)
type Server struct {
	Addr string
}

// make TCP connection and handle exchange logic
func (s *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", ":"+s.Addr)

	if err != nil {
		return errors.New("could not bind to port: " + s.Addr)
	}

	fmt.Println("Successfully binded to port ", s.Addr)

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Could not accept incoming connection")
			continue
		}

		go handleConnection(conn)
	}
}

// handle individual incoming connections
func handleConnection(conn net.Conn) {
	fmt.Println("Handling connection from ", conn.RemoteAddr())

	defer conn.Close()

	//read, parse, respond

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Could not read from connection")
	}

	truncated := buffer[:n]
	req, err := ParseRequest(truncated)

	if err != nil {
		//return error response
	}

	fmt.Println(req)

}
