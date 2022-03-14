package main

import (
    "fmt"
    //"math/rand"
)

func nwd(a int, b int) int{
	if b != 0{
		return nwd(b, a%b) 
	}
	return a
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
}
