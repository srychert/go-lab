package old

import (
	"fmt"
	"math/rand"
)

func Poziom1() {
	computerNumber := rand.Intn(100) + 1
	var userGuess int

	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")

	for userGuess != computerNumber {
		fmt.Println("Podaj liczbę: ")
		fmt.Scan(&userGuess)
		if userGuess < computerNumber {
			fmt.Println("Za mała")
		}
		if userGuess > computerNumber {
			fmt.Println("Za duża")
		}
	}
	fmt.Println("Brawo wygrałeś!")
}
