package solution

import (
	"fmt"
	"io/ioutil"
	"log"
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

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	template := map[string]string{}
	for _, step := range input[2:] {
		split := strings.Split(step, " -> ")
		template[split[0]] = split[1]
	}
	polymer := map[string]int{}
	for i, char := range input[0] {
		l1 := string(char)
		if i == len(input[0])-1 {
			break
		}
		l2 := string(input[0][i+1])
		polymer[l1+l2]++
	}
	i := 0
	for i < 40 {
		newPolymer := map[string]int{}
		for k, v := range polymer {
			letter := template[k]
			l1 := string(k[0])
			l2 := string(k[1])
			newPolymer[l1+letter] += v
			newPolymer[letter+l2] += v
		}
		polymer = newPolymer
		i++
	}
	totals := []int{}
	counts := map[string]int{}
	for k, v := range polymer {
		l1 := string(k[0])
		l2 := string(k[1])
		counts[l1] += v
		counts[l2] += v
	}
	counts[string(input[0][0])]++
	counts[string(input[0][len(input[0])-1])]++
	for k := range counts {
		counts[k] /= 2
	}
	for _, v := range counts {
		totals = append(totals, v)
	}
	sort.Ints(totals)
	fmt.Println(polymer, counts, totals)
	return totals[len(totals)-1] - totals[0]
}

// Solution2 i/o
func Solution2(path string) int {
	// same as above, 10 -> 40
	return 0
}
