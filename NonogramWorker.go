package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

type Worker struct {
	Id       int
	Board    Board
	MovesIn  chan Move
	MovesOut chan Move
	Tasks    []Method

	WaitGroup *sync.WaitGroup
}

func (w Worker) Solve() (bool, error) {
	defer w.WaitGroup.Done()

	for _, val := range w.Tasks {
		switch val {
		case Boxes:
			// w.SolveByBoxes()
		case Spaces:
			// w.SolveBySpaces()
		case Forcing:
			// w.SolveByForcing()
		case Glue:
			w.SolveByGlue()
		case Joining:
			// w.SolveByJoining()
		case Splitting:
			// w.SolveBySplitting()
		case Punctuating:
			// w.SolveByPunctuating()
		case Mercury:
			// w.SolveByMercury()
		default:
			return false, errors.New("worker error: Unknown task assigned")
		}
	}

	fmt.Printf("Worker[%d] is done working\n", w.Id)
	return true, nil
}

func (w Worker) LeftFill(hints []int, length int) (A []Mark) {
	A = make([]Mark, length)
	current := 0
	for i := 0; i < len(hints); i++ {
		for j := 0; j < hints[i]; j, current = j+1, current+1 {
			A[current] = Fill
		}
		current++
	}
	return
}

func (w Worker) RightFill(hints []int, length int) (A []Mark) {
	A = make([]Mark, length)
	current := length - 1
	for i := len(hints) - 1; i >= 0; i-- {
		for j, stop := current, current-hints[i]; j > stop; j, current = j-1, current-1 {
			A[current] = Fill
		}
		current--
	}
	return
}

func (w Worker) BoxOverlap(A []Mark, B []Mark, length int) (C []Mark) {
	C = make([]Mark, length)
	for i := 0; i < length; i++ {
		if A[i] == B[i] && A[i] == Fill {
			C[i] = Fill
		}
	}
	return
}

func (w Worker) SolveByBoxes() {
	fmt.Printf("Worker[%d] working on Boxes\n", w.Id)

	// Run on rows
	for i := 0; i < w.Board.ColumnCount; i++ {
		A := w.LeftFill(w.Board.RowHints[i], w.Board.ColumnCount)
		//fmt.Printf("Left: %+v\n", A)
		B := w.RightFill(w.Board.RowHints[i], w.Board.ColumnCount)
		//fmt.Printf("Right: %+v\n", B)
		C := w.BoxOverlap(A, B, w.Board.ColumnCount)
		//fmt.Printf("Overlap: %+v\n", C)
		for j := 0; j < w.Board.ColumnCount; j++ {
			if C[j] == Fill {
				w.MovesOut <- Move{
					WorkerId: w.Id,
					X:        j,
					Y:        i,
					Mark:     Fill,
				}
			}
		}
	}

	// Run on columns
	for i := 0; i < w.Board.RowCount; i++ {
		A := w.LeftFill(w.Board.ColumnHints[i], w.Board.RowCount)
		//fmt.Printf("Top: %+v\n", A)
		B := w.RightFill(w.Board.ColumnHints[i], w.Board.RowCount)
		//fmt.Printf("Bottom: %+v\n", B)
		C := w.BoxOverlap(A, B, w.Board.RowCount)
		//fmt.Printf("Overlap: %+v\n", C)
		for j := 0; j < w.Board.RowCount; j++ {
			if C[j] == Fill {
				w.MovesOut <- Move{
					WorkerId: w.Id,
					X:        i,
					Y:        j,
					Mark:     Fill,
				}
			}
		}
	}
}

func (w Worker) SolveBySpaces() {
	fmt.Printf("Worker[%d] working on Spaces\n", w.Id)
	w.RandomMoves()
}

