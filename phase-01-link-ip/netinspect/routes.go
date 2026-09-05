package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	Interface   string
	Destination net.IP
	Gateway     net.IP
	Metric      uint32
	Mask        net.IP
}

func printRoutes() error {
	file, err := os.Open("/proc/net/route")
	if err != nil {
		return fmt.Errorf("failed to open /proc/net/route: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skip the header line.
	if !scanner.Scan() {
		return fmt.Errorf("failed to read /proc/net/route header")
	}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		// We expect:
		// Iface Destination Gateway Flags RefCnt Use Metric Mask MTU Window IRTT
		if len(fields) < 8 {
			return fmt.Errorf("unexpected route line: %q", scanner.Text())
		}

		iface := fields[0]

		destination, err := decodeRouteIPv4(fields[1])
		if err != nil {
			return fmt.Errorf("invalid destination %q: %w", fields[1], err)
		}

		gateway, err := decodeRouteIPv4(fields[2])
		if err != nil {
			return fmt.Errorf("invalid gateway %q: %w", fields[2], err)
		}

		metric, err := strconv.ParseUint(fields[6], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid metric %q: %w", fields[6], err)
		}

		mask, err := decodeRouteIPv4(fields[7])
		if err != nil {
			return fmt.Errorf("invalid mask %q: %w", fields[7], err)
		}

		destination = applyMask(destination, mask)

		fmt.Printf("Interface: %s\n", iface)
		fmt.Printf("  Destination: %s\n", formatDestination(destination, mask))
		fmt.Printf("  Gateway: %s\n", formatGateway(gateway))
		fmt.Printf("  Metric: %d\n", metric)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading /proc/net/route: %w", err)
	}

	return nil
}

func decodeRouteIPv4(hexValue string) (net.IP, error) {
	value, err := strconv.ParseUint(hexValue, 16, 32)
	if err != nil {
		return nil, err
	}

	b0 := byte(value)
	b1 := byte(value >> 8)
	b2 := byte(value >> 16)
	b3 := byte(value >> 24)

	return net.IPv4(b0, b1, b2, b3), nil
}

func applyMask(ip net.IP, mask net.IP) net.IP {
	ip4 := ip.To4()
	mask4 := mask.To4()

	return net.IPv4(
		ip4[0]&mask4[0],
		ip4[1]&mask4[1],
		ip4[2]&mask4[2],
		ip4[3]&mask4[3],
	)
}

func formatDestination(destination net.IP, mask net.IP) string {
	mask4 := mask.To4()

	if mask4[0] == 0 &&
		mask4[1] == 0 &&
		mask4[2] == 0 &&
		mask4[3] == 0 {
		return "default"
	}

	prefixLength, _ := net.IPMask(mask4).Size()

	return fmt.Sprintf("%s/%d", destination, prefixLength)
}

func formatGateway(gateway net.IP) string {
	gateway4 := gateway.To4()

	if gateway4[0] == 0 &&
		gateway4[1] == 0 &&
		gateway4[2] == 0 &&
		gateway4[3] == 0 {
		return "-"
	}

	return gateway4.String()
}
