package jolt

import "testing"

func TestPartOneSolution(t *testing.T) {
	solution := PartOneSolution("./test-input.txt")
	expected := 220
	if solution != expected {
		t.Fatalf("expected %d to equal %d", solution, expected)
	}

	solution = PartOneSolution("./input.txt")
	t.Fatal(solution)
}

func TestPartTwoSolution(t *testing.T) {
	solution := PartTwoSolution("./small-test-input.txt")
	expected := 8
	if solution != expected {
		t.Fatalf("expected %d to equal %d", solution, expected)
	}

	solution = PartTwoSolution("./test-input.txt")
	expected = 19208
	if solution != expected {
		t.Fatalf("expected %d to equal %d", solution, expected)
	}

	solution = PartTwoSolution("./input.txt")
	t.Fatal(solution)
}
