package solution

import "testing"

func TestRunSolution(t *testing.T) {
	solution := Solution("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}
func TestSolution(t *testing.T) {
	output := Solution("./test-input.txt")
	expected := 0

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
}
