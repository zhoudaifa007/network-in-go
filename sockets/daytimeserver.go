package main

import (
	"log"
	"net"
	"time"

	"github.com/kalimatas/network-in-go"
)

func main() {
	service := "127.0.0.1:1222"

	lAddr, err := net.ResolveTCPAddr("tcp4", service)
	gonet.CheckError(err)

	ln, err := net.ListenTCP("tcp4", lAddr)
	gonet.CheckError(err)

	log.Println("Listening on", lAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
