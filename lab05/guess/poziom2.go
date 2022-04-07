package guess

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Poziom2() {
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
}
