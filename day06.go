package main

import (
	"os"
	"fmt"
	"bufio"
)

const DEFAULT_FILENAME = "./input/06"

func countPresent(arr []int) int {
	count := 0
	for _, val := range arr {
		if val > 0 { count++ }
	}
	return count;
}

func countEqual(arr []int, target int) int {
	count := 0
	for _, val := range arr {
		if val == target { count++ }
	}
	return count;
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()


	var group_answers = make([]int, 26)

	any_count := 0
	all_count := 0

	var line string
	var group_size int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		// End of group
		if line == "" {
			any_count += countPresent(group_answers)
			all_count += countEqual(group_answers, group_size)

			// Reset vars
			for i, _ := range group_answers {
				group_answers[i] = 0;
			}
			group_size = 0;

			continue
		}

		for _, char := range line {
			group_answers[int(char) - int('a')] += 1
		}

		group_size++
	}

	// Edge case for final line
	if line != "" {
		any_count += countPresent(group_answers)
		all_count += countEqual(group_answers, group_size)
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	fmt.Printf("Part 1: %d\n", any_count)
	fmt.Printf("Part 2: %d\n", all_count)
}
