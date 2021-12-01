package solution

import (
	"io/ioutil"
	"log"
	"os"
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

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	print(input)
	return 0
}
