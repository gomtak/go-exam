package main

import (
	"fmt"
	"time"
)

func main() {
	ping := time.Tick(100 * time.Millisecond)
	pong := time.After(500 * time.Millisecond)
	for {
		select {
		case <-ping:
			fmt.Println("ping")
		case <-pong:
			fmt.Println("pong")
			return //주석처리후 테스트
		default:
			fmt.Println("-")
			// time.Sleep(50 * time.Millisecond)
		}
	}
}
