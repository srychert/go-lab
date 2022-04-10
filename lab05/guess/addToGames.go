package guess

import (
	"fmt"
)

func addToGames(games []Result, score int, results *map[Date][]Result) []Result {
	fmt.Println("Brawo wygrałeś!")
	var name string
	fmt.Println("Podaj imie:")
	fmt.Scanln(&name)
	// msg personal or global best; returns true if new entry in map (date)
	n := checkScore(name, score, results)
	// add to results
	r := Result{Name: name, Score: score}
	games = append(games, r)
	d := todayDate()
	if n {
		// filter for best score for each player and then sort by score
		fg := filterGames(games)
		(*results)[d] = sortGames(fg)
	} else {
		// add to existing games and then filter and sort
		(*results)[d] = append((*results)[d], games...)
		fg := filterGames((*results)[d])
		(*results)[d] = sortGames(fg)
	}
	return games
}
