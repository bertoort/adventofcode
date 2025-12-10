package solution

// This file explains different approaches to findConnections

import (
	"container/heap"
	"math"
	"sort"
)

// ============================================================================
// BASIC IMPLEMENTATION: Find All Connections, Then Sort
// ============================================================================

// Euclidean distance between two 3D points (straight line distance)
// Formula: sqrt((x1-x2)² + (y1-y2)² + (z1-z2)²)
func euclideanDistance(a, b LightBox) int {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	distance := math.Sqrt(dx*dx + dy*dy + dz*dz)
	return int(distance) // Convert to int for storage
}

// Alternative: Store squared distance for faster comparison (avoid sqrt)
// Only use if you don't need the actual distance value
func euclideanDistanceSquared(a, b LightBox) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

/*
// findConnectionsBasic: O(n²) to generate, O(n² log n²) to sort
// Time: O(n² log n) where n = number of lightboxes
// Space: O(n²) to store all connections
func findConnectionsBasic(lights []LightBox) []Connection {
	n := len(lights)
	// Total pairs: n * (n-1) / 2
	connections := make([]Connection, 0, n*(n-1)/2)

	// Generate all pairs (avoid duplicates: i < j)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := euclideanDistance(lights[i], lights[j])
			connections = append(connections, Connection{
				a:        lights[i],
				b:        lights[j],
				distance: dist,
			})
		}
	}

	// Sort by distance (ascending: shortest first)
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})

	return connections
}
*/
// ============================================================================
// OPTIMIZATION 1: Use Heap to Find Only k Shortest (Without Full Sort)
// ============================================================================

// MinHeap for connections (smallest distance at top)
type ConnectionHeap []Connection

func (h ConnectionHeap) Len() int           { return len(h) }
func (h ConnectionHeap) Less(i, j int) bool { return h[i].distance > h[j].distance } // Max heap (inverted)
func (h ConnectionHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ConnectionHeap) Push(x interface{}) {
	*h = append(*h, x.(Connection))
}

func (h *ConnectionHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
// findConnectionsHeap: O(n² log k) where k = number of connections needed
// Better when k << n² (you only need a small subset)
// Space: O(k) instead of O(n²)
func findConnectionsHeap(lights []LightBox, k int) []Connection {
	n := len(lights)
	h := &ConnectionHeap{}
	heap.Init(h)

	// Generate all pairs, but only keep k smallest
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := euclideanDistance(lights[i], lights[j])
			conn := Connection{
				a:        lights[i],
				b:        lights[j],
				distance: dist,
			}

			if h.Len() < k {
				// Heap not full yet, add it
				heap.Push(h, conn)
			} else if dist < (*h)[0].distance {
				// New connection is smaller than largest in heap
				heap.Pop(h)        // Remove largest
				heap.Push(h, conn) // Add new one
			}
			// Otherwise, ignore (too large)
		}
	}

	// Extract and sort the k connections
	result := make([]Connection, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Connection)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].distance < result[j].distance
	})

	return result
}
*/
// ============================================================================
// OPTIMIZATION 2: Quickselect for Partial Sorting
// ============================================================================
/*
// findConnectionsQuickselect: O(n²) average case to find k smallest
// Worst case: O(n²) but can degrade to O(n² log n) if unlucky
// Space: O(n²) but doesn't need to sort everything
func findConnectionsQuickselect(lights []LightBox, k int) []Connection {
	n := len(lights)
	connections := make([]Connection, 0, n*(n-1)/2)

	// Generate all pairs
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := euclideanDistance(lights[i], lights[j])
			connections = append(connections, Connection{
				a:        lights[i],
				b:        lights[j],
				distance: dist,
			})
		}
	}

	// Quickselect: partition to find k smallest
	// This is more complex to implement correctly
	// For simplicity, just sort the first k after partitioning
	// (Full quickselect implementation would be more efficient)

	return connections[:k] // Simplified - you'd need actual quickselect
}
*/
// ============================================================================
// COMPARISON OF APPROACHES
// ============================================================================

/*
For n lightboxes, you have n*(n-1)/2 ≈ n²/2 pairs

Scenario 1: You need ALL connections sorted
  - findConnectionsBasic: O(n² log n) time, O(n²) space
  - Best choice: You need everything anyway

Scenario 2: You only need k shortest connections (k << n²)
  - findConnectionsBasic: O(n² log n) time, O(n²) space
  - findConnectionsHeap: O(n² log k) time, O(k) space
  - Best choice: Heap approach saves time and space

Example with 1000 lightboxes:
  - Total pairs: ~500,000
  - If you need 1000 connections:
    - Basic: Generate 500k, sort 500k = O(500k * log(500k)) ≈ 9.5M operations
    - Heap: Generate 500k, maintain heap of 1k = O(500k * log(1k)) ≈ 5M operations
    - Space: 500k vs 1k connections stored

Scenario 3: You need exactly n-1 connections (MST)
  - For MST, you typically need to check many connections before finding valid ones
  - Heap might not help much if you need to process many rejected connections
  - Basic approach is often fine

Scenario 4: Very large n (10,000+ lightboxes)
  - n² = 100M pairs - this is getting expensive
  - Consider spatial data structures (k-d trees, etc.) for geometric problems
  - But for Manhattan distance in 3D, brute force is often still practical
*/

