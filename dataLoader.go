package testdataloader

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

func GetBasePath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for _, err := ioutil.ReadFile(dir + "/go.mod"); err != nil && len(dir) > 1; {
		println(dir)
		dir = filepath.Dir(dir)
		_, err = ioutil.ReadFile(dir + "/go.mod")
	}
	if len(dir) < 2 {
		panic("No go.mod found")
	}
	return dir
}

// GetRandTestData returns a random file matching the pattern parameter and panics on error.
// The pattern is relative to the "go.mod".
// Example: test files are json files in the "data" subdirectory:
//
// 		- go.mod
// 		- data
//			- test1.json
//			- test2.json
// Here you can provide pattern="data/test*.json"
func GetRandTestData(pattern string) []byte {
	absolutePattern := GetBasePath() + "/" + pattern
	files, err := filepath.Glob(absolutePattern)
	if err != nil {
		panic(err)
	}
	if files == nil {
		panic(fmt.Errorf("Pattern \"%s\" matches no files", absolutePattern))
	}
	filename := files[rand.Intn(len(files))]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return file
}
