package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	fmt.Println("connected:", addr)

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			if err == io.EOF {
				fmt.Println("closed by client:", addr)
			} else {
				fmt.Println("read error:", err)
			}
			return
		}

		if n == 0 {
			continue
		}

		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("write error:", err)
			return
		}
	}
}
