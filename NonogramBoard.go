package CodeCamp2018

import (
	"errors"
	"github.com/dennwc/dom/svg"
)

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

func (b Board) MarkCell(m Move) (bool, error) {
	if m.X > b.ColumnCount {
		return false, errors.New("Board Error: X value larger than ColumnCount")
	}
	if m.Y > b.RowCount {
		return false, errors.New("Board Error: Y value larger than Rowcount")
	}
	b.BoardMarks[m.X][m.Y] = m.Mark
	return true, nil
}