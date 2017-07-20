package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Fprintln(os.Stderr, "Invalid address")
		os.Exit(1)
	}

	mask := addr.DefaultMask()
	ones, bits := mask.Size()
	fmt.Println("The address is ", addr.String(),
		" Default mask length is ", bits,
		" Ones are ", ones,
		" Mask is ", mask.String(),
		" Network is ", addr.Mask(mask))
}
