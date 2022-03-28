package main

import (
	"flag"
	"fmt"
)

func main() {
	day := flag.Int("d", 0, "birthDay")
	name := flag.String("n", "", "First letter of name")
	surname := flag.String("s", "", "First letter of surname")
	flag.Parse()
	var nameLetter string
	var surnameLetter string
	for *day <= 0 || *name == "" || *surname == "" {
		if *day <= 0 {
			fmt.Print("Specify day: ")
			fmt.Scanf("%d\n", day)
		}

		if *name == "" {
			fmt.Print("Type name: ")
			fmt.Scanf("%s\n", name)
		}

		if *surname == "" {
			fmt.Print("Type surname: ")
			fmt.Scanf("%s\n", surname)
		}
	}

	nameLetter = (*name)[:1]
	surnameLetter = (*surname)[:1]

	fmt.Println(*day, nameLetter, surnameLetter)
	fmt.Println(*day, *name, *surname)

	// daysMap :=

}
