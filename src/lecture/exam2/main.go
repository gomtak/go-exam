package main

import (
	"fmt"
	"sync"
	"time"
)

func calc(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			wg.Done()
			fmt.Println(n + n)
			time.Sleep(200 * time.Millisecond)
		case <-quit:
			wg.Wait()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	go calc(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		ch <- i
	}

	quit <- true
	//~
}
