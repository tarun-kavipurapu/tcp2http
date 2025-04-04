package internal

import (
	"fmt"
	"net"
)

type UdpServer struct {
	listner *net.UDPConn
	addr    net.UDPAddr
}

func NewUdpServer(addr net.UDPAddr) *UdpServer {
	return &UdpServer{
		addr: addr,
	}

}

func (u *UdpServer) Listen() error {
	listn, err := net.ListenUDP("udp", &u.addr)

	if err != nil {
		return fmt.Errorf("error listening UDP %v", err)
	}

	u.listner = listn

	return u.acceptLoop()
}
func (u *UdpServer) acceptLoop() error {
	for {
		/*
			In UDP there is no need to fucking accept a connection before sending the data the data can be directly sent
				conn, err := u.listner.A
				if err != nil {
					return fmt.Errorf("error accepting connection %s", err)
				}
		*/
		// buff := make([]byte, 256)
		// n, err := u.listner.Read(buff)
		// if err != nil {
		// 	fmt.Errorf(err.Error())
		// }
		// fmt.Println(string(buff[:n]))
		lines := readStream(u.listner)
		for line := range lines {
			fmt.Println(line)
		}

	}
}

// func (u *UdpServer) handleConn(conn net.Conn) {
// 	defer conn.Close()
// 	for {
// 		lines := readStream(conn)
// 		for line := range lines {
// 			fmt.Println(line)
// 		}
// 		_, err := conn.Write([]byte("Received the data yo"))
// 		if err != nil {
// 			fmt.Printf("Error writing %s", err)
// 		}
// 	}

// }
