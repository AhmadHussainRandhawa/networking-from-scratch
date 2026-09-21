package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server:", conn.RemoteAddr())

	_, err = conn.Write([]byte("Hello\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: HELLO")
}
