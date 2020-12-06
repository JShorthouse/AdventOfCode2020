package main

import (
	"testing"
)

func TestFindSeat(t *testing.T) {
	var tests = []struct {
		input string
		e_row int
		e_col int
	}{
		{ "FBFBBFFRLR", 44, 5 },
		{ "BFFFBBFRRR", 70, 7 },
		{ "FFFBBBFRRR", 14, 7 },
		{ "BBFFBBFRLL", 102, 4 },
	}

	for _, test := range tests {
		row, col, err := FindSeat( test.input )
		if err != nil {
			t.Error(err)
			continue
		}

		if row != test.e_row || col != test.e_col {
			t.Errorf("Expected (%d, %d) but got (%d, %d)", test.e_row, test.e_col, row, col)
		}
	}
}

