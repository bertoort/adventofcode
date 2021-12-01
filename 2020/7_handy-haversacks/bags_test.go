package bags

import "testing"

func TestParseRule(t *testing.T) {
	rule := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	bag, others := ParseRule(rule)
	expectedBag := "light red"
	expectedOther := []string{"bright white", "muted yellow"}
	if bag != expectedBag {
		t.Fatalf("expected %s to equal %s", bag, expectedBag)
	}
	if len(others) != len(expectedOther) {
		t.Fatalf("expected %d other, got %d", len(expectedOther), len(others))
	}
	if others[0].Bag != expectedOther[0] || others[1].Bag != expectedOther[1] {
		t.Fatalf("expected %v to equal %v", others, expectedOther)
	}
}

func TestPartOneSolution(t *testing.T) {
	output := PartOneSolution("./test-input.txt")
	expected := 4
	if output != expected {
		t.Fatalf("expected %v to equal %v", output, expected)
	}

	solution := PartOneSolution("./input.txt")
	t.Fatal(solution)
}

func TestPartTwoSolution(t *testing.T) {
	output := PartTwoSolution("./test-input.txt")
	expected := 32
	if output != expected {
		t.Fatalf("expected %v to equal %v", output, expected)
	}

	solution := PartTwoSolution("./input.txt")
	t.Fatal(solution)
}
