package lib

import (
	"fmt"
	"unicode/utf8"
)


func printMatrix(matrix [][]int) {
	fmt.Println("Matrix:")
	fmt.Println("[")
	last := matrix[len(matrix) -1][len(matrix) -1]
	last_str := fmt.Sprintf("%v", last)
	lenOfEle := utf8.RuneCountInString(last_str) + 1
// 	Printing every slice inside matrix
	for _, value := range matrix{
	    for _, v2 := range value{
	        fmt.Printf("%*d", lenOfEle, v2)
	    }
		fmt.Println()
	}
	fmt.Print("]")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MatrixDraw() {
	var maxValue int
	fmt.Print("Give value greater than 0 : ")
	_, err := fmt.Scan(&maxValue)
	// row and column = len maxValue * 2 - 1
	n := maxValue*2 - 1
	// x,y of center
	mid := maxValue - 1

	if err == nil && maxValue > 0 {
		// make empty matrix
		matrix := make([][]int, n)
		for i := range matrix {
			// make empty row
			row := make([]int, n)
			for j := range row {
				// set every value in row
				value := max(abs(mid-i), abs(mid-j)) + 1
				row[j] = value
			}
			matrix[i] = row
		}
		printMatrix(matrix)
	} else {
		fmt.Println("Wrong value!")
	}

}
