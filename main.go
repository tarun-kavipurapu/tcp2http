package main

import (
	"fmt"
	"http-go/internal"
	"net"
)

func main() {

	addr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 4555,
	}

	udpServer := internal.NewUdpServer(addr)
	err := udpServer.Listen()
	if err != nil {
		fmt.Println(err)
	}
}
