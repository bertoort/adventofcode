package handheld

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Action types for instuctions
type Action string

var (
	nop Action = "nop"
	acc Action = "acc"
	jmp Action = "jmp"
)

// Instruction holds handheld instruction info
type Instruction struct {
	Action
	Value int
}

func parseInstruction(instruction string) Instruction {
	split := strings.Split(instruction, " ")
	action := nop
	switch split[0] {
	case "acc":
		action = acc
	case "jmp":
		action = jmp
	}
	number, _ := strconv.Atoi(split[1][1:])
	if string(split[1][0]) == "-" {
		number *= -1
	}
	return Instruction{Action: action, Value: number}
}

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}

// PartOneSolution io
func PartOneSolution(path string) int {
	instructions := setup(path)
	accumulator := 0
	visited := map[int]bool{}
	index := 0
	for {
		if ok := visited[index]; !ok {
			visited[index] = true
		} else {
			return accumulator
		}
		instruction := parseInstruction(instructions[index])
		switch instruction.Action {
		case nop:
			index++
		case acc:
			index++
			accumulator += instruction.Value
		case jmp:
			index += instruction.Value
		}
	}
}

func solve(instructions []string, tried map[int]bool) int {
	accumulator := 0
	visited := map[int]bool{}
	index := 0
	changed := false
	for {
		if index == len(instructions) {
			return accumulator
		}
		if ok := visited[index]; !ok {
			visited[index] = true
		} else {
			return solve(instructions, tried)
		}
		instruction := parseInstruction(instructions[index])
		if ok := tried[index]; !ok && !changed {
			if instruction.Action == nop {
				instruction.Action = jmp
			} else {
				instruction.Action = nop
			}
			tried[index] = true
			changed = true
		}
		switch instruction.Action {
		case nop:
			index++
		case acc:
			index++
			accumulator += instruction.Value
		case jmp:
			index += instruction.Value
		}
	}
}

// PartTwoSolution io
func PartTwoSolution(path string) int {
	instructions := setup(path)
	return solve(instructions, map[int]bool{})
}
