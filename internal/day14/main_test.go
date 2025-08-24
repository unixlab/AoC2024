package day14

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aocinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 12
	actualResult := RunPart1(aocinput.Read("../../", "day14", true), 11, 7)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 219512160
	actualResult := RunPart1(aocinput.Read("../../", "day14", false), 101, 103)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 0
	actualResult := RunPart2(aocinput.Read("../../", "day14", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 0
	actualResult := RunPart2(aocinput.Read("../../", "day14", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
