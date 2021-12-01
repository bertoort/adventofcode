package customs

import "testing"

func TestCustomCount(t *testing.T) {
	expected := 11
	count := PartOneSolution("./test-input.txt")
	if expected != count {
		t.Fatalf("expected %d to equal %d", count, expected)
	}
}

func TestCustomCountTwo(t *testing.T) {
	expected := 6
	count := PartTwoSolution("./test-input.txt")
	if expected != count {
		t.Fatalf("expected %d to equal %d", count, expected)
	}
}

func TestPartOne(t *testing.T) {
	solution := PartOneSolution("./input.txt")
	t.Fatal(solution)
}

func TestPartTwo(t *testing.T) {
	solution := PartTwoSolution("./input.txt")
	t.Fatal(solution)
}
