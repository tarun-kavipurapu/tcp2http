package main

func main() {
	server := NewServer()
	server.Listener().Error()
}
