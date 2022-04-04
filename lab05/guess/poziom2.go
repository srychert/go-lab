package guess

import (
	"fmt"
	"math/rand"
	"os"
)

func Poziom2() {
	computerNumber := rand.Intn(100) + 1
	var userGuess int
	var command string

	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")

	for userGuess != computerNumber {
		fmt.Println("Podaj liczbę: ")
		_, err := fmt.Scan(&userGuess)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Scan(&command)
		if command == "koniec" {
			os.Exit(0)
		}
		if userGuess < computerNumber {
			fmt.Println("Za mała")
		}
		if userGuess > computerNumber {
			fmt.Println("Za duża")
		}
	}
	fmt.Println("Brawo wygrałeś!")
}