func (w Worker) SolveByForcing() {
	fmt.Printf("Worker[%d] working on Forcing\n", w.Id)
	for row, hints := range w.Board.RowHints {
		chunk, offset := getRowChunk(w.Board.BoardMarks[row])
		for i, hint := range hints {
			if len(chunk) <= i {
				continue
			}
			diff := chunk[i] - hint
			if float64(diff) < float64(chunk[i])/2 {
				//We gota match boys!
				start := offset[i]
				offset := diff
				fillIn := hint - diff
				for k := 0; k < fillIn; k++ {
					//fmt.Printf("fill in X:%d, Y:%d\n", start+offset+k, row)
					w.MovesOut <- Move{
						WorkerId: w.Id,
						X:        start + offset + k,
						Y:        row,
						Mark:     Fill,
					}
				}
			}
		}
	}

	for col, hints := range w.Board.ColumnHints {
		chunk, offset := getColumnChunk(w.Board.BoardMarks, col)
		for i, hint := range hints {
			if len(chunk) <= i {
				continue
			}
			diff := chunk[i] - hint
			if float64(diff) < float64(chunk[i])/2 {
				//We got a macth boys!
				start := offset[i]
				offset := diff
				fillIn := hint - diff
				for k := 0; k < fillIn; k++ {
					//fmt.Printf("fill in X:%d, Y:%d\n", col, start+offset+k)
					w.MovesOut <- Move{
						WorkerId: w.Id,
						X:        col,
						Y:        start + offset + k,
						Mark:     Fill,
					}
				}
			}
		}
	}
}

func getRowChunk(row []Mark) ([]int, []int) {
	count := 0
	ret := []int{}
	offset := []int{}
	for i, cell := range row {
		if cell == Empty {
			count++
		} else if count > 0 {
			ret = append(ret, count)
			offset = append(offset, i-count)
			count = 0
		}
	}
	if count > 0 {
		ret = append(ret, count)
		offset = append(offset, len(row)-count)
	}
	//fmt.Printf("Row Chunks: %d, %d\n", ret, offset)
	return ret, offset
}

func getColumnChunk(column [][]Mark, col int) ([]int, []int) {
	count := 0
	ret := []int{}
	offset := []int{}
	for r := 0; r < len(column[0]); r++ {
		if column[r][col] == Empty {
			count++
		} else if count > 0 {
			ret = append(ret, count)
			offset = append(ret, r)
			count = 0
		}
	}
	if count > 0 {
		ret = append(ret, count)
		offset = append(offset, len(column[0])-count)
	}
	//fmt.Printf("Col Chunks: %d, %d\n", ret, offset)
	return ret, offset
}

func getChunkRow(m []Mark, i int) (int, int) {
	start := false
	l := 0
	index := 0
	for k, mark := range m {
		if start && mark != Fill {
			if index == i {
				return l, k - l
			}
			index++
			start = false
			l = 0
		}
		if mark == Fill {
			if !start {
				start = true
			}
		}
		if start {
			l++
		}
	}
	if start {
		return l, len(m) - l
	}
	return 0, 0
}

func getChunkCol(m [][]Mark, col, i int) (int, int) {
	start := false
	l := 0
	index := 0
	for k, m := range m {
		mark := m[col]
		if start && mark != 1 {
			if index == i {
				return l, k - l
			}
			index++
			start = false
			l = 0
		}
		if mark == 1 {
			if !start {
				start = true
			}
		}
		if start {
			l++
		}
	}
	if start {
		return l, len(m) - l
	}
	return 0, 0
}

