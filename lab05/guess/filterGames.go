package guess

func filterGames(games []Result) []Result {
	mfg := make(map[string]Result)
	var fg []Result

	for _, v := range games {
		if s, ok := mfg[v.Name]; ok {
			if v.Score < s.Score {
				mfg[v.Name] = v
			}
		} else {
			mfg[v.Name] = v
		}
	}

	for _, v := range mfg {
		fg = append(fg, v)
	}

	return fg
}
