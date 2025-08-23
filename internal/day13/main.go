// Package day13 is our package for the 13th AoC day
package day13

import (
	"regexp"
	"strconv"
	"strings"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	part1Sum := 0
	var aX, aY, bX, bY, pX, pY float64
	for _, line := range input {
		if strings.HasPrefix(line, "Button A") {
			re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
			matches := re.FindStringSubmatch(line)
			aX, _ = strconv.ParseFloat(matches[1], 64)
			aY, _ = strconv.ParseFloat(matches[2], 64)
		}
		if strings.HasPrefix(line, "Button B") {
			re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
			matches := re.FindStringSubmatch(line)
			bX, _ = strconv.ParseFloat(matches[1], 64)
			bY, _ = strconv.ParseFloat(matches[2], 64)
		}
		if strings.HasPrefix(line, "Prize") {
			re := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
			matches := re.FindStringSubmatch(line)
			pX, _ = strconv.ParseFloat(matches[1], 64)
			pY, _ = strconv.ParseFloat(matches[2], 64)
		}
		if pX != 0 {
			// after some googling I found we can use Cramer's rule
			// https://en.wikipedia.org/wiki/Cramer%27s_rule#2x2_System
			d := aX*bY - aY*bX
			if d != 0 {
				a := (pX*bY - pY*bX) / d
				b := (aX*pY - aY*pX) / d
				aX, aY, bX, bY, pX, pY = 0, 0, 0, 0, 0, 0
				// boundary check a between 0 and 100
				if a <= 0 || a > 100 {
					continue
				}
				// boundary check b between 0 and 100
				if b <= 0 || b > 100 {
					continue
				}
				// check if a is not a whole number
				if a != float64(int(a)) {
					continue
				}
				// check if b is not a whole number
				if b != float64(int(b)) {
					continue
				}
				part1Sum += int(3*a + b)
			}
			aX, aY, bX, bY, pX, pY = 0, 0, 0, 0, 0, 0
		}
	}
	return part1Sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	part2Sum := 0
	var aX, aY, bX, bY, pX, pY float64
	for _, line := range input {
		if strings.HasPrefix(line, "Button A") {
			re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
			matches := re.FindStringSubmatch(line)
			aX, _ = strconv.ParseFloat(matches[1], 64)
			aY, _ = strconv.ParseFloat(matches[2], 64)
		}
		if strings.HasPrefix(line, "Button B") {
			re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
			matches := re.FindStringSubmatch(line)
			bX, _ = strconv.ParseFloat(matches[1], 64)
			bY, _ = strconv.ParseFloat(matches[2], 64)
		}
		if strings.HasPrefix(line, "Prize") {
			re := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
			matches := re.FindStringSubmatch(line)
			pX, _ = strconv.ParseFloat(matches[1], 64)
			pX += 10000000000000
			pY, _ = strconv.ParseFloat(matches[2], 64)
			pY += 10000000000000
		}
		if pX != 0 {
			// after some googling I found we can use Cramer's rule
			// https://en.wikipedia.org/wiki/Cramer%27s_rule#2x2_System
			d := aX*bY - aY*bX
			if d != 0 {
				a := (pX*bY - pY*bX) / d
				b := (aX*pY - aY*pX) / d
				aX, aY, bX, bY, pX, pY = 0, 0, 0, 0, 0, 0
				// boundary check if a greater 0
				if a <= 0 {
					continue
				}
				// boundary check if b greater 0
				if b <= 0 {
					continue
				}
				// check if a is not a whole number
				if a != float64(int(a)) {
					continue
				}
				// check if b is not a whole number
				if b != float64(int(b)) {
					continue
				}
				part2Sum += int(3*a + b)
			}
			aX, aY, bX, bY, pX, pY = 0, 0, 0, 0, 0, 0
		}
	}
	return part2Sum
}
