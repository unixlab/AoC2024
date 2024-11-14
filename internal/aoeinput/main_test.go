package aoeinput

import (
	"os"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	tempDirectoryName, err := os.MkdirTemp("/tmp/", "aoeinput")
	if err != nil {
		t.Fatalf("file creation failed: %v", err)
	}

	var file *os.File
	for _, path := range []string{"inputs", "examples"} {
		newDir := tempDirectoryName + "/" + path
		err = os.Mkdir(newDir, 0755)
		if err != nil {
			t.Fatalf("file creation failed: %v", err)
		}
		file, err = os.Create(newDir + "/test")
		if err != nil {
			t.Fatalf("file creation failed: %v", err)
		}
		_, err = file.WriteString(path)
		if err != nil {
			t.Fatalf("file creation failed: %v", err)
		}
		err = file.Close()
		if err != nil {
			t.Fatalf("file creation failed: %v", err)
		}
	}

	defer os.RemoveAll(tempDirectoryName)

	tempDirectoryName += "/"

	expectedResult := []string{"inputs"}
	actualResult := Read(tempDirectoryName, "test", false)
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("Read() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = []string{"examples"}
	actualResult = Read(tempDirectoryName, "test", true)
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("Read() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestReadError(t *testing.T) {
	expectedResult := []string{"error while reading input"}
	actualResult := Read("/tmp", "NotADay", true)
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("Read() = expected %v, got %v", expectedResult, actualResult)
	}
}
