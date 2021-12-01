package encoding

import "testing"

func TestPartOneSolution(t *testing.T) {
	preambleLength := 5
	solution := PartOneSolution(preambleLength, "./test-input.txt")
	expected := 127
	if solution != expected {
		t.Fatalf("expected %d to equal %d", solution, expected)
	}

	solution = PartOneSolution(25, "./input.txt")
	t.Fatal(solution)
}

func TestPartTwoSolution(t *testing.T) {
	preambleLength := 5
	solution := PartTwoSolution(preambleLength, "./test-input.txt")
	expected := 62
	if solution != expected {
		t.Fatalf("expected %d to equal %d", solution, expected)
	}

	solution = PartTwoSolution(25, "./input.txt")
	t.Fatal(solution)
}

func TestSumSmallestAndLargest(t *testing.T) {
	numbers := []int{15, 25, 47, 40}
	expected := 62
	solution := sumSmallestAndLargest(numbers)
	if solution != expected {
		t.Fatalf("expected %d to equal %d", solution, expected)
	}
}
