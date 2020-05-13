package main

import "testing"

func TestSoma(t *testing.T) {
	testTables := []struct {
		x int
		y int
		r int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 2},
		{1, 2, 3},
		{2, 1, 3},
	}

	for _, table := range testTables {
		resultado := soma(table.x, table.y)
		if resultado != table.r {
			t.Errorf("Error: Got: %d, Expected: %d", resultado, table.r)
		}
	}

}
