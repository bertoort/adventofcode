package solution

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}


func parseInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("failed to parse input line", err)
	}
	return number
}


// Solution i/o
func Solution(path string) int {
	input := setup(path)
	horizontal := 0
	depth := 0
	for _, line := range input {
		split := strings.Split(line, " ")
		direction := split[0]
		count := parseInt(split[1])
		if direction == "forward" {
			horizontal += count
		}
		if direction == "down" {
			depth += count
		}
		if direction == "up" {
			depth -= count
		}
	}
	return horizontal * depth
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	aim := 0
	depth := 0
	horizontal := 0
	for _, line := range input {
		split := strings.Split(line, " ")
		direction := split[0]
		count := parseInt(split[1])
		if direction == "forward" {
			horizontal += count
			depth += aim * count
		}
		if direction == "down" {
			aim += count
		}
		if direction == "up" {
			aim -= count
		}
	}
	return horizontal * depth
}
