package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const DEFAULT_FILENAME = "./input/01"
const DEFAULT_TARGET    = 2020

func FindSumPair( values []int, goal int ) (int, int, error) {
	for i := 0; i < len(values)-1; i++ {
		needed := goal - values[i]

		slice := values[i+1:]
		index := sort.SearchInts(slice, needed)

		if index < len(slice) && slice[index] == needed {
			return values[i], slice[index], nil
		}
	}
	return 0, 0, errors.New("Could not find pair")
}

func FindSumTriple( values []int, goal int ) (int, int, int, error) {
	for i := 0; i < len(values)-2; i++ {
		for j := i; j < len(values)-1; j++ {
			two_sum := values[i] + values[j]
			if two_sum > goal { break }

			needed := goal - two_sum;
			slice := values[j+1:]

			index := sort.SearchInts(slice, needed)

			if index < len(slice) && slice[index] == needed {
				return values[i], values[j], slice[index], nil
			}
		}
	}
	return 0, 0, 0, errors.New("Could not find triple")
}

func readIntInput(file_name string) []int {
	file, err := os.Open(file_name)
	if err != nil { panic(err) }

	defer file.Close()

	values := make([]int, 0, 500)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil { panic(err) }

		values = append(values, val)
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	return values
}

func main() {
	var file_n string
	var target int

	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }
	if len(os.Args) > 2 { target, _ = strconv.Atoi(os.Args[2]) } else { target = DEFAULT_TARGET }

	values := readIntInput(file_n)
	sort.Ints(values)
	
	val1, val2, err := FindSumPair(values, target)
	if err != nil { panic(err) }

	p1_ans := val1 * val2
	fmt.Printf("Part 1: %d\n", p1_ans)

	val1, val2, val3, err := FindSumTriple(values, target)
	if err != nil { panic(err) }
	
	p2_ans := val1 * val2 * val3
	fmt.Printf("Part 2: %d\n", p2_ans)
}
