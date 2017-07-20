package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	digits := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	buf := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	encoder.Write(digits)
	encoder.Close()

	fmt.Println(buf)

	dbuf := make([]byte, 12)
	decoder := base64.NewDecoder(base64.StdEncoding, buf)
	decoder.Read(dbuf)

	for _, ch := range dbuf {
		fmt.Print(ch)
	}
	//fmt.Println(dbuf)
}
