package main

import "fmt"
import "packages/math"

// NOTE: Go wants your project to look like this:
// project_name
// 	- src
// 		- packages
// 			- math/math.go
// 			- main.go
//
// With GOROOT=/path/to/project_name
//
// 1. cd ingo packages/math & run `go install`
// 2. cd into packages/ & run `go run main`

func main() {
	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
