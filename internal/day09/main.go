// Package day09 is our package for the 9th AoC day
package day09

import (
	"container/list"
	"strconv"
	"strings"
)

// Chunk holds all important data of one chunk of data of free space
type Chunk struct {
	Idx    int
	Length int
	Free   bool
	Moved  bool
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	sum := 0
	for _, line := range input {
		disk := list.New()
		for idx, c := range strings.Split(line, "") {
			n, _ := strconv.Atoi(c)
			if idx%2 == 0 {
				// file
				for i := 0; i < n; i++ {
					disk.PushBack(idx / 2)
				}
			} else {
				// free
				for i := 0; i < n; i++ {
					disk.PushBack(-1)
				}
			}
		}
		head := disk.Front()
		for i := 0; i < disk.Len(); i++ {
			if head.Value.(int) == -1 {
				back := disk.Back()
				for back.Value.(int) < 0 {
					back = back.Prev()
				}
				disk.MoveBefore(back, head)
				disk.MoveToBack(head)
				head = back
			}
			head = head.Next()
		}
		head = disk.Front()
		for i := 0; i < disk.Len(); i++ {
			val := head.Value.(int)
			if val > 0 {
				sum += i * val
			}
			head = head.Next()
		}
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	sum := 0
	var chunks []Chunk
	for _, line := range input {
		disk := list.New()
		for idx, c := range strings.Split(line, "") {
			var chunk Chunk
			chunk.Idx = idx / 2
			chunk.Length, _ = strconv.Atoi(c)
			if idx%2 == 0 {
				chunk.Free = false
			} else {
				chunk.Free = true
			}
			chunks = append(chunks, chunk)
			disk.PushBack(idx)
		}
		change := true
		for change {
			change = false
			head := disk.Back()
			for i := disk.Len(); i > 0; i-- {
				if chunks[head.Value.(int)].Free {
					head = head.Prev()
					continue
				}
				if chunks[head.Value.(int)].Moved {
					head = head.Prev()
					continue
				}
				neededSpace := chunks[head.Value.(int)].Length
				chunker := disk.Front()
				for chunker != disk.Back() {
					if chunks[head.Value.(int)] == chunks[chunker.Value.(int)] {
						break
					}
					idx := chunker.Value.(int)
					if !chunks[idx].Free {
						chunker = chunker.Next()
						continue
					}
					if chunks[idx].Length < neededSpace {
						chunker = chunker.Next()
						continue
					}
					var newChunk Chunk
					newChunk.Idx = chunks[idx].Idx
					newChunk.Length = neededSpace
					newChunk.Free = true
					newChunk.Moved = true
					chunks[idx].Length -= neededSpace
					chunks[idx].Moved = true
					newPos := disk.Front()
					for newPos.Value.(int) != idx {
						newPos = newPos.Next()
					}
					beforeHead := head.Prev()
					disk.MoveBefore(head, newPos)
					chunks = append(chunks, newChunk)
					newLast := len(chunks) - 1
					disk.InsertAfter(newLast, beforeHead)
					head = disk.Back()
					change = true
					break
				}
				head = head.Prev()
			}

		}
		var numbers []int
		sumer := disk.Front()
		for i := 0; i < disk.Len(); i++ {
			idx := sumer.Value.(int)
			if chunks[idx].Free {
				for j := 0; j < chunks[idx].Length; j++ {
					numbers = append(numbers, 0)
				}
			} else {
				for j := 0; j < chunks[idx].Length; j++ {
					numbers = append(numbers, chunks[idx].Idx)
				}
			}
			sumer = sumer.Next()
		}
		for idx, num := range numbers {
			sum += idx * num
		}
	}
	return sum
}
