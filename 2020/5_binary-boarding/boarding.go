package boarding

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

const rows = 127
const columns = 7

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}

func particion(total int, particion string) int {
	start := 0
	end := total
	for _, l := range particion {
		letter := string(l)
		start, end = next(start, end, letter == "F" || letter == "L")
	}
	return start
}

func next(start, end int, lower bool) (int, int) {
	if end-start == 1 {
		if lower {
			return start, 0
		}
		return end, 0
	}
	half := (start + end + 2) / 2
	if lower {
		return start, (half - 1)
	}
	return half, end
}

// BinaryParticion particions row and column of seat
func BinaryParticion(seat string) int {
	row := particion(rows, seat[:7])
	column := particion(columns, seat[7:])
	return row*8 + column
}

// PartOneSolution io
func PartOneSolution(path string) int {
	file := setup(path)
	highest := 0
	for _, seat := range file {
		position := BinaryParticion(seat)
		if position > highest {
			highest = position
		}
	}
	return highest
}

// PartTwoSolution io
func PartTwoSolution(path string) int {
	file := setup(path)
	seats := []int{}
	for _, seat := range file {
		position := BinaryParticion(seat)
		seats = append(seats, position)
	}
	sort.Ints(seats)
	for i, seat := range seats {
		if i > len(seats)-3 {
			break
		}
		if seats[i+1] != seat+1 {
			return seat + 1
		}
	}
	return 0
}
