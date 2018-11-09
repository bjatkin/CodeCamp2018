package main

import "github.com/dennwc/dom/svg"

func drawUI() {
	w, h := 1409, 773

	doc := svg.NewFullscreen()
	root := doc.NewG()

	size := 500
	cols := 10
	rows := 10
	board := NewUIBoard(root, w/2-size/2, h/2-(5*(size/4))/2, size, 5*(size/4), cols, rows)
	board.UpdateCoord(Fill, 0, 0)
	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			board.UpdateCoord(Fill, i, j)
		}
	}
}
