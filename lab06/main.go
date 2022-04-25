package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const dim = 10

type Position struct {
	x    int
	y    int
	name string
}

func (p *Position) update(f func() (int, int)) {
	p.x, p.y = f()

}

func genPositions(indices *[]Position, name string, number int) ([]Position, error) {
	positions := []Position{}

	if len(*indices) < number {
		msg := fmt.Sprintf("Not enough positions to choose for %ss", name)
		return positions, errors.New(msg)
	}

	for i := 0; i < number; i++ {
		randomIndex := rand.Intn(len(*indices))
		pick := (*indices)[randomIndex]
		pick.name = name

		*indices = append((*indices)[:randomIndex], (*indices)[randomIndex+1:]...)

		positions = append(positions, pick)

	}
	return positions, nil
}

func drawBoard(board [dim][dim]string, elements ...[]Position) {
	for _, ele := range elements {
		for _, v := range ele {
			switch v.name {
			case "leaf":
				board[v.x][v.y] = "l"
			case "ant":
				board[v.x][v.y] = "a"
			case "empty":
				board[v.x][v.y] = "-"
			}
		}
	}
	for _, v := range board {
		fmt.Println(v)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	board := [dim][dim]string{}
	indices := []Position{}

	for i, row := range board {
		for j := range row {
			indices = append(indices, Position{x: i, y: j, name: "empty"})
		}
	}

	leafs, errL := genPositions(&indices, "leaf", 30)
	ants, errA := genPositions(&indices, "ant", 20)

	if errL != nil || errA != nil {
		log.Fatal(errL, errA)
	}

	// fmt.Printf("%v\n%v\n", leafs, ants)

	drawBoard(board, leafs, ants, indices)

}
