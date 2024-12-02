package day02

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aoeinput"
)

func TestCheckIncreasing(t *testing.T) {
	expectedResult := true
	actualResult := checkIncreasing([]int{1, 2, 3, 4, 5}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkIncreasing([]int{1, 3, 5, 6, 7}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkIncreasing([]int{1, 4, 5, 6, 9}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkIncreasing([]int{1, 5, 10}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkIncreasing([]int{1, 10, 20}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkIncreasing([]int{1, 2, 3, 4, 3, 4, 5, 6}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkIncreasing([]int{1, 2, 3, 3, 4, 5, 6}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkIncreasing([]int{1, 5, 10}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkIncreasing([]int{1, 10, 15}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkIncreasing([]int{1, 2, 3, 4, 3, 4, 5, 6}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkIncreasing([]int{1, 2, 3, 3, 4, 5, 6}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkIncreasing([]int{1, 2, 3, 13, 4, 5, 6}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkIncreasing([]int{1, 1, 1}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkIncreasing([]int{1, 4, 4, 7}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkIncreasing([]int{1, 6, 7, 8}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkIncreasing([]int{6, 7, 8, 1}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkIncreasing([]int{9, 6, 9, 10}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkIncreasing() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestCheckDecreasing(t *testing.T) {
	expectedResult := true
	actualResult := checkDecreasing([]int{5, 4, 3, 2, 1}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{7, 6, 5, 3, 1}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{9, 6, 5, 4, 1}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkDecreasing([]int{10, 5, 1}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkDecreasing([]int{20, 10, 1}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkDecreasing([]int{10, 9, 8, 7, 8, 7, 6, 5}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkDecreasing([]int{10, 9, 8, 7, 7, 6, 5}, false)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkDecreasing([]int{10, 5, 1}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkDecreasing([]int{20, 10, 1}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkDecreasing([]int{10, 9, 8, 7, 8, 7, 6, 5}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{10, 9, 8, 7, 7, 6, 5}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{10, 9, 8, 14, 7, 6, 5}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = false
	actualResult = checkDecreasing([]int{9, 9, 9}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{9, 6, 3, 3, 2}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{9, 3, 2, 1}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{3, 2, 1, 9}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{3, 9, 2, 1}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = true
	actualResult = checkDecreasing([]int{9, 3, 2, 1}, true)
	if actualResult != expectedResult {
		t.Fatalf("checkDecreasing() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1Example(t *testing.T) {
	expectedResult := 2
	actualResult := RunPart1(aoeinput.Read("../../", "day02", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 526
	actualResult := RunPart1(aoeinput.Read("../../", "day02", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 4
	actualResult := RunPart2(aoeinput.Read("../../", "day02", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 566
	actualResult := RunPart2(aoeinput.Read("../../", "day02", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
