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

	_, err = conn.Write([]byte("HELLO\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sent: HELLO")

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buffer[:n]))

	_, err = conn.Write([]byte("MESSAGE hello server\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: MESSAGE hello server")

	// wait for ACK
	conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("received: %q\n", string(buffer[:n]))

}
