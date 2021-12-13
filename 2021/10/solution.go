package solution

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}

func parseInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("failed to parse input line", err)
	}
	return number
}

var pairs = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
	">": "<",
}

var reversePairs = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
	"<": ">",
}

func findError(line string) (string, []string) {
	stack := []string{}
	for _, c := range line {
		char := string(c)
		if char == "(" || char == "<" || char == "[" || char == "{" {
			stack = append(stack, char)
		} else if stack[len(stack)-1] == pairs[char] {
			stack = stack[:len(stack)-1]
		} else {
			return char, stack
		}
	}
	return "", stack
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)

	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	counts := map[string]int{
		")": 0,
		"}": 0,
		"]": 0,
		">": 0,
	}
	for _, line := range input {
		char, _ := findError(line)
		counts[char]++
	}
	total := 0
	for key, value := range counts {
		total += points[key] * value
	}
	return total
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)

	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	list := []int{}
	for _, line := range input {
		char, stack := findError(line)
		if char != "" {
			continue
		}
		total := 0
		fmt.Println(stack)
		for i := len(stack)-1; i >= 0; i-- {
			total = ((total*5) + points[reversePairs[stack[i]]])
		}
		list = append(list, total)
	}
	fmt.Println(list)
	sort.Ints(list)
	return list[int(math.Floor(float64(len(list))/2))]
}
