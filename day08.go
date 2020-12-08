package main

import (
	"fmt"
	"os"
	"bufio"
	"errors"
	"strconv"
)

const DEFAULT_FILENAME = "./input/08"

type Operation int
const (
	NOP Operation = iota
	ACC
	JMP
)

type VM struct {
	PC int
	ACC int
	Instructions []Instruction
}

type Instruction struct {
	OP Operation
	Val int
}

func ParseInstructions(lines []string) ([]Instruction, error) {
	ins := make([]Instruction, len(lines))

	for i, line := range lines {
		op := line[:3]
		num := line[4:]

		switch op {
			case "nop":
				ins[i].OP = NOP
			case "acc":
				ins[i].OP = ACC
			case "jmp":
				ins[i].OP = JMP
			default:
				return nil, errors.New("Unrecognised operation " + op)
		}

		var err error
		ins[i].Val, err = strconv.Atoi(num)
		if err != nil {
			return nil, errors.New("Malformed input")
		}
	}

	return ins, nil
}

func (vm *VM) RunStep() {
	ins := vm.Instructions[vm.PC]

	switch ins.OP {
		case NOP:
			vm.PC++
		case ACC:
			vm.ACC += ins.Val
			vm.PC++
		case JMP:
			vm.PC += ins.Val
	}
}

func (vm *VM) CheckLoop() bool {
	instructions_run := make([]bool, len(vm.Instructions))

	for {
		if vm.PC >= len(vm.Instructions) {
			return false
		}

		if instructions_run[vm.PC] {
			return true;
		}
		instructions_run[vm.PC] = true
		vm.RunStep()
	}
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
		lines = append(lines, scanner.Text() )
	}

	err = scanner.Err()
	if err != nil { panic(err) }

	ins, err := ParseInstructions( lines )
	if err != nil { panic(err) }

	var vm VM
	vm.Instructions = ins;

	vm.CheckLoop()
	fmt.Printf("Part 1: %d\n", vm.ACC);

	// Bruteforce
	found := false
	for i, ins := range vm.Instructions {
		if ins.OP != NOP && ins.OP != JMP {
			continue
		}
		vm.PC = 0
		vm.ACC = 0

		// Switch, check, then change back
		loop := false
		if ins.OP == NOP {
			vm.Instructions[i].OP = JMP
			loop = vm.CheckLoop()
			vm.Instructions[i].OP = NOP
		} else {
			vm.Instructions[i].OP = NOP
			loop = vm.CheckLoop()
			vm.Instructions[i].OP = JMP
		}

		if !loop {
			found = true
			break
		}
	}

	if !found {
		panic("Could not find answer for p2")
	}

	fmt.Printf("Part 2: %d\n", vm.ACC);
}
