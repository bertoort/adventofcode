package solution

// This file is for explanation only - showing how union-find works

// ============================================================================
// UNION-FIND (DISJOINT SET) DATA STRUCTURE
// ============================================================================

// UnionFind tracks which elements are in the same group
type UnionFind struct {
	parent map[string]string // id -> parent id
	rank   map[string]int    // id -> rank (for optimization)
}

func NewUnionFind(ids []string) *UnionFind {
	uf := &UnionFind{
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}
	// Initialize: each element is its own parent
	for _, id := range ids {
		uf.parent[id] = id
		uf.rank[id] = 0
	}
	return uf
}

// Find returns the root (representative) of the group containing id
// Uses path compression: flattens the tree as it traverses
func (uf *UnionFind) Find(id string) string {
	if uf.parent[id] != id {
		// Path compression: make parent point directly to root
		uf.parent[id] = uf.Find(uf.parent[id])
	}
	return uf.parent[id]
}

// Union merges the groups containing id1 and id2
// Returns true if they were in different groups (merge happened)
// Returns false if they were already in the same group (would create cycle)
func (uf *UnionFind) Union(id1, id2 string) bool {
	root1 := uf.Find(id1)
	root2 := uf.Find(id2)

	// Already in same group - would create cycle!
	if root1 == root2 {
		return false
	}

	// Union by rank: attach smaller tree under larger tree
	if uf.rank[root1] < uf.rank[root2] {
		uf.parent[root1] = root2
	} else if uf.rank[root1] > uf.rank[root2] {
		uf.parent[root2] = root1
	} else {
		// Same rank: pick one as root and increment its rank
		uf.parent[root2] = root1
		uf.rank[root1]++
	}
	return true
}

// ============================================================================
// HOW IT WORKS WITH YOUR PROBLEM
// ============================================================================

/*
Example with 4 lights: A, B, C, D

Initial state:
  parent: {A: A, B: B, C: C, D: D}
  rank:   {A: 0, B: 0, C: 0, D: 0}

Connections sorted by distance:
  1. A-B (distance 5)
  2. C-D (distance 8)
  3. A-C (distance 10)
  4. B-D (distance 15)

Step 1: Connect A-B
  Find(A) = A, Find(B) = B (different groups)
  Union(A, B) -> merge them
  parent: {A: A, B: A, C: C, D: D}
  Circuit created: A and B are connected

Step 2: Connect C-D
  Find(C) = C, Find(D) = D (different groups)
  Union(C, D) -> merge them
  parent: {A: A, B: A, C: C, D: C}
  Circuit created: C and D are connected

Step 3: Connect A-C
  Find(A) = A, Find(C) = C (different groups)
  Union(A, C) -> merge them
  parent: {A: A, B: A, C: A, D: C}
  Circuits merged: Now A, B, C, D are all in one circuit

Step 4: Connect B-D
  Find(B) = A, Find(D) = A (same group!)
  Union(B, D) returns false
  Skip this connection (would create a cycle)

Time Complexity:
  - Find: O(α(n)) ≈ O(1) amortized (inverse Ackermann function)
  - Union: O(α(n)) ≈ O(1) amortized
  - Total for k connections: O(k) ≈ O(n²) if k = n²
*/

// ============================================================================
// ALTERNATIVE: YOUR STRING-BASED APPROACH
// ============================================================================

/*
Your approach uses string union like "id1|id2|id3"

Pros:
  - Simple to understand
  - Easy to see which lights are in a circuit

Cons:
  - Checking if two lights are in same circuit: O(n) per check
    Need to search through all circuit strings
  - Merging circuits: O(n) to concatenate strings
  - Finding which circuit a light belongs to: O(n) to search all circuits

Time complexity with your approach:
  - For each connection: O(n) to check if both in same circuit
  - For k connections: O(k * n) = O(n³) worst case

Time complexity with union-find:
  - For each connection: O(1) amortized to check and merge
  - For k connections: O(k) = O(n²) worst case

For 1000 lights:
  - Your approach: ~1,000,000,000 operations
  - Union-find: ~1,000,000 operations
*/

// ============================================================================
// HOW TO USE UNION-FIND IN YOUR CODE
// ============================================================================

/*
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

	// Build circuits using Kruskal's algorithm
	circuitBoard := make(CircuitBoard)
	connectionsAdded := 0

	for _, conn := range connections {
		if connectionsAdded >= circuitCount {
			break
		}

		// Try to merge: returns false if would create cycle
		if uf.Union(conn.a.id, conn.b.id) {
			// Successfully merged - add to circuit board
			// You'd need to rebuild the circuit strings from union-find
			connectionsAdded++
		}
	}

	// Convert union-find groups back to your string format
	// Group all lights by their root
	groups := make(map[string][]string)
	for _, light := range lights {
		root := uf.Find(light.id)
		groups[root] = append(groups[root], light.id)
	}

	// Build circuit strings
	for _, group := range groups {
		if len(group) > 1 { // Only circuits with 2+ lights
			circuitStr := strings.Join(group, "|")
			circuitBoard[circuitStr] = true
		}
	}

	return circuitBoard
}
*/
