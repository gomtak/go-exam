package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(sum(1, 2, 3))
}

func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}
