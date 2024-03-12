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
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmoployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "Programador",
					}, nil
				}

				GetPersonByDNI = func(id string) (Person, error) {
					return Person{
						Name: "Mario Espinoza",
						Age:  22,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  22,
					DNI:  "1",
					Name: "Mario Espinoza",
				},
				Employee: Employee{
					Id:       1,
					Position: "Programador",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmoployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {

			t.Errorf("Error When Getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, Got %d Expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmoployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
