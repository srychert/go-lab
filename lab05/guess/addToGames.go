package guess

import (
	"fmt"
	"sort"
)

func sortGames(games []Result) []Result {
	sort.Slice(games, func(i, j int) bool {
		if games[i].Score != games[j].Score {
			return games[i].Score < games[j].Score
		}
		return games[i].Name < games[j].Name
	})

	return games
}

func addToGames(games []Result, score int, results *map[Date][]Result) []Result {
	fmt.Println("Brawo wygrałeś!")
	var name string
	fmt.Print("Podaj imie: ")
	fmt.Scanln(&name)
	// msg personal or global best; returns if new entry in map (date)
	n := checkScore(name, score, results)
	// add to results
	r := Result{Name: name, Score: score}
	games = append(games, r)
	d := todayDate()
	if n {
		// add only best?
		fg := filterGames(games)
		(*results)[d] = sortGames(fg)
	} else {
		(*results)[d] = append((*results)[d], games...)
		fg := filterGames((*results)[d])
		(*results)[d] = sortGames(fg)
	}
	return games
}
