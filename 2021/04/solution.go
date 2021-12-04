package solution

import (
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
	return strings.Split(string(b), "\n\n")
}

func parseInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("failed to parse input line", err)
	}
	return number
}

type Box struct {
	value  string
	marked bool
}

type Board struct {
	numbers []Box
	size    int
	won	 bool
}

func newBoards(input []string) []Board {
	boards := []Board{}
	for _, board := range input {
		boards = append(boards, newBoard(board))
	}
	return boards
}

func newBoard(input string) Board {
	rows := strings.Split(input, "\n")
	board := Board{numbers: []Box{}, size: 5}
	for _, row := range rows {
		for i := range row {
			if (i == 0 || (i+2)%3 != 0) && i != len(row)-1 {
				continue
			}
			num := string(row[i-1]) + string(row[i])
			board.numbers = append(board.numbers, Box{value: strings.Trim(num, " ")})
		}
	}
	return board
}

func (b *Board) mark(number string) {
	for i := range b.numbers {
		if b.numbers[i].value == number {
			b.numbers[i].marked = true
		}
	}
}

func (b Board) check() bool {
	verticalMarks := make([]int, b.size)
	horizontalMarks := 0
	for i, box := range b.numbers {
		column := i % b.size
		if column == 0 {
			horizontalMarks = 0
		}
		if box.marked {
			horizontalMarks++
			verticalMarks[column]++
		}
		// check for wins
		if verticalMarks[column] == b.size || horizontalMarks == b.size {
			return true
		}
	}

	return false
}

func (b Board) sumUnmarked() int {
	total := 0
	for _, box := range b.numbers {
		if !box.marked {
			total += parseInt(box.value)
		}
	}
	return total
}

func (b Board) ToString() string {
	result := ""
	for i, row := range b.numbers {
		column := i % b.size
		if i != 0 && column == 0 {
			result += "\n"
		}
		if row.marked {
			result += " " + "*" + row.value
		} else {
			result += " " + row.value
		}
	}
	return result
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	calledNumbers := strings.Split(input[0], ",")
	boards := newBoards(input[1:])
	for _, number := range calledNumbers {
		for _, board := range boards {
			board.mark(number)
			win := board.check()
			if win {
				return parseInt(number) * board.sumUnmarked()
			}
		}
	}
	return 0
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	calledNumbers := strings.Split(input[0], ",")
	boards := newBoards(input[1:])
	result := 0
	for _, number := range calledNumbers {
		for i, board := range boards {
			if board.won {
				continue
			}
			boards[i].mark(number)
			win := board.check()
			if win {
				boards[i].won = true
				result = parseInt(number) * board.sumUnmarked()
			}
		}
	}
	return result
}
