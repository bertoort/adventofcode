package customs

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n\n")
}

// CustomCount sums custom unique question count
func CustomCount(group string) int {
	count := 0
	questionIndex := map[rune]bool{}
	questions := strings.Split(group, "\n")
	for _, question := range questions {
		for _, letter := range question {
			if ok, _ := questionIndex[letter]; !ok {
				count++
				questionIndex[letter] = true
			}
		}
	}
	return count
}

// CustomCountTwo sums custom repeated question count
func CustomCountTwo(group string) int {
	questionIndex := map[rune]int{}
	questions := strings.Split(group, "\n")
	for _, question := range questions {
		for _, letter := range question {
			if _, ok := questionIndex[letter]; !ok {
				questionIndex[letter] = 1
			} else {
				questionIndex[letter]++
			}
		}
	}
	count := 0
	for _, value := range questionIndex {
		if value == len(questions) {
			count++
		}
	}
	return count
}

// PartOneSolution io
func PartOneSolution(path string) int {
	file := setup(path)
	total := 0
	for _, group := range file {
		sum := CustomCount(group)
		total += sum
	}
	return total
}

// PartTwoSolution io
func PartTwoSolution(path string) int {
	file := setup(path)
	total := 0
	for _, group := range file {
		sum := CustomCountTwo(group)
		total += sum
	}
	return total
}
