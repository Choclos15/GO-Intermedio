package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "Name"
	fmt.Println(e)
	e.SetId(5)
	e.SetName("Name 2")
	fmt.Println(e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e Employee) GetId() int {
	return e.id
}

func (e Employee) GetName() string {
	return e.name
}
