package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {

	var wg sync.WaitGroup


	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			fmt.Println("goroutine", id, "done.")
			wg.Done()
		}(i)
		time.Sleep(time.Second)
	}

	fmt.Println("main...")
	wg.Wait()
	fmt.Println("main exit.")

}
