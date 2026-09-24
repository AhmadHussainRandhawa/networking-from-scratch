package main

import (
	"bufio"
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

	reader := bufio.NewReader(conn)

	// -------------------------
	// Handshake
	// -------------------------

	line, err := readFrame(reader)
	if err != nil {
		log.Fatal(err)
	}

	command := ParseCommand(line)

	if command.Type != CommandHello {
		fmt.Println("expected HELLO")
		return
	}

	if err := writeResponse(conn, "WELCOME\n"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("sent: WELCOME")

	// -------------------------
	// Command loop
	// -------------------------

	for {
		line, err := readFrame(reader)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client closed connection")
				return
			}

			log.Fatal(err)
		}

		command := ParseCommand(line)

		switch command.Type {
		case CommandMessage:
			fmt.Println("application message:", command.Payload)

			if err := writeResponse(conn, "ACK\n"); err != nil {
				log.Fatal(err)
			}

			fmt.Println("sent: ACK")

		case CommandQuit:
			if err := writeResponse(conn, "BYE\n"); err != nil {
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

func readFrame(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(line, "\n"), nil
}

func writeResponse(conn net.Conn, response string) error {
	_, err := conn.Write([]byte(response))
	return err
}
