package main

import (
	"fmt"
	"io/ioutil"
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

	conn, err := net.Dial("tcp", service)
	gonet.CheckError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	gonet.CheckError(err)

	result, err := ioutil.ReadAll(conn)
	gonet.CheckError(err)

	fmt.Println(string(result))
}
