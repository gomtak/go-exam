package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func question1() {
	var s string
	fmt.Scan(&s)
	fmt.Print(s)
	fmt.Println(s)
	fmt.Printf("hello %s", s)
}
func question2() {
	var name string
	var age int
	fmt.Println("이름을 입력하세요 : ")
	fmt.Scanln(&name)
	fmt.Println("나이를 입력하세요 : ")
	fmt.Scanln(&age)
	fmt.Println()
	fmt.Println("결과")
	fmt.Printf("이름 : %s, 나이 : %d", name, age)
}

func question3() {
	var i int
	fmt.Println(reflect.TypeOf(i))
}
func resInt(n int) int {
	return n
}

func resStr(s string) string {
	return s
}

func question4() {

	var l int = resInt(30)        // int 형식의 인자를 받고 그것을 return하는 이름 resInt인 함수 작성 필요
	var m string = resStr("codz") // string 형식의 인자를 받고 그것을 return하는 이름 resInt인 함수 작성 필요
	fmt.Println(l)
	fmt.Println(m)

}
func multiCalc(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return x, y
	// return a + b, a * b
}
func question5() {
	var a, b int = multiCalc(1, 2)
	fmt.Printf("합의 값 %d, 곱의 값 %d", a, b)
}
func tripleReturn() (a, b, c int) {
	return 1, 2, 3
}
func question6() {
	var a, b, c = tripleReturn()
	fmt.Printf("첫번쨰 : %d, 두번째 : %d,세번째 : %d", a, b, c)
}
func question7() {
	_, a, b := funcA()
	var c, _, d = funcB()
	var e, f, _ = funcC()
	var g, _, _ = funcD()

	fmt.Println(a, b, c, d, e, f, g)
}

func funcD() (a, b, c string) {
	return "D1", "공백", "공백"
}

func funcC() (a, b, c string) {
	return "C1", "C2", "공백"
}

func funcB() (a, b, c string) {
	return "B1", "공백", "B3"
}

func funcA() (a, b, c string) {
	return "공백", "A2", "A3"
}
func one() {
	fmt.Println("1111")
}

func two() {
	fmt.Println("2222")
}
func tree() {
	fmt.Println("3333")
}
func question8() {
	defer two()
	defer tree()
	one()
}
func question9() {
	file, err := os.OpenFile(
		"res.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	w.WriteString("Hello World!")
	w.Flush()

	r := bufio.NewReader(file)

	fi, _ := file.Stat()
	b := make([]byte, fi.Size())

	file.Seek(0, os.SEEK_SET)
	r.Read(b)

	fmt.Println(string(b))
}
func tempFunc(s ...string) {
	fmt.Println(s)
}
func question10() {
	tempFunc("a")
	tempFunc("a", "b")
	tempFunc("a", "b", "c")
}
func main() {
	// question1()
	// question2()
	// question3()
	// question4()
	// question5()
	// question6()
	// question7()
	// question8()
	// question9()
	// question10()
	var s = "abc"
	fmt.Println(s)
	d := "dbc"
	fmt.Println(d)
	// var s string = "codz"
	// fmt.Println("Hellow", "World", s)

	// s := "name" 새로운 변수가 추가되지 않는이상 이런식으로 재선언 되지 않음.
	// ./main.go:11:4: no new variables on left side of :=

	// s := "name"
	// fmt.Println("Hello", s)
	// print(s)

	// s := "name"
	// fmt.Printf("Hello %s", s)

	// s := 4567
	// msg := fmt.Sprint("Hello", "world", s)
	// fmt.Print(msg)

	// var s string
	// fmt.Scan(&s)
	// fmt.Print(s)

	a := 5
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 6

	fmt.Println(*b)
	fmt.Println(&a)
	fmt.Println(&b)

	// 포인터를 통해서 값 변경이 가능. 아니면 그냥 초기값이 뱉어짐.
	// var data Data

	// ChangeData(&data)

	// fmt.Println(data.value)
	// fmt.Println(data.data[100])

	// i := resInt(10)
	// s := resStr("name")

	// fmt.Print(i)

	// a, b := multiReturn()
	// fmt.Println(a)
	// fmt.Println(b)
}

//	func resInt(s int) int {
//		return 101
//	}
// func multiReturn() (n, m int) {
// 	return 123, 10
// }

// func multiReturn() (n string, m int) {
// 	return "ten", 10
// }

// type Data struct {
// 	value int
// 	data  [200]int // 200개의 원소를 가진 int형 배열
// }

// func ChangeData(d *Data) {
// 	d.value = 999
// 	d.data[100] = 999
// }
