// Package day04 is our package for the 4th AoC day
package day04

import (
	"strings"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	var grid [][]string
	sum := 0
	// fill grid
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	// find Xes
	var xmas strings.Builder
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "X" {
				// horizontal forward
				if x+3 < len(grid[y]) {
					for j := x; j <= x+3; j++ {
						xmas.WriteString(grid[y][j])
					}
					sum += strings.Count(xmas.String(), "XMAS")
					xmas.Reset()
				}
				// horizontal backward
				if x-3 >= 0 {
					for j := x; j >= x-3; j-- {
						xmas.WriteString(grid[y][j])
					}
					sum += strings.Count(xmas.String(), "XMAS")
					xmas.Reset()
				}
				// vertical forward
				if y+3 < len(grid) {
					for i := y; i < y+4; i++ {
						xmas.WriteString(grid[i][x])
					}
					sum += strings.Count(xmas.String(), "XMAS")
					xmas.Reset()
				}
				// vertical backward
				if y-3 >= 0 {
					for i := y; i >= y-3; i-- {
						xmas.WriteString(grid[i][x])
					}
					sum += strings.Count(xmas.String(), "XMAS")
					xmas.Reset()
				}
				// diagonal left->bottom
				if x-3 >= 0 && y+3 < len(grid) {
					for i := 0; i <= 3; i++ {
						xmas.WriteString(grid[y+i][x-i])
					}
					sum += strings.Count(xmas.String(), "XMAS")
					xmas.Reset()
				}
				// diagonal left->top
				if x-3 >= 0 && y-3 >= 0 {
					for i := 0; i <= 3; i++ {
						xmas.WriteString(grid[y-i][x-i])
					}
					sum += strings.Count(xmas.String(), "XMAS")
					xmas.Reset()
				}
				// diagonal right->bottom
				if x+3 < len(grid[y]) && y+3 < len(grid) {
					for i := 0; i <= 3; i++ {
						xmas.WriteString(grid[y+i][x+i])
					}
					sum += strings.Count(xmas.String(), "XMAS")
					xmas.Reset()
				}
				// diagonal right->top
				if x+3 < len(grid[y]) && y-3 >= 0 {
					for i := 0; i <= 3; i++ {
						xmas.WriteString(grid[y-i][x+i])
					}
					sum += strings.Count(xmas.String(), "XMAS")
					xmas.Reset()
				}
			}
		}
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	var grid [][]string
	sum := 0
	// fill grid
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	// find As
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] == "A" {
				if grid[y-1][x-1] == grid[y+1][x+1] {
					continue
				}
				var countM, countS int
				for i := y - 1; i <= y+1; i += 2 {
					for j := x - 1; j <= x+1; j += 2 {
						countS += strings.Count(grid[i][j], "S")
						countM += strings.Count(grid[i][j], "M")
					}
				}
				if countS == 2 && countM == 2 {
					sum++
				}
			}
		}
	}
	return sum
}
