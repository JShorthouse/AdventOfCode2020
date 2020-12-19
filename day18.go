package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
)

const DEFAULT_FILENAME = "input/18"

type Token struct {
	Type int
	Value int64
}

const (
	NUMBER int = iota
	OP_ADD
	OP_SUB
	OP_MUL
	BRAC_OPEN
	BRAC_CLOSE
)

func Tokenize(eq_str string) ([]Token, error) {
	i := 0
	tokens := make([]Token, 0, 20)

	for i < len(eq_str) {
		var tk Token 

		switch {
			case eq_str[i] == '+': {
				tk.Type = OP_ADD
				i += 2
			}
			case eq_str[i] == '-': {
				tk.Type = OP_SUB
				i += 2
			}
			case eq_str[i] == '*': {
				tk.Type = OP_MUL
				i += 2
			}
			case eq_str[i] == '(': {
				tk.Type = BRAC_OPEN 
				i++
				if i < len(eq_str) && eq_str[i] == ' ' {
					i++
				}
			}
			case eq_str[i] == ')': {
				tk.Type = BRAC_CLOSE 
				i++
				if i < len(eq_str) && eq_str[i] == ' ' {
						i++
				}
			}
			case eq_str[i] >= '0' && eq_str[i] <= '9': {
				// Find end of number
				end_idx := i+1;
				for end_idx < len(eq_str) &&
					eq_str[end_idx] >= '0' &&
					eq_str[end_idx] <= '9' {
						end_idx++
				}

				tk.Type = NUMBER
				
				val, err := strconv.Atoi(eq_str[i:end_idx])
				if err != nil { return nil, errors.New("Malformed input") }

				tk.Value = int64(val)

				// Advance past number
				i = end_idx
				if i < len(eq_str) && eq_str[i] == ' ' {
					i++
				}
			}

			default: {
				return nil, errors.New("Unrecognised operator '" + string(eq_str[i]) + "'")
			}
		}

		tokens = append(tokens, tk)
	}

	return tokens, nil
}

func calc(tokens []Token) (int64, error) {
		i := 0
		cur_sum := int64(0)
		cur_op := OP_ADD

		for {
		var cur_val int64

		if(tokens[i].Type == NUMBER) {
			cur_val = tokens[i].Value
		} else if tokens[i].Type == BRAC_OPEN {
			i += 1

			depth := 1
			end_idx := i

			// Find inside of bracket and process
			for {
				if tokens[end_idx].Type == BRAC_OPEN {
					depth++
				} else if tokens[end_idx].Type == BRAC_CLOSE {
					depth--

					if depth == 0 {
						break
					}
				}
				end_idx++
			}

			var err error
			cur_val, err = calc(tokens[i:end_idx])
			if err != nil { return 0, err }

			i = end_idx
		} else {
			return 0, errors.New("Malformed token sequence")
		}

		switch cur_op {
			case OP_ADD: { cur_sum += cur_val }
			case OP_SUB: { cur_sum -= cur_val }
			case OP_MUL: { cur_sum *= cur_val }
		}

		if i + 1 >= len(tokens) {
			break
		}

		cur_op = tokens[i+1].Type;
		i += 2
	}

	return cur_sum, nil
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	tokens := make([][]Token, 0, 200)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tk, err := Tokenize(scanner.Text())
		if err != nil { panic(err) }
		tokens = append(tokens, tk)
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	p1_ans := int64(0)

	for _, tk := range tokens {
		val, err := calc(tk)
		if err != nil { panic(err) }
		p1_ans += val
	}

	fmt.Printf("Part 1: %d\n", p1_ans);
}
