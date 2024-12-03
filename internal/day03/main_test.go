package day03

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 161
	actualResult := RunPart1(aoeinput.Read("../../", "day03p1", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 182619815
	actualResult := RunPart1(aoeinput.Read("../../", "day03", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 48
	actualResult := RunPart2(aoeinput.Read("../../", "day03p2", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 80747545
	actualResult := RunPart2(aoeinput.Read("../../", "day03", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
