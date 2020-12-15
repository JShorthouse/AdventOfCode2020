package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"errors"
)

const DEFAULT_FILENAME = "./input/12"

type Point struct {
	X int
	Y int
}

func PosMod(val int, mod int) int {
	res := val % mod;
	if res >= 0 {
		return res;
	} else {
		return res + mod;
	}
}

func Abs(i int) int {
	if i > 0 { return i } else { return -i }
}

var trig_table = [4][2]int{ {0, 1}, {1, 0}, {0, -1}, {-1, 0} }

func rotatePoint(point *Point, origin Point, angle int) {
	p_angle := PosMod(angle, 360)

	i := 0
	for p_angle >= 90 {
		i++
		p_angle -= 90
	}

	sin := trig_table[i][0]
	cos := trig_table[i][1]

	// Translate to centre
	point.X -= origin.X
	point.Y -= origin.Y

	new_x := point.X * cos - point.Y * sin
	new_y := point.X * sin + point.Y * cos

	// Translate back
	point.X = new_x + origin.X
	point.Y = new_y + origin.Y
}


type Instruction struct {
	Dir byte
	Value int
}

const (
	NORTH int = iota
	EAST
	SOUTH
	WEST
)

var Offsets = [4]Point{ {0, -1}, {1, 0}, {0, 1}, {-1, 0} }

func ProcessMoveShip(instructions []Instruction) (int, error) {
	var ship = Point{ 0, 0 }
	var ship_dir = EAST;

	for _, ins := range instructions {
		switch(ins.Dir) {
		case 'N':
			ship.Y -= ins.Value;
		case 'S':
			ship.Y += ins.Value;
		case 'E':
			ship.X += ins.Value;
		case 'W':
			ship.X -= ins.Value;
		case 'L':
			angle := ins.Value / 90;
			ship_dir = PosMod( ship_dir - angle, 4);
		case 'R':
			angle := ins.Value / 90;
			ship_dir = PosMod( ship_dir + angle, 4);
		case 'F':
			ship.X += Offsets[ship_dir].X * ins.Value;
			ship.Y += Offsets[ship_dir].Y * ins.Value;
		default:
			return 0, errors.New("Unrecognised instruction " + string(ins.Dir))
		}
	}

	return Abs(ship.X) + Abs(ship.Y), nil
}

func ProcessMoveWaypoint(instructions []Instruction) (int, error) {
	var ship = Point{ 0, 0 }
	var waypoint = Point{ 10, -1 }

	for _, ins := range instructions {
		switch(ins.Dir) {
		case 'N':
			waypoint.Y -= ins.Value;
		case 'S':
			waypoint.Y += ins.Value;
		case 'E':
			waypoint.X += ins.Value;
		case 'W':
			waypoint.X -= ins.Value;
		case 'L':
			rotatePoint(&waypoint, ship, -ins.Value)
		case 'R':
			rotatePoint(&waypoint, ship, ins.Value)
		case 'F':
			offset := Point{ waypoint.X - ship.X, waypoint.Y - ship.Y }
			ship.X += offset.X * ins.Value
			ship.Y += offset.Y * ins.Value
			waypoint.X = ship.X + offset.X
			waypoint.Y = ship.Y + offset.Y
		default:
			return 0, errors.New("Unrecognised instruction " + string(ins.Dir))
		}
	}

	return Abs(ship.X) + Abs(ship.Y), nil
}


func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	instructions := make([]Instruction, 0, 500);

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text();

		var ins Instruction
		ins.Dir = line[0];

		val, err := strconv.Atoi(line[1:])
		if err != nil { panic(err) }

		ins.Value = val

		instructions = append(instructions, ins)
	}

	p1_ans, err := ProcessMoveShip(instructions)
	if err != nil { panic(err) }

	fmt.Printf("Part 1: %d\n", p1_ans);


	p2_ans, err := ProcessMoveWaypoint(instructions)
	if err != nil { panic(err) }

	fmt.Printf("Part 2: %d\n", p2_ans);
}
