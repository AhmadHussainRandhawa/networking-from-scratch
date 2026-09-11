package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.23:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Client Connected to:", conn.RemoteAddr())

	message := "HELLo\n"

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(conn)

	message, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("received: %q\n", message)

}
