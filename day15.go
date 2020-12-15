package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

const DEFAULT_FILENAME = "./input/15"


func FindNumber( initial_nums []int, target int ) int {
	last_spoken := make([]int, target, target)

	i := 1
	cur_num := 0
	prev_num := initial_nums[0]

	for _, num := range initial_nums {
		last_spoken[prev_num] = i-1
		prev_num = num
		i++
	}

	for i <= target {
		if last_spoken[prev_num] > 0 {
			cur_num = (i-1) - last_spoken[prev_num]
		} else {
			cur_num = 0
		}
		last_spoken[prev_num] = i-1
		prev_num = cur_num

		i++
	}

	return cur_num
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	initial_nums := make([]int, 0, 10)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	for _, str := range strings.Split(line, ",") {
		val, err := strconv.Atoi(str)
		if err != nil { panic(err) }

		initial_nums = append(initial_nums, val)
	}

	err = scanner.Err()
	if err != nil { panic(err) }


	p1_ans := FindNumber(initial_nums, 2020)
	fmt.Printf("Part 1: %d\n", p1_ans)

	p2_ans := FindNumber(initial_nums, 30000000)
	fmt.Printf("Part 2: %d\n", p2_ans)
}
