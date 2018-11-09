package CodeCamp2018

import "github.com/dennwc/dom/svg"

type Board struct {
	SVG         *svg.Rect
	ColumnCount int
	RowCount    int
	BoardMarks  [][]Mark
	ColumnHints [][]int
	RowHints    [][]int
}
