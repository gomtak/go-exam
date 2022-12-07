package main

import (
	"fmt"
	"reflect"
)

func qeustion1() {
	var strArray1 [5]string
	strArray1[0] = "a"
	strArray1[1] = "b"
	strArray1[2] = "c"
	strArray1[3] = "d"
	strArray1[4] = "e"
	fmt.Println(strArray1)
}

func qeustion2() {
	var strArray2 [5]string = [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(strArray2)
}

func qeustion3() {
	strArray3 := [5]string{"10", "20", "30"}
	fmt.Println(strArray3)
}

func qeustion4() {
	strArray4 := [...]string{"100", "200", "300"}
	fmt.Println(strArray4)
	fmt.Println(len(strArray4))
}

func qeustion5() {
	strArray5 := [5]string{2: "codz", 4: "states"}
	fmt.Println(strArray5)
	fmt.Printf("%#v\n", strArray5)
	fmt.Printf("%v\n", strArray5)
}
func qeustion6() {
	var a1 []int
	var a2 [5]int

	fmt.Println(reflect.ValueOf(a1).Kind())
	fmt.Println(reflect.ValueOf(a1))
	fmt.Println(reflect.TypeOf(a1))

	fmt.Println(reflect.ValueOf(a2).Kind())
	fmt.Println(reflect.TypeOf(a2))
}
func qeustion7() {
	var intSlice = []int{0, 1, 2, 3}
	fmt.Printf("%#v\n", intSlice)
	fmt.Printf("%v\n", intSlice)
	fmt.Printf(":2 %v\n", intSlice[:2])
	fmt.Printf("1:3 %v\n", intSlice[1:3])
	fmt.Printf("2: %v\n", intSlice[2:])
	fmt.Printf("0:3 %v\n", intSlice[0:3])
}
func qeustion8() {
	// var map1 map[string]int

	// map2 := make(map[string]string)
	// map3 := make(map[int]string, 1000)

	var timeZone = map[string]int{
		"UTC": 0,
		"EST": 5,
		"CST": 6,
		"MST": 7,
		"PST": 8,
	}
	for s, _ := range timeZone {
		println(s)
	}

	// timeSet := timeZone["CST"]
}

func qeustion9() {
	map1 := make(map[string]int)
	map1 = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range map1 {
		fmt.Printf("키 : %s , 값 : %d \n", k, v)
	}
	fmt.Println(map1)
}

func qeustion10() {
	type Persons struct {
		Name string
		Age  int
		Pnum string
	}
	var pers Persons
	pers.Name = "codz"
	pers.Age = 23
	pers.Pnum = "01098765432"

	fmt.Println(pers)
	fmt.Printf("%#v\n", pers)
	fmt.Println("name : ", pers.Name, " age : ", pers.Age, " pnum : ", pers.Pnum)
}
func qeustion11() {
	type Persons struct {
		Name string
		Age  int
		Pnum string
	}

	p1 := new(Persons)
	p1.Name = "codz"
	p1.Age = 29
	p1.Pnum = "01098765432"
	fmt.Println("Person 1 : ", p1)
	fmt.Printf("%#v\n", p1)

	var p2 = new(Persons)
	p2.Name = "states"
	p2.Age = 33
	p2.Pnum = "01054329876"
	fmt.Println("Person 2 : ", p2)
	fmt.Printf("%#v\n", p2)

	//선언과 동시에 초기화
	var p3 = &Persons{"test", 17, "01045673210"}
	fmt.Println("Person 3 : ", p3)
	fmt.Printf("%#v\n", p3)

	//구조체 배열 형식
	p4 := []Persons{
		Persons{
			Name: "p1",
			Age:  10,
			Pnum: "9876",
		},
		Persons{
			Name: "p2",
			Age:  9,
			Pnum: "8765",
		},
	}
	fmt.Println(p4)
	fmt.Printf("%#v\n", p4)
}
func qeustion12() {

	//예제1 : 인터페이스 데이터 선언
	var x interface{}
	x = 1
	fmt.Println(x)

	x = "Hello World"
	fmt.Println(x)

	//예제2 : 인터페이스 데이터 타입 변환과 반환값
	// var x interface{}
	x = "10"

	sVal := x.(string)
	fmt.Println(sVal)
	x = 10
	nVal, bCnv := x.(int)
	fmt.Println(nVal, bCnv)

	//예제3 : 타입에 따른 핸들링 기법
	// var x interface{}
	x = "hello"
	xType := reflect.TypeOf(x)

	if xType.Kind() == reflect.Int {
		fmt.Println("Int ", x)
	} else if xType.Kind() == reflect.String {
		fmt.Println("string ", x)
	}
	type Persons struct {
		Name string
		Age  int
	}
	//예제4 : 구조체를 선언하고, 구조체별 인터페이스 핸들링
	codz := Persons{Name: "codz", Age: 20}

	// var x interface{}
	x = codz

	xPerson := x.(Persons)
	xPrsAge := x.(Persons).Age
	ty := reflect.TypeOf(xPerson)
	fmt.Printf("%v, %d, %s\n", xPerson, xPrsAge, xPerson.Name)
	fmt.Println(ty.String())

}
func main() {
	// qeustion1()
	// qeustion2()
	// qeustion3()
	// qeustion4()
	// qeustion5()
	// qeustion6()
	// qeustion7()
	// qeustion8()
	// qeustion9()
	// qeustion10()
	// qeustion11()
	qeustion12()
}
