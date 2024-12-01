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
