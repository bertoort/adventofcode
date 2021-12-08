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
	total := 0
	for _, line := range input {
		split := strings.Split(line, " | ")
		fourLetters := strings.Split(split[1], " ")
		for _, letters := range fourLetters {
			if len(letters) == 2 || len(letters) == 3 || len(letters) == 4 || len(letters) == 7 {
				total++
			}
		}
	}
	return total
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func Trim(s string, letters string) string {
	r := []rune(s)
	result := ""
	for _, c := range r {
		letter := string(c)
		if !strings.Contains(letters, letter) {
			result += letter
		}
	}
	return string(result)
}

func Contains(word, s string) bool {
	result := 0
	for _, c := range s {
		letter := string(c)
		if strings.Contains(word, letter) {
			result++
		}
	}
	return len(s) == result
}

func decodeNumbers(numbers []string) map[string]string {
	decoder := map[string]string{}
	cheatSheet := map[string]string{}
	for _, letters := range numbers {
		switch len(letters) {
		case 2:
			decoder[SortString(letters)] = "1"
			cheatSheet["1"] = SortString(letters)
		case 3:
			decoder[SortString(letters)] = "7"
			cheatSheet["7"] = SortString(letters)
		case 4:
			decoder[SortString(letters)] = "4"
			cheatSheet["4"] = SortString(letters)
		}
	}
	for _, l := range numbers {
		letters := SortString(l)
		switch len(letters) {
		case 5:
			// 2: a,c,d,e,g
			// 3: a,c,d,f,g
			// 5: a,b,d,f,g
			if Contains(letters, cheatSheet["1"]) {
				decoder[SortString(letters)] = "3"
			} else if !Contains(letters, Trim(cheatSheet["4"], cheatSheet["1"])) {
				decoder[SortString(letters)] = "2"
			} else {
				decoder[SortString(letters)] = "5"
			}
		case 6:
			// 0: a,b,c,e,f,g
			// 6: a,b,d,e,f,g
			// 9: a,b,c,d,f,g
			if !Contains(letters, cheatSheet["7"]) {
				decoder[SortString(letters)] = "6"
			} else if Contains(letters, cheatSheet["4"]) {
				decoder[SortString(letters)] = "9"
			} else {
				decoder[SortString(letters)] = "0"
			}
		case 7:
			decoder[SortString(letters)] = "8"
		}
	}
	return decoder
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	total := 0
	for _, line := range input {
		split := strings.Split(line, " | ")
		fourLetters := strings.Split(split[1], " ")
		numbers := strings.Split(split[0], " ")
		decoder := decodeNumbers(numbers)
		result := ""
		for _, letters := range fourLetters {
			result += decoder[SortString(letters)]
		}
		fmt.Println(decoder, result)
		total += parseInt(result)
	}
	return total
}
