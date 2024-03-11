package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func main() {
	//forma 1
	e := Employee{}
	fmt.Println(e)
	//Forma 2
	e2 := Employee{id: 1, name: "Name", vacation: true}
	fmt.Println(e2)
	//Forma 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 1
	e3.name = "Name"
	fmt.Printf("%v\n", *e3)
	//Forma 4
	e4 := NewEmployee(2, "Name 2", true)
	fmt.Println(*e4)

}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}
