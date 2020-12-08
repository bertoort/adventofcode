package bags

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Other is other bag inside a bag
type Other struct {
	Bag   string
	Count int
}

const shinyBag = "shiny gold"

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}

func splitBag(bag string) string {
	return strings.Join(strings.Split(strings.Trim(bag, " "), " ")[:2], " ")
}

func splitOthers(other string) []Other {
	bags := strings.Split(strings.Trim(other, " "), ",")
	result := []Other{}
	for _, bag := range bags {
		split := strings.Split(strings.Trim(bag, " "), " ")
		if split[0] == "no" {
			return result
		}
		count := 0
		if num, err := strconv.Atoi(split[0]); err == nil {
			count = num
		}
		result = append(result, Other{Bag: strings.Join(split[1:3], " "), Count: count})
	}
	return result
}

// ParseRule parses a rule sentence
func ParseRule(rule string) (string, []Other) {
	split := strings.Split(rule, "contain")
	bag := splitBag(split[0])
	others := splitOthers(split[1])
	return bag, others
}

func traverseBags(bags map[string][]Other, others []Other) bool {
	for _, other := range others {
		if other.Bag == shinyBag {
			return true
		}
		found := traverseBags(bags, bags[other.Bag])
		if found {
			return found
		}
	}
	return false
}

func traverseCountBags(count int, bags map[string][]Other, others []Other) int {
	total := 1
	for _, other := range others {
		sum := traverseCountBags(other.Count, bags, bags[other.Bag])
		total += sum
	}
	return total * count
}

// PartOneSolution io
func PartOneSolution(path string) int {
	file := setup(path)
	bags := map[string][]Other{}
	for _, rule := range file {
		bag, others := ParseRule(rule)
		bags[bag] = others
	}
	count := 0
	for _, bag := range bags {
		found := traverseBags(bags, bag)
		if found {
			count++
		}
	}
	return count
}

// PartTwoSolution io
func PartTwoSolution(path string) int {
	file := setup(path)
	bags := map[string][]Other{}
	for _, rule := range file {
		bag, others := ParseRule(rule)
		bags[bag] = others
	}
	return traverseCountBags(1, bags, bags[shinyBag]) - 1
}
