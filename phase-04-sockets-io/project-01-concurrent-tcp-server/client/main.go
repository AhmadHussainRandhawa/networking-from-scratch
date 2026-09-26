package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println("dial error:", err)
				return
			}
			defer conn.Close()

			message := fmt.Sprintf("hello from client %d\n", id)

			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Println("write error:", err)
				return
			}

			buffer := make([]byte, 1024)

			_, err = conn.Read(buffer)
			if err != nil {
				fmt.Println("read error:", err)
				return
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("all clients finished")
}
