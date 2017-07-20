package main

import (
	"fmt"
	"net"
	"os"

	"github.com/kalimatas/network-in-go"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	addr, err := net.ResolveUDPAddr("udp4", service)
	gonet.CheckError(err)

	conn, err := net.DialUDP("udp4", nil, addr)
	gonet.CheckError(err)

	_, err = conn.Write([]byte("whatever"))
	gonet.CheckError(err)

	var buf [512]byte

	n, err := conn.Read(buf[0:])
	gonet.CheckError(err)

	fmt.Println(string(buf[0:n]))
}
