package jolt

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var maxJolt = 3

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}

func sortAdapters(file []string) []int {
	adapters := []int{0}
	for _, adapter := range file {
		a, _ := strconv.Atoi(adapter)
		adapters = append(adapters, a)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	return adapters
}

// PartOneSolution io
func PartOneSolution(path string) int {
	file := setup(path)
	adapters := sortAdapters(file)
	jolt := 0
	oneDiff := 0
	threeDiff := 0
	for _, adapter := range adapters {
		diff := adapter - jolt
		jolt = adapter
		if diff == 1 {
			oneDiff++
		}
		if diff == 3 {
			threeDiff++
		}
	}
	return oneDiff * threeDiff
}

func checkNext(adapters []int, index, jolt int) bool {
	if index+1 < len(adapters) {
		return jolt-adapters[index+1] <= maxJolt
	}
	return false
}

func countPossibleAdapters(fromIndex int, adapters []int, visited map[int]int) int {
	if fromIndex >= len(adapters)-maxJolt {
		return 1
	}

	num := adapters[fromIndex]
	if res, ok := visited[num]; ok {
		return res
	}

	var count int
	for i := fromIndex + 1; i < fromIndex+4; i++ {
		n := adapters[i]
		if areCompatible(num, n) {
			count += countPossibleAdapters(i, adapters, visited)
		}
	}

	visited[num] = count
	return count
}

func areCompatible(low, high int) bool {
	return low+1 == high || low+2 == high || low+3 == high
}

// PartTwoSolution io
func PartTwoSolution(path string) int {
	file := setup(path)
	adapters := sortAdapters(file)
	return countPossibleAdapters(0, adapters, make(map[int]int))
}
