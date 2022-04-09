package guess

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Result struct {
	Name  string
	Score int
}

func makeJSON(games []Result) string {
	fmt.Println(games)
	sort.Slice(games, func(i, j int) bool {
		if games[i].Score != games[j].Score {
			return games[i].Score < games[j].Score
		}
		return games[i].Name < games[j].Name
	})

	r_json, _ := json.MarshalIndent(games, "", "  ")
	return string(r_json)
}

func newGame(poziom func(), games []Result) bool {
	var newGame string
NEWGAME:
	fmt.Println(time.Now())
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
