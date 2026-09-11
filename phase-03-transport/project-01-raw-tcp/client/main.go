package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.23:9000")
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()

	fmt.Println("Client Connected to:", conn.RemoteAddr())

}
