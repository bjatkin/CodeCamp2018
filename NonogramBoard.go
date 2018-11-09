package CodeCamp2018

import "github.com/dennwc/dom/svg"

type Board struct {
	ColumnCount int
	RowCount    int
	BoardMarks  [][]Mark
	ColumnHints [][]int
	RowHints    [][]int
}

func newBoard(rowCount, columnCount int, rowHints, columnHints [][]int) (b *Board) {
	b.RowCount = rowCount
	b.ColumnCount = columnCount
	b.RowHints = rowHints
	b.ColumnHints = columnHints
	return
}

