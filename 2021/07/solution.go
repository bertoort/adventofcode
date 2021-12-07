package solution

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
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

func average(xs []string) int {
	total := 0
	for _, v := range xs {
		total += parseInt(v)
	}
	return int(math.Round(float64(total) / float64(len(xs))))
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	data := []float64{}
	for _, num := range input {
		data = append(data, float64(parseInt(num)))
	}
	mean, _ := stats.Median(data)
	total := 0.0
	for _, crab := range data {
		total += math.Abs(crab-mean)
	}
	return int(total)
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	data := []float64{}
	for _, num := range input {
		data = append(data, float64(parseInt(num)))
	}
	mean, _ := stats.Mean(data)
	total := 0.0
	for _, crab := range data {
		n := math.Abs(crab-math.Floor(mean))
		total += (n*(n+1))/2
	}
	return int(total)
}
