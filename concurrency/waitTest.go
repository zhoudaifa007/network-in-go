package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		fmt.Println("goroutine done")
		close(exit)
	}()

	fmt.Println("main ...")
	<-exit
	fmt.Println("make exit")
}
