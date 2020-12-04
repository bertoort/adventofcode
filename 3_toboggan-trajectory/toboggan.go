package toboggan

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const tree = 35

func setup() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}

func hitsTree(line string, location int) bool {
	index := location % len(line)
	return line[index] == tree
}

// PartOneSolution i/o
func PartOneSolution() {
	file := setup()
	location := 0
	total := 0
	rightMove := 3
	for _, line := range file {
		if hit := hitsTree(line, location); hit {
			total++
		}
		location += rightMove
	}
	log.Fatalf("part one total: %d trees", total)
}

// PartTwoSolution i/o
func PartTwoSolution() {
	file := setup()
	total := 1
	input := []struct {
		right int
		down  int
	}{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, path := range input {
		location := 0
		trees := 0
		for i, line := range file {
			if i%path.down == 0 {
				if hit := hitsTree(line, location); hit {
					trees++
				}
				location += path.right
			}
		}
		total *= trees
	}
	log.Fatalf("part two total: %d", total)
}
