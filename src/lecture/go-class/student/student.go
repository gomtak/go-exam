package student

import (
	"fmt"
	"lecture/go-class/persons"
)

type Students struct {
	persons.Persons
}

func (s Students) Add() {
	fmt.Println("Add Student")
}
func (s Students) Delete() {
	fmt.Println("Delete Student")
}
