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
	return strings.Split(string(b), "\n")
}

func parseInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("failed to parse input line", err)
	}
	return number
}

func k(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func explore(x, y int, path []int, board []string, visited map[string]struct{}) [][]int {
	paths := [][]int{}
	checks := [][]int{{x, y + 1}, {x + 1, y}}
	visited[k(x, y)] = struct{}{}
	path = append(path, parseInt(string(board[y][x])))
	if x == len(board[0])-1 && y == len(board)-1 {
		return append(paths, path)
	}
	for _, check := range checks {
		id := k(check[0], check[1])
		if _, ok := visited[id]; !ok && check[0] < len(board[0]) && check[1] < len(board) && check[0] >= 0 && check[1] >= 0 {
			newVisited := map[string]struct{}{}
			for key, value := range visited {
				newVisited[key] = value
			}
			paths = append(paths, explore(check[0], check[1], path, board, newVisited)...)
		}
	}
	return paths
}

// Solution i/o
func Solution(pathName string) int {
	input := setup(pathName)
	visited := map[string]struct{}{}
	paths := explore(0, 0, []int{}, input, visited)
	shortestPath := 0
	for _, path := range paths {
		sum := 0
		for _, v := range path[1:] {
			sum += v
		}
		if shortestPath == 0 || sum < shortestPath {
			shortestPath = sum
		}
	}
	return shortestPath
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	print(input)
	return 0
}
