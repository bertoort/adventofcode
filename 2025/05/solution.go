package solution

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func setup(path string) ([]string, []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	var separatorIndex int
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			separatorIndex = i
			break
		}
	}
	return lines[:separatorIndex], lines[separatorIndex+1:]
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
	ranges, numbers := setup(path)
	// validIDs := getValidIDs(ranges)
	// count := getFreshIDs(numbers, validIDs)
	count := 0
	for _, number := range numbers {
		if isInRange(number, ranges) {
			count++
		}
	}
	return count
}

// Solution2 i/o
func Solution2(path string) int {
	ranges, _ := setup(path)
	validRanges := getValidRanges(ranges)
	count := countRangeValues(validRanges)
	return count
}
