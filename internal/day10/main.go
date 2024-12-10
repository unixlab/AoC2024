// Package day10 is our package for the 10th AoC day
package day10

import (
	"strconv"
	"strings"

	"github.com/unixlab/AoC2024/internal/aoegeneric"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/dominikbraun/graph"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	var grid [][]int
	idx := 0
	dijkstraGraph := dijkstra.NewGraph()
	var starts []int
	var ends []int
	for _, line := range input {
		var row []int
		for _, c := range strings.Split(line, "") {
			n, _ := strconv.Atoi(c)
			if n == 0 {
				starts = append(starts, idx)
			}
			if n == 9 {
				ends = append(ends, idx)
			}
			row = append(row, n)
			_ = dijkstraGraph.AddEmptyVertex(idx)
			idx++
		}
		grid = append(grid, row)
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			for i := y - 1; i <= y+1; i++ {
				for j := x - 1; j <= x+1; j++ {
					if i < 0 || i >= len(grid) {
						continue
					}
					if j < 0 || j >= len(grid[i]) {
						continue
					}
					if y == i && x == j {
						continue
					}
					if grid[y][x]+1 == grid[i][j] {
						sourceArc := len(grid[y])*y + x
						destinationArc := len(grid[i])*i + j
						if aoegeneric.GetDistance(x, y, j, i) == 1 {
							_ = dijkstraGraph.AddArc(sourceArc, destinationArc, 1)
						}
					}
				}
			}
		}
	}
	sum := 0
	for _, start := range starts {
		for _, end := range ends {
			_, err := dijkstraGraph.Longest(start, end)
			if err == nil {
				sum++
			}
		}
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	var grid [][]int
	idx := 0
	dfsGraph := graph.New(graph.IntHash, graph.Directed())
	var starts []int
	var ends []int
	for _, line := range input {
		var row []int
		for _, c := range strings.Split(line, "") {
			n, _ := strconv.Atoi(c)
			if n == 0 {
				starts = append(starts, idx)
			}
			if n == 9 {
				ends = append(ends, idx)
			}
			row = append(row, n)
			_ = dfsGraph.AddVertex(idx)
			idx++
		}
		grid = append(grid, row)
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			for i := y - 1; i <= y+1; i++ {
				for j := x - 1; j <= x+1; j++ {
					if i < 0 || i >= len(grid) {
						continue
					}
					if j < 0 || j >= len(grid[i]) {
						continue
					}
					if y == i && x == j {
						continue
					}
					if grid[y][x]+1 == grid[i][j] {
						sourceArc := len(grid[y])*y + x
						destinationArc := len(grid[i])*i + j
						if aoegeneric.GetDistance(x, y, j, i) == 1 {
							_ = dfsGraph.AddEdge(sourceArc, destinationArc)
						}
					}
				}
			}
		}
	}
	sum := 0
	for _, start := range starts {
		for _, end := range ends {
			paths, _ := graph.AllPathsBetween(dfsGraph, start, end)
			sum += len(paths)
		}
	}
	return sum
}
