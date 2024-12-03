// Package day03 is our package for the 3th AoC day
package day03

import (
	"regexp"
	"strconv"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	sum := 0
	for _, line := range input {
		mulSearch, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
		finds := mulSearch.FindAllStringSubmatch(line, -1)
		for _, find := range finds {
			num1, _ := strconv.Atoi(find[1])
			num2, _ := strconv.Atoi(find[2])
			sum += num1 * num2
		}
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	sum := 0
	mul := true
	for _, line := range input {
		mulSearch, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)|(don't)|(do)`)
		finds := mulSearch.FindAllStringSubmatch(line, -1)
		for _, find := range finds {
			if find[0] == "do" {
				mul = true
				continue
			}
			if find[0] == "don't" {
				mul = false
				continue
			}
			if !mul {
				continue
			}
			num1, _ := strconv.Atoi(find[1])
			num2, _ := strconv.Atoi(find[2])
			sum += num1 * num2
		}
	}
	return sum
}
