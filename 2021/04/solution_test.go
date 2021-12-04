package solution

import "testing"

func TestRunSolution(t *testing.T) {
	solution := Solution("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}
func TestRunSolution2(t *testing.T) {
	solution := Solution2("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}
func TestSolution(t *testing.T) {
	output := Solution("./test-input.txt")
	expected := 4512

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
}
func TestSolution2(t *testing.T) {
	output := Solution2("./test-input.txt")
	expected := 1924

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
}
