package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	conn.Write([]byte("Hey, I got your message."))

	resp, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Server Said :", resp)
}
