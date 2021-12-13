package solution

import (
	"fmt"
	"io/ioutil"
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

func id(x, y int) string {
	return fmt.Sprintf("%d^%d", x, y)
}

func getXY(id string) (int, int) {
	split := strings.Split(id, "^")
	return parseInt(split[0]), parseInt(split[1])
}

func flash(dumbos map[string]int, flashed map[string]bool, point string) int {
	total := 1
	flashes := []string{point}
	dumbos[point] = 0
	for len(flashes) > 0 {
		flash := flashes[0]
		flashes = flashes[1:]
		x, y := getXY(flash)
		adjacent := [][]int{
			{x - 1, y + 1}, {x - 1, y}, {x - 1, y - 1},
			{x + 1, y + 1}, {x + 1, y}, {x + 1, y - 1},
			{x, y + 1}, {x, y - 1},
		}
		for _, coords := range adjacent {
			x := coords[0]
			y := coords[1]
			if x < 0 || x > 9 || y < 0 || y > 9 {
				continue
			}
			id := id(x, y)
			if dumbos[id] == 9 {
				flashes = append(flashes, id)
				dumbos[id] = 0
				flashed[id] = true
				total++
			} else if !flashed[id] {
				dumbos[id]++
			}
		}
	}
	return total
}

func step(dumbos map[string]int) int {
	total := 0
	x := 0
	flashed := map[string]bool{}
	for x < 10 {
		y := 0
		for y < 10 {
			if dumbos[id(x, y)] == 9 {
				flashed[id(x, y)] = true
				total += flash(dumbos, flashed, id(x, y))
			} else if !flashed[id(x, y)] {
				dumbos[id(x, y)]++
			}
			y++
		}
		x++
	}
	return total
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	dumbos := map[string]int{}
	for y, line := range input {
		for x, char := range line {
			dumbos[id(x, y)] = parseInt(string(char))
		}
	}
	i := 0
	total := 0
	for i < 100 {
		total += step(dumbos)
		i++
	}
	return total
}

func step2(dumbos map[string]int) bool {
	x := 0
	flashed := map[string]bool{}
	for x < 10 {
		y := 0
		for y < 10 {
			if dumbos[id(x, y)] == 9 {
				flashed[id(x, y)] = true
				flash(dumbos, flashed, id(x, y))
			} else if !flashed[id(x, y)] {
				dumbos[id(x, y)]++
			}
			y++
		}
		x++
	}
	for _, count := range dumbos {
		if count != 0 {
			return false
		}
	}
	return true
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	dumbos := map[string]int{}
	for y, line := range input {
		for x, char := range line {
			dumbos[id(x, y)] = parseInt(string(char))
		}
	}
	i := 0
	for true {
		sync := step2(dumbos)
		if sync {
			break
		}
		i++
	}
	return i + 1
}
