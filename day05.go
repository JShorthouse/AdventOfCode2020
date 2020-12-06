package main

import (
	"fmt"
	"errors"
	"os"
	"bufio"
)

const DEFAULT_FILENAME = "./input/05"
const SEAT_ROWS = 128
const SEAT_COLS = 8


func FindSeat(line string) (int, int, error) {
	row_lower, row_upper := 0, SEAT_ROWS-1
	col_lower, col_upper := 0, SEAT_COLS-1

	for _, ins := range line {
		if row_lower == row_upper && col_lower == col_upper {
			return 0, 0, errors.New("Too many instructions in " + line)
		}

		row_middle := (row_lower + row_upper) / 2;
		col_middle := (col_lower + col_upper) / 2;

		switch ins {
			case 'F':
				row_upper = row_middle;
			case 'B':
				row_lower = row_middle + 1;
			case 'L':
				col_upper = col_middle;
			case 'R':
				col_lower = col_middle + 1;
			default:
				return 0, 0, errors.New("Invalid instruction " + string(ins))
		}
	}

	if row_lower != row_upper || col_lower != col_upper {
		return 0, 0, errors.New("Not enough instructions in " + line)
	}

	return row_lower, col_lower, nil
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	lines := make([]string, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text() )
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	var seats_occupied [SEAT_ROWS * SEAT_COLS]bool
	max_seat := 0

	for _, line := range lines {
		row, col, err := FindSeat( line )
		if err != nil { panic(err) }

		seat_id := (row * SEAT_COLS) + col;

		if seat_id > max_seat {
			max_seat = seat_id;
		}

		seats_occupied[seat_id] = true;
	}

	fmt.Printf("Part 1: %d\n", max_seat);


	// Find first unnocupied seat with occupied seats either side
	for i := 2; i<len(seats_occupied); i++ {
		if seats_occupied[i] && !seats_occupied[i-1] && seats_occupied[i-2] {
			fmt.Printf("Part 2: %d\n", i-1);
			return
		}
	}

	panic("Could not find seat for part 2")
}
