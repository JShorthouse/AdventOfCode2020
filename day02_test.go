package main

import (
	"testing"
)

func TestCheckPasswordCount(t *testing.T) {
	var tests = []struct {
		input []string
		expected int
	}{
		{
			[]string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			2,
		},
	}

	for _, test := range tests {
		num, err := CheckPasswordCount( test.input )
		if err != nil {
			t.Error(err)
			continue
		}

		if num != test.expected {
			t.Errorf("Expected %d but got %d", test.expected, num)
		}
	}
}

func TestCheckPasswordIndex(t *testing.T) {
	var tests = []struct {
		input []string
		expected int
	}{
		{
			[]string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			1,
		},
	}

	for _, test := range tests {
		num, err := CheckPasswordIndex( test.input )
		if err != nil {
			t.Error(err)
			continue
		}

		if num != test.expected {
			t.Errorf("Expected %d but got %d", test.expected, num)
		}
	}
}
