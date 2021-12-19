package solution

import "testing"

func TestRunSolution(t *testing.T) {
	// not 4006
	// should 4517
	solution := Solution("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}

func TestSolution(t *testing.T) {
	output := Solution("./test-input.txt")
	expected := 1588

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
}
func TestRunSolution2(t *testing.T) {
	solution := Solution2("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}
func TestSolution2(t *testing.T) {
	output := Solution2("./test-input.txt")
	expected := 0

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
}
