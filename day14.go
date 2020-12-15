
package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"math"
)

const DEFAULT_FILENAME = "./input/14"

const (
	MASK int = iota
	WRITE
)

type Instruction struct {
	Type int
	First uint
	Second uint
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func RunProgram(instructions []Instruction) int {
	mem_map := make(map[uint]uint)

	var zero_mask uint
	var one_mask uint

	for _, ins := range instructions {
		if ins.Type == MASK {
			zero_mask = ins.First
			one_mask = ins.Second
		} else if ins.Type == WRITE {
			val := ins.Second
			val = val | one_mask
			val = val & ^zero_mask

			mem_map[ins.First] = val
		}
	}

	total := 0
	for _, val := range mem_map {
		total += int(val)
	}

	return total
}

const BitMask36 = 0xFFFFFFFFF

func RunProgramFloating(instructions []Instruction) uint {
	mem_map := make(map[uint]uint)

	var zero_mask uint
	var one_mask uint

	var float_mask uint
	var num_floating int
	var float_offsets []int
	var final_offset int

	for _, ins := range instructions {
		if ins.Type == MASK {
			zero_mask = ins.First
			one_mask = ins.Second
			float_mask = (^( zero_mask | one_mask )) & BitMask36

			num_floating = 0
			float_offsets = make([]int, 0, 36)
			final_offset = 0

			in_padding := true
			cur_offset := 0


			// Iterate over binary digits as chars and calculate offsets
			for _, c := range strconv.FormatInt(int64(float_mask), 2) {
				if c == '0' && in_padding { continue } else { in_padding = false }

				cur_offset++
				if c == '1' {
					num_floating++
					float_offsets = append(float_offsets, cur_offset)
					cur_offset = 0
				}

			}
			
			if cur_offset > 0 {
				final_offset = cur_offset
			}


		} else if ins.Type == WRITE {
			base_address := ins.First
			base_address = base_address | one_mask
			base_address = base_address & ^float_mask


			// Calculate all combinations of N floating values by generating
			// numbers up to 2^N and shifting each digit by the float offsets
			for i := 0; i < Pow(2, num_floating); i++ {
				var float_values uint = 0
				for j, offset := range float_offsets {
					float_values = float_values << offset
					float_values += (uint(i) >> j) & 1
				}

				float_values = float_values << final_offset

				address := base_address | float_values
				mem_map[address] = ins.Second
			}
		}
	}

	total := uint(0)
	for _, val := range mem_map {
		total += val
	}

	return total
}


func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	instructions := make([]Instruction, 0, 500)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var ins Instruction

		if strings.HasPrefix(line, "mask") {
			mask_str := line[len("mask = "):]
			var zero_mask uint = 0
			var one_mask uint = 0
			for _, ch := range mask_str {
				zero_mask = zero_mask << 1
				one_mask = one_mask << 1

				if ch == '0' {
					zero_mask += 1
				} else if ch == '1' {
					one_mask += 1
				}
			}

			ins.Type = MASK
			ins.First = zero_mask
			ins.Second = one_mask

			instructions = append(instructions, ins)

		} else if strings.HasPrefix(line, "mem") {
			var index uint
			var value uint
			found, err := fmt.Sscanf(line, "mem[%d] = %d", &index, &value)
			if err != nil { panic(err) }
			if found != 2 { panic("Malformed input") }

			ins.Type = WRITE
			ins.First = index
			ins.Second = value

			instructions = append(instructions, ins)
		} else {
			panic("Malformed input")
		}
	}

	err = scanner.Err()
	if err != nil { panic(err) }


	p1_ans := RunProgram(instructions)
	fmt.Printf("Part 1: %d\n", p1_ans)

	p2_ans := RunProgramFloating(instructions)
	fmt.Printf("Part 2: %d\n", p2_ans)
}
