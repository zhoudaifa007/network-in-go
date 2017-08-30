package main

import "fmt"

type tester interface {
	test()
	string() string
}

type data struct {
	age int
}

func (h *data) test() {
	fmt.Println("helo")
	h.age = 1
}
func (data) string() string { return "boy" }

func main() {
	var d data

	//var t tester = d
	var t tester = &d

	t.test()
	fmt.Println(d.age)

	fmt.Println(t.string())

}
