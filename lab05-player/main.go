package main

import (
	"bufio"
	"fmt"
	"os"
)

func printMsg(msg any) {
	fmt.Fprintln(os.Stderr, msg)
}

func main() {
	// read bounds and info
	reader := bufio.NewReader(os.Stdin)
	boundsMsg, _ := reader.ReadString('\n')
	exitMsg, _ := reader.ReadString('\n')

	lowerBound := 0
	upperBound := 100
	exitCommand := "koniec"

	setUpperBound(boundsMsg, &upperBound)
	setExitCommand(exitMsg, &exitCommand)

	var input string

	play := true
	lowerBoundLoop := lowerBound
	upperBoundLoop := upperBound
	guess := upperBound / 2

	i := 0
	stop := 10

	// max log2(upperbound) rounded up, e.g.:
	// upperVound = 100 => iterations = 7
	// upperVound = 131 => iterations = 8
	for play {
		input, _ = reader.ReadString('\n')

		switch input {
		case "Podaj liczbę:\n":
			// start
			fmt.Println(guess)
		case "Za mała\n":
			// logic
			lowerBoundLoop = guess
			if (upperBoundLoop-guess)/2 == 0 {
				guess = guess + (upperBoundLoop-guess)/2 + 1
			} else {
				guess = guess + (upperBoundLoop-guess)/2
			}
			continue
		case "Za duża\n":
			// logic
			upperBoundLoop = guess
			guess = guess - (guess-lowerBoundLoop)/2
			continue
		case "Brawo wygrałeś!\n":
			lowerBoundLoop = lowerBound
			upperBoundLoop = upperBound
			guess = upperBound / 2
			continue
		case "Podaj imie:\n":
			fmt.Println("mx2")
		case "Nowy rekord osobisty!\n":
			continue
		case "Nowy rekord rankingu!\n":
			continue
		case "Gramy jeszcze raz? [T/N]:\n":
			// when to stop?
			if i >= stop {
				fmt.Println("n")
				play = false
				input, _ = reader.ReadString(']')
				printMsg(input)
			}
			i++
			fmt.Println("t")
		}
	}

	printMsg("Playing completed!\n")
}
