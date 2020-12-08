package handheld

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

var testInstructions = []struct {
	In  string
	Out Instruction
}{
	{In: "nop +0", Out: Instruction{Action: nop, Value: 0}},
	{In: "jmp +3", Out: Instruction{Action: jmp, Value: 3}},
	{In: "acc -99", Out: Instruction{Action: acc, Value: -99}},
}

func TestParseInstruction(t *testing.T) {
	for _, tt := range testInstructions {
		output := parseInstruction(tt.In)
		assert.Equal(t, output.Action, tt.Out.Action)
		assert.Equal(t, output.Value, tt.Out.Value)
	}
}

func TestPartOneSolution(t *testing.T) {
	solution := PartOneSolution("./test-input.txt")
	expected := 5
	if solution != expected {
		t.Fatalf("expected %d to equal %d", solution, expected)
	}

	solution = PartOneSolution("./input.txt")
	t.Fatal(solution)
}

func TestPartTwoSolution(t *testing.T) {
	solution := PartTwoSolution("./test-input.txt")
	expected := 8
	if solution != expected {
		t.Fatalf("expected %d to equal %d", solution, expected)
	}

	solution = PartTwoSolution("./input.txt")
	t.Fatal(solution)
}
