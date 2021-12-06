package solution

import (
	"io/ioutil"
	"log"
	"math"
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findDistance(startS, endS string) []string {
	start := parseInt(startS)
	end := parseInt(endS)
	lower := min(start, end)
	distance := int(math.Abs(float64(start - end)))
	i := 0
	result := []string{}
	for i <= distance {
		result = append(result, strconv.Itoa(lower+i))
		i++
	}
	return result
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	counts := map[string]int{}
	for _, row := range input {
		coords := strings.Split(row, " -> ")
		start := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")
		if start[0] == end[0] {
			points := findDistance(start[1], end[1])
			x := start[0]
			for _, y := range points {
				counts[x+","+y]++
			}
		} else if start[1] == end[1] {
			points := findDistance(start[0], end[0])
			y := start[1]
			for _, point := range points {
				counts[point+","+y]++
			}
		}
	}
	intersections := 0
	for _, count := range counts {
		if count > 1 {
			intersections++
		}
	}
	return intersections
}

func findPoints(start, end []string) []string {
	xInc := true
	if parseInt(start[0]) > parseInt(end[0]) {
		xInc = false
	}
	yInc := true
	if parseInt(start[1]) > parseInt(end[1]) {
		yInc = false
	}
	distance := int(math.Abs(float64(parseInt(start[0]) - parseInt(end[0]))))
	index := 0
	result := []string{}
	for index <= distance {
		x := parseInt(start[0]) - index
		if xInc {
			x = parseInt(start[0]) + index
		}
		y := parseInt(start[1]) - index
		if yInc {
			y = parseInt(start[1]) + index
		}
		result = append(result, strconv.Itoa(x)+","+strconv.Itoa(y))
		index++
	}
	return result
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	counts := map[string]int{}
	for _, row := range input {
		coords := strings.Split(row, " -> ")
		start := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")
		if start[0] != end[0] && start[1] != end[1] {
			points := findPoints(start, end)
			for _, point := range points {
				counts[point]++
			}
		} else if start[0] == end[0] {
			points := findDistance(start[1], end[1])
			x := start[0]
			for _, y := range points {
				counts[x+","+y]++
			}
		} else if start[1] == end[1] {
			points := findDistance(start[0], end[0])
			y := start[1]
			for _, point := range points {
				counts[point+","+y]++
			}
		}
	}
	intersections := 0
	for _, count := range counts {
		if count > 1 {
			intersections++
		}
	}
	return intersections
}
