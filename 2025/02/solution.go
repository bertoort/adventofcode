package solution

import (
	"io"
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
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), ",")
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
	sum := 0
	for _, idRange := range input {
		idRange := NewIDRange(idRange)
		sum += idRange.GetSumOfInvalidIDs()
	}
	return sum
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	sum := 0
	for _, idRange := range input {
		idRange := NewIDRange(idRange)
		sum += idRange.GetSumOfSuperInvalidIDs()
	}
	return sum
}
