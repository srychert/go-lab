package main

import (
	"fmt"
	lib "simple-programs/lib"
)

func main() {
	for {
		packageNames := []string{"Draw Rectangle", "Second highest Int", "My Pi"}
		var userOption int

		fmt.Println("Choose what program to run (int): ")
		for index, value := range packageNames {
			fmt.Printf("%v) %v\n", index+1, value)
		}

		fmt.Scan(&userOption)

		switch userOption {
		case 1:
			lib.DrawRec()
		case 2:
			lib.Second()
		case 3:
			lib.MyPi()
		default:
			fmt.Println("Wrong option")
			continue
		}
		break
	}
}
