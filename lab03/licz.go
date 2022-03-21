package main

import (
	"flag"
	"fmt"
)

const ERR = -100000

func licz() {
	liczba1 := flag.Int("liczba1", ERR, "pierwsza")
	liczba2 := flag.Int("liczba2", ERR, "druga")
	flag.Parse()
	a := *liczba1
	b := *liczba2

	if a == ERR {
		fmt.Print("Podaj pierwszą liczbe: ")
		fmt.Scanf("%d", &b)
	}

	if b == ERR {
		fmt.Print("Podaj drugą liczbe: ")
		fmt.Scanf("%d", &b)
	}

	fmt.Println("Suma liczba:", a, "+", b, "wynosi", a+b)
}
