package main

import (
	"errors"
	"fmt"
)

var res = 1

func Calc(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if n == 0 {
		panic(errors.New("zero calc"))
	} else {
		res *= n
		fmt.Println("Result: ", res)
	}
}

func main() {
	Calc(1)
	Calc(2)
	Calc(0)
	Calc(3)
}
