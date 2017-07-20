package main

import (
	"log"
	"net"

	"time"

	"github.com/kalimatas/network-in-go"
)

func main() {
	service := "127.0.0.1:1222"

	lAddr, err := net.ResolveUDPAddr("udp4", service)
	gonet.CheckError(err)

	ln, err := net.ListenUDP("udp4", lAddr)
	gonet.CheckError(err)

	log.Println("Listening on", lAddr)

	for {
		handleUDPClient(ln)
	}
}

func handleUDPClient(conn *net.UDPConn) {
	var buf [512]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()

	log.Println("Writing", daytime, "to", addr.String())
	conn.WriteToUDP([]byte(daytime), addr)
}
