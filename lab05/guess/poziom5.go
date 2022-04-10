package guess

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
	"os"
)

func Poziom5() {
	// defined in newGame.go
	var games []Result
	// read results form file
	results := readResults()
GAME:
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	upperBound := 100
	computerNumber := math_rand.Intn(upperBound) + 1
	input := Input{0, ""}

	var score int = 0
	// fmt.Println(computerNumber)
	fmt.Printf("Teraz będziesz zgadywać liczbę, którą wylosowałem z przedziału [1, %v]\n", upperBound)
	fmt.Println("Napisz 'koniec' aby wyjść")

	for input.guess != computerNumber {
		fmt.Println("Podaj liczbę:")
		input = getUserInput()
		if input.guess != 0 {	
			score++
		}
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

	games = addToGames(games, score, &results)
	if newGame(games, results) {
		goto GAME
	}
}
