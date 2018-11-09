package CodeCamp2018

import "github.com/dennwc/dom/svg"

type Mark int

const (
	Fill  Mark = 0
	Cross Mark = 1
)

type Move struct {
	X    int
	Y    int
	Mark Mark
}

type Board struct {
	SVG         *svg.Rect
	ColumnCount int
	RowCount    int
	BoardMarks  [][]Mark
	ColumnHints [][]int
	RowHints    [][]int
}
