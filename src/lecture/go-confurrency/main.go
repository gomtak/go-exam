package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(s string, i int) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s : %d \n", s, i)
	}
}

func main() {
	var wg sync.WaitGroup
	// goroutine 에서 사용한 함수 사용
	wg.Add(1)
	go counter("a", 10)
	wg.Done()
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter("b", 10)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println("c : ", i)
		}
	}()
	wg.Wait()
}
