package main

import (
    "fmt"
	"strconv"
    //"math/rand"
)

func nwd(a int, b int) int{
	if b != 0{
		return nwd(b, a%b) 
	}
	return a
}

// Given a number n, print n-th Fibonacci Number
func fib(n int) int{
	if n <= 1{
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibSequencePrint(n int){
	for i := 1; i <= n; i++{
		fmt.Println(strconv.Itoa(i) + ")", fib(i))
	}
}

func main() {
	// Szybki test czy działa funkcja
	// stop := 100
	// for i := 15; i <=stop; i++{
	// 	j := rand.Intn(stop)
	// 	fmt.Println(i, j, nwd(i, j))
	// }

	var first int;
	var second int;

	fmt.Println("Podaj dwie liczby całkowite")
	fmt.Scan(&first, &second)
	fmt.Println("Nwd:", nwd(first, second))
	
	var numb int;
	fmt.Println("Która liczba ciągu fibonacciego?")
	fmt.Scan(&numb)
	fmt.Println(fib(numb))
	fibSequencePrint(numb)
}
