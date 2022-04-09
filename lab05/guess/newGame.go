package guess

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Result struct {
	Name  string
	Score int
}

func makeJSON(games []Result) string {
	sort.Slice(games, func(i, j int) bool {
		if games[i].Score != games[j].Score {
			return games[i].Score < games[j].Score
		}
		return games[i].Name < games[j].Name
	})

	r_json, _ := json.MarshalIndent(games, "", "  ")
	return string(r_json)
}

func newGame(games []Result) bool {
	var newGame string
NEWGAME:
	fmt.Print("Gramy jeszcze raz? [T/N]: ")
	fmt.Scanln(&newGame)
	if strings.ToUpper(newGame) == "T" {
		return true
	} else if strings.ToUpper(newGame) == "N" {
		results := makeJSON(games)
		fmt.Println(results)
		return false
	} else {
		goto NEWGAME
	}
}
