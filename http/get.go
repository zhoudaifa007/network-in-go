package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if response.Status != "200 OK" {
		log.Fatal(response.Status)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	contentTypes := response.Header["Content-Type"]
	if !isAcceptableHeader(contentTypes) {
		fmt.Fprintln(os.Stderr, "Cannot handle", contentTypes)
		os.Exit(1)
	}

	var buf [512]byte
	bReader := response.Body
	for {
		n, err := bReader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
}

func isAcceptableHeader(contentTypes []string) bool {
	for _, cType := range contentTypes {
		if strings.Index(cType, "UTF-8") != -1 {
			return true
		}
	}
	return false
}
