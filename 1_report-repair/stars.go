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

func isSolutionTwo(n1, n2 int) {
	if n1+n2 == goal {
		log.Fatalf("solution two found %d, from %d and %d", n1*n2, n1, n2)
	}
}

func isSolutionThree(n1, n2, n3 int) {
	if n1+n2+n3 == goal {
		log.Fatalf("solution three found %d, from %d and %d and %d", n1*n2*n3, n1, n2, n3)
	}
}

// SolutionTwo i/o
func SolutionTwo() {
	input := setup()
	largeNumbers := []int{}
	smallNumbers := []int{}
	for _, line := range input[1:] {
		number := parseInt(line)
		isSolutionTwo(parseInt(input[0]), number)
		if number >= (goal / 2) {
			largeNumbers = append(largeNumbers, number)
		} else {
			smallNumbers = append(smallNumbers, number)
		}
	}
	for _, large := range largeNumbers {
		for _, small := range smallNumbers {
			isSolutionTwo(large, small)
		}
	}
	log.Fatal("no solution two found")
}

// SolutionThree i/o
func SolutionThree() {
	input := setup()
	largeNumbers := []int{}
	smallNumbers := []int{}
	for _, line := range input {
		number := parseInt(line)
		if number >= (goal / 2) {
			largeNumbers = append(largeNumbers, number)
		} else {
			smallNumbers = append(smallNumbers, number)
		}
	}
	for _, large := range largeNumbers {
		second := 0
		for _, small := range smallNumbers {
			if second == 0 && large+small < goal {
				second = small
				continue
			}
			isSolutionThree(large, second, small)
		}
	}
	log.Fatal("no solution three found")
}
