package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
		fmt.Println("end")
	}()
	ch <- 1
	wg.Wait()
}
