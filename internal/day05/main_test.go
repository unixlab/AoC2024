package day05

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 143
	actualResult := RunPart1(aoeinput.Read("../../", "day05", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 6051
	actualResult := RunPart1(aoeinput.Read("../../", "day05", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 123
	actualResult := RunPart2(aoeinput.Read("../../", "day05", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 5093
	actualResult := RunPart2(aoeinput.Read("../../", "day05", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}