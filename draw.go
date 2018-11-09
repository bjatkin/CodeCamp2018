package main

import "github.com/dennwc/dom/svg"

func drawUI() {
	w, h := 1409, 773

	doc := svg.NewFullscreen()
	root := doc.NewG()

	size := 500
	cols := 30
	rows := 30
	board := NewUIBoard(root, w/2-size/2, h/2-(5*(size/4))/2, size, 5*(size/4), cols, rows,
		[][]int{[]int{1, 2, 3}, []int{1, 2, 44, 5}, []int{1, 2}},
		[][]int{[]int{3, 4, 5}, []int{5, 6, 71}, []int{3, 5}})
	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			board.UpdateCoord(Fill, i, j)
		}
	}
}
