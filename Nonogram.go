package CodeCamp2018

type Mark int

const (
	Fill Mark = 0
	Cross Mark = 1
)

type Move struct {
	X int
	Y int
	Mark Mark
}

type Board struct {
	ColumnCount int
	RowCount int
	BoardMarks [][]Mark
	ColumnHints [][]int
	RowHints [][]int
}


