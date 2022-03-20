package lib

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

func Second() {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 0)

	itemsLength := 9
	maxValue := 100

	for i := 0; i < itemsLength; i++ {
		numbers = append(numbers, rand.Intn(maxValue))
	}

	fmt.Println(numbers, "\n", getSecondHighestInt(numbers))
}
