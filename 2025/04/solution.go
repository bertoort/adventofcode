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
	wall := Wall{grid: [][]string{}}
	for _, row := range input {
		wall.grid = append(wall.grid, strings.Split(row, ""))
	}
	return wall.countAccessibleRolls()
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	wall := Wall{grid: [][]string{}}
	for _, row := range input {
		wall.grid = append(wall.grid, strings.Split(row, ""))
	}
	return wall.removeAllAccessibleRolls()
}
