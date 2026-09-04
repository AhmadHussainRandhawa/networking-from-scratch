package main

import (
	"fmt"
	"net"
)

func printInterfaces() error {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, iface := range interfaces {
		fmt.Printf("Interface: %s\n", iface.Name)
		fmt.Printf("  Index: %d\n", iface.Index)
		fmt.Printf("  MTU: %d\n", iface.MTU)
		fmt.Printf("  MAC: %s\n", iface.HardwareAddr)
		fmt.Printf("  Flags: %s\n", iface.Flags)
		fmt.Println()
	}

	return nil
}
