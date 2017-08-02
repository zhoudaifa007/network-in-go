package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.StatusCode)
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}
}
