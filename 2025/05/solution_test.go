package solution

import (
	"testing"
)

func TestRunSolution(t *testing.T) {
	solution := Solution("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}
func TestSolution(t *testing.T) {
	output := Solution("./test-input.txt")
	expected := 3

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
}
func TestRunSolution2(t *testing.T) {
	solution := Solution2("./input.txt")
	t.Fatalf("part one solution: %d", solution)
}
func TestSolution2(t *testing.T) {
	output := Solution2("./test-input.txt")
	expected := 14

	if output != expected {
		t.Fatalf("part one fail: expected %d to equal %d", output, expected)
	}
}

// BenchmarkSortRanges benchmarks the selection sort implementation (O(n²))
// Results: ~12,160 ns/op, 3072 B/op, 1 allocs/op
func BenchmarkSortRanges(b *testing.B) {
	ranges, _ := setup("./input.txt")
	validRanges := parseRanges(ranges)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		testRanges := make([]Range, len(validRanges))
		copy(testRanges, validRanges)
		sortRanges(testRanges)
	}
}

// BenchmarkSortRanges2 benchmarks the sort.Slice implementation (O(n log n))
// Results: ~7,279 ns/op, 3160 B/op, 4 allocs/op
// Approximately 1.67x faster than selection sort
func BenchmarkSortRanges2(b *testing.B) {
	ranges, _ := setup("./input.txt")
	validRanges := parseRanges(ranges)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		testRanges := make([]Range, len(validRanges))
		copy(testRanges, validRanges)
		sortRanges2(testRanges)
	}
}
