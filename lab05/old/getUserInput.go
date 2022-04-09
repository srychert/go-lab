package old

import (
	"fmt"
	"strconv"
)

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
