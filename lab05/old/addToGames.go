package old

import "fmt"

func addToGames(games []Result, score int) []Result {
	fmt.Println("Brawo wygrałeś!")
	var name string
	fmt.Print("Podaj imie: ")
	fmt.Scanln(&name)
	r := Result{Name: name, Score: score}
	games = append(games, r)
	return games
}
