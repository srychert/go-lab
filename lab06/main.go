package main

import (
	"fmt"
	"math/rand"
	"time"
)

const dim = 10

type Position struct {
	x int
	y int
}

func (p *Position) update(f func() (int, int)) {
	p.x, p.y = f()

}

func checkIfUniqe(leaf Position, leafs *[]Position) bool {
	for _, v := range *leafs {
		if v == leaf {
			return false
		}
	}
	return true
}

func main() {
	board := [dim][dim]int{}
	leafs := []Position{}
	// ants := []Position{}
	rand.Seed(time.Now().Unix())

	for len(leafs) < dim {
		x := rand.Intn(dim)
		y := rand.Intn(dim)
		leaf := Position{x, y}
		if checkIfUniqe(leaf, &leafs) {
			leafs = append(leafs, leaf)
		}
	}

	for _, v := range leafs {
		board[v.x][v.y] = 1
	}

	for _, v := range board {
		fmt.Println(v)
	}

	fmt.Println(leafs)

	// p := Position{0, 0}
	// fmt.Println(p)
	// p.update(func() (int, int) { return 1, 1 })
	// fmt.Println(p)
	// fmt.Println("")

}
