// Package day05 is our package for the 5th AoC day
package day05

import (
	"strconv"
	"strings"
)

// isInList checks if an element is in an int array
func isInList(l []int, e int) bool {
	for _, n := range l {
		if n == e {
			return true
		}
	}
	return false
}

// isInOrder checks if the update conforms to all rules
func isInOrder(updates []int, rules map[int][]int) bool {
	for k, u := range updates {
		before := rules[u]
		for i := k; i >= 0; i-- {
			if isInList(before, updates[i]) {
				return false
			}
		}
	}
	return true
}

// runPart is for both parts
func runPart(input []string) (int, int) {
	sumPart1 := 0
	sumPart2 := 0
	rules := make(map[int][]int, 2000)
	for _, line := range input {
		if line == "" {
			continue
		}
		if strings.Count(line, "|") == 1 {
			rule := strings.Split(line, "|")
			page, _ := strconv.Atoi(rule[0])
			before, _ := strconv.Atoi(rule[1])
			rules[page] = append(rules[page], before)
			continue
		}
		var updates []int
		if strings.Count(line, ",") > 0 {
			for _, s := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(s)
				updates = append(updates, n)
			}
		}
		if isInOrder(updates, rules) {
			middle := len(updates) / 2
			sumPart1 += updates[middle]
			continue
		}
	Sort:
		for !isInOrder(updates, rules) {
			for k, u := range updates {
				before := rules[u]
				for i := k; i >= 0; i-- {
					if isInList(before, updates[i]) {
						var newUpdates []int
						newUpdates = append(newUpdates, updates[:i]...)
						newUpdates = append(newUpdates, updates[i+1])
						newUpdates = append(newUpdates, updates[i])
						newUpdates = append(newUpdates, updates[i+2:]...)
						updates = newUpdates
						continue Sort
					}
				}
			}
		}
		middle := len(updates) / 2
		sumPart2 += updates[middle]
	}
	return sumPart1, sumPart2
}

// RunPart1 is for the second star of the day
func RunPart1(input []string) int {
	result, _ := runPart(input)
	return result
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	_, result := runPart(input)
	return result
}
