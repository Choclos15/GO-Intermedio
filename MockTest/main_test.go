package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id: 1,
			dni: "1",
			mockFunc: func() {
				GetEmoployeeById = func (id ind) (Employee, error) {
					return Employee{
						Id: 1,
						Position: "Programador",
					}, nil
				}

				GetPersonByDNI = func (id string) (Person, error)  {
					return Person{
						Name: "Mario Espinoza",
						Age: 22,
						DNI: "1",
					}, nil
				}
			},
		}
	}
}
