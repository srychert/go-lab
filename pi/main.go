package main

import (
	"fmt"
	"math"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func piSum(n int) float64 {
	var result float64 = 0.0
	operation := true

	for i := 1; i < n*2; i += 2 {
		if operation {
			result = result + 1.0/float64(i)
		} else {
			result = result - 1.0/float64(i)
		}
		operation = !operation
	}
	return 4 * result
}

func main() {
	myPi := piSum(10_000_000_000)
	fmt.Println(myPi)
	fmt.Println(math.Pi)
	fmt.Println(almostEqual(math.Pi, myPi))
}
