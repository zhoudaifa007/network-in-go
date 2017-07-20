package main

import (
	"log"
	"net"

	"os"

	"github.com/kalimatas/network-in-go"
)

const (
	CD  = "CD"
	DIR = "DIR"
	PWD = "PWD"
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

			conn.Close()
			return
		}

		cmd := string(buf[0:n])
		log.Println("Got cmd from client:", cmd)

		// decode request
		if cmd[0:2] == CD {
			chdir(conn, cmd[3:])
		} else if cmd[0:3] == DIR {
			dir(conn)
		} else if cmd[0:3] == PWD {
			pwd(conn)
		}
	}
}

func chdir(conn net.Conn, dir string) {
	if err := os.Chdir(dir); err == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR: " + err.Error()))
	}
}

func dir(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}

	for _, n := range names {
		conn.Write([]byte(n + "\r\n"))
	}
}

func pwd(conn net.Conn) {
	d, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(d))
}
