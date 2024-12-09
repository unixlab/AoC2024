package day09

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 1928
	actualResult := RunPart1(aoeinput.Read("../../", "day09", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 6607511583593
	actualResult := RunPart1(aoeinput.Read("../../", "day09", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 2858
	actualResult := RunPart2(aoeinput.Read("../../", "day09", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 6636608781232
	actualResult := RunPart2(aoeinput.Read("../../", "day09", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
