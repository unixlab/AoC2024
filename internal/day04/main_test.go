package day04

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aocinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 18
	actualResult := RunPart1(aocinput.Read("../../", "day04", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 2530
	actualResult := RunPart1(aocinput.Read("../../", "day04", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 9
	actualResult := RunPart2(aocinput.Read("../../", "day04", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 1921
	actualResult := RunPart2(aocinput.Read("../../", "day04", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
