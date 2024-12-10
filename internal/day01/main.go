// Package day01 is our package for the 1th AoC day
package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/unixlab/AoC2024/internal/aocgeneric"
)

// getSortedColumns reads the input and returns the columns as arrays
func getSortedColumns(input []string) [2][]int {
	var columns [2][]int
	for _, line := range input {
		for k, nString := range strings.Fields(line) {
			n, _ := strconv.Atoi(nString)
			columns[k] = append(columns[k], n)
		}
	}
	sort.Ints(columns[0])
	sort.Ints(columns[1])
	return columns
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	columns := getSortedColumns(input)
	var sum int
	for i := 0; i < len(columns[0]); i++ {
		sum += aocgeneric.AbsDiff(columns[0][i], columns[1][i])
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	columns := getSortedColumns(input)
	var sum int
	for _, numberColumn1 := range columns[0] {
		count := 0
		for _, numberColumn2 := range columns[1] {
			if numberColumn1 == numberColumn2 {
				count++
			}
		}
		sum += numberColumn1 * count
	}
	return sum
}
