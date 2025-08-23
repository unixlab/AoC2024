package day12

import (
	"testing"

	"github.com/unixlab/AoC2024/internal/aocinput"
)

func TestRunPart1Example1(t *testing.T) {
	expectedResult := 140
	actualResult := RunPart1(aocinput.Read("../../", "day12e1", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1Example2(t *testing.T) {
	expectedResult := 772
	actualResult := RunPart1(aocinput.Read("../../", "day12e2", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1Example3(t *testing.T) {
	expectedResult := 1930
	actualResult := RunPart1(aocinput.Read("../../", "day12e3", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 1363484
	actualResult := RunPart1(aocinput.Read("../../", "day12", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

//func TestRunPart2Example1(t *testing.T) {
//	expectedResult := 80
//	actualResult := RunPart2(aocinput.Read("../../", "day12e1", true))
//	if actualResult != expectedResult {
//		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
//	}
//}

//func TestRunPart2Example2(t *testing.T) {
//	expectedResult := 436
//	actualResult := RunPart2(aocinput.Read("../../", "day12e2", true))
//	if actualResult != expectedResult {
//		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
//	}
//}

//func TestRunPart2Example4(t *testing.T) {
//	expectedResult := 236
//	actualResult := RunPart2(aocinput.Read("../../", "day12e4", true))
//	if actualResult != expectedResult {
//		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
//	}
//}

//func TestRunPart2Example5(t *testing.T) {
//	expectedResult := 368
//	actualResult := RunPart2(aocinput.Read("../../", "day12e5", true))
//	if actualResult != expectedResult {
//		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
//	}
//}

//func TestRunPart2(t *testing.T) {
//	expectedResult := 0
//	actualResult := RunPart2(aocinput.Read("../../", "day12", false))
//	if actualResult != expectedResult {
//		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
//	}
//}
