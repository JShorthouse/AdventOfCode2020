package main

import (
	"os"
	"bufio"
	"fmt"
)

const DEFAULT_FILENAME = "./input/11"

type Offset struct {
	x int
	y int
}

var offsets = []Offset{
	{ 1,  1}, { 1,  0}, {0,  1},
	{-1, -1}, {-1,  0}, {0, -1},
	{-1,  1}, { 1, -1},
}
func ProcessNormalStep(grid [][]byte) ([][]byte, int) {
	new_grid := make([][]byte, len(grid))
	for i := range new_grid {
		new_grid[i] = make([]byte, len(grid[0]))
	}

	num_updated := 0

	for y := 0; y<len(grid); y++ {
		for x:=0; x<len(grid[0]); x++ {
			if grid[y][x] == '.' {
				new_grid[y][x] = '.'
				continue
			}

			num_occupied := 0
			for _, offset := range offsets {
				test_x := x + offset.x
				test_y := y + offset.y
				if test_x >= 0 && test_y >= 0 && test_x < len(grid[0]) && test_y < len(grid) {
					if grid[test_y][test_x] == '#' {
						num_occupied++
					}
				}
			}

			if grid[y][x] == 'L' && num_occupied == 0 {
				new_grid[y][x] = '#'
				num_updated++
			} else if grid[y][x] == '#' && num_occupied >= 4 {
				new_grid[y][x] = 'L'
				num_updated++
			} else {
				new_grid[y][x] = grid[y][x]
			}
		}
	}

	return new_grid, num_updated
}

func ProcessVisionStep(grid [][]byte) ([][]byte, int) {
	new_grid := make([][]byte, len(grid))
	for i := range new_grid {
		new_grid[i] = make([]byte, len(grid[0]))
	}

	num_updated := 0

	for y := 0; y<len(grid); y++ {
		for x:=0; x<len(grid[0]); x++ {
			if grid[y][x] == '.' {
				new_grid[y][x] = '.'
				continue
			}

			num_occupied := 0
			for _, offset := range offsets {
				test_x := x + offset.x
				test_y := y + offset.y

				for test_x >= 0 && test_y >= 0 && test_x < len(grid[0]) && test_y < len(grid) {
					if grid[test_y][test_x] == '#' {
						num_occupied++
						break
					} else if grid[test_y][test_x] == 'L' {
						break
					}

					test_x += offset.x
					test_y += offset.y
				}
			}

			if grid[y][x] == 'L' && num_occupied == 0 {
				new_grid[y][x] = '#'
				num_updated++
			} else if grid[y][x] == '#' && num_occupied >= 5 {
				new_grid[y][x] = 'L'
				num_updated++
			} else {
				new_grid[y][x] = grid[y][x]
			}
		}
	}

	return new_grid, num_updated
}


func CountChars(grid [][]byte, char byte) int {
	count := 0
	for y := 0; y<len(grid); y++ {
		for x:=0; x<len(grid[0]); x++ {
			if grid[y][x] == char {
				count++
			}
		}
	}
	return count
}


func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	grid := make([][]byte, 0, 100)
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, make([]byte, 0, 50))

		line := scanner.Text()

		for _, c := range line {
			grid[y] = append(grid[y], byte(c))
		}
		y++
	}

	err = scanner.Err()
	if err != nil { panic(err) }


	p1_grid, _ := ProcessNormalStep(grid)
	num_changed := 1

	for num_changed != 0 {
		p1_grid, num_changed = ProcessNormalStep(p1_grid)
	}

	fmt.Printf("Part 1: %d\n", CountChars(p1_grid, '#'))


	p2_grid, _ := ProcessVisionStep(grid)
	num_changed = 1

	for num_changed != 0 {
		p2_grid, num_changed = ProcessVisionStep(p2_grid)
	}

	fmt.Printf("Part 2: %d\n", CountChars(p2_grid, '#'))
}


