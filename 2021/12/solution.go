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

type Cave struct {
	name  string
	large bool
	edges map[string]bool
}

func newCave(name, edge string) Cave {
	return Cave{
		name:  name,
		large: name == strings.ToUpper(name),
		edges: map[string]bool{edge: true},
	}
}

func addCaves(nodes map[string]Cave, c1, c2 string) {
	if cave, ok := nodes[c1]; ok {
		cave.edges[c2] = true
	} else {
		nodes[c1] = newCave(c1, c2)
	}
}

func findPath(nodes map[string]Cave, visited map[string]struct{}, path, cave string) []string {
	paths := []string{}
	if cave == "end" {
		return append(paths, path)
	}
	currentCave := nodes[cave]
	if !currentCave.large {
		visited[cave] = struct{}{}
	}
	for edge := range currentCave.edges {
		_, ok := visited[edge]
		if edge != "start" && !ok {
			newVisited := make(map[string]struct{})
			for k, v := range visited {
				newVisited[k] = v
			}
			newPath := fmt.Sprintf("%s,%s", path, edge)
			paths = append(paths, findPath(nodes, newVisited, newPath, edge)...)
		}
	}
	return paths
}

func findPaths(nodes map[string]Cave) int {
	paths := []string{}
	for edge := range nodes["start"].edges {
		paths = append(paths, findPath(nodes, make(map[string]struct{}), "start", edge)...)
	}
	return len(paths)
}

// Solution i/o
func Solution(path string) int {
	input := setup(path)
	nodes := map[string]Cave{}
	for _, line := range input {
		caves := strings.Split(line, "-")
		addCaves(nodes, caves[0], caves[1])
		addCaves(nodes, caves[1], caves[0])
	}
	return findPaths(nodes)
}

func findPath2(nodes map[string]Cave, visited map[string]struct{}, buffer bool, path, cave string) []string {
	paths := []string{}
	if cave == "end" {
		return append(paths, path)
	}
	currentCave := nodes[cave]
	if !currentCave.large {
		visited[cave] = struct{}{}
	}
	for edge := range currentCave.edges {
		_, ok := visited[edge]
		singleSmall := (!ok || buffer)
		if edge != "start" && singleSmall {
			newVisited := make(map[string]struct{})
			for k, v := range visited {
				newVisited[k] = v
			}
			newPath := fmt.Sprintf("%s,%s", path, edge)
			paths = append(paths, findPath2(nodes, newVisited, buffer && !ok, newPath, edge)...)
		}
	}
	return paths
}

func findPaths2(nodes map[string]Cave) int {
	paths := []string{}
	for edge := range nodes["start"].edges {
		paths = append(paths, findPath2(nodes, make(map[string]struct{}), true, "start", edge)...)
	}
	return len(paths)
}

// Solution2 i/o
func Solution2(path string) int {
	input := setup(path)
	nodes := map[string]Cave{}
	for _, line := range input {
		caves := strings.Split(line, "-")
		addCaves(nodes, caves[0], caves[1])
		addCaves(nodes, caves[1], caves[0])
	}
	return findPaths2(nodes)
}
