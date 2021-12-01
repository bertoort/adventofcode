package stars

import "testing"

func TestSolutionPartOne(t *testing.T) {
	solution := SolutionPartOne("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}

func TestSolutionPartTwo(t *testing.T) {
	solution := SolutionPartTwo("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}

func TestSolutions(t *testing.T) {
	output := SolutionPartOne("./test-input.txt")
	expected := 514579

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
	output2 := SolutionPartTwo("./test-input.txt")
	expected2 := 241861950
	if output != expected {
		t.Fatalf("part two fail: expected %d to equal %d", output2, expected2)
	}
}
