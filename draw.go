package main

import (
	"time"

	"github.com/dennwc/dom/svg"
)

func drawUI(size int, rowHints, columnHints [][]int) UIBoard {
	w, h := 1409, 773

	doc := svg.NewFullscreen()
	root := doc.NewG()

	cols := len(columnHints)
	rows := len(rowHints)
	board := NewUIBoard(root, w/2-(5*(size/4))/2, h/2-(size/2), size, 5*(size/4), cols, rows, columnHints, rowHints)

	//Fade the GUI in
	go func() {
		interval := time.Millisecond * 10
		for i := 0.0; i < 1; i += 0.01 {
			board.root.SetAttribute("opacity", i)
			time.Sleep(interval)
		}
	}()

	return board
}
