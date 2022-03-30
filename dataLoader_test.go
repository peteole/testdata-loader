package testdataloader_test

import (
	"testing"

	testdataloader "github.com/peteole/testdata-loader"
)

func TestLoadTest(t *testing.T) {
	test1 := testdataloader.GetTestFile("testdata/test1.txt")
	if string(test1) != "test content 1" {
		t.Errorf("Test data \"%s\" do not match the file content \"test content 1\"", test1)
	}
}

func TestLoadRandTest(t *testing.T) {
	test1 := testdataloader.GetRandTestFile("testdata/test*.txt")
	if string(test1) != "test content 1" {
		t.Errorf("Test data \"%s\" do not match the file content \"test content 1\"", test1)
	}
}
