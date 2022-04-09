package guess

import "fmt"

func checkScore(name string, score int, results *map[Date][]Result) bool {
	var userScores []int
	var best int
	new := true

	date := todayDate()

	for d, r := range *results {
		if d.toString() == date.toString() {
			new = false
		}
		best = r[0].Score
		for _, v := range r {
			if v.Score < best {
				best = v.Score
			}
			if v.Name == name {
				userScores = append(userScores, v.Score)
			}
		}
	}

	var personal int

	for i, e := range userScores {
		if i == 0 || e < personal {
			personal = e
		}
	}

	if personal == 0 || score < personal {
		fmt.Println("Nowy rekord osobisty!")
	}

	if best == 0 || score < best {
		fmt.Println("Nowy rekord rankingu!")
	}

	return new

}
