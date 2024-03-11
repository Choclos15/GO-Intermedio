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

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.person.name = "Mario"
	ftEmployee.person.age = 22
	fmt.Println(ftEmployee)
	GetMessage(ftEmployee.person)
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}
