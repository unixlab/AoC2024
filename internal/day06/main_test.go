package day06

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 41
	actualResult := RunPart1(aoeinput.Read("../../", "day06", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 5199
	actualResult := RunPart1(aoeinput.Read("../../", "day06", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 6
	actualResult := RunPart2(aoeinput.Read("../../", "day06", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 1915
	actualResult := RunPart2(aoeinput.Read("../../", "day06", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
