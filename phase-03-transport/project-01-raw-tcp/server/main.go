package main

import (
	"bufio"
	"fmt"
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

	reader := bufio.NewReader(conn)

	// Read HELLO.
	message, err := reader.ReadString('\n') // Read till \n, and return including \n.
	if err != nil {
		log.Fatal(err)
	}

	message = strings.TrimSuffix(message, "\n")
	fmt.Printf("received: %q\n", message)

	if message != "HELLO" {
		fmt.Println("protocol error: expected HELLO")
		return
	}

	response := "WELCOME\n"

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sent: %q\n", response)

	// Process commands until the client sends QUIT.
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		message = strings.TrimSuffix(message, "\n")
		fmt.Printf("received: %q\n", message)

		switch {

		case strings.HasPrefix(message, "MSG "):
			text := strings.TrimPrefix(message, "MSG ")

			response = "MSG " + text + "\n"

			if _, err = conn.Write([]byte(response)); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("sent: %q\n", response)

		case message == "QUIT":
			if _, err = conn.Write([]byte("Bye\n")); err != nil {
				log.Fatal(err)
			}

			fmt.Println("sent: \"BYE\\n\"")
			fmt.Println("client requested disconnect")
			return

		default:
			if _, err := conn.Write([]byte("ERROR unknown command\n")); err != nil {
				log.Fatal(err)
			}
		}
	}
}
