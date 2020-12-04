package toboggan

import "testing"

func TestPassword(t *testing.T) {
	total := PartOneSolution("./test-input.txt")
	expected := 7
	if total != expected {
		t.Fatalf("expected %d to equal %d trees", total, expected)
	}

	total2 := PartTwoSolution("./test-input.txt")
	expected2 := 336
	if total2 != expected2 {
		t.Fatalf("expected %d to equal %d", total2, expected2)
	}
}

func TestPartOne(t *testing.T) {
	total := PartOneSolution("./input.txt")
	t.Fatalf("part one total: %d", total)
}

func TestPartTwo(t *testing.T) {
	total := PartTwoSolution("./input.txt")
	t.Fatalf("part two total: %d", total)
}
