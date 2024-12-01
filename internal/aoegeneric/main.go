// Package aoegeneric has generic function for AoC
package aoegeneric

import "math"

// AbsDiff returns the absolut difference between two integers
func AbsDiff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}
