package main

import (
	"fmt"
	"os"
	lib "simple-programs/lib"
)

func drawMenu(menu []string) {
	fmt.Println()
	for index, value := range menu {
		fmt.Printf("%v) %v\n", index, value)
	}
	fmt.Print("\nChoose what program to run (int): ")
}

func main() {
	for {
		pNames := []string{"Exit", "Draw Rectangle", "Second highest Int", "My Pi",
			"Word Counter", "Draw Matrix"}

		p := [...]func(){
			func() { os.Exit(0) }, lib.DrawRec, lib.Second, lib.MyPi,
			lib.WordCount, lib.MatrixDraw,
		}

		drawMenu(pNames)
		var userOption int
		fmt.Scan(&userOption)

		if userOption >= 0 && userOption < len(p) {
			p[userOption]()
		} else {
			fmt.Println("Invalid option")
		}
		continue
	}
}
