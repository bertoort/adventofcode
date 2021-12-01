package password

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Content information from parsed line
type Content struct {
	Min      int
	Max      int
	Letter   rune
	Password string
}

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}

func parseLine(line string) Content {
	content := strings.Split(line, ": ")
	policy := strings.Split(content[0], " ")
	ranges := strings.Split(policy[0], "-")
	min, err := strconv.Atoi(ranges[0])
	if err != nil {
		log.Fatal("parsing min", min)
	}
	max, err := strconv.Atoi(ranges[1])
	if err != nil {
		log.Fatal("parsing max", max)
	}
	letter := policy[1]
	password := content[1]
	return Content{
		Min:      min,
		Max:      max,
		Letter:   []rune(letter)[0],
		Password: password,
	}
}

func isValidPassword(content Content) bool {
	count := 0
	for _, l := range content.Password {
		if l == content.Letter {
			count++
		}
	}
	return count >= content.Min && count <= content.Max
}

func isValidPartTwoPassword(content Content) bool {
	letter := []byte(string(content.Letter))[0]
	first := content.Password[content.Min-1] == letter
	second := content.Password[content.Max-1] == letter
	return (first || second) && !(first && second)
}

// PartOneSolution i/o
func PartOneSolution(path string) int {
	file := setup(path)
	total := 0
	for _, line := range file {
		content := parseLine(line)
		if valid := isValidPassword(content); valid {
			total++
		}
	}
	return total
}

// PartTwoSolution i/o
func PartTwoSolution(path string) int {
	file := setup(path)
	total := 0
	for _, line := range file {
		content := parseLine(line)
		if valid := isValidPartTwoPassword(content); valid {
			total++
		}
	}
	return total
}
