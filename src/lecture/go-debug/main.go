package main

import (
	"fmt"
	clt "lecture/go-debug/controller" // user package import
)

func main() {
	sum := clt.Sum(3, 6)
	fmt.Println(sum)
}
