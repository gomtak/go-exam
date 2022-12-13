package persons

import "fmt"

type Persons struct {
}

func (s Persons) Add() {
	fmt.Println("Add Persons")
}
func (s Persons) Delete() {
	fmt.Println("Delete Persons")
}
