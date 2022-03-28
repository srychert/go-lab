package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	day := flag.Int("d", 0, "birthDay")
	name := flag.String("n", "", "First letter of name")
	surname := flag.String("s", "", "First letter of surname")
	flag.Parse()
	var nameLetter string
	var surnameLetter string
	for *day <= 0 || *day > 31 || *name == "" || *surname == "" {
		if *day <= 0 || *day > 31 {
			fmt.Print("Specify day in range[1, 31]: ")
			fmt.Scanf("%d\n", day)
		}

		if *name == "" {
			fmt.Print("Type name (at least one character): ")
			fmt.Scanf("%s\n", name)
		}

		if *surname == "" {
			fmt.Print("Type surname (at least one character): ")
			fmt.Scanf("%s\n", surname)
		}
	}

	nameLetter = strings.ToLower((*name)[:1])
	surnameLetter = strings.ToLower((*surname)[:1])

	daysList := [...]string{
		"",
		"Occupational Therapist", "Paralegal", "Social Worker",
		"Court Reporter", "Fitness Trainer", "Food Scientist",
		"Medical Secretary", "Musician", "Actuary",
		"Lawyer", "Recreational Therapist", "Event Planner",
		"Historian", "Referee", "Janitor",
		"Cost Estimator", "Construction Manager", "Respiratory Therapist",
		"Real Estate Agent", "Chef", "Sports Coach",
		"Computer Systems Analyst", "Teacher Assistant", "Middle School Teacher",
		"Physicist", "Desktop publisher", "Telemarketer",
		"Environmental scientist", "Customer Service Representative", "Psychologist",
		"Dancer",
	}

	namesMap := map[string]string{
		"a": "valuable", "b": "silent", "c": "sparkling",
		"d": "broad", "e": "spotless", "f": "itchy",
		"g": "ludicrous", "h": "domineering", "i": "hallowed",
		"j": "chilly", "k": "skillful", "l": "boring",
		"m": "lively", "n": "dangerous", "o": "nutritious",
		"p": "exotic", "q": "scientific", "r": "incompetent",
		"s": "mighty", "t": "hesitant", "u": "lovely",
		"v": "mindless", "w": "grandiose", "x": "unwritten",
		"y": "verdant", "z": "abaft",
	}

	var result string

	result = fmt.Sprintf("%s %s %s", daysList[*day],
		namesMap[nameLetter], namesMap[surnameLetter])
	fmt.Println(result)

}
