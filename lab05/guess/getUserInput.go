package guess

import (
	"fmt"
	"strconv"
)

type Input struct {
	guess   int
	command string
}

func getUserInput() Input {
	var scanedString string
	fmt.Scanln(&scanedString)

	input := Input{0, scanedString}

	guess, err := strconv.Atoi(scanedString)

	if err == nil {
		input.guess = guess
	}

	return input
}
