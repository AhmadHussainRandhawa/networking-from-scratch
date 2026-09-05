package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("=== Interfaces ===")
	if err := printInterfaces(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("=== Addresses ===")
	if err := printAddresses(interfaces); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("=== Routes ===")
	if err := printRoutes(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
