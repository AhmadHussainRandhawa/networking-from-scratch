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

	fmt.Println("server listening on :9000")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	// Step 1: Receive HELLO
	n, err := conn.Read(buffer)
	if err != nil {
		if err == io.EOF {
			fmt.Println("client closed connection")
			return
		}

		log.Fatal(err)
	}

	message := string(buffer[:n])
	fmt.Printf("received: %q\n", message)

	if message != "HELLO\n" {
		fmt.Println("unexpected message")
		return
	}

	// Step 2: Send WELCOME
	_, err = conn.Write([]byte("WELCOME\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: WELCOME")

	// Step 3: Receive MESSAGE
	n, err = conn.Read(buffer)
	if err != nil {
		if err == io.EOF {
			fmt.Println("client closed connection")
			return
		}

		log.Fatal(err)
	}

	message = string(buffer[:n])
	fmt.Printf("received: %q\n", message)

	// Step 4: Acknowledge MESSAGE
	if len(message) >= len("MESSAGE ") &&
		message[:len("MESSAGE ")] == "MESSAGE " {

		fmt.Println("application message:", message[len("MESSAGE "):])

		_, err = conn.Write([]byte("ACK\n"))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("sent: ACK")
	}
}
