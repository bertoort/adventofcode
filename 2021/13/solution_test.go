package solution

import "testing"

func TestRunSolution(t *testing.T) {
	// not 959
	solution := Solution("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}
func TestSolution(t *testing.T) {
	output := Solution("./test-input.txt")
	expected := 17

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
	expected := 16

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
}
