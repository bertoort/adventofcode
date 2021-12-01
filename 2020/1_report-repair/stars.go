package stars

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const goal = 2020

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

func isSolutionTwo(n1, n2 int) int {
	if n1+n2 == goal {
		return n1 * n2
	}
	return 0
}

func isSolutionThree(n1, n2, n3 int) int {
	if n1+n2+n3 == goal {
		return n1 * n2 * n3
	}
	return 0
}

// SolutionPartOne i/o
func SolutionPartOne(path string) int {
	input := setup(path)
	largeNumbers := []int{}
	smallNumbers := []int{}
	for _, line := range input[1:] {
		number := parseInt(line)
		if result := isSolutionTwo(parseInt(input[0]), number); result != 0 {
			return result
		}
		if number >= (goal / 2) {
			largeNumbers = append(largeNumbers, number)
		} else {
			smallNumbers = append(smallNumbers, number)
		}
	}
	for _, large := range largeNumbers {
		for _, small := range smallNumbers {
			if result := isSolutionTwo(large, small); result != 0 {
				return result
			}
		}
	}
	return 0
}

// SolutionPartTwo i/o
func SolutionPartTwo(path string) int {
	input := setup(path)
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
			if result := isSolutionThree(large, second, small); result != 0 {
				return result
			}
		}
	}
	return 0
}
