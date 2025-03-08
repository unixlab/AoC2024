// Package aocgeneric has generic function for AoC
package aocgeneric

import (
	"errors"
	"math"
)

// Coordinate hold x y coordinates
type Coordinate struct {
	X int
	Y int
}

type CoordinateQueue struct {
	Coordinates []Coordinate
}

func (q *CoordinateQueue) Pop() (Coordinate, error) {
	if len(q.Coordinates) == 0 {
		return Coordinate{}, errors.New("queue empty")
	}
	c := q.Coordinates[0]
	q.Coordinates = q.Coordinates[1:]
	return c, nil
}

func (q *CoordinateQueue) Push(c Coordinate) {
	q.Coordinates = append(q.Coordinates, c)
}

// AbsDiff returns the absolut difference between two integers
func AbsDiff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}

// GetDistance return the manhattan distance between to coordinates
func GetDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

// Pow calculates the value of a raised to the power of b as integers
func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
