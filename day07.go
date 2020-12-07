package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"errors"
)

type BagChild struct {
	Bag *Bag
	Count int
}

type Bag struct {
	Color string
	Children []BagChild
}

var BagMap = make(map[string]*Bag)


const DEFAULT_FILENAME = "./input/07"
const TARGET_COLOR     = "shiny gold"

func CreateBags(lines []string) error {
	for _, line := range lines {
		i := strings.Index(line, "bags")
		if i == -1 { return errors.New("Malformed input") }

		parent_col := line[:i-1]
		child_list := line[i + len("bags contain "):]

		// Process parent
		parent_bag, exists := BagMap[parent_col];
		if !exists {
			parent_bag = new(Bag)
			parent_bag.Color = parent_col
			parent_bag.Children = make([]BagChild, 0, 5)
			BagMap[parent_col] = parent_bag
		}


		// Process children
		if child_list == "no other bags." {
			continue
		}
		for {
			child_num := int(child_list[0]) - int('0') // Convert to number
			if child_num < 0 || child_num > 10 { return errors.New("Malformed input") }

			i = strings.Index(child_list, "bag")
			if i == -1 { return errors.New("Malformed input") }

			child_col := child_list[2:i-1]

			child_bag, exists := BagMap[child_col]

			if !exists {
				child_bag = new(Bag)
				child_bag.Color = child_col
				child_bag.Children = make([]BagChild, 0, 5)
				BagMap[child_col] = child_bag
			}

			var child_node BagChild
			child_node.Bag = child_bag
			child_node.Count = child_num

			parent_bag.Children = append(parent_bag.Children, child_node)

			// Move to next child
			i = strings.Index(child_list, ",")
			if i == -1 {
				// End of children
				break
			} else {
				// Advance to next child in list
				child_list = child_list[i+2:]
			}
		}
	}

	return nil
}

func BagContains( bag *Bag, target string ) bool {
	if len(bag.Children) == 0 { return false; }

	for _, child := range bag.Children {
		if child.Bag.Color == target { return true; }
	}

	for _, child := range bag.Children {
		if BagContains( child.Bag, target ) { return true; }
	}
	
	return false;
}

func CountBags( bag *Bag ) int {
	if len(bag.Children) == 0 { return 1 };

	num_bags := 1
	for _, child := range bag.Children {
		num_bags += child.Count * CountBags(child.Bag)
	}

	return num_bags
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

	err = CreateBags(lines)
	if err != nil { panic(err) }

	num_parents := 0;

	// Bruteforce
	for _, bag := range BagMap {
		if bag.Color == TARGET_COLOR { continue }

		if BagContains( bag, TARGET_COLOR ) { 
			num_parents++
		}
	}

	fmt.Printf("Part 1: %d\n", num_parents)

	target_bag := BagMap[ TARGET_COLOR ];
	num_children := CountBags( target_bag ) - 1

	fmt.Printf("Part 2: %d\n", num_children)
}
