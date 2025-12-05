package solution

import (
	"sort"
	"strings"
)

type Range struct {
	start int
	end   int
}

func getValidIDs(ranges []string) map[int]bool {
	// Use a map to track unique IDs
	uniqueIDs := make(map[int]bool)
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start := parseInt(parts[0])
		end := parseInt(parts[1])
		for i := start; i <= end; i++ {
			uniqueIDs[i] = true
		}
	}
	return uniqueIDs
}

func getFreshIDs(numbers []string, validIDs map[int]bool) int {
	sum := 0
	for _, number := range numbers {
		if validIDs[parseInt(number)] {
			sum++
		}
	}
	return sum
}

// Approach 2
func isInRange(number string, ranges []string) bool {
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start := parseInt(parts[0])
		end := parseInt(parts[1])
		if parseInt(number) >= start && parseInt(number) <= end {
			return true
		}
	}
	return false
}

// Solution 2
func getValidRanges(ranges []string) []Range {
	// Parse ranges into Range structs
	validRanges := parseRanges(ranges)

	// Sort ranges by start value
	validRanges = sortRanges2(validRanges)

	// Merge overlapping or adjacent ranges
	return mergeRanges(validRanges)
}


func parseRanges(ranges []string) []Range {
	validRanges := []Range{}
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start := parseInt(parts[0])
		end := parseInt(parts[1])
		validRanges = append(validRanges, Range{start: start, end: end})
	}
	return validRanges
}

func sortRanges(ranges []Range) []Range {
	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[i].start > ranges[j].start {
				ranges[i], ranges[j] = ranges[j], ranges[i]
			}
		}
	}
	return ranges
}

func sortRanges2(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
	return ranges
}

func mergeRanges(ranges []Range) []Range {
	result := []Range{}
	for i := 0; i < len(ranges); {
		// Try to merge with subsequent ranges
		j := i + 1;
		for ; j < len(ranges); j++ {
			// If ranges[i] overlaps or is adjacent to ranges[j], merge them
			if ranges[i].end >= ranges[j].start-1 {
				// Merge: extend ranges[i] to cover both
				if ranges[j].end > ranges[i].end {
					ranges[i].end = ranges[j].end
				}
				continue
			} 
			// No overlap, stop merging
			break
		}
		// Add the merged range to result
		result = append(result, ranges[i])
		// Move to the next unmerged range
		i = j
	}
	return result
}

func countRangeValues(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		sum += r.end - r.start + 1
	}
	return sum
}