package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id   int
	name string
}

type FullTimeEmployee struct {
	person   Person
	Employee Employee
}

type TemporaryEmpleyee struct {
	person   Person
	employee Employee
	taxRate  int
}

type PrintInfo interface {
	getMessage() string
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.person.name = "Mario"
	ftEmployee.person.age = 22
	fmt.Println(ftEmployee)
	GetMessage(ftEmployee.person)
	tMessage := TemporaryEmpleyee{}
	getMessage(tMessage)
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func (fTEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func (tEmployee TemporaryEmpleyee) getMessage() string {
	return "Temporary Time Employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}
