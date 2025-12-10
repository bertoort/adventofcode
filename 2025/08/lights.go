package solution

import (
	"sort"
	"strings"
)

type LightBox struct {
	id string
	x  int
	y  int
	z  int
}

// CircuitBoard is a list of circuits
type CircuitBoard map[string][]string

// Connection (edge) is two circuits and the distance between them
type Connection struct {
	a        LightBox
	b        LightBox
	distance int
}

func NewLightBox(input string) LightBox {
	parts := strings.Split(input, ",")
	x := parseInt(parts[0])
	y := parseInt(parts[1])
	z := parseInt(parts[2])
	return LightBox{id: input, x: x, y: y, z: z}
}

func parseLights(input []string) []LightBox {
	lights := make([]LightBox, len(input))
	for i, light := range input {
		lights[i] = NewLightBox(light)
	}
	return lights
}

func connectLights(input []string, circuitCount int) CircuitBoard {
	lights := parseLights(input)
	// Initialize union-find
	ids := make([]string, len(lights))
	for i, light := range lights {
		ids[i] = light.id
	}
	uf := NewUnionFind(ids)
	// Get all connections sorted by distance (ascending)
	connections := findConnections(lights)

	for _, conn := range connections[:circuitCount] {
		uf.Union(conn.a.id, conn.b.id)
	}

	// Group all lights by their root
	circuitBoard := make(CircuitBoard)
	for _, light := range lights {
		root := uf.Find(light.id)
		circuitBoard[root] = append(circuitBoard[root], light.id)
	}

	return circuitBoard
}

func multiplyTopCircuits(circuitBoard CircuitBoard, topCount int) int {
	total := 1
	// Extract keys for sorting
	circuitLengths := make([]int, 0, len(circuitBoard))
	for _, v := range circuitBoard {
		circuitLengths = append(circuitLengths, len(v))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(circuitLengths)))
	// Get the top lengths for multiplication
	topLengths := circuitLengths[:topCount]
	for _, length := range topLengths {
		total *= length
	}
	return total
}

// Solution 2

func connectAllLights(input []string) Connection {
	lights := parseLights(input)
	// Initialize union-find
	ids := make([]string, len(lights))
	for i, light := range lights {
		ids[i] = light.id
	}
	uf := NewUnionFind(ids)
	// Get all connections sorted by distance (ascending)
	connections := findConnections(lights)
	lastConnection := Connection{}
	for _, conn := range connections {
		merged := uf.Union(conn.a.id, conn.b.id)
		if merged {
			lastConnection = conn
		}
	}

	return lastConnection
}

func multiplyLastConnection(lastConnection Connection) int {
	c1 := lastConnection.a
	c2 := lastConnection.b
	return c1.x * c2.x
}
