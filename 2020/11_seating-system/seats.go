package seats

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	floor  string = "."
	chair  string = "L"
	person string = "#"
)

func setup(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	lines := strings.Split(string(b), "\n")
	output := [][]string{}
	for _, line := range lines {
		row := []string{}
		for _, place := range line {
			row = append(row, string(place))
		}
		output = append(output, row)
	}
	return output
}

func isOccupied(seats [][]string, row, col int) int {
	if seats[row][col] == person {
		return 1
	}
	return 0
}

func isOccupiedTwo(seats [][]string, row, col int) int {
	seat := seats[row][col]
	switch seat {
	case person:
		return 1
	case chair:
		return 0
	default:
		return -1
	}
}

func countSurrounding(seats [][]string, row, col int) int {
	count := 0

	if row != 0 {
		if col != 0 {
			//check top-left
			count += isOccupied(seats, row-1, col-1)
		}
		//check top
		count += isOccupied(seats, row-1, col)
		if col < len(seats[row])-1 {
			//check top-right
			count += isOccupied(seats, row-1, col+1)
		}
	}
	if col != 0 {
		//check left
		count += isOccupied(seats, row, col-1)
	}
	if col < len(seats[row])-1 {
		//check right
		count += isOccupied(seats, row, col+1)
	}
	if row < len(seats)-1 {
		if col != 0 {
			//check bottom-left
			count += isOccupied(seats, row+1, col-1)
		}
		//check bottom
		count += isOccupied(seats, row+1, col)
		if col < len(seats[row])-1 {
			//check bottom-right
			count += isOccupied(seats, row+1, col+1)
		}
	}
	return count
}

func countSurroundingTwo(seats [][]string, row, col int) int {
	count := 0
	colsLength := len(seats[row]) - 1
	rowsLength := len(seats) - 1

	//check top-left
	found := false
	r := row
	c := col
	for r != 0 && c != 0 && !found {
		r--
		c--
		occupied := isOccupiedTwo(seats, r, c)
		if occupied != -1 {
			found = true
			count += occupied
		}
	}

	//check top
	found = false
	r = row
	for r != 0 && !found {
		r--
		occupied := isOccupiedTwo(seats, r, col)
		if occupied != -1 {
			found = true
			count += occupied
		}
	}

	//check top-right
	found = false
	r = row
	c = col
	for r != 0 && c < colsLength && !found {
		r--
		c++
		occupied := isOccupiedTwo(seats, r, c)
		if occupied != -1 {
			found = true
			count += occupied
		}
	}

	// check left
	found = false
	c = col
	for c != 0 && !found {
		c--
		occupied := isOccupiedTwo(seats, row, c)
		if occupied != -1 {
			found = true
			count += occupied
		}
	}

	//check right
	found = false
	c = col
	for c < colsLength && !found {
		c++
		occupied := isOccupiedTwo(seats, row, c)
		if occupied != -1 {
			found = true
			count += occupied
		}
	}

	// check bottom-left
	found = false
	r = row
	c = col
	for r < rowsLength && c != 0 && !found {
		r++
		c--
		occupied := isOccupiedTwo(seats, r, c)
		if occupied != -1 {
			found = true
			count += occupied
		}
	}

	//check bottom
	found = false
	r = row
	for r < rowsLength && !found {
		r++
		occupied := isOccupiedTwo(seats, r, col)
		if occupied != -1 {
			found = true
			count += occupied
		}
	}

	//check bottom-right
	found = false
	r = row
	c = col
	for r < rowsLength && c < colsLength && !found {
		r++
		c++
		occupied := isOccupiedTwo(seats, r, c)
		if occupied != -1 {
			found = true
			count += occupied
		}
	}

	return count
}

// OccupySeats fills seats and returns seated count with new arrangement
func OccupySeats(seats [][]string) (int, [][]string) {
	newSeats := [][]string{}
	newCount := 0
	for i, row := range seats {
		newRow := []string{}
		count := 0
		for j, seat := range row {
			newSeat := seat
			filled := countSurrounding(seats, i, j)
			if seat == chair && filled == 0 {
				newSeat = person
			}
			if seat == person && filled >= 4 {
				newSeat = chair
			}
			if newSeat == person {
				count++
			}
			newRow = append(newRow, newSeat)
		}
		newSeats = append(newSeats, newRow)
		newCount += count
	}
	return newCount, newSeats
}

// OccupySeatsTwo fills seats and returns seated count with new arrangement
func OccupySeatsTwo(seats [][]string) (int, [][]string) {
	newSeats := [][]string{}
	newCount := 0
	for i, row := range seats {
		newRow := []string{}
		count := 0
		for j, seat := range row {
			newSeat := seat
			filled := countSurroundingTwo(seats, i, j)
			if seat == chair && filled == 0 {
				newSeat = person
			}
			if seat == person && filled >= 5 {
				newSeat = chair
			}
			if newSeat == person {
				count++
			}
			newRow = append(newRow, newSeat)
		}
		newSeats = append(newSeats, newRow)
		newCount += count
	}
	return newCount, newSeats
}

// PartOneSolution io
func PartOneSolution(path string) int {
	seats := setup(path)
	arrangements := map[int]bool{0: true}
	for {
		count, newSeats := OccupySeats(seats)
		if ok := arrangements[count]; !ok {
			arrangements[count] = true
			seats = newSeats
			continue
		}
		return count
	}
}

// PartTwoSolution io
func PartTwoSolution(path string) int {
	seats := setup(path)
	arrangements := map[int]bool{0: true}
	for {
		count, newSeats := OccupySeatsTwo(seats)
		if ok := arrangements[count]; !ok {
			arrangements[count] = true
			seats = newSeats
			continue
		}
		return count
	}
}
