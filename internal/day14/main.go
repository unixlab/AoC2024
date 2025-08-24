// Package day14 is our package for the 14th AoC day
package day14

import "fmt"

type Robot struct {
	X  int
	Y  int
	vX int
	vY int
}

// RunPart1 is for the first star of the day
func RunPart1(input []string, wide int, tall int) int {
	var robots []Robot
	for _, line := range input {
		var x, y, vx, vy int
		_, parseError := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		if parseError != nil {
			panic(parseError)
		}
		robots = append(robots, Robot{X: x, Y: y, vX: vx, vY: vy})
	}
	for i := 1; i <= 100; i++ {
		for j := range robots {
			robots[j].X += robots[j].vX
			if robots[j].X < 0 {
				robots[j].X += wide
			}
			if robots[j].X >= wide {
				robots[j].X -= wide
			}
			robots[j].Y += robots[j].vY
			if robots[j].Y < 0 {
				robots[j].Y += tall
			}
			if robots[j].Y >= tall {
				robots[j].Y -= tall
			}
		}
	}
	var q1, q2, q3, q4 int
	for _, r := range robots {
		if r.X < wide/2 && r.Y < tall/2 {
			q1++
		} else if r.X > wide/2 && r.Y < tall/2 {
			q2++
		} else if r.X < wide/2 && r.Y > tall/2 {
			q3++
		} else if r.X > wide/2 && r.Y > tall/2 {
			q4++
		}
	}
	//fmt.Printf("Q1: %d, Q2: %d, Q3: %d, Q4: %d\n", q1, q2, q3, q4)
	//for y := 0; y < tall; y++ {
	//	for x := 0; x < wide; x++ {
	//		found := 0
	//		for _, r := range robots {
	//			if r.X == x && r.Y == y {
	//				found++
	//			}
	//		}
	//		if found > 0 {
	//			fmt.Print(found)
	//		} else {
	//			fmt.Print(".")
	//		}
	//	}
	//	fmt.Println()
	//}
	//fmt.Println(robots)
	return q1 * q2 * q3 * q4
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	for _, line := range input {
		fmt.Println(line)
	}
	return 0
}
