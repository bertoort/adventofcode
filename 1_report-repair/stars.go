package stars

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const goal = 2020

func setup() []string {
	file, err := os.Open("./input.txt")
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

func isSolution(n1, n2 int) {
	if n1+n2 == goal {
		log.Fatalf("solution found %d, from %d and %d", n1*n2, n1, n2)
	}
}

// Solution i/o
func Solution() {
	input := setup()
	largeNumbers := []int{}
	smallNumbers := []int{}
	start := parseInt(input[0])
	for _, line := range input[1:] {
		number := parseInt(line)
		isSolution(start, number)
		if number >= (goal / 2) {
			largeNumbers = append(largeNumbers, number)
		} else {
			smallNumbers = append(smallNumbers, number)
		}
	}
	for _, large := range largeNumbers {
		for _, small := range smallNumbers {
			isSolution(large, small)
		}
	}
	log.Fatal("no solution found")
}
