package solution

import (
	"math"
	"strings"
)

type Point struct {
	x int
	y int
}

func newPoint(input string) Point {
	parts := strings.Split(input, ",")
	x := parseInt(parts[0])
	y := parseInt(parts[1])
	return Point{x: x, y: y}
}

type Rectangle struct {
	p1 Point
	p2 Point
}

func NewRectangle(p1, p2 Point) *Rectangle {
	return &Rectangle{p1: p1, p2: p2}
}

func (r *Rectangle) Area() int {
	return int(math.Abs(float64(r.p2.x-r.p1.x)+1) * math.Abs(float64(r.p2.y-r.p1.y)+1))
}

func findLargestRectangleArea(input []string) int {
	largestArea := 0
	for _, line := range input {
		p1 := newPoint(line)
		for _, line2 := range input {
			p2 := newPoint(line2)
			rectangle := NewRectangle(p1, p2)
			area := rectangle.Area()
			if area > largestArea {
				largestArea = area
			}
		}
	}
	return largestArea
}
