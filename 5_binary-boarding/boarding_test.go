package boarding

import (
	"testing"
)

var boardingTests = []struct {
	In  string
	Out int
}{{"FBFBBFFRLR", 357}, {"BFFFBBFRRR", 567}, {"FFFBBBFRRR", 119}, {"BBFFBBFRLL", 820}}

func TestBinaryParticion(t *testing.T) {
	for _, tt := range boardingTests {
		output := BinaryParticion(tt.In)
		if output != tt.Out {
			t.Fatalf("expected %d to equal %d for %s", output, tt.Out, tt.In)
		}
	}
}

func TestPartOneSolution(t *testing.T) {
	solution := PartOneSolution("./test-input.txt")
	highest := 820
	if solution != highest {
		t.Fatalf("expected %d to bet the highest %d", solution, highest)
	}

	answer := PartOneSolution("./input.txt")
	t.Fatal(answer)
}

func TestPartTwoSolution(t *testing.T) {
	answer := PartTwoSolution("./input.txt")
	t.Fatal(answer)
}
