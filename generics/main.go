package main

import "fmt"

type Number interface {
	int8 | int16 | int32 | int | int64 | float32 | float64
}

func printGeneric[v any](variable v) {
	fmt.Println(variable)
}

func increment[age Number](myAge age) any {
	val := myAge + 1
	fmt.Println(val)
	return val
}

func bubbleSort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {
	printGeneric("test")
	printGeneric(34)

	var age1 int = 40
	var age2 float32 = 69.420

	newAge1 := increment(age1)
	newAge2 := increment(age2)

	fmt.Printf("old: %-5v type: %-7T new: %-5v type: %T\n", age1, age1, newAge1, newAge1)
	fmt.Printf("old: %-5v type: %-7T new: %-5v type: %T\n", age2, age2, newAge2, newAge2)

	list := []int32{4, 2, 78, 9, 21, 60}
	listFloat := []float64{4.3, 57.2, 20.1, 1.1, 5.9}

	sorted := bubbleSort(list)
	sortedFloat := bubbleSort(listFloat)

	fmt.Println(sorted)
	fmt.Println(sortedFloat)

}
