// Package aoegeneric has generic function for AoC
package aoegeneric

import "math"

// AbsDiff returns the absolut difference between two integers
func AbsDiff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}

// GetDistance return the manhattan distance between to coordinates
func GetDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}
