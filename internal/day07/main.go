// Package day07 is our package for the 7th AoC day
package day07

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/unixlab/AoC2024/internal/aoegeneric"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	var equations [][]int
	for _, line := range input {
		var equation []int
		for _, c := range strings.Fields(line) {
			c = strings.ReplaceAll(c, ":", "")
			n, _ := strconv.Atoi(c)
			equation = append(equation, n)
		}
		equations = append(equations, equation)
	}
	sum := 0
	for _, equation := range equations {
		for i := 0; i < aoegeneric.Pow(2, len(equation)-2); i++ {
			target := 0
			mutations := i
			mutationSum := 0
			for idx, number := range equation {
				if idx == 0 {
					target = number
					continue
				}
				if idx == 1 {
					mutationSum = number
					continue
				}
				if mutations&1 == 0 {
					mutationSum += number
				} else {
					mutationSum *= number
				}
				mutations >>= 1
			}
			if mutationSum == target {
				sum += target
				break
			}
			if mutationSum > target {
				continue
			}
		}
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	var equations [][]string
	for _, line := range input {
		var equation []string
		for _, c := range strings.Fields(line) {
			c = strings.ReplaceAll(c, ":", "")
			equation = append(equation, c)
		}
		equations = append(equations, equation)
	}
	sum := big.NewInt(0)
	for _, equation := range equations {
		for i := 0; i < aoegeneric.Pow(4, len(equation)-2); i++ {
			target := big.NewInt(0)
			mutations := i
			if mutations&3 == 3 {
				continue
			}
			mutationSum := big.NewInt(0)
			numbersUsed := 0
			for idx, number := range equation {
				bigIntNum, _ := big.NewInt(0).SetString(number, 10)
				if idx == 0 {
					target, _ = target.SetString(number, 10)
					continue
				}
				if idx == 1 {
					mutationSum, _ = mutationSum.SetString(number, 10)
					continue
				}
				if mutations&3 == 0 {
					mutationSum = mutationSum.Add(mutationSum, bigIntNum)
				} else if mutations&3 == 1 {
					mutationSum = mutationSum.Mul(mutationSum, bigIntNum)
				} else if mutations&3 == 2 {
					mutationSum, _ = mutationSum.SetString(mutationSum.String()+bigIntNum.String(), 10)
				} else if mutations&3 == 3 {
					continue
				}
				mutations >>= 2
				numbersUsed++
			}
			if mutationSum.Cmp(target) == 0 && numbersUsed == len(equation)-2 {
				sum = sum.Add(sum, target)
				break
			}
			if mutationSum.Cmp(target) > 0 {
				continue
			}
		}
	}
	return int(sum.Int64())
}
