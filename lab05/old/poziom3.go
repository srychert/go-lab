package old

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Poziom3() {
	rand.Seed(time.Now().Unix())
	upperBound := 100
	computerNumber := rand.Intn(upperBound) + 1
	input := Input{0, ""}

	fmt.Printf("Teraz będziesz zgadywać liczbę, którą wylosowałem z przedziału [1, %v]\n", upperBound)
	fmt.Println("Napisz 'koniec' aby wyjść")

	for input.guess != computerNumber {
		fmt.Print("Podaj liczbę: ")
		input = getUserInput()
		if input.command == "koniec" {
			fmt.Println("Żegnaj")
			os.Exit(0)
		}
		if input.guess < computerNumber && input.guess != 0 {
			fmt.Println("Za mała")
		}
		if input.guess > computerNumber && input.guess != 0 {
			fmt.Println("Za duża")
		}
	}

	fmt.Println("Brawo wygrałeś!")
	var newGame string
NEWGAME:
	fmt.Print("Gramy jeszcze raz? [T/N]: ")
	fmt.Scanln(&newGame)
	if strings.ToUpper(newGame) == "T" {
		Poziom3()
	} else if strings.ToUpper(newGame) == "N" {
		os.Exit(0)
	}
	goto NEWGAME
}
