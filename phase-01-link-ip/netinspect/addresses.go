package main

import (
	"fmt"
	"net"
)

func printAddresses(interfaces []net.Interface) error {
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return fmt.Errorf("failed to get addresses for %s: %w", iface.Name, err)
		}

		fmt.Printf("Interface: %s\n", iface.Name)

		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				return fmt.Errorf(
					"failed to parse address %q on %s: %w",
					addr.String(),
					iface.Name,
					err,
				)
			}

			version := "IPv6"
			if ip.To4() != nil {
				version = "IPv4"
			}

			fmt.Printf("  %s: %s\n", version, addr)
		}

		fmt.Println()
	}

	return nil
}
