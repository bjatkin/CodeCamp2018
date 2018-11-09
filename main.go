package main

import (
	"fmt"

	"github.com/dennwc/dom/svg"
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
	fmt.Println("test")

	w, h := 1409, 773

	doc := svg.NewFullscreen()
	root := doc.NewG()

	size := 500
	board := NewUIBoard(root, w/2, h/2, size, 5*(size/4))
	board.AddUIBox(Fill, 10, 10)

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
