 package solution

import (
	"fmt"
	"sort"
	"strings"
)

const start byte = 'S'
const splitter byte = '^'
const beam byte = '|'

type Tachyon []string

func (t Tachyon) String() string {
	return strings.Join(t, "\n")
}

// Solution 1

func drawBeam(tachyon Tachyon, row int) (Tachyon, int) {
	newTachyon := make(Tachyon, len(tachyon))
	copy(newTachyon, tachyon)
	splits := 0
	for i := range tachyon[row] {
		previousRow := tachyon[row-1]
		currentRow := []byte(newTachyon[row])
		if previousRow[i] == beam || previousRow[i] == start {
			if newTachyon[row][i] == splitter {
				currentRow[i-1] = beam
				currentRow[i+1] = beam
				splits++
			} else {
				currentRow[i] = beam
			}
		}
		newTachyon[row] = string(currentRow)
	}
	return newTachyon, splits
}

func shootBeam(tachyon Tachyon) (Tachyon, int) {
	newTachyon := make(Tachyon, len(tachyon))
	copy(newTachyon, tachyon)
	totalSplits := 0
	for i := 1; i < len(tachyon); i++ {
		t, splits := drawBeam(newTachyon, i)
		newTachyon = t
		totalSplits += splits
	}
	return newTachyon, totalSplits
}

// Solution 2 - Original (exponential time complexity, hangs on large inputs)
/*
func drawBeams(tachyon Tachyon, index int) (Tachyon, Tachyon) {
	t1 := make(Tachyon, len(tachyon))
	t2 := make(Tachyon, len(tachyon))
	copy(t1, tachyon)
	copy(t2, tachyon)
	split := false
	for i := range tachyon[index] {
		previousRow := tachyon[index-1]
		currentRow1 := []byte(t1[index])
		currentRow2 := []byte(t2[index])
		if previousRow[i] == beam || previousRow[i] == start {
			if tachyon[index][i] == splitter {
				split = true
				currentRow1[i-1] = beam
				currentRow2[i+1] = beam
			} else {
				currentRow1[i] = beam
			}
		}
		t1[index] = string(currentRow1)
		t2[index] = string(currentRow2)
	}
	if (split) {
		return t1, t2
	}
	return t1, nil
}

func drawTimeline(tachyon Tachyon, index int) int {
	if (tachyon == nil) {
		return 0
	}
	// End of a timeline
	if (index >= len(tachyon)) {
		return 1
	}
	t1, t2 := drawBeams(tachyon, index)
	// No split
	if t2 == nil {
		return drawTimeline(t1, index+1)
	}
	// Split occurred
	return drawTimeline(t1, index+1) + drawTimeline(t2, index+1)
}

func countTimelines(tachyon Tachyon) int {
	currentIndex := 1
	timelines := drawTimeline(tachyon, currentIndex)
	return timelines
}
*/

// Solution 2 - Optimized with memoization

// State represents which positions in a row have beams
type beamState map[int]bool

// getBeamPosition extracts which positions have beams from a row
func getBeamPosition(row string) beamState {
	state := make(beamState)
	for i := range row {
		if row[i] == beam || row[i] == start {
			state[i] = true
			return state
		}
	}
	return state
}

// applyBeamsToRow applies beams from previous row to current row, returns new beam positions and whether split occurred
func applyBeamsToRow(tachyon Tachyon, rowIndex int, prevBeams beamState) (beamState, beamState) {
	if rowIndex >= len(tachyon) {
		return nil, nil
	}
	
	currentRow := tachyon[rowIndex]
	newBeams1 := make(beamState)
	newBeams2 := make(beamState)
	split := false
	
	for i := range currentRow {
		// Check if beam comes from above
		if prevBeams[i] {
			if currentRow[i] == splitter {
				// Split: beam goes left and right
				if i > 0 {
					newBeams1[i-1] = true
				}
				if i < len(currentRow)-1 {
					newBeams2[i+1] = true
				}
				split = true
			} else {
				// Normal beam propagation
				newBeams1[i] = true
			}
		}
	}
	
	if split {
		return newBeams1, newBeams2
	}
	return newBeams1, nil
}

// memoKey represents a memoization key
type memoKey struct {
	rowIndex int
	beams    string // serialized beam positions
}

// serializeBeams converts beam state to a UNIQUE string for memoization
func serializeBeams(beams beamState) string {
	if len(beams) == 0 {
		return ""
	}
	// Collect positions and sort for deterministic serialization
	positions := make([]int, 0, len(beams))
	for pos := range beams {
		positions = append(positions, pos)
	}
	sort.Ints(positions)
	
	// Build string representation, i.e. "5,10,15"
	result := ""
	for i, pos := range positions {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", pos)
	}
	return result
}

func drawTimelineMemo(tachyon Tachyon, rowIndex int, prevBeams beamState, memo map[memoKey]int) int {
	// End of timeline
	if rowIndex >= len(tachyon) {
		return 1
	}
	
	// Check memoization
	key := memoKey{
		rowIndex: rowIndex,
		beams:    serializeBeams(prevBeams),
	}
	if count, exists := memo[key]; exists {
		return count
	}
	
	// Compute new beam positions
	beams1, beams2 := applyBeamsToRow(tachyon, rowIndex, prevBeams)
	
	var count int
	if beams2 == nil {
		// No split
		count = drawTimelineMemo(tachyon, rowIndex+1, beams1, memo)
	} else {
		// Split occurred
		count = drawTimelineMemo(tachyon, rowIndex+1, beams1, memo) + 
		        drawTimelineMemo(tachyon, rowIndex+1, beams2, memo)
	}
	
	// Memoize result
	memo[key] = count
	return count
}

func countTimelines(tachyon Tachyon) int {
	if len(tachyon) == 0 {
		return 0
	}
	
	// Get initial beam position from first row
	initialBeam := getBeamPosition(tachyon[0])
	
	// Use memoization to avoid recomputing same states
	memo := make(map[memoKey]int)
	
	return drawTimelineMemo(tachyon, 1, initialBeam, memo)
}