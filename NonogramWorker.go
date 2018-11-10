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
	Done chan bool
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
				//w.SolveByForcing()
			case Glue:
				// w.SolveByGlue()
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
		case move, ok := <- w.MovesIn:
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
		S += len(w.Board.RowHints[i])-1
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
					w.MovesOut <- Move {
						WorkerId: w.Id,
						X: k,
						Y: i,
						Mark: Fill,
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
		S += len(w.Board.ColumnHints[i])-1
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
					w.MovesOut <- Move {
						WorkerId: w.Id,
						X: i,
						Y: k,
						Mark: Fill,
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

func (w Worker) SolveByGlue() {
	fmt.Printf("Worker[%d] working on Glue\n", w.Id)
	w.RandomMoves()
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
