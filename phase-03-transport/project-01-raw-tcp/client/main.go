package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("connected to:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	// HELLO
	if _, err := conn.Write([]byte("HELLO\n")); err != nil {
		log.Fatal(err)
	}

	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server: %q\n", response)

	// MSG
	if _, err := conn.Write([]byte("MSG hello server\n")); err != nil {
		log.Fatal(err)
	}

	response, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server: %q\n", response)

	// Another MSG
	if _, err := conn.Write([]byte("MSG TCP is a byte stream\n")); err != nil {
		log.Fatal(err)
	}

	response, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server: %q\n", response)

	// QUIT
	if _, err := conn.Write([]byte("QUIT\n")); err != nil {
		log.Fatal(err)
	}

	response, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server: %q\n", response)
}
