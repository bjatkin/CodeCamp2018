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
	Done     chan bool
	Tasks    []Method

	WaitGroup *sync.WaitGroup
}

func (w Worker) Solve() (bool, error) {
	defer w.WaitGroup.Done()

	go w.ProcessInbox()

	for {
		for _, val := range w.Tasks {
			switch val {
			case Boxes:
				w.SolveByBoxes()
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

		select {
		case <-w.Done:
			break
		default:
		}
	}

	fmt.Printf("Worker[%d] is done working\n", w.Id)
	return true, nil
}

func (w Worker) ProcessInbox() {
	for {
		select {
		case move, ok := <-w.MovesIn:
			err := w.Board.MarkCell(move)
			if err != nil {
				fmt.Printf("%+v\n", err)
				fmt.Printf("%+v\n", move)
			}
			if !ok {
				break
			}
		}
	}
}

func (w Worker) SolveByBoxes() {
	fmt.Printf("Worker[%d] working on Boxes\n", w.Id)

	// Run on rows
	for i := 0; i < w.Board.ColumnCount; i++ {
		S := 0
		// H is the highest value of Hints
		H := 0
		for _, value := range w.Board.RowHints[i] {
			S += value
			if value > H {
				H = value
			}
		}
		// S is the total value of the hints plus the spaces between them
		S += len(w.Board.RowHints[i]) - 1
		//fmt.Printf("S = %d = %d + %d\n", S, S - len(w.Board.RowHints[i])+1, len(w.Board.RowHints[i])-1)
		// D is the difference between the sum and the total length
		D := w.Board.ColumnCount - S
		//fmt.Printf("D = %d = %d - %d\n", D, w.Board.ColumnCount, S)

		E := 0
		for j := 0; j < len(w.Board.RowHints[i]); j++ {
			// F is the number of cells that can be filled
			F := w.Board.RowHints[i][j] - D
			//fmt.Printf("F = %d = %d - %d\n", F, w.Board.RowHints[i][j], D)

			E += w.Board.RowHints[i][j]
			//fmt.Printf("E = %d = %d + %d\n", E, E - w.Board.RowHints[i][j], w.Board.RowHints[i][j])

			if F > 0 {
				// Index to begin marking as filled
				B := E - F
				//fmt.Printf("B = %d = %d - %d\n", B, B + F, F)

				// Fill all cells between the values B and E
				for k := B; k < E; k++ {
					w.MovesOut <- Move{
						WorkerId: w.Id,
						MethodId: Boxes,
						X:        k,
						Y:        i,
						Mark:     Fill,
					}
				}
			}

			// Increment for space
			E += 1
			//fmt.Printf("E = %d = %d + 1\n", E, E + 1)
		}
		//fmt.Printf("\n")
	}

	// Run on columns
	for i := 0; i < w.Board.RowCount; i++ {
		S := 0
		// H is the highest value of Hints
		H := 0
		for _, value := range w.Board.ColumnHints[i] {
			S += value
			if value > H {
				H = value
			}
		}
		// S is the total value of the hints plus the spaces between them
		S += len(w.Board.ColumnHints[i]) - 1
		//fmt.Printf("S = %d = %d + %d\n", S, S - len(w.Board.ColumnHints[i])+1, len(w.Board.ColumnHints[i])-1)
		// D is the difference between the sum and the total length
		D := w.Board.RowCount - S
		//fmt.Printf("D = %d = %d - %d\n", D, w.Board.RowCount, S)

		E := 0
		for j := 0; j < len(w.Board.ColumnHints[i]); j++ {
			// F is the number of cells that can be filled
			F := w.Board.ColumnHints[i][j] - D
			//fmt.Printf("F = %d = %d - %d\n", F, w.Board.ColumnHints[i][j], D)

			E += w.Board.ColumnHints[i][j]
			//fmt.Printf("E = %d = %d + %d\n", E, E - w.Board.ColumnHints[i][j], w.Board.ColumnHints[i][j])

			if F > 0 {
				// Index to begin marking as filled
				B := E - F
				//fmt.Printf("B = %d = %d - %d\n", B, B + F, F)

				// Fill all cells between the values B and E
				for k := B; k < E; k++ {
					w.MovesOut <- Move{
						WorkerId: w.Id,
						MethodId: Boxes,
						X:        i,
						Y:        k,
						Mark:     Fill,
					}
				}
			}

			// Increment for space
			E += 1
			//fmt.Printf("E = %d = %d + 1\n", E, E + 1)
		}
		//fmt.Printf("\n")
	}
}

func (w Worker) SolveBySpaces() {
	fmt.Printf("Worker[%d] working on Spaces\n", w.Id)
	w.RandomMoves()
}

func (w Worker) SolveByForcing() {
	// fmt.Printf("Worker[%d] working on Forcing\n", w.Id)
	for row, hints := range w.Board.RowHints {
		chunk, offset := getRowChunk(w.Board.BoardMarks[row])
		for i, hint := range hints {
			if len(chunk) <= i {
				continue
			}
			diff := hint - chunk[i]
			if float64(diff) < float64(chunk[i])/2 {
				//We gota match boys!
				start := offset[i]
				off := diff
				fillIn := hint - diff
				for k := 0; k < fillIn; k++ {
					fmt.Printf("fill in a X:%d, Y:%d\nchunk: %v, offset: %v\nfillIn: %d\nhint:%d\n", start+off+k-1, row, chunk, offset, fillIn, hint)
					w.MovesOut <- Move{
						WorkerId: w.Id,
						MethodId: Forcing,
						X:        start + off + k,
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
			diff := hint - chunk[i]
			if float64(diff) < float64(chunk[i])/2 {
				//We got a macth boys!
				start := offset[i]
				off := diff
				fillIn := hint - diff
				for k := 0; k < fillIn; k++ {
					fmt.Printf("fill in b X:%d, Y:%d\nchunk: %v\n offset: %v\n", col, start+off+k, chunk, offset)
					w.MovesOut <- Move{
						WorkerId: w.Id,
						MethodId: Forcing,
						X:        col,
						Y:        start + off + k,
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
	fmt.Printf("Worker[%d] working on Glue\n", w.Id)
	for r, row := range w.Board.BoardMarks {
		farLeft := make([]Mark, len(row))
		farRight := make([]Mark, len(row))
		for i, hint := range w.Board.RowHints[r] {
			_, prev := getChunkRow(row, i-1)
			l, start := getChunkRow(row, i)
			// Logf("b", "1: %d, %d, %v, %d", l, start, row, i)
			if l > hint || l == 0 {
				continue
			}
			end := start + l
			leftStart := end - hint
			rightEnd := start + hint - 1
			if i == 1 {
				// Logf("b", "1: %d, %d, %v, %d", l, start, row, i)
				// Logf("b", "2: %d, %d", leftStart, rightEnd)
			}
			for b := leftStart; b < leftStart+hint; b++ {
				if leftStart < 0 {
					// fmt.Printf("%d\n", b-leftStart)
					farLeft[b-leftStart] = Fill
					continue
				}

				if leftStart < prev {
					farLeft[b+(prev-leftStart)] = Fill
					continue
				}
				farLeft[b] = Fill
			}
			for j := rightEnd; j > rightEnd-hint; j-- {
				if rightEnd >= len(farRight) {
					// fmt.Printf("%d\n", j-(rightEnd-len(farRight))-1)
					farRight[j-(rightEnd-len(farRight))-1] = Fill
					continue
				}
				farRight[j] = Fill
			}
		}

		// Logf("b", "fill\n%v\n%v", farLeft, farRight)
		for i, leftFill := range farLeft {
			if leftFill == farRight[i] && leftFill == Fill {
				// Logf("b", "X: %d, Y: %d", i, row)
				w.MovesOut <- Move{
					WorkerId: w.Id,
					MethodId: Glue,
					X:        i,
					Y:        r,
					Mark:     Fill,
				}
			}
		}
	}

	for c := 0; c < len(w.Board.BoardMarks); c++ {
		farLeft := make([]Mark, len(w.Board.BoardMarks))
		farRight := make([]Mark, len(w.Board.BoardMarks))
		for i, hint := range w.Board.ColumnHints[c] {
			_, prev := getChunkCol(w.Board.BoardMarks, c, i-1)
			l, start := getChunkCol(w.Board.BoardMarks, c, i)
			// Logf("b", "1: %d, %d, %d", l, start, i)
			if l > hint || l == 0 {
				continue
			}
			end := start + l
			leftStart := end - hint
			rightEnd := start + hint
			for b := leftStart; b < leftStart+hint; b++ {
				if leftStart < 0 {
					// fmt.Printf("%d\n", b-leftStart)
					farLeft[b-leftStart] = Fill
					continue
				}

				if leftStart < prev {
					farLeft[b+(prev-leftStart)] = Fill
					continue
				}
				farLeft[b] = Fill
			}
			for j := rightEnd; j > rightEnd-hint; j-- {
				if rightEnd >= len(farRight) {
					// fmt.Printf("%d\n", j-(rightEnd-len(farRight))-1)
					farRight[j-(rightEnd-len(farRight))-1] = Fill
					continue
				}
				farRight[j] = Fill
			}
		}

		// Logf("b", "fill\n%v\n%v", farLeft, farRight)
		for i, leftFill := range farLeft {
			if leftFill == farRight[i] && leftFill == Fill {
				w.MovesOut <- Move{
					WorkerId: w.Id,
					MethodId: Glue,
					X:        c,
					Y:        i,
					Mark:     Fill,
				}
			}
		}
	}
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
