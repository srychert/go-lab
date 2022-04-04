package guess

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Poziom3() {
	rand.Seed(time.Now().Unix())
	computerNumber := rand.Intn(100) + 1
	var userGuess int
	var command string
	var score int = 0

	type Games struct {
		names  []string
		scores []int
	}

	var games Games

	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")

	for userGuess != computerNumber {
		score += 1
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
	var name string
	fmt.Println("Jak masz na imie?")
	fmt.Scan(&name)
	games.names = append(games.names, name)
	games.scores = append(games.scores, score)

	var newGame string
	fmt.Print("Gramy jeszcze raz? [T/N]: ")
	fmt.Scan(&newGame)
	if strings.ToUpper(newGame) == "T" {
		Poziom3()
	}
	fmt.Println(games)
}
