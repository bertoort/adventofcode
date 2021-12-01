package encoding

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Preamble list of numbers
type Preamble struct {
	List []int
}

// Check looks at the last 25 numbers for proper sum
func (p *Preamble) Check(preambleLength, index int) bool {
	list := p.List[index-preambleLength : index]
	goal := p.List[index]
	for i, num := range list {
		for j, n := range list[i:] {
			if j == 0 {
				continue
			}
			if num+n == goal {
				return false
			}
		}
	}
	return true
}

func sumSmallestAndLargest(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0] + numbers[len(numbers)-1]
}

// FindSum checks for valid previous contiguous sum
func (p *Preamble) FindSum(preambleLength, index int) int {
	list := p.List[0:index]
	goal := p.List[index]
	for i, number := range list {
		sum := number
		currentNumbers := []int{number}
		for j, n := range list[i:] {
			if j == 0 {
				continue
			}
			if sum+n == goal {
				return sumSmallestAndLargest(append(currentNumbers, n))
			}
			if sum+n < goal {
				sum += n
				currentNumbers = append(currentNumbers, n)
			} else {
				break
			}
		}
	}
	return 0
}

func setup(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	numbers := []int{}
	for _, line := range strings.Split(string(b), "\n") {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}
	return numbers
}

// PartOneSolution io
func PartOneSolution(preambleLength int, path string) int {
	numbers := setup(path)
	preamble := Preamble{List: numbers}
	for i := preambleLength; i <= len(preamble.List); i++ {
		invalid := preamble.Check(preambleLength, i)
		if invalid {
			return preamble.List[i]
		}
	}
	return 0
}

// PartTwoSolution io
func PartTwoSolution(preambleLength int, path string) int {
	numbers := setup(path)
	preamble := Preamble{List: numbers}
	index := 0
	for i := preambleLength; i <= len(preamble.List); i++ {
		invalid := preamble.Check(preambleLength, i)
		if invalid {
			index = i
			break
		}
	}
	return preamble.FindSum(preambleLength, index)
}
