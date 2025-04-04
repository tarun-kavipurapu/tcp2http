package internal

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func readStream(read io.ReadCloser) <-chan string {
	lineChan := make(chan string)
	// fmt.Println("debug")
	go func() {
		defer close(lineChan)
		defer read.Close()
		currLine := ""

		for {
			buff := make([]byte, 8)
			n, err := read.Read(buff)
			if err != nil {
				if currLine != "" {
					lineChan <- currLine
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error:%s\n", err.Error())
				return
			}
			str := string(buff[:n])
			fmt.Println(str)
			parts := strings.Split(str, "\n")
			fmt.Println(parts, "after split")

			for i := 0; i < len(parts)-1; i++ {
				lineChan <- fmt.Sprintf("%s%s", currLine, parts[i])
				currLine = ""
			}
			currLine += parts[len(parts)-1]
		}

	}()
	return lineChan
}

func readUDP(conn *net.UDPConn) <-chan string {
	packetChan := make(chan string)
	go func() {
		defer close(packetChan)
		buff := make([]byte, 1024) // Buffer size for UDP packets

		for {
			n, addr, err := conn.ReadFromUDP(buff)
			if err != nil {
				fmt.Printf("error reading from UDP: %s\n", err.Error())
				return
			}
			fmt.Printf("received packet from %s\n", addr.String())
			packetChan <- string(buff[:n])
		}
	}()
	return packetChan
}
