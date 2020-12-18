package main

import (
	"os"
	"fmt"
	"bufio"
)

const DEFAULT_FILENAME = "./input/17"

type Point4 struct {
	x int
	y int
	z int
	w int
}

var offsets []Point4

func GenOffsets() {
	offsets = make([]Point4, 0, 26);

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			for z := 0; z < 3; z++ {
				for w := 0; w < 3; w++ {
					rel_x := 1 - x
					rel_y := 1 - y
					rel_z := 1 - z
					rel_w := 1 - w

					if rel_x == 0 && rel_y == 0 && rel_z == 0 && rel_w == 0 { continue }
					
					offsets = append(offsets, Point4{rel_x, rel_y, rel_z, rel_w})
				}
			}
		}
	}
	fmt.Printf("Offsets: %d\n", len(offsets))
}

type InfGrid map[int]map[int]map[int]map[int]byte

func (grid InfGrid) Get(x, y, z, w int) (byte, bool) {
	y_data, ok := grid[x]
	if !ok { return 0, false }

	z_data, ok := y_data[y]
	if !ok { return 0, false }

	w_data, ok := z_data[z]
	if !ok { return 0, false }

	value, ok := w_data[w]
	if !ok { return 0, false }

	return value, true
}

func (grid InfGrid) Set(x, y, z, w int, val byte) {
	y_data, ok := grid[x]
	if !ok { 
		y_data = make(map[int]map[int]map[int]byte) 
		grid[x] = y_data
	}

	z_data, ok := y_data[y]
	if !ok { 
		z_data = make(map[int]map[int]byte)
		y_data[y] = z_data
	}

	w_data, ok := z_data[z]
	if !ok { 
		w_data = make(map[int]byte)
		z_data[z] = w_data
	}

	w_data[w] = val
}

func (grid InfGrid) Iterate(fn func(x, y, z, w int, val byte)) {
	for x, y_data := range grid {
		for y, z_data := range y_data {
			for z, w_data := range z_data {
				for w, val := range w_data {
					fn(x,y,z,w, val)
				}
			}
		}
	}
}

func (grid InfGrid) Count(target byte) int {
	total := 0

	grid.Iterate( func(_,_,_,_ int, val byte) {
		if val == target {
			total++
		}
	})

	return total
}


func RunCycles(grid InfGrid, num_cycles int, four_d bool) InfGrid {
	var cur_grid = grid
	for i := 0; i < num_cycles; i++ {
		var count_grid = make(InfGrid)
		var new_grid = make(InfGrid)

		cur_grid.Iterate( func(x,y,z,w int, val byte) {
			for _, offset := range offsets {
				if offset.w != 0 && !four_d { continue }

				if val == '#' {
					target_x := x + offset.x
					target_y := y + offset.y
					target_z := z + offset.z
					target_w := w + offset.w

					cur_val, _ := count_grid.Get(target_x, target_y, target_z, target_w)
					count_grid.Set(target_x, target_y, target_z, target_w, cur_val + 1)
				}
			}
		})

		// Update values based on counts
		count_grid.Iterate( func(x,y,z,w int, count byte) {
			val, _ := cur_grid.Get(x, y, z, w)
			if val != '#' && count == 3 {
				new_grid.Set(x,y,z,w, '#')
			} else if val == '#' && (count == 2 || count == 3) {
				new_grid.Set(x,y,z,w, '#')
			}
		})

		cur_grid = new_grid
	}

	return cur_grid
}


func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	var grid InfGrid
	grid = make(InfGrid)

	x := 0
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, rn := range line {
			grid.Set(x, y, 0, 0, byte(rn))
			x++
		}
		x = 0
		y++
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	GenOffsets()

	var dup_grid = make(InfGrid)
	grid.Iterate(func(x,y,z,w int, val byte) {
		dup_grid.Set(x,y,z,w, val)
	})

	p1_grid := RunCycles( grid, 6, false )
	fmt.Printf("Part 1: %d\n", p1_grid.Count('#'))

	p2_grid := RunCycles( dup_grid, 6, true )
	fmt.Printf("Part 1: %d\n", p2_grid.Count('#'))
}
