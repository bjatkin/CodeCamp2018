package CodeCamp2018

type Board struct {
	SVG         *svg.Rect
	ColumnCount int
	RowCount    int
	BoardMarks  [][]Mark
	ColumnHints [][]int
	RowHints    [][]int
}
