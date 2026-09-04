package main

import (
	"fmt"
	"os"
)

func main() {
	err := printInterfaces()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
