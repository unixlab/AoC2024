// Package day08 is our package for the 8th AoC day
package day08

import (
	"strings"

	"github.com/unixlab/AoC2024/internal/aocgeneric"

	"gonum.org/v1/gonum/stat/combin"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	antennas := make(map[string][]aocgeneric.Coordinate, 1000)
	maxY := len(input)
	maxX := 0
	for y, row := range input {
		maxX = len(row)
		for x, cell := range strings.Split(row, "") {
			if cell != "." {
				antennas[cell] = append(antennas[cell], aocgeneric.Coordinate{X: x, Y: y})
			}
		}
	}
	for name, coord := range antennas {
		for _, comb := range combin.Combinations(len(coord), 2) {
			if name == "#" {
				continue
			}
			xDiff := coord[comb[0]].X - coord[comb[1]].X
			yDiff := coord[comb[0]].Y - coord[comb[1]].Y
			newX := coord[comb[0]].X + xDiff
			newY := coord[comb[0]].Y + yDiff
			if newX >= 0 && newY >= 0 && newX < maxX && newY < maxY {
				antennas["#"] = append(antennas["#"], aocgeneric.Coordinate{X: newX, Y: newY})
			}
			xDiff = coord[comb[1]].X - coord[comb[0]].X
			yDiff = coord[comb[1]].Y - coord[comb[0]].Y
			newX = coord[comb[1]].X + xDiff
			newY = coord[comb[1]].Y + yDiff
			if newX >= 0 && newY >= 0 && newX < maxX && newY < maxY {
				antennas["#"] = append(antennas["#"], aocgeneric.Coordinate{X: newX, Y: newY})
			}
		}
	}
	uniqCoords := make(map[aocgeneric.Coordinate]struct{}, len(antennas["#"]))
	for _, coord := range antennas["#"] {
		uniqCoords[coord] = struct{}{}
	}
	return len(uniqCoords)
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	antennas := make(map[string][]aocgeneric.Coordinate, 1000)
	maxY := len(input)
	maxX := 0
	for y, row := range input {
		maxX = len(row)
		for x, cell := range strings.Split(row, "") {
			if cell != "." {
				antennas[cell] = append(antennas[cell], aocgeneric.Coordinate{X: x, Y: y})
			}
		}
	}
	for name, coord := range antennas {
		for _, comb := range combin.Combinations(len(coord), 2) {
			if name == "#" {
				continue
			}
			xDiff := coord[comb[0]].X - coord[comb[1]].X
			yDiff := coord[comb[0]].Y - coord[comb[1]].Y
			newX := coord[comb[0]].X + xDiff
			newY := coord[comb[0]].Y + yDiff
			for newX >= 0 && newY >= 0 && newX < maxX && newY < maxY {
				antennas["#"] = append(antennas["#"], aocgeneric.Coordinate{X: newX, Y: newY})
				newX = newX + xDiff
				newY = newY + yDiff
			}
			xDiff = coord[comb[1]].X - coord[comb[0]].X
			yDiff = coord[comb[1]].Y - coord[comb[0]].Y
			newX = coord[comb[1]].X + xDiff
			newY = coord[comb[1]].Y + yDiff
			for newX >= 0 && newY >= 0 && newX < maxX && newY < maxY {
				antennas["#"] = append(antennas["#"], aocgeneric.Coordinate{X: newX, Y: newY})
				newX = newX + xDiff
				newY = newY + yDiff
			}
		}
	}
	uniqCoords := make(map[aocgeneric.Coordinate]struct{}, len(antennas["#"]))
	for _, antenna := range antennas {
		for _, coord := range antenna {
			uniqCoords[coord] = struct{}{}
		}
	}
	return len(uniqCoords)
}
