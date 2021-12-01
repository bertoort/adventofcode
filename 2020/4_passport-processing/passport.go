package passport

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Key for passport info
type Key string

var (
	byr Key = "byr"
	iyr Key = "iyr"
	eyr Key = "eyr"
	hgt Key = "hgt"
	hcl Key = "hcl"
	ecl Key = "ecl"
	pid Key = "pid"
	cid Key = "cid"
)

// Passport information
type Passport map[Key]string

// ValidNorthPoleCredentails checks for valid North Pole passport
func (p Passport) ValidNorthPoleCredentails() bool {
	requiredKeys := []Key{
		byr,
		iyr,
		eyr,
		hgt,
		hcl,
		ecl,
		pid,
	}
	for _, key := range requiredKeys {
		if value, ok := p[key]; !ok || value == "" {
			return false
		}
	}
	return true
}

// Validate checks for valid North Pole passport fields
func (p Passport) Validate() bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	birthYear, err := strconv.Atoi(p[byr])
	if err != nil || birthYear < 1920 || birthYear > 2002 {
		return false
	}
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	issueYear, err := strconv.Atoi(p[iyr])
	if err != nil || issueYear < 2010 || issueYear > 2020 {
		return false
	}
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	expiration, err := strconv.Atoi(p[eyr])
	if err != nil || expiration < 2020 || expiration > 2030 {
		return false
	}
	// hgt (Height) - a number followed by either cm or in:
	validHeight := false
	cms := strings.Split(p[hgt], "cm")
	if len(cms) == 2 {
		// If cm, the number must be at least 150 and at most 193.
		if height, err := strconv.Atoi(cms[0]); err == nil && height >= 150 && height <= 193 {
			validHeight = true
		}
	}
	ins := strings.Split(p[hgt], "in")
	if len(ins) == 2 {
		// If in, the number must be at least 59 and at most 76.
		if height, err := strconv.Atoi(ins[0]); err == nil && height >= 59 && height <= 76 {
			validHeight = true
		}
	}
	if !validHeight {
		return false
	}
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hexcolorRegexString := "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	if matched, err := regexp.Match(hexcolorRegexString, []byte(p[hcl])); err != nil || !matched {
		return false
	}
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	validEyeOptions := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	validEyes := false
	for _, color := range validEyeOptions {
		if color == p[ecl] {
			validEyes = true
			break
		}
	}
	if !validEyes {
		return false
	}
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	_, err = strconv.Atoi(p[pid])
	if err != nil || len(p[pid]) != 9 {
		return false
	}
	return true
}

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n\n")
}

func parsePassport(input string) Passport {
	passport := Passport{}
	info := []string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		spaces := strings.Split(line, " ")
		info = append(info, spaces...)
	}
	for _, pair := range info {
		values := strings.Split(pair, ":")
		switch Key(values[0]) {
		case byr:
			passport[byr] = values[1]
		case iyr:
			passport[iyr] = values[1]
		case eyr:
			passport[eyr] = values[1]
		case hgt:
			passport[hgt] = values[1]
		case hcl:
			passport[hcl] = values[1]
		case ecl:
			passport[ecl] = values[1]
		case pid:
			passport[pid] = values[1]
		case cid:
			passport[cid] = values[1]
		}
	}
	return passport
}

// PartOneSolution i/o
func PartOneSolution(path string) int {
	file := setup(path)
	total := 0
	for _, passport := range file {
		p := parsePassport(passport)
		if valid := p.ValidNorthPoleCredentails(); valid {
			total++
		}
	}
	return total
}

// PartTwoSolution i/o
func PartTwoSolution(path string) int {
	file := setup(path)
	total := 0
	for _, passport := range file {
		p := parsePassport(passport)
		if valid := p.Validate(); valid {
			total++
		}
	}
	return total
}
