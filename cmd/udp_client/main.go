package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:4555")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	msg := " fucker"
	fmt.Printf("Ping: %v\n", msg)
	if _, err = conn.Write([]byte(msg)); err != nil {
		fmt.Printf("Write err %v", err)
		os.Exit(-1)
	}
}
