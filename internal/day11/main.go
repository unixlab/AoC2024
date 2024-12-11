// Package day11 is our package for the 11th AoC day
package day11

import (
	"strconv"
	"strings"
)

// Stone is a linked list of our stones
type Stone struct {
	Next        *Stone
	Prev        *Stone
	Value       int
	ValueLength int
}

// runWithLinkedList uses a linked list
func runWithLinkedList(input []string, iterations int) int {
	head := &Stone{
		Next:  nil,
		Prev:  nil,
		Value: 0,
	}
	var last *Stone
	for _, line := range input {
		for idx, c := range strings.Fields(line) {
			n, _ := strconv.Atoi(c)
			if idx == 0 {
				head.Value = n
				head.ValueLength = len(c)
				last = head
				continue
			}
			newStone := &Stone{
				Next:        nil,
				Prev:        last,
				Value:       n,
				ValueLength: len(c),
			}
			last.Next = newStone
			last = newStone
		}
	}
	last.Next = head
	head.Prev = last
	for i := 0; i < iterations; i++ {
		iter := head
		for {
			if iter.Value == 0 {
				iter.Value = 1
			} else if iter.ValueLength%2 == 0 {
				s := strconv.Itoa(iter.Value)
				s1 := s[:len(s)/2]
				s2 := s[len(s)/2:]
				n1, _ := strconv.Atoi(s1)
				n2, _ := strconv.Atoi(s2)
				newStone := &Stone{
					Next:        iter,
					Prev:        iter.Prev,
					Value:       n1,
					ValueLength: len(strconv.Itoa(n1)),
				}
				iter.Prev.Next = newStone
				iter.Prev = newStone
				iter.Value = n2
				iter.ValueLength = len(strconv.Itoa(n2))
				if iter == head {
					head = newStone
				}
			} else {
				iter.Value *= 2024
				iter.ValueLength = len(strconv.Itoa(iter.Value))
			}
			iter = iter.Next
			if iter == head {
				break
			}
		}
	}
	sum := 0
	printer := head
	for {
		sum++
		printer = printer.Next
		if printer == head {
			break
		}
	}
	return sum
}

// runWithArray uses an array
func runWithArray(input []string, iterations int) int {
	var stones []int
	for _, line := range input {
		for _, c := range strings.Fields(line) {
			n, _ := strconv.Atoi(c)
			stones = append(stones, n)
		}
	}
	for i := 0; i < iterations; i++ {
		for s := 0; s < len(stones); s++ {
			if stones[s] == 0 {
				stones[s] = 1
			} else if len(strconv.Itoa(stones[s]))%2 == 0 {
				s0 := strconv.Itoa(stones[s])
				s1 := s0[:len(s0)/2]
				s2 := s0[len(s0)/2:]
				n1, _ := strconv.Atoi(s1)
				n2, _ := strconv.Atoi(s2)
				newStones := make([]int, len(stones[:s]))
				_ = copy(newStones, stones[:s])
				newStones = append(newStones, n1, n2)
				newStones = append(newStones, stones[s+1:]...)
				stones = newStones
				s++
			} else {
				stones[s] *= 2024
			}
		}
	}
	return len(stones)
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	return runWithLinkedList(input, 25)
}

// RunPart2 is for the second star of the day
func RunPart2(input []string, example bool) int {
	iterations := 75
	if example {
		iterations = 25
	}
	return runWithArray(input, iterations)
}
