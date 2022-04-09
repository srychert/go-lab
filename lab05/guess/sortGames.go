package guess

import "sort"

func sortGames(games []Result) []Result {
	sort.Slice(games, func(i, j int) bool {
		if games[i].Score != games[j].Score {
			return games[i].Score < games[j].Score
		}
		return games[i].Name < games[j].Name
	})

	return games
}
