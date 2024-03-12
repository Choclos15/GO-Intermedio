package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)

	// if total != 9 {
	// 	t.Errorf("Sum Was Incorrect, Got %d Expected %d", total, 9)
	// }
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum Was Incorrect, Got %d Expected %d", total, item.n)
		}
	}
}
