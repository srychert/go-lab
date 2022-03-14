package main

import "fmt"

func nwd(a int, b int) int {
	if b != 0 {
		return nwd(b, a%b)
	}
	return a
}

func main() {
	var first int
	var second int

	fmt.Println("Podaj dwie liczby całkowite")
	fmt.Scan(&first, &second)
	fmt.Println("Nwd:", nwd(first, second))
}
