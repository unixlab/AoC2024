package day10

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aocinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 36
	actualResult := RunPart1(aocinput.Read("../../", "day10", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 786
	actualResult := RunPart1(aocinput.Read("../../", "day10", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 81
	actualResult := RunPart2(aocinput.Read("../../", "day10", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 1722
	actualResult := RunPart2(aocinput.Read("../../", "day10", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
