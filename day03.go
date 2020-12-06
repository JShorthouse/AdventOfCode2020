package main

import (
	"fmt"
	"os"
	"bufio"
)

const DEFAULT_FILENAME = "./input/03"
const TREE = '#'

func CountTrees(lines []string, right int, down int) int {
	trees := 0
	x_pos := 0

	for i, line := range lines {
		if (i % down != 0) { continue }

		if line[x_pos] == TREE {
			trees++
		}

		x_pos = (x_pos + right) % len(line)
	}

	return trees;
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	lines := make([]string, 0, 500)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	p1_ans := CountTrees(lines, 3, 1)
	fmt.Printf("Part 1: %d\n", p1_ans)

	p2_ans := p1_ans

	var slopes = []struct {
		right int
		down int
	}{ {1,1}, {5,1}, {7,1}, {1,2} }

	for _, s := range slopes {
		p2_ans *= CountTrees(lines, s.right, s.down)
	}

	fmt.Printf("Part 2: %d\n", p2_ans)
}
