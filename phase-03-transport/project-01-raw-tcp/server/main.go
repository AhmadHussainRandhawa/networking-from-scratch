package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	// -------------------------
	// Handshake
	// -------------------------

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
		fmt.Println("expected HELLO")
		return
	}

	_, err = conn.Write([]byte("WELCOME\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: WELCOME")

	// -------------------------
	// Command loop
	// -------------------------

	for {
		n, err = conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client closed connection")
				return
			}

			log.Fatal(err)
		}

		message = string(buffer[:n])
		message = strings.TrimSuffix(message, "\n")

		fmt.Printf("received: %q\n", message)

		switch {
		case strings.HasPrefix(message, "MESSAGE "):
			text := strings.TrimPrefix(message, "MESSAGE ")

			fmt.Println("application message:", text)

			_, err = conn.Write([]byte("ACK\n"))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("sent: ACK")

		case message == "QUIT":
			_, err = conn.Write([]byte("BYE\n"))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("sent: BYE")
			fmt.Println("closing connection")

			return

		default:
			fmt.Println("unknown command")
		}
	}
}
