package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"sort"
	"errors"
)


const DEFAULT_FILENAME = "./input/09"
const PREAMBLE_SIZE    = 25

func FindWithoutSum(nums []int, prev_size int) (int, error) {
	prev := make([]int, prev_size)

	for i:=prev_size; i<len(nums); i++ {
		copy(prev, nums[i-prev_size:i])
		sort.Ints(prev)

		cur_num := nums[i]
		found := false

		for j:=0; j<prev_size-1; j++ {
			if prev[j] > cur_num { break }

			for k:=j; k<prev_size; k++ {
				sum := prev[j] + prev[k]

				if sum == cur_num {
					found = true
					goto end_loop
				} else if sum > cur_num {
					break
				}
			}
		}

end_loop:
		if !found {
			return cur_num, nil
		}
	}

	return 0, errors.New("Could not find sum")
}

func FindContSum(nums []int, target int) ([]int, error) {
	for i:=0; i<len(nums); i++ {
		slice := nums[i:]

		sum := 0
		// Binary search for slice from i with sum equal to target
		r := sort.Search(len(slice), func(i int) bool {
			sum = 0
			for j:=0; j<i; j++ {
				sum += slice[j]
			}
			return sum >= target
		})

		// Discard no match or only single number
		if r == -1 { continue }
		if r == 1 { continue }

		if sum == target {
			return nums[i:i+r], nil
		}
	}

	return []int{}, errors.New("Slice not found")
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	numbers := make([]int, 0, 500)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil { panic(err) }
		numbers = append(numbers, num)
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	p1_ans, err  := FindWithoutSum( numbers, PREAMBLE_SIZE )
	if err != nil { panic(err) }

	fmt.Printf("Part 1: %d\n", p1_ans)

	slice, err := FindContSum(numbers, p1_ans)
	if err != nil { panic(err) }

	sort.Ints(slice)

	p2_ans := slice[0] + slice[len(slice)-1]

	fmt.Printf("Part 2: %d\n", p2_ans)
}
