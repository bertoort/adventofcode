package navigation

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Coordinates for ship
type Coordinates struct {
	Lat     int
	Long    int
	Degrees int
	Waypoint
}

// Waypoint location
type Waypoint struct {
	Lat  int
	Long int
}

// Sum total coordinate change
func (c Coordinates) Sum() int {
	return int(math.Abs(float64(c.Lat))) + int(math.Abs(float64(c.Long)))
}

// Rotate changes direction the ship is facing
func (c *Coordinates) Rotate(degrees int, left bool) {
	if left {
		degrees *= -1
	}
	c.Degrees += degrees
	if c.Degrees < 0 {
		c.Degrees = 360 + c.Degrees
	}
	c.Degrees = c.Degrees % 360
}

// Rotate changes direction of waypoint
func (w *Waypoint) Rotate(degrees int, left bool) {
	switch degrees {
	case 90, 270:
		temp := w.Lat
		w.Lat = w.Long
		w.Long = temp
		w.Long *= -1
		if degrees == 270 {
			left = !left
		}
		if left {
			w.Lat *= -1
			w.Long *= -1
		}
	case 180:
		w.Lat *= -1
		w.Long *= -1
	}

}

// Forward move the ship in the facing direction
func (c *Coordinates) Forward(value int) {
	switch c.Degrees {
	case 0:
		// east
		c.Lat += value
	case 90:
		// south:
		c.Long -= value
	case 180:
		// west
		c.Lat -= value
	case 270:
		// north
		c.Long += value
	}
}

// WaypointForward move the ship in the facing direction according to waypoint
func (c *Coordinates) WaypointForward(value int) {
	c.Lat += c.Waypoint.Lat * value
	c.Long += c.Waypoint.Long * value
}

func setup(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return strings.Split(string(b), "\n")
}

// PartOneSolution io
func PartOneSolution(path string) int {
	instructions := setup(path)
	coordinates := Coordinates{}
	for _, instruction := range instructions {
		letter := string(instruction[0])
		value, _ := strconv.Atoi(instruction[1:])

		switch letter {
		case "N":
			// move north
			coordinates.Long += value
		case "S":
			// move south
			coordinates.Long -= value
		case "E":
			// move east
			coordinates.Lat += value
		case "W":
			// move west
			coordinates.Lat -= value
		case "L", "R":
			// rotate left
			coordinates.Rotate(value, letter == "L")
		case "F":
			// move forward
			coordinates.Forward(value)
		}
	}
	return coordinates.Sum()
}

// PartTwoSolution io
func PartTwoSolution(path string) int {
	instructions := setup(path)
	coordinates := Coordinates{Waypoint: Waypoint{Lat: 10, Long: 1}}
	for _, instruction := range instructions {
		letter := string(instruction[0])
		value, _ := strconv.Atoi(instruction[1:])
		fmt.Println(coordinates)
		switch letter {
		case "N":
			// move north
			coordinates.Waypoint.Long += value
		case "S":
			// move south
			coordinates.Waypoint.Long -= value
		case "E":
			// move east
			coordinates.Waypoint.Lat += value
		case "W":
			// move west
			coordinates.Waypoint.Lat -= value
		case "L", "R":
			// rotate left
			coordinates.Waypoint.Rotate(value, letter == "L")
		case "F":
			// move forward
			coordinates.WaypointForward(value)
		}
	}
	return coordinates.Sum()
}
