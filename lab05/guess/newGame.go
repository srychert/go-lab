package guess

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Result struct {
	Name  string
	Score int
}

// json of all games from session
func makeJSON(games []Result) string {
	g := sortGames(games)

	r_json, _ := json.MarshalIndent(g, "", "  ")
	return string(r_json)
}

func newGame(games []Result, results map[Date][]Result) bool {
	var newGame string
NEWGAME:
	fmt.Print("Gramy jeszcze raz? [T/N]: ")
	fmt.Scanln(&newGame)
	if strings.ToUpper(newGame) == "T" {
		return true
	} else if strings.ToUpper(newGame) == "N" {
		results_string := makeJSON(games)
		fmt.Println(results_string)
		saveResults(results)
		return false
	} else {
		goto NEWGAME
	}
}
