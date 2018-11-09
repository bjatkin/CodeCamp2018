package main

import "github.com/dennwc/dom/svg"

func drawUI() {
	w, h := 1409, 773

	doc := svg.NewFullscreen()
	root := doc.NewG()

	size := 500
	board := NewUIBoard(root, w/2, h/2, size, 5*(size/4))
	board.AddUIBox(Fill, 10, 10)
}
