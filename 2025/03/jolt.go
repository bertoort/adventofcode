package solution

import (
	"strconv"
	"strings"
)

type Jolt struct {
	value string
}

func NewJolt(value string) *Jolt {
	return &Jolt{value: value}
}

func splitIntoInts(value string) []int {
	numbers := []int{}
	for _, ch := range strings.Split(value, "") {
		n, _ := strconv.Atoi(ch)
		numbers = append(numbers, n)
	}
	return numbers
}

func findLargestInt(numbers []int) (int, int) {
	index := 0
	largest := 0
	for i, number := range numbers {
		if number > largest {
			largest = number
			index = i
		}
	}
	return largest, index
}

func (j Jolt) GetMax(digits int) int {
	numbers := splitIntoInts(j.value)
	max := ""
	start := 0
	for i := digits; i > 0; i-- {
		f, fIndex := findLargestInt(numbers[start : len(numbers)-i+1])
		max += strconv.Itoa(f)
		start += fIndex + 1
	}
	return parseInt(max)
}
