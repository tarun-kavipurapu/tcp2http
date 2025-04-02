package main

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
}

func NewServer() *Server {
	return &Server{}
}
func (t *Server) Listener() error {
	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		return fmt.Errorf("error starting listening %v", err)
	}
	t.listener = ln
	// fmt.Printf("Listening on 42069")
	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("error Accepting %v", err)
		}
		go t.handleConnections(conn)
	}
}

func (t *Server) handleConnections(conn net.Conn) {
	defer conn.Close()
	stream := ReadStream(conn)
	for line := range stream {
		fmt.Println(line)
	}
}