func (w Worker) SolveByGlue() {
	for r, row := range w.Board.BoardMarks {
		farLeft := make([]Mark, len(row))
		farRight := make([]Mark, len(row))
		for i, hint := range w.Board.RowHints[r] {
			l, start := getChunkRow(row, i)
			Logf("b", "1: %d, %d, %v, %d", l, start, row, i)
			if l > hint {
				continue
			}
			end := start + l
			leftStart := end - hint
			rightEnd := start + hint
			if leftStart < 0 {
				leftStart = 0
			}
			for b := leftStart; b < leftStart+hint; b++ {
				farLeft[b] = Fill
			}
			if rightEnd > len(row)-1 {
				rightEnd = len(row) - 1
			}
			Logf("b", "2: %d, %d", leftStart, rightEnd)
			for j := rightEnd; j > rightEnd-hint; j-- {
				Logf("b", "%d", j)
				farRight[j] = Fill
			}
		}

		Logf("b", "fill\n%v\n%v", farLeft, farRight)
		for i, leftFill := range farLeft {
			if leftFill == farRight[i] && leftFill == Fill {
				Logf("b", "X: %d, Y: %d", i, row)
				w.MovesOut <- Move{
					WorkerId: w.Id,
					X:        i,
					Y:        r,
					Mark:     Fill,
				}
			}
		}
	}

	return
	// for row, hints := range w.Board.RowHints {
	// 	for r, hint := range hints {
	// 		chunks, offsets := getRowChunk(w.Board.BoardMarks[r])
	// 		//Prepend a zero cause go is dumb
	// 		offsets = flipArray(offsets)
	// 		offsets = append(offsets, 0)
	// 		offsets = flipArray(offsets)
	// 		for i := offsets[r], i <//i, row := range w.Board.BoardMarks[r] {

	// 		}
	// 	}
	// }
	left, right := false, false
	for row, hints := range w.Board.RowHints {
		for _, hint := range hints {
			farLeft := make([]Mark, len(w.Board.BoardMarks[row]))
			farRight := make([]Mark, len(w.Board.BoardMarks[row]))
			left, right = false, false
			for m, mark := range w.Board.BoardMarks[row] {
				Log("b", "here")
				if mark == Fill {
					for k := m; k < m+hint; k++ {
						if m-1 >= 0 && w.Board.BoardMarks[row][m-1] == Fill {
							break
						}
						if m+hint-1 < len(w.Board.BoardMarks[row]) {
							break
						}
						Log("b", "far right")
						if len(farRight) < m+hint {
							continue
						}
						farRight[k] = Fill
						left = true
					}
					if len(w.Board.BoardMarks[row]) > m+1 && w.Board.BoardMarks[row][m+1] == Fill {
						continue
					}
					for j := m - hint; j < m; j++ {
						Logf("b", "far left %v, m: %d, hint: %d", w.Board.BoardMarks[row], m, hint)
						if j < 0 {
							continue
						}
						farLeft[j] = Fill
						right = true
					}
				}
			}
			if left && right {
				Logf("b", "fill\n%v\n%v", farLeft, farRight)
				for i, leftFill := range farLeft {
					if leftFill == farRight[i] && leftFill == Fill {
						Logf("b", "X: %d, Y: %d", i, row)
						w.MovesOut <- Move{
							WorkerId: w.Id,
							X:        i,
							Y:        row,
							Mark:     Fill,
						}
					}
				}
			}
		}
	}
	Log("b", "did rows")

	//Now do the columns
	top, bot := false, false
	for col, hints := range w.Board.ColumnHints {
		for _, hint := range hints {
			farTop := make([]Mark, len(w.Board.BoardMarks))
			farBot := make([]Mark, len(w.Board.BoardMarks))
			top, bot = false, false
			for m, mark := range w.Board.BoardMarks {
				if mark[col] == Fill {
					for k := m; k < m+hint; k++ {
						if len(farTop) < m+hint {
							continue
						}
						farTop[k] = Fill
						left = true
					}
					if len(w.Board.BoardMarks) > m+1 && w.Board.BoardMarks[m+1][col] == Fill {
						continue
					}
					for j := m; j > m-hint; j-- {
						farBot[j] = Fill
						right = true
					}
				}
			}
			if top && bot {
				for i, botFill := range farBot {
					if botFill == farTop[i] && botFill == Fill {
						w.MovesOut <- Move{
							WorkerId: w.Id,
							X:        col,
							Y:        i,
							Mark:     Fill,
						}
					}
				}
			}
		}
	}
	Log("b", "did columns")
}

func (w Worker) SolveByJoining() {
	fmt.Printf("Worker[%d] working on Joining\n", w.Id)
	w.RandomMoves()
}

func (w Worker) SolveBySplitting() {
	fmt.Printf("Worker[%d] working on Splitting\n", w.Id)
	w.RandomMoves()
}

func (w Worker) SolveByPunctuating() {
	fmt.Printf("Worker[%d] working on Punctuating\n", w.Id)
	w.RandomMoves()
}

func (w Worker) SolveByMercury() {
	fmt.Printf("Worker[%d] working on Mercury\n", w.Id)
	w.RandomMoves()
}

func (w Worker) RandomMoves() {
	for i := 0; i < rand.Int()%100; i++ {
		w.MovesOut <- Move{
			WorkerId: w.Id,
			X:        rand.Int() % w.Board.ColumnCount,
			Y:        rand.Int() % w.Board.RowCount,
			Mark:     Mark(rand.Int() % int(MarkCount)),
		}
	}
}
