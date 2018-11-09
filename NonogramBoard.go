package main

type Board struct {
	ColumnCount int
	RowCount    int
	BoardMarks  [][]Mark
	ColumnHints [][]int
	RowHints    [][]int
}
