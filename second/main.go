package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func getSecondHighestInt(arr []int) int {
	high := math.MinInt
	secondHighest := math.MinInt

	for _, value := range arr {
		if value > high {
			high, secondHighest = value, high
		} else if value > secondHighest {
			secondHighest = value
		}
	}
	return secondHighest
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 0)

	for i := 0; i < 9; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	fmt.Println(numbers, "\n", getSecondHighestInt(numbers))
}
