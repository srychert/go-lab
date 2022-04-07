package guess

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Games struct {
	names  []string
	scores []int
}

type Slice struct {
	sort.IntSlice
	idx []int
}

func (s Slice) Swap(i, j int) {
	s.IntSlice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func NewSlice(n ...int) *Slice {
	s := &Slice{IntSlice: sort.IntSlice(n), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

type Result struct {
	Name  string
	Score int
}

func makeResults(games Games) string {
	// make a map of best score for each name
	m := make(map[string]int)
	for i, name := range games.names {
		v, exist := m[name]
		if exist {
			if v > games.scores[i] {
				m[name] = games.scores[i]
			}
		} else {
			m[name] = games.scores[i]
		}
	}

	// slices of keys and values of map -> names and scores
	keys := make([]string, 0, len(m))
	values := make([]int, 0, len(m))
	for k, v := range m {
		values = append(values, v)
		keys = append(keys, k)
	}

	// sort scores and keep indices
	scores := NewSlice(values...)
	sort.Sort(scores)
	// slice of soreted results (asc)
	results := make([]Result, 0, len(m))
	for j, idx := range scores.idx {
		results = append(results, Result{Name: keys[idx], Score: scores.IntSlice[j]})
	}

	// convert to json
	r_json, _ := json.MarshalIndent(results, "", "  ")

	return string(r_json)
}

func newGame(poziom func(), games Games) bool {
	var newGame string
NEWGAME:
	fmt.Print("Gramy jeszcze raz? [T/N]: ")
	fmt.Scanln(&newGame)
	if strings.ToUpper(newGame) == "T" {
		return true
	} else if strings.ToUpper(newGame) == "N" {
		results := makeResults(games)
		fmt.Println(results)
		return false
	} else {
		goto NEWGAME
	}
}
