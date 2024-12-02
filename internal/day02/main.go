// Package day02 is our package for the 2nd AoC day
package day02

import (
	"strconv"
	"strings"
)

// checkIncreasing checks if a list is ever-increasing by less then 3
func checkIncreasing(list []int, ignoreOneError bool) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] || list[i]+3 < list[i+1] {
			if !ignoreOneError {
				return false
			}
			for k := 0; k < len(list); k++ {
				var temp []int
				temp = append(temp, list[:k]...)
				temp = append(temp, list[k+1:]...)
				if checkIncreasing(temp, false) {
					return true
				}
			}
			return false
		}
	}
	return true
}

// checkDecreasing checks if a list is ever-decreasing by less then 3
func checkDecreasing(list []int, ignoreOneError bool) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] <= list[i+1] || list[i]-3 > list[i+1] {
			if !ignoreOneError {
				return false
			}
			for k := 0; k < len(list); k++ {
				var temp []int
				temp = append(temp, list[:k]...)
				temp = append(temp, list[k+1:]...)
				if checkDecreasing(temp, false) {
					return true
				}
			}
			return false
		}
	}
	return true
}

// runPart has logic that can be reused between part1 and part2
func runPart(input []string) (int, int) {
	numberOfSafeLevelsPart1 := 0
	numberOfSafeLevelsPart2 := 0
	for _, line := range input {
		var numbers []int
		for _, v := range strings.Fields(line) {
			n, _ := strconv.Atoi(v)
			numbers = append(numbers, n)
		}
		if checkIncreasing(numbers, false) || checkDecreasing(numbers, false) {
			numberOfSafeLevelsPart1++
		}
		if checkIncreasing(numbers, true) || checkDecreasing(numbers, true) {
			numberOfSafeLevelsPart2++
		}
	}
	return numberOfSafeLevelsPart1, numberOfSafeLevelsPart2
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	result, _ := runPart(input)
	return result
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	_, result := runPart(input)
	return result
}
