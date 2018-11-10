package main

import (
	"github.com/dennwc/dom/svg"
)

func drawUI(rowHints, columnHints [][]int) chan Move {
	w, h := 1409, 773

	doc := svg.NewFullscreen()
	root := doc.NewG()

	size := 500
	cols := len(columnHints)
	rows := len(rowHints)
	board := NewUIBoard(root, w/2-size/2, h/2-(5*(size/4))/2, size, 5*(size/4), cols, rows, columnHints, rowHints)
	// for j := 0; j < cols; j++ {
	// 	for i := 0; i < rows; i++ {
	// 		if rand.Float64() > 0.5 {
	// 			board.UpdateCoord(Cross, i, j)
	// 		} else {
	// 			board.UpdateCoord(Fill, i, j)
	// 		}
	// 	}
	// }

	//Create a channel to listen for moves on
	channel := make(chan Move, WorkerCount)
	go func() {
		for {
			move := <-channel
			board.UpdateCoord(move.Mark, move.X, move.Y)
		}
	}()
	return channel
}
