// Package aoeinput handles all the AoC input
package aoeinput

import (
	"bufio"
	"os"
	"strings"
)

// Read is our global function to read AoC input
func Read(path string, day string, example bool) []string {
	var inputPath strings.Builder
	inputPath.WriteString(path)
	if example {
		inputPath.WriteString("examples")
	} else {
		inputPath.WriteString("inputs")
	}
	inputPath.WriteString("/")
	inputPath.WriteString(day)
	file, err := os.Open(inputPath.String())
	if err != nil {
		return []string{"error while reading input"}
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
