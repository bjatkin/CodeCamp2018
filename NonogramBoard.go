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
	b.BoardMarks = make([][]Mark, rowCount)
	for r := 0; r < rowCount; r++ {
		b.BoardMarks[r] = make([]Mark, columnCount)
	}
	return
}

func (b Board) MarkCell(m Move) error {
	if m.X > b.ColumnCount {
		return errors.New("board error: X value larger than ColumnCount")
	}
	if m.X < 0 {
		return errors.New("board error: X value less than 0")
	}
	if m.Y > b.RowCount {
		return errors.New("board error: Y value larger than RowCount")
	}
	if m.Y < 0 {
		return errors.New("board error: Y value less than 0")
	}
	b.BoardMarks[m.Y][m.X] = m.Mark
	return nil
}

func (b Board) CellValue(x, y int) Mark {
	return b.BoardMarks[x][y]
}
