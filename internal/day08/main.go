// Package day08 is our package for the 8th AoC day
package day08

import (
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

// Coordinate hold x y coordinates in a grid
type Coordinate struct {
	x int
	y int
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	antennas := make(map[string][]Coordinate, 1000)
	maxY := len(input)
	maxX := 0
	for y, row := range input {
		maxX = len(row)
		for x, cell := range strings.Split(row, "") {
			if cell != "." {
				antennas[cell] = append(antennas[cell], Coordinate{x: x, y: y})
			}
		}
	}
	for name, coord := range antennas {
		for _, comb := range combin.Combinations(len(coord), 2) {
			if name == "#" {
				continue
			}
			xDiff := coord[comb[0]].x - coord[comb[1]].x
			yDiff := coord[comb[0]].y - coord[comb[1]].y
			newX := coord[comb[0]].x + xDiff
			newY := coord[comb[0]].y + yDiff
			if newX >= 0 && newY >= 0 && newX < maxX && newY < maxY {
				antennas["#"] = append(antennas["#"], Coordinate{x: newX, y: newY})
			}
			xDiff = coord[comb[1]].x - coord[comb[0]].x
			yDiff = coord[comb[1]].y - coord[comb[0]].y
			newX = coord[comb[1]].x + xDiff
			newY = coord[comb[1]].y + yDiff
			if newX >= 0 && newY >= 0 && newX < maxX && newY < maxY {
				antennas["#"] = append(antennas["#"], Coordinate{x: newX, y: newY})
			}
		}
	}
	uniqCoords := make(map[Coordinate]struct{}, len(antennas["#"]))
	for _, coord := range antennas["#"] {
		uniqCoords[coord] = struct{}{}
	}
	return len(uniqCoords)
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	antennas := make(map[string][]Coordinate, 1000)
	maxY := len(input)
	maxX := 0
	for y, row := range input {
		maxX = len(row)
		for x, cell := range strings.Split(row, "") {
			if cell != "." {
				antennas[cell] = append(antennas[cell], Coordinate{x: x, y: y})
			}
		}
	}
	for name, coord := range antennas {
		for _, comb := range combin.Combinations(len(coord), 2) {
			if name == "#" {
				continue
			}
			xDiff := coord[comb[0]].x - coord[comb[1]].x
			yDiff := coord[comb[0]].y - coord[comb[1]].y
			newX := coord[comb[0]].x + xDiff
			newY := coord[comb[0]].y + yDiff
			for newX >= 0 && newY >= 0 && newX < maxX && newY < maxY {
				antennas["#"] = append(antennas["#"], Coordinate{x: newX, y: newY})
				newX = newX + xDiff
				newY = newY + yDiff
			}
			xDiff = coord[comb[1]].x - coord[comb[0]].x
			yDiff = coord[comb[1]].y - coord[comb[0]].y
			newX = coord[comb[1]].x + xDiff
			newY = coord[comb[1]].y + yDiff
			for newX >= 0 && newY >= 0 && newX < maxX && newY < maxY {
				antennas["#"] = append(antennas["#"], Coordinate{x: newX, y: newY})
				newX = newX + xDiff
				newY = newY + yDiff
			}
		}
	}
	uniqCoords := make(map[Coordinate]struct{}, len(antennas["#"]))
	for _, antenna := range antennas {
		for _, coord := range antenna {
			uniqCoords[coord] = struct{}{}
		}
	}
	return len(uniqCoords)
}
