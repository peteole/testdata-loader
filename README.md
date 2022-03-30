# testdata-loader
Loading data for your tests is always a pain in go since where you do not specify paths relative to the test itself, but relative to the entrypoint of the test.
This brings lots of confusion.

In order to release this pain, this package allowas you to load test data relative to the your go.mod.
Like this you no more need to worry about the entrypoint of your program and concentrate on the logic you want to do:

## Example
Consider the structure
```txt
- go.mod
- main_test.go
- data
  - test1.json
  - test2.json
  ```
  Now run `go get github.com/peteole/testdata-loader` and modify your test:
  
  ```go
  package main_test
  import (
	"testing"

	testdataloader "github.com/peteole/testdata-loader"
  )
func TestMain(t *testing.T){
    test1:=testdataloader.GetTestFile("data/test1.json")
    randomTest:=testdataloader.GetRandTestFile("data/test*.json")
    //now do something using your test data
}
```
