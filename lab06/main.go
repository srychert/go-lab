package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

// size of the board
const dim = 10

type Position struct {
	x     int
	y     int
	name  string
	count int
}

func (p *Position) update(x, y int) {
	p.x, p.y = x, y

}

func (p *Position) changeCount(c int) {
	p.count = c

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

		if name == "leaf" {
			pick.count = 1
		}

		*indices = append((*indices)[:randomIndex], (*indices)[randomIndex+1:]...)

		positions = append(positions, pick)

	}
	return positions, nil
}

func drawBoard(elements ...[]Position) {
	board := [dim][dim]string{}

	for i, row := range board {
		for j := range row {
			board[i][j] = "--"
		}
	}

	for _, ele := range elements {
		for _, v := range ele {
			switch v.name {
			case "leaf":
				if board[v.x][v.y] != "--" {
					board[v.x][v.y] = board[v.x][v.y] + fmt.Sprintf("l%v", v.count)
				} else {
					l := fmt.Sprintf("l%v", v.count)
					l = color.Colorize(color.Green, l)
					board[v.x][v.y] = l
				}
			case "ant":
				a := fmt.Sprintf("a%v", v.count)
				a = color.Colorize(color.Red, a)
				if board[v.x][v.y] != "--" {
					board[v.x][v.y] = board[v.x][v.y] + a
				} else {
					board[v.x][v.y] = a
				}
			}
		}
	}

	for _, v := range board {
		fmt.Println(v)
	}
	fmt.Println()
}

func updatePosition(e *Position) {
	positions := []Position{
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: -1, y: 0},
	}

	newCord := rand.Intn(len(positions))
	x, y := e.x+positions[newCord].x, e.y+positions[newCord].y

	// if new position is outside the board move to the opposite edge
	if x < 0 {
		x = dim - 1
	}

	if x > dim-1 {
		x = 0
	}

	if y < 0 {
		y = dim - 1
	}

	if y > dim-1 {
		y = 0
	}

	e.update(x, y)
}

func move(e *Position, l []Position) []Position {
	for i, v := range l {
		if v.x == e.x && v.y == e.y {
			// if ant carries an leaf drop it
			if e.count > 0 {
				e.changeCount(e.count - 1)
				v.changeCount(v.count + 1)
				l[i] = v
				return l
			}

			// ant picks up a leaf
			e.changeCount(e.count + 1)
			if v.count > 1 {
				v.changeCount(v.count - 1)
				l[i] = v
				return l
			} else {
				l[i] = l[len(l)-1]
				return l[:len(l)-1]
			}
		}
	}
	return l
}

func main() {
	rand.Seed(time.Now().UnixNano())
	antP := flag.Int("a", 10, "percentage of ants")
	leafP := flag.Int("l", 90, "percentage of leafs")
	iter := flag.Int("i", 1000, "number of iterations")
	draw := flag.Bool("draw", true, "draw every frame")
	flag.Parse()

	board := [dim][dim]string{}
	indices := []Position{}

	for i, row := range board {
		for j := range row {
			indices = append(indices, Position{x: i, y: j, name: "empty", count: 0})
		}
	}

	leafs, errL := genPositions(&indices, "leaf", dim*dim**leafP/100)
	ants, errA := genPositions(&indices, "ant", dim*dim**antP/100)

	if errL != nil || errA != nil {
		log.Fatal(errL, errA)
	}

	// initial state
	drawBoard(leafs, ants)

	// main loop
	for i := 0; i < *iter; i++ {
		for i, a := range ants {
			updatePosition(&a)
			ants[i] = a
		}

		for i, a := range ants {
			leafs = move(&a, leafs)
			ants[i] = a
		}

		if *draw {
			fmt.Println("ants:", len(ants), "leafs:", len(leafs))
			drawBoard(leafs, ants)
		} else {
			if i == *iter-2 {
				fmt.Println("ants:", len(ants), "leafs:", len(leafs))
				drawBoard(leafs, ants)
			}
		}
	}

	// count how many ants carry a leaf at the end
	cnt := 0
	for _, a := range ants {
		if a.count != 0 {
			cnt += 1
		}
	}

	fmt.Printf("%v ants carry a leaf\n\n", cnt)
	fmt.Printf("%v\n\n%v", leafs, ants)

}
