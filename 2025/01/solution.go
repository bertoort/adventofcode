package solution

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("failed to open input file", err)
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("failed to read input file", err)
	}
	return strings.Split(string(b), "\n")
}

func parseInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("failed to parse int", err)
	}
	return number
}

type Safe struct {
	dial   int
	zero   int
	zeroes int
	max    int
}

func (s *Safe) Add(count int) {
	if (s.dial + count) > s.max {
		s.dial = count - (s.max + 1 - s.dial)
		if s.dial != 0 {
			s.zeroes++
		}
		return
	}
	s.dial += count
}

func (s *Safe) Subtract(count int) {
	if s.dial-count < 0 {
		if s.dial != 0 {
			s.zeroes++
		}
		s.dial = (s.max + 1) - (count - s.dial)
		return
	}
	s.dial -= count
}

func (s *Safe) Move(code string) {
	direction := string(code[0])
	count := parseInt(code[1:])
	clicks := count % (s.max + 1)
	s.zeroes += count / (s.max + 1)
	if direction == "L" {
		s.Subtract(clicks)
	} else {
		s.Add(clicks)
	}
	if s.dial == 0 {
		s.zero++
		s.zeroes++
	}
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	safe := Safe{dial: 50, max: 99}
	for _, code := range input {
		safe.Move(code)
	}
	return safe.zero
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	safe := Safe{dial: 50, max: 99}
	for _, code := range input {
		safe.Move(code)
	}
	return safe.zeroes
}
