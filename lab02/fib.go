package main

import (
	"fmt"
	"strconv"
)

// Given a number n, print n-th Fibonacci Number
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibSequencePrint(n int) {
	var a int64 = 0
	var b int64 = 1
	for i := 0; i <= n; i++ {
		fmt.Println(strconv.Itoa(i)+")", a)
		a, b = b, a+b
	}
}

func main() {
	var n int

	fmt.Println("Ile liczb fibonacciego?")
	fmt.Scan(&n)
	fibSequencePrint(n)
}