// ============================================================================
// RECOMMENDATION FOR YOUR USE CASE
// ============================================================================

/*
Your algorithm:
1. Find connections (sorted by distance)
2. Process connections one by one until you have circuitCount connections
3. Skip connections that would create cycles

Analysis:
- circuitCount < total lightboxes (you confirmed this)
- You don't know in advance how many connections you'll need to process
- You might process many connections before finding valid ones (cycle detection)
- You need connections in sorted order

Recommendation:
Since circuitCount is smaller, you have two good options:

Option 1: Heap-based (RECOMMENDED for large n)
- Use findConnectionsHeap with k = circuitCount * buffer (e.g., 2-5x)
- Why? If circuitCount << n², you save both time and space
- Example: 1000 lightboxes, circuitCount = 100
  - Total pairs: ~500,000
  - Heap with k=500: O(500k * log(500)) ≈ 4.5M operations
  - Full sort: O(500k * log(500k)) ≈ 9.5M operations
  - Space: 500 vs 500k connections

Option 2: Basic approach (simpler, still good)
- Use findConnectionsBasic (generate all, sort all)
- Why? Simpler code, easier to maintain
- The sorting overhead is acceptable for moderate n
- Better if you're not sure how many connections you'll need to process

Choose Option 1 if:
- You have many lightboxes (n > 500)
- circuitCount is much smaller than n
- You want to optimize for both time and space

Choose Option 2 if:
- You have moderate number of lightboxes (n < 500)
- You prefer simpler code
- You're not sure about the buffer size needed
*/

// ============================================================================
// CAN YOU AVOID COMPUTING ALL DISTANCES?
// ============================================================================

/*
Short answer: Not really, for this problem.

Why you need all pairs:
1. You're building circuits by connecting lightboxes
2. You need to know which connections are shortest
3. There's no way to know which pairs are shortest without comparing distances
4. Even with spatial data structures, you'd still need to consider many pairs

When you COULD avoid:
- If you only needed connections to nearest neighbors (k-nearest neighbor)
- If you had constraints (e.g., only connect within certain distance)
- If the problem had geometric structure you could exploit

For your problem:
- You need to consider all pairs to find the globally shortest connections
- O(n²) is the theoretical lower bound for this problem
- The good news: distance calculation is O(1), so O(n²) is just the pair generation
*/

// ============================================================================
// PRACTICAL IMPLEMENTATION
// ============================================================================

// ============================================================================
// RECOMMENDED IMPLEMENTATIONS FOR YOUR CODE
// ============================================================================

// Option 1: Basic approach (simpler, good for moderate n)
func findConnectionsBasic(lights []LightBox) []Connection {
	n := len(lights)
	connections := make([]Connection, 0, n*(n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := float64(lights[i].x - lights[j].x)
			dy := float64(lights[i].y - lights[j].y)
			dz := float64(lights[i].z - lights[j].z)
			distance := int(math.Sqrt(dx*dx + dy*dy + dz*dz))

			connections = append(connections, Connection{
				a:        lights[i],
				b:        lights[j],
				distance: distance,
			})
		}
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})

	return connections
}

// Option 2: Heap-based (optimized for circuitCount << n²)
// Use this when circuitCount is much smaller than total pairs
func findConnectionsHeap(lights []LightBox, maxConnections int) []Connection {
	n := len(lights)
	h := &ConnectionHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := float64(lights[i].x - lights[j].x)
			dy := float64(lights[i].y - lights[j].y)
			dz := float64(lights[i].z - lights[j].z)
			distance := int(math.Sqrt(dx*dx + dy*dy + dz*dz))

			conn := Connection{
				a:        lights[i],
				b:        lights[j],
				distance: distance,
			}

			if h.Len() < maxConnections {
				heap.Push(h, conn)
			} else if distance < (*h)[0].distance {
				heap.Pop(h)
				heap.Push(h, conn)
			}
		}
	}

	// Extract and sort
	result := make([]Connection, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Connection)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].distance < result[j].distance
	})

	return result
}

// Recommended: Use heap if circuitCount is known and small
// Otherwise use basic approach
func findConnections(lights []LightBox) []Connection {
	// For now, use basic approach
	// Switch to heap version if you have many lightboxes
	return findConnectionsBasic(lights)
}
