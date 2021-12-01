package passport

import "testing"

func TestPartOneValidate(t *testing.T) {
	total := PartOneSolution("./passports.txt")
	expected := 2
	if total != expected {
		t.Fatalf("expected %d to equal %d valid passorts", total, expected)
	}
}

func TestPartTwoValidate(t *testing.T) {
	totalInvalid := PartTwoSolution("./invalid-passports.txt")
	if totalInvalid != 0 {
		t.Fatalf("should have zero valid passports, not %d", totalInvalid)
	}
	totalValid := PartTwoSolution("./valid-passports.txt")
	expected := 4
	if totalValid != expected {
		t.Fatalf("expected %d to equal %d valid passorts", totalValid, expected)
	}
}

func TestPartOneSolution(t *testing.T) {
	total := PartOneSolution("./input.txt")
	t.Fatalf("part one total: %d", total)
}

func TestPartTwoSolution(t *testing.T) {
	total := PartTwoSolution("./input.txt")
	t.Fatalf("part two total: %d", total)
}
