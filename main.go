package main

import (
	"fmt"
)

//All the colors you'll need
const (
	Pcolor = "#041630"
	Scolor = "#5a8bbc"
	Black  = "#000020"
	White  = "#f4f4ff"

	WorkerCount = 4
)

func main() {
	drawUI()

	rowHints := [][]int{
		{2, 2},
		{1, 1, 1},
		{0},
		{3},
		{1, 3},
	}
	columnHints := [][]int{
		{2, 1},
		{1},
		{1, 2},
		{1, 2},
		{2, 2},
	}
	mainBoard := NewBoard(5, 5, rowHints, columnHints)
	nonogramMaster := NewMaster(mainBoard)
	fmt.Print(nonogramMaster.Board.RowCount)
}
