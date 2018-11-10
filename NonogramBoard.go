package main

import (
	"errors"
)

type Board struct {
	ColumnCount int
	RowCount    int
	BoardMarks  [][]Mark
	ColumnHints [][]int
	RowHints    [][]int
}

func NewBoard(rowCount, columnCount int, rowHints, columnHints [][]int) (b Board) {
	b = Board{}
	b.RowCount = rowCount
	b.ColumnCount = columnCount
	b.RowHints = rowHints
	b.ColumnHints = columnHints
	return
}

func (b Board) MarkCell(m Move) (bool, error) {
	if m.X > b.ColumnCount {
		return false, errors.New("board error: X value larger than ColumnCount")
	}
	if m.Y > b.RowCount {
		return false, errors.New("board error: Y value larger than RowCount")
	}
	b.BoardMarks[m.X][m.Y] = m.Mark
	return true, nil
}
