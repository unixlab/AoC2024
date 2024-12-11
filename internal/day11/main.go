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

// run is for both star of the day
func run(input []string, iterations int) int {
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

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	return run(input, 25)
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	return run(input, 75)
}
