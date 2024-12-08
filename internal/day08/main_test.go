package day08

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 14
	actualResult := RunPart1(aoeinput.Read("../../", "day08", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 291
	actualResult := RunPart1(aoeinput.Read("../../", "day08", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 34
	actualResult := RunPart2(aoeinput.Read("../../", "day08", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 1015
	actualResult := RunPart2(aoeinput.Read("../../", "day08", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
