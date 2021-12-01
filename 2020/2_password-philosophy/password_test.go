package password

import "testing"

func TestPassword(t *testing.T) {
	total := PartOneSolution("./test-input.txt")
	expected := 2
	if total != expected {
		t.Fatalf("expected %d to equal %d valid passwords", total, expected)
	}

	total2 := PartTwoSolution("./test-input.txt")
	expected2 := 1
	if total2 != expected2 {
		t.Fatalf("expected %d to equal %d valid passwords", total2, expected2)
	}
}

func TestPasswordPartOne(t *testing.T) {
	total := PartOneSolution("./input.txt")
	t.Fatalf("part one total: %d", total)
}

func TestPasswordPartTwo(t *testing.T) {
	total := PartTwoSolution("./input.txt")
	t.Fatalf("part two total: %d", total)
}
