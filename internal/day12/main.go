// Package day12 is our package for the 12th AoC day
package day12

import (
	"fmt"
	"github.com/unixlab/AoC2024/internal/aocgeneric"
	"strings"
)

// RunPart is for both parts
func RunPart(input []string) (int, int) {
	var grid [][]string
	part1Price := 0
	part2Price := 0
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			lowerPlant := strings.ToLower(grid[y][x])
			if grid[y][x] == lowerPlant {
				continue
			}
			plant := grid[y][x]
			plantCounter := 0
			borders := 0
			var borderCoords []aocgeneric.Coordinate
			var queue aocgeneric.CoordinateQueue
			queue.Push(aocgeneric.Coordinate{X: x, Y: y})
			for c, e := queue.Pop(); e == nil; c, e = queue.Pop() {
				if grid[c.Y][c.X] == lowerPlant {
					continue
				}
				grid[c.Y][c.X] = lowerPlant
				plantCounter++
				if c.Y+1 < len(grid) && grid[c.Y+1][c.X] == plant {
					queue.Push(aocgeneric.Coordinate{X: c.X, Y: c.Y + 1})
				} else if c.Y+1 < len(grid) && grid[c.Y+1][c.X] == lowerPlant {
					// noop
				} else {
					borders++
					borderCoords = append(borderCoords, aocgeneric.Coordinate{X: c.X, Y: c.Y + 1})
				}
				if c.Y-1 >= 0 && grid[c.Y-1][c.X] == plant {
					queue.Push(aocgeneric.Coordinate{X: c.X, Y: c.Y - 1})
				} else if c.Y-1 >= 0 && grid[c.Y-1][c.X] == lowerPlant {
					// noop
				} else {
					borders++
					borderCoords = append(borderCoords, aocgeneric.Coordinate{X: c.X, Y: c.Y - 1})
				}
				if c.X+1 < len(grid) && grid[c.Y][c.X+1] == plant {
					queue.Push(aocgeneric.Coordinate{X: c.X + 1, Y: c.Y})
				} else if c.X+1 < len(grid) && grid[c.Y][c.X+1] == lowerPlant {
					// noop
				} else {
					borders++
					borderCoords = append(borderCoords, aocgeneric.Coordinate{X: c.X + 1, Y: c.Y})
				}
				if c.X-1 >= 0 && grid[c.Y][c.X-1] == plant {
					queue.Push(aocgeneric.Coordinate{X: c.X - 1, Y: c.Y})
				} else if c.X-1 >= 0 && grid[c.Y][c.X-1] == lowerPlant {
					// noop
				} else {
					borders++
					borderCoords = append(borderCoords, aocgeneric.Coordinate{X: c.X - 1, Y: c.Y})
				}
			}
			part1Price += plantCounter * borders
			oldX := -1
			oldY := -1
			edgeCounter := 0
			for _, c := range borderCoords {
				if c.X != oldX && c.Y != oldY {
					edgeCounter++
				}
				oldX = c.X
				oldY = c.Y
			}
			fmt.Println(edgeCounter)
		}
	}
	return part1Price, part2Price
}

// RunPart1 is for the second star of the day
func RunPart1(input []string) int {
	result, _ := RunPart(input)
	return result
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	_, result := RunPart(input)
	return result
}
