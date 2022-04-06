package guess

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type input2 interface {
	int | string
}

type Input struct {
	guess   int
	command string
}

func setInput(i interface{}) Input {
	var input Input
	switch v := i.(type) {
	case nil:
		fmt.Println("x is nil") // here v has type interface{}
	case int:
		fmt.Println("x is", v) // here v has type int
		input.guess = i.(int)
	case string:
		fmt.Println("x is string") // here v has type interface{}
		input.command = i.(string)
	}
	return input
}

func getUserInput() (int, string) {

	var test any
	var input interface{}
	fmt.Scan(&input)
	fmt.Println(input, test)
	i := setInput(input)

	return i.guess, i.command

}

func Poziom3() {
	rand.Seed(time.Now().Unix())
	upperBound := 100
	computerNumber := rand.Intn(upperBound) + 1
	var userGuess int

	var score int = 0

	type Games struct {
		names  []string
		scores []int
	}

	var games Games

	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")
	fmt.Println("Napisz 'koniec' aby wyjść")

	for userGuess != computerNumber {
		score += 1
		fmt.Print("Podaj liczbę: ")
		userGuess, command := getUserInput()

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
