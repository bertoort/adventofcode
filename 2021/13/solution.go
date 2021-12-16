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

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	board := make(map[string]struct{})
	for i, line := range input {
		if line == "" {
			input = input[i+1:]
			break
		}
		board[line] = struct{}{}
	}
	for i, line := range input {
		if i > 0 {
			break
		}
		newBoard := make(map[string]struct{})
		s := strings.Split(line, " ")
		split := strings.Split(s[2], "=")
		fold := split[0]
		foldIndex := parseInt(split[1])
		for coords := range board {
			split := strings.Split(coords, ",")
			x := parseInt(split[0])
			y := parseInt(split[1])
			if fold == "x" && x > foldIndex {
				x = (foldIndex * 2) - x
			}
			if fold == "y" && y > foldIndex {
				y = (foldIndex * 2) - y
			}
			newBoard[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
		}
		board = newBoard
	}
	return len(board)
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	board := make(map[string]struct{})
	for i, line := range input {
		if line == "" {
			input = input[i+1:]
			break
		}
		board[line] = struct{}{}
	}
	maxX := 0
	maxY := 0
	for _, line := range input {
		newBoard := make(map[string]struct{})
		s := strings.Split(line, " ")
		split := strings.Split(s[2], "=")
		fold := split[0]
		foldIndex := parseInt(split[1])
		for coords := range board {
			split := strings.Split(coords, ",")
			x := parseInt(split[0])
			y := parseInt(split[1])
			if fold == "x" && x > foldIndex {
				x = (foldIndex * 2) - x
				maxX = foldIndex
			}
			if fold == "y" && y > foldIndex {
				y = (foldIndex * 2) - y
				maxY = foldIndex
			}
			newBoard[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
		}
		board = newBoard
	}
	y := 0
	for y < maxY {
		x := 0
		line := ""
		for x < maxX {
			_, ok := board[fmt.Sprintf("%d,%d", x, y)]
			if ok {
				line += "#"
			} else {
				line += "."
			}
			x++
		}
		fmt.Println(line)
		y++
	}
	fmt.Println(maxX, maxY)
	return len(board)
}
