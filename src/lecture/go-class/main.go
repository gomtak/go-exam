package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
	Pnum string
}

func main() {
	p := Person{Name: "codz", Age: 23, Pnum: "01098765432"}

	encoder := json.NewEncoder(os.Stdout)     //인코드 선언
	if err := encoder.Encode(p); err != nil { //인코드 실행
		panic(err)
	}

	fmt.Println(p)
	//json 문자열로 반환
	jdata, _ := json.Marshal(p)
	fmt.Println(jdata)

}
