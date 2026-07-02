package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	fmt.Println("Listening on Port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	message := string(buf[:n])
	fmt.Println("Received:", message)

	if message == "Hey, I got your message." {
		conn.Write([]byte("Ok, I've accepted this message."))
		return
	}

	conn.Write([]byte("Unknown message"))
}
