package solution

import (
	"strconv"
	"strings"
)

type IDRange struct {
	start int
	end   int
}

func NewIDRange(input string) *IDRange {
	split := strings.Split(input, "-")
	return &IDRange{
		start: parseInt(split[0]),
		end:   parseInt(split[1]),
	}
}

func isOdd(id int) bool {
	return id%2 != 0
}

func validateID(id int) bool {
	stringID := strconv.Itoa(id)
	length := len(stringID)
	if isOdd(length) {
		return true
	}
	firstHalf := stringID[:length/2]
	secondHalf := stringID[length/2:]
	return firstHalf != secondHalf
}

func (id *IDRange) GetInvalidIDs() []int {
	result := []int{}
	for id.start <= id.end {
		isValid := validateID(id.start)
		if !isValid {
			result = append(result, id.start)
		}
		id.start++
	}
	return result
}

func (id *IDRange) GetSumOfInvalidIDs() int {
	sum := 0
	for _, id := range id.GetInvalidIDs() {
		sum += id
	}
	return sum
}

// Part Two

func isRepeating(split, stringID string) bool {
	concat := split
	for len(concat) < len(stringID) {
		concat += split
	}
	return concat == stringID
}

func validateSuperID(id int) bool {
	stringID := strconv.Itoa(id)
	length := len(stringID)
	for i := 0; i < length/2; i++ {
		split := stringID[:i+1]
		repeats := isRepeating(split, stringID)
		if repeats {
			return false
		}
	}
	return true
}

func (id *IDRange) GetSuperInvalidIDs() []int {
	result := []int{}
	for id.start <= id.end {
		isValid := validateSuperID(id.start)
		if !isValid {
			result = append(result, id.start)
		}
		id.start++
	}
	return result
}

func (id *IDRange) GetSumOfSuperInvalidIDs() int {
	sum := 0
	for _, id := range id.GetSuperInvalidIDs() {
		sum += id
	}
	return sum
}
