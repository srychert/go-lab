package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var liczba1 int64
	var liczba2 int64

	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		liczba1, _ = strconv.ParseInt(os.Args[2], 10, 64)
		liczba2, _ = strconv.ParseInt(os.Args[4], 10, 64)
	} else {
		fmt.Scan(&liczba1)
		fmt.Scan(&liczba2)
	}
	fmt.Println(liczba1 + liczba2)
}
