package main

import (
	"time"
	"fmt"
)

var c int

func counter() int {
	c++
	return c
}

func main() {
	a := 100


	a += 100

	fmt.Println("main:",a, counter())

	go func(x,y int) {
		time.Sleep(time.Second)
		fmt.Println("go:",x,y)
	}(a,counter())



	time.Sleep(time.Second * 3)
}
