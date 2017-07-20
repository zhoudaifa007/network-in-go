package main

import (
	"log"
	"net"

	"github.com/kalimatas/network-in-go"
)

func main() {
	service := "127.0.0.1:1223"

	ln, err := net.Listen("tcp", service)
	gonet.CheckError(err)

	log.Println("Listening on", service)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println("error reading from client:", err)
			return
		}

		log.Println("Got from client:", string(buf[0:]))

		_, err = conn.Write(buf[0:n])
		if err != nil {
			log.Println("error writing to client:", err)
			return
		}
	}
}
