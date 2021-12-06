package solution

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	return strings.Split(string(b), ",")
}

func parseInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("failed to parse input line", err)
	}
	return number
}

func tick(fishies []int, day, limit int) int {
	if day == limit {
		return len(fishies)
	}
	pool := []int{}
	for _, fish := range fishies {
		if fish == 0 {
			pool = append(pool, 6)
			pool = append(pool, 8)
			continue
		}
		pool = append(pool, fish-1)
	}
	return tick(pool, day+1, limit)
}

// Solution i/o
func Solution(path string) int {
	limit := 80
	input := setup(path)
	fishies := []int{}
	for _, num := range input {
		fishies = append(fishies, parseInt(num))
	}
	return tick(fishies, 0, limit)
}

func sumFish(fishies map[int]int) int {
	total := 0
	fmt.Println(fishies)
	for _, fish := range fishies {
		total += fish
	}
	return total
}

func tick2(fishies map[int]int, limit int) int {
	day := 0
	for day < limit {
		pool := map[int]int{}
		for fish, count := range fishies {
			if fish == 0 {
				pool[6] += count
				pool[8] += count
				continue
			}
			pool[fish-1] += count
		}
		fishies = pool
		day++
	}
	return sumFish(fishies)
}

// Solution2 i/o
func Solution2(path string) int {
	limit := 256
	input := setup(path)
	fishies := map[int]int{}
	for _, num := range input {
		fishies[parseInt(num)]++
	}
	return tick2(fishies, limit)
}
