package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"math"
)

const DEFAULT_FILENAME = "./input/13"

func FindNextBus(bus_times []int, target int) (int, int) {
	closest_id := 0
	closest_total := math.MaxInt32

	for _, id := range bus_times {
		sum := 0
		for sum < target {
			sum += id
		}
		if sum < closest_total {
			closest_total = sum
			closest_id = id
		}
	}

	return closest_id, closest_total
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	time_line := scanner.Text()
	scanner.Scan()
	bus_line := scanner.Text()

	err = scanner.Err()
	if err != nil { panic(err) }


	bus_times := make([]int, 0, 50)

	for _, str := range strings.Split(bus_line, ",") {
		if str == "x" { continue }
		time, err := strconv.Atoi(str)
		if err != nil { panic(err) }

		bus_times = append(bus_times, time)
	}

	earliest_time, err := strconv.Atoi(time_line)
	if err != nil { panic(err) }

	id, time := FindNextBus(bus_times, earliest_time)

	p1_ans := id * (time - earliest_time)

	fmt.Printf("Part 1: %d\n", p1_ans)
}
