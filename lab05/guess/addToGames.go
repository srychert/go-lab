package guess

import "fmt"

func addToGames(games Games, score int) Games {
	fmt.Println("Brawo wygrałeś!")
	var name string
	fmt.Print("Podaj imie: ")
	fmt.Scanln(&name)
	games.names = append(games.names, name)
	games.scores = append(games.scores, score)
	return games
}
