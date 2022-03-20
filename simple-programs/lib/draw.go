package lib

import (
	"fmt"
	"strings"
)

// Comment
func DrawRec() {

	var rWidth int
	var rHeight int
	for {

		fmt.Println("Width: ")
		fmt.Scan(&rWidth)
		fmt.Println("Height: ")
		fmt.Scan(&rHeight)

		if rWidth < 1 || rHeight < 1 {
			fmt.Println("Must be positive")
			continue
		}

		spacesCount := 0

		if rWidth > 2 {
			spacesCount = rWidth - 2
		}

		for i := 0; i < rHeight; i++ {
			switch i {
			case 0, rHeight - 1:
				s := strings.Repeat("*", rWidth)
				fmt.Printf("%v\n", s)
			default:
				s := strings.Repeat(" ", spacesCount)
				fmt.Printf("*%s*\n", s)
			}
		}
		break
	}

}
