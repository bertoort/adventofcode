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
	total := 0
	for i, line := range input {
		if i == 0 {
			continue
		}
		number := parseInt(line)
		prevNumber := parseInt(input[i-1])
		if number > prevNumber {
			total++
		}
	}
	return total
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	total := 0
	for i, line := range input {
		if i == len(input)-3 {
			break
		}
		n1 := parseInt(line)
		n2 := parseInt(input[i+1])
		n3 := parseInt(input[i+2])
		n4 := parseInt(input[i+3])
		prevNumber := n1 + n2 + n3
		number := n2 + n3 + n4
		if number > prevNumber {
			total++
		}
	}
	return total
}
