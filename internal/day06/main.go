// Package day06 is our package for the 6th AoC day
package day06

import (
	"strings"
)

type Coordinate struct {
	x int
	y int
}

// 0 - up
// 1 - left
// 2 - down
// 3 - right
type Guard struct {
	InitX int
	InitY int
	PosX  int
	PosY  int
	MoveX int
	MoveY int
	State int
}

func (g Guard) Init(y int, x int) Guard {
	var newGuard Guard
	newGuard.InitX = x
	newGuard.InitY = y
	newGuard.PosX = x
	newGuard.PosY = y
	newGuard.MoveX = 0
	newGuard.MoveY = -1
	newGuard.State = 0
	return newGuard
}

func (g Guard) Move() Guard {
	g.PosX += g.MoveX
	g.PosY += g.MoveY
	return g
}

func (g Guard) MoveBack() Guard {
	g.PosX += g.MoveX * -1
	g.PosY += g.MoveY * -1
	return g
}

func (g Guard) Rotate() Guard {
	switch g.State {
	case 0:
		g.MoveX = 1
		g.MoveY = 0
		g.State++
	case 1:
		g.MoveX = 0
		g.MoveY = 1
		g.State++
	case 2:
		g.MoveX = -1
		g.MoveY = 0
		g.State++
	case 3:
		g.MoveX = 0
		g.MoveY = -1
		g.State = 0
	}
	return g
}

func (g Guard) Reset() Guard {
	g.PosX = g.InitX
	g.PosY = g.InitY
	g.MoveX = 0
	g.MoveY = -1
	g.State = 0
	return g
}

func (g Guard) InGrid(grid [][]string) bool {
	if g.PosX < 0 {
		return false
	}
	if g.PosY < 0 {
		return false
	}
	if g.PosX >= len(grid) {
		return false
	}
	if g.PosY >= len(grid[g.PosX]) {
		return false
	}
	return true
}

// runPart is for the first star of the day
func runPart(input []string, earlyExit bool) (int, int) {
	var grid [][]string
	var guard Guard
	visitedP1 := make(map[Coordinate]struct{}, 2000)
	for y, line := range input {
		var row []string
		for x, field := range strings.Split(line, "") {
			if field == "^" {
				guard = guard.Init(y, x)
				field = "."
			}
			row = append(row, field)
		}
		grid = append(grid, row)
	}
	for {
		guard = guard.Move()
		if guard.InGrid(grid) {
			if grid[guard.PosY][guard.PosX] == "#" {
				guard = guard.MoveBack()
				guard = guard.Rotate()
			} else {
				visitedP1[Coordinate{guard.PosX, guard.PosY}] = struct{}{}
			}
		} else {
			break
		}
	}
	if earlyExit {
		return len(visitedP1), 0
	}
	guard = guard.Reset()
	loop := 0
	for cord := range visitedP1 {
		y := cord.y
		x := cord.x
		if grid[y][x] == "#" {
			continue
		}
		grid[y][x] = "#"
		visitedP2 := make(map[Coordinate]int, 2000)
		for {
			guard = guard.Move()
			if guard.InGrid(grid) {
				if grid[guard.PosY][guard.PosX] == "#" {
					guard = guard.MoveBack()
					guard = guard.Rotate()
				} else {
					if visitedP2[Coordinate{guard.PosX, guard.PosY}] > 3 {
						loop++
						guard = guard.Reset()
						break
					} else {
						visitedP2[Coordinate{guard.PosX, guard.PosY}]++
					}
				}
			} else {
				guard = guard.Reset()
				break
			}
		}
		grid[y][x] = "."
	}
	return len(visitedP1), loop
}

// RunPart1 is for the second star of the day
func RunPart1(input []string) int {
	result, _ := runPart(input, true)
	return result
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	_, result := runPart(input, false)
	return result
}
