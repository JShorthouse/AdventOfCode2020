package main

import (
	"testing"
	"sort"
)

func TestFindSumPair(t *testing.T) {
	var tests = []struct {
		input []int
		expected int
	}{
		{
			[]int{ 1721, 979, 366, 299, 675, 1456 },
			514579,
		},
	}

	for _, test := range tests {
		sort.Ints(test.input)
		val1, val2, err := FindSumPair( test.input, 2020 )

		if err != nil {
			t.Error(err)
			continue
		}

		product := val1 * val2
		if product != test.expected {
			t.Errorf("Expected %d but got %d", test.expected, product)
		}
	}
}

func TestFindSumTriple(t *testing.T) {
	var tests = []struct {
		input []int
		expected int
	}{
		{
			[]int{ 1721, 979, 366, 299, 675, 1456 },
			241861950,
		},
	}

	for _, test := range tests {
		sort.Ints(test.input)
		val1, val2, val3, err := FindSumTriple( test.input, 2020 )

		if err != nil {
			t.Error(err)
			continue
		}

		product := val1 * val2 * val3
		if product != test.expected {
			t.Errorf("Expected %d but got %d", test.expected, product)
		}
	}
}
