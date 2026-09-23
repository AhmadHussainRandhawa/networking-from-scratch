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

	fmt.Println("connected to server:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	// -------------------------
	// Handshake
	// -------------------------

	_, err = conn.Write([]byte("HELLO\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: HELLO")

	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("received: %q\n", buffer[:n])

	// -------------------------
	// Message 1
	// -------------------------

	_, err = conn.Write([]byte("MESSAGE hello server\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: MESSAGE hello server")

	n, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("received: %q\n", buffer[:n])

	// -------------------------
	// Message 2
	// -------------------------

	_, err = conn.Write([]byte("MESSAGE networking is interesting\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: MESSAGE networking is interesting")

	n, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("received: %q\n", buffer[:n])

	// -------------------------
	// Quit
	// -------------------------

	_, err = conn.Write([]byte("QUIT\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: QUIT")

	n, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("received: %q\n", buffer[:n])
}
