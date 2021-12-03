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

func parseBinary(input string) int {
	i, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		log.Fatal("failed to parse input line", err)
	}
	return int(i)
}

func findCounts(input []string) []int {
	counts := make([]int, len(input[0]))
	for _, line := range input {
		for i, letter := range line {
			if string(letter) == "1" {
				counts[i] += 1
			}
		}
	}
	return counts
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	counts := findCounts(input)
	length := len(input)
	gamma := ""
	epsilon := ""
	for _, count := range counts {
		if count > length/2 {
			gamma += "1"
			epsilon += "0"
			continue
		}
		gamma += "0"
		epsilon += "1"
	}
	return parseBinary(gamma) * parseBinary(epsilon)
}

func filterInput(input []string, index int, mostCommon bool) []string {
	if len(input) == 1 {
		return input
	}
	newInput := []string{}
	counts := findCounts(input)
	desired := "0"
	oneIsMostCommon := counts[index] >= len(input)-counts[index] && mostCommon
	oneIsLeastCommon := counts[index] < len(input)-counts[index] && !mostCommon
	if oneIsMostCommon || oneIsLeastCommon {
		desired = "1"
	}
	for _, line := range input {
		if string(line[index]) == desired {
			newInput = append(newInput, line)
		}
	}
	return filterInput(newInput, index+1, mostCommon)
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	oxygen := filterInput(input, 0, true)
	co2 := filterInput(input, 0, false)
	return parseBinary(oxygen[0]) * parseBinary(co2[0])
}
