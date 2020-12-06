package main

import (
	"fmt"
	"os"
	"bufio"
)

const FILE_NAME = "./input/02"

func ExtractPassFields(line string) (int, int, rune, string, error) {
	var min, max int
	var p_char rune
	var password string
	_, err := fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &p_char, &password)

	if err != nil { return 0, 0, 0, "", err }

	return min, max, p_char, password, nil
}

func CheckPasswordCount(lines []string) (int, error) {
	num_valid := 0
	for _, line := range lines {
		min, max, p_char, password, err := ExtractPassFields(line)
		if err != nil { return 0, err }

		count := 0
		for _, c := range password {
			if c == p_char {
				count++
			}
		}

		if count >= min && count <= max {
			num_valid++
		}
	}

	return num_valid, nil
}

func CheckPasswordIndex(lines []string) (int, error) {
	num_valid := 0
	for _, line := range lines {
		first_pos, second_pos, p_char, password, err := ExtractPassFields(line)
		if err != nil { return 0, err }

		first  := rune(password[first_pos-1])  == p_char
		second := rune(password[second_pos-1]) == p_char

		if (first || second) && !(first && second) {
			num_valid++
		}
	}

	return num_valid, nil
}

func main() {
	file, err := os.Open(FILE_NAME)
	if err != nil { panic(err) }
	defer file.Close()

	lines := make([]string, 0, 500)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	p1_ans, err := CheckPasswordCount( lines )
	if err != nil { panic(err) }

	fmt.Printf("Part 1: %d\n", p1_ans )

	p2_ans, err := CheckPasswordIndex( lines )
	if err != nil { panic(err) }

	fmt.Printf("Part 2: %d\n", p2_ans )
}
