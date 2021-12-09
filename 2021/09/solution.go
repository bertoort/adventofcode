package solution

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
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

func c(point byte) int {
	return parseInt(string(point))
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	total := 0
	for i, row := range input {
		isTop := i == 0
		isBottom := i == len(input)-1
		for j, r := range row {
			point := parseInt(string(r))
			isLeft := j == 0
			isRight := j == len(row)-1
			if !isTop && c(input[i-1][j]) <= point {
				continue
			}
			if !isLeft && c(row[j-1]) <= point {
				continue
			}
			if !isRight && c(row[j+1]) <= point {
				continue
			}
			if !isBottom && c(input[i+1][j]) <= point {
				continue
			}
			total += point + 1
		}
	}
	return total
}

func totalBasin(board []string, y, x int) int {
	checked := map[string]bool{}
	points := [][]int{{y, x}}
	total := 0
	for len(points) > 0 {
		total += 1
		current := points[0]
		y := current[0]
		x := current[1]
		// check left
		if x != 0 && c(board[y][x-1]) != 9 && !checked[strconv.Itoa(y)+strconv.Itoa(x-1)] {
			points = append(points, []int{y, x - 1})
			checked[strconv.Itoa(y)+strconv.Itoa(x-1)] = true
		}
		if y != 0 && c(board[y-1][x]) != 9 && !checked[strconv.Itoa(y-1)+strconv.Itoa(x)] {
			points = append(points, []int{y - 1, x})
			checked[strconv.Itoa(y-1)+strconv.Itoa(x)] = true
		}
		if y != len(board)-1 && c(board[y+1][x]) != 9 && !checked[strconv.Itoa(y+1)+strconv.Itoa(x)] {
			points = append(points, []int{y + 1, x})
			checked[strconv.Itoa(y+1)+strconv.Itoa(x)] = true
		}
		if x != len(board[0])-1 && c(board[y][x+1]) != 9 && !checked[strconv.Itoa(y)+strconv.Itoa(x+1)] {
			points = append(points, []int{y, x + 1})
			checked[strconv.Itoa(y)+strconv.Itoa(x+1)] = true
		}
		points = points[1:]
	}
	return total - 1
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	basins := []int{}
	for i, row := range input {
		isTop := i == 0
		isBottom := i == len(input)-1
		for j, r := range row {
			point := parseInt(string(r))
			isLeft := j == 0
			isRight := j == len(row)-1
			if !isTop && c(input[i-1][j]) <= point {
				continue
			}
			if !isLeft && c(row[j-1]) <= point {
				continue
			}
			if !isRight && c(row[j+1]) <= point {
				continue
			}
			if !isBottom && c(input[i+1][j]) <= point {
				continue
			}
			basins = append(basins, totalBasin(input, i, j))
		}
	}
	sort.Ints(basins)
	total := 1
	for _, value := range basins[len(basins)-3:] {
		total *= value
	}
	return total
}
