package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"regexp"
)

type Passport struct {
	BirthYear string
	IssueYear string
	ExpYear   string
	Height    string
	HairCol   string
	EyeCol    string
	ID        string
	CountryID string
}

const DEFAULT_FILENAME = "./input/04"

func ParsePassports(lines []string) []Passport {
	passports := make([]Passport, len(lines))

	for _, line := range lines {
		var pass Passport

		for _, field := range strings.Fields(line) {
			i := strings.Index(field, ":")
			if i == -1 {
				panic("Error - Malformed Input")
			}

			key := field[:i]
			value := field[i+1:]

			switch key {
				case "byr":
					pass.BirthYear = value;
				case "iyr":
					pass.IssueYear = value;
				case "eyr":
					pass.ExpYear = value;
				case "hgt":
					pass.Height = value;
				case "hcl":
					pass.HairCol = value;
				case "ecl":
					pass.EyeCol = value;
				case "pid":
					pass.ID = value;
				case "cid":
					pass.CountryID = value;
				default: 
					panic("Error - Unexpected key " + key)
			}
		}
		passports = append(passports, pass)

	}

	return passports;
}

func CountValidPassports( passports []Passport) (int, int) {
	num_present := 0
	num_valid := 0

	for _, pass := range passports {
		if pass.BirthYear != "" &&
		   pass.IssueYear != "" &&
		   pass.ExpYear   != "" &&
		   pass.Height    != "" &&
		   pass.HairCol   != "" &&
		   pass.EyeCol    != "" &&
		   pass.ID        != "" {
			   num_present++
		} else {
			continue
		}

		b_year, err := strconv.Atoi(pass.BirthYear)
		if err != nil || b_year < 1920 || b_year > 2002 {
			continue
		}

		i_year, err := strconv.Atoi(pass.IssueYear)
		if err != nil || i_year < 2010 || i_year > 2020 {
			continue
		}

		e_year, err := strconv.Atoi(pass.ExpYear)
		if err != nil || e_year < 2020 || e_year > 2030 {
			continue
		}

		height_num, err := strconv.Atoi(pass.Height[:len(pass.Height)-2])
		if err != nil { continue }

		height_unit := pass.Height[len(pass.Height)-2:]

		switch height_unit {
			case "cm":
				if height_num < 150 || height_num > 193 { continue }
			case "in":
				if height_num < 59 || height_num > 76 { continue }
			default:
				continue
		}


		hair_valid, err := regexp.Match(`^#[a-f0-9]{6}$`, []byte(pass.HairCol))
		if err != nil { panic(err) }

		if !hair_valid {
			continue
		}


		eye_valid := false
		for _, col := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if pass.EyeCol == col {
				eye_valid = true
				break
			}
		}
		if !eye_valid {
			continue
		}

		if len(pass.ID) != 9 {
			continue
		}

		id_valid, err := regexp.Match(`^0*[0-9]+$`, []byte(pass.ID))
		if err != nil { panic(err) }

		if !id_valid {
			continue
		}

		num_valid++
	}

	return num_present, num_valid;
}

func main() {
	var file_n string
	if len(os.Args) > 1 { file_n = os.Args[1] } else { file_n = DEFAULT_FILENAME }

	file, err := os.Open(file_n)
	if err != nil { panic(err) }
	defer file.Close()

	lines := make([]string, 0, 500)
	var builder strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			lines = append(lines, builder.String() )
			builder.Reset()
		} else {
			builder.WriteString(" ")
			builder.WriteString( line )
		}
	}

	if( builder.Len() > 0) {
		lines = append(lines, builder.String() )
	}

	passports := ParsePassports( lines )

	p1_ans, p2_ans := CountValidPassports( passports )

	fmt.Printf("Part 1: %d\n", p1_ans);
	fmt.Printf("Part 2: %d\n", p2_ans);
}
