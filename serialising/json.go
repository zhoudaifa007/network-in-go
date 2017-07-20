package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type person struct {
	Name  name
	Email []email
}

type name struct {
	First string
	Last  string
}

type email struct {
	Kind    string
	Address string
}

var encode *bool

func init() {
	encode = flag.Bool("encode", true, "whether encode or decode")
	flag.Parse()
}

func main() {
	if *encode {
		person := person{
			Name: name{First: "Alex", Last: "Guz"},
			Email: []email{
				{Kind: "personal", Address: "email@fef.com"},
				{Kind: "not-personal", Address: "another@fei.com"},
			},
		}

		encodedPerson, err := json.Marshal(person)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error encoding %s", err)
			os.Exit(1)
		}

		fmt.Println(string(encodedPerson))

		//encoder := json.NewEncoder(os.Stdout)
		//if err := encoder.Encode(person); err != nil {
		//	fmt.Fprintf(os.Stderr, "error encoding %s", err)
		//	os.Exit(1)
		//}
	} else {
		var p person
		decoder := json.NewDecoder(os.Stdin)
		if err := decoder.Decode(&p); err != nil {
			fmt.Fprintf(os.Stderr, "error decoding %s", err)
			os.Exit(1)
		}

		fmt.Println(p)
	}
}
