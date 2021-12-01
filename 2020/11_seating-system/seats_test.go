package seats

import "testing"

func TestPartOneSolution(t *testing.T) {
	output := PartOneSolution("./test-input.txt")
	expected := 37
	if output != expected {
		t.Fatalf("expected %v to equal %v", output, expected)
	}

	solution := PartOneSolution("./input.txt")
	t.Fatal(solution)
}

func TestPartTwoSolution(t *testing.T) {
	output := PartTwoSolution("./test-input.txt")
	expected := 26
	if output != expected {
		t.Fatalf("expected %v to equal %v", output, expected)
	}

	solution := PartTwoSolution("./input.txt")
	t.Fatal(solution)
}
