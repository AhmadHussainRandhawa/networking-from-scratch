package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		if err == io.EOF {
			fmt.Println("client closed connection.")
			return
		}
		log.Fatal(err)
	}

	message := string(buffer[:n])

	fmt.Printf("received %d bytes: %q\n", n, message)

	if message == "HELLO\n" {
		_, err := conn.Write([]byte("WELCOME\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("sent: WELCOME")
}
