package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := printInterfaces(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := printAddresses(interfaces); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
