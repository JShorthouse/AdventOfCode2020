package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

const DEFAULT_FILENAME = "./input/16"

type Bound struct {
	Upper int
	Lower int
}

type FieldType struct {
	Name string
	Bound1 Bound
	Bound2 Bound
}

type Ticket struct {
	Fields []int
}

func FindValidTickets(tickets []Ticket, fields_types []FieldType) (int, []Ticket) {
	error_rate := 0
	valid_tickets := make([]Ticket, 0, len(tickets))

	for _, ticket := range tickets {
		ticket_valid := true
		for _, val := range ticket.Fields {
			value_valid := false
			for _, field_t := range fields_types {
				if ( val >= field_t.Bound1.Lower && val <= field_t.Bound1.Upper) ||
				   ( val >= field_t.Bound2.Lower && val <= field_t.Bound2.Upper) {
					value_valid = true
					break;
				}
			}

			if !value_valid {
				error_rate += val
				ticket_valid = false
			}
		}
		if ticket_valid {
			valid_tickets = append(valid_tickets, ticket)
		}
	}
	return error_rate, valid_tickets
}

// Return grid of [field_idx][field_t] = valid
func ValidFieldGrid(v_tickets []Ticket, field_types []FieldType) ([][]bool) {

	valid_grid := make([][]bool, len(field_types))
	for i, _ := range valid_grid {
		valid_grid[i] = make([]bool, len(field_types))
	}

	for type_i, field_t := range field_types {
		valid_field_count := make([]int, len(v_tickets[0].Fields))

		// Count field positions that are valid for this field type
		for _, ticket := range v_tickets {
			for i, val := range ticket.Fields {
				if ( val >= field_t.Bound1.Lower && val <= field_t.Bound1.Upper) ||
				   ( val >= field_t.Bound2.Lower && val <= field_t.Bound2.Upper) {
					   valid_field_count[i] += 1
				}
			}
		}

		// Find all valid positions
		for field_i, num := range valid_field_count {
			if num == len(v_tickets) {
				valid_grid[field_i][type_i] = true
			}
		}
	}

	return valid_grid
}

func DetermineFields(grid [][]bool, field_types []FieldType) []FieldType {
	ordered_field_types := make([]FieldType, len(field_types))

	resolved_count := 0

	for resolved_count < len(field_types) {
		for field_idx, field := range grid {

			// Find if field type is only possible for this field
			// out of remaining fields
			for t_idx, possible := range field {
				if possible == false { continue }

				others_possible := false
				for other_idx, other_field := range grid {
					if other_idx == field_idx { continue }
					if other_field[t_idx] == true {
						others_possible = true
						break
					}
				}

				if !others_possible { 
					ordered_field_types[field_idx] = field_types[t_idx]

					// Clear out field
					for i := range field {
						field[i] = false
					}

					resolved_count++
					break
				}
			}
		}
	}

	return ordered_field_types
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	field_types := make([]FieldType, 0, 50)

	scanner := bufio.NewScanner(file)

	// Process field constraints
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { break }

		var field FieldType

		i := strings.Index(line, ":")
		field.Name = line[:i]

		found, err := fmt.Sscanf(line[i+2:], "%d-%d or %d-%d", &field.Bound1.Lower,
		                &field.Bound1.Upper, &field.Bound2.Lower, &field.Bound2.Upper)
		if err != nil { panic(err) }
		if found != 4 { panic("Malformed input") }

		field_types = append(field_types, field)
	}

	// Skip "your ticket:" header
	scanner.Scan()

	scanner.Scan()
	my_ticket_line := scanner.Text()

	var my_ticket Ticket
	for _, str := range strings.Split(my_ticket_line, ",") {
		val, err := strconv.Atoi(str)
		if err != nil { panic(err) }

		my_ticket.Fields = append(my_ticket.Fields, val)
	}

	// Skip blank lines and "nearby tickets:" header
	scanner.Scan()
	scanner.Scan()

	other_tickets := make([]Ticket, 0, 500)

	// Process other tickets
	for scanner.Scan() {
		line := scanner.Text()


		var ticket Ticket
		for _, str := range strings.Split(line, ",") {
			val, err := strconv.Atoi(str)
			if err != nil { panic(err) }

			ticket.Fields = append(ticket.Fields, val)
		}

		other_tickets = append(other_tickets, ticket)
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	error_rate, valid_tickets := FindValidTickets( other_tickets, field_types )
	fmt.Printf("Part 1: %d\n", error_rate)

	grid := ValidFieldGrid(valid_tickets, field_types)
	field_order := DetermineFields(grid, field_types)

	p2_ans := 1
	for i, field := range field_order {
		if strings.HasPrefix(field.Name, "departure") {
			p2_ans *= my_ticket.Fields[i]
		}
	}

	fmt.Printf("Part 2: %d\n", p2_ans)
}
