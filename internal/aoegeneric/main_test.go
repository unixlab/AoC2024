package aoegeneric

import (
	"reflect"
	"testing"
)

func TestAbsDiff(t *testing.T) {
	actualResult := AbsDiff(1, 2)
	expectedResult := 1
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("AbsDiff() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = AbsDiff(7, 2)
	expectedResult = 5
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("AbsDiff() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = AbsDiff(97, 27)
	expectedResult = 70
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("AbsDiff() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = AbsDiff(-10, 10)
	expectedResult = 20
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("AbsDiff() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = AbsDiff(10, -10)
	expectedResult = 20
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("AbsDiff() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestGetDistance(t *testing.T) {
	actualResult := GetDistance(0, 0, 0, 1)
	expectedResult := 1
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("GetDistance() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = GetDistance(0, 0, 1, 1)
	expectedResult = 2
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("GetDistance() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = GetDistance(0, 1, 3, 5)
	expectedResult = 7
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("GetDistance() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = GetDistance(1, 1, 0, 0)
	expectedResult = 2
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("GetDistance() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestPow(t *testing.T) {
	actualResult := Pow(2, 4)
	expectedResult := 16
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("Pow() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = Pow(10, 3)
	expectedResult = 1000
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("Pow() = expected %v, got %v", expectedResult, actualResult)
	}
	actualResult = Pow(5, 2)
	expectedResult = 25
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("Pow() = expected %v, got %v", expectedResult, actualResult)
	}
}
