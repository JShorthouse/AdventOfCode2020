package main

import (
	"os"
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

const DEFAULT_FILENAME = "./input/10"
const MAX_DIFFERENCE = 3


func CountDiffs(numbers []int) ([MAX_DIFFERENCE+1]int, bool) {
	last_joltage := 0
	var difference_count [MAX_DIFFERENCE+1]int

	for _, jolts := range numbers[1:] {
		diff := jolts - last_joltage
		if diff > MAX_DIFFERENCE {
			return difference_count, false
		}
		difference_count[diff]++

		last_joltage = jolts
	}
	return difference_count, true
}

type Chunk struct {
	start_idx int
	end_idx int
	valid_combinations int
}


// Find all groups containing numbers with differences of 1 between them
func FindRemovableChunks(numbers []int) []Chunk {
	chunks := make([]Chunk, 0, 10)

	last_joltage := 0
	last_diff := 0

	chunk_start := 0
	chunk_end := 0
	chunk_active := false

	for i, jolts := range numbers {
		diff := jolts - last_joltage
		
		if diff == 1 && last_diff == 1 {
			if chunk_active {
				chunk_end = i-1
			} else {
				chunk_start = i-1
				chunk_end = i-1
				chunk_active = true
			}
		} else {
			if chunk_active {
				var c Chunk
				c.start_idx = chunk_start
				c.end_idx = chunk_end
				chunks = append(chunks, c)

				chunk_active = false
			}
		}

		last_joltage = jolts
		last_diff = diff
	}

	// Edge case for end
	if chunk_active {
		var c Chunk
		c.start_idx = chunk_start
		c.end_idx = chunk_end
		chunks = append(chunks, c)
	}

	return chunks
}

// Count all valid subsets of the chunk
// e.g. every combination of value removals where the difference between values
// is still <= MAX_DIFFERENCE
func ValidChunkCombinations(numbers []int, chunk Chunk) int {
	chunk_len := (chunk.end_idx - chunk.start_idx) + 1

	// Chunk with surrounding number at either end
	chunk_context := numbers[chunk.start_idx-1:chunk.end_idx+2]

	valid_combos := 0
	total_combinations := 1 << (chunk_len) // chunk_len^2

	for combination := 0; combination < total_combinations; combination++ {
		valid := true
		last_joltage := chunk_context[0]

		for i, jolts := range chunk_context[1:] {

			// Binary combination generation
			// Use binary digit presence to determine which indices of chunk to
			// keep for this combination
			if (1 << i) & combination > 0 {
				continue
			}

			diff := jolts - last_joltage
			if diff > MAX_DIFFERENCE {
				valid = false
				break
			}
			last_joltage = jolts
		}

		if valid {
			valid_combos++
		}
	}

	return valid_combos
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	numbers := make([]int, 0, 500)

	numbers = append(numbers, 0) // Start joltage

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil { panic(err) }
		numbers = append(numbers, num)
	}

	err = scanner.Err()
	if err != nil { panic(err) }


	sort.Ints(numbers)

	numbers = append(numbers, numbers[len(numbers)-1] + 3) // Device joltage

	difference_count, valid := CountDiffs(numbers)
	if !valid { panic("Invalid input") }

	p1_ans := difference_count[1] * difference_count[3]
	fmt.Printf("Part 1: %d\n", p1_ans)


	chunks := FindRemovableChunks(numbers)

	p2_ans := 1
	for _, chunk := range chunks {
		p2_ans *= ValidChunkCombinations(numbers, chunk)
	}

	fmt.Printf("Part 2: %d\n", p2_ans)
}
