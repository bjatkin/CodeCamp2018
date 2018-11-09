package main

import (
	"errors"
	"fmt"
)

type Worker struct {
	Id int
	Board *Board
	MovesIn <-chan Move
	MovesOut chan<- Move
	Tasks []Method
}

func (w Worker) Solve() (bool, error) {

	for _, val := range w.Tasks {
		switch val {
		case Boxes:
			w.SolveByBoxes()
		case Spaces:
			w.SolveBySpaces()
		case Forcing:
			w.SolveByForcing()
		case Glue:
			w.SolveByGlue()
		case Joining:
			w.SolveByJoining()
		case Splitting:
			w.SolveBySplitting()
		case Punctuating:
			w.SolveByPunctuating()
		case Mercury:
			w.SolveByMercury()
		default:
			return false, errors.New("worker error: Unknown task assigned")
		}
	}

	return true, nil
}

func (w Worker) SolveByBoxes() {
	fmt.Printf("Worker[%d] working on Boxes\n", w.Id)
}

func (w Worker) SolveBySpaces() {
	fmt.Printf("Worker[%d] working on Spaces\n", w.Id)
}

func (w Worker) SolveByForcing() {
	fmt.Printf("Worker[%d] working on Forcing\n", w.Id)
}

func (w Worker) SolveByGlue() {
	fmt.Printf("Worker[%d] working on Glue\n", w.Id)
}

func (w Worker) SolveByJoining() {
	fmt.Printf("Worker[%d] working on Joining\n", w.Id)
}

func (w Worker) SolveBySplitting() {
	fmt.Printf("Worker[%d] working on Splitting\n", w.Id)
}

func (w Worker) SolveByPunctuating() {
	fmt.Printf("Worker[%d] working on Punctuating\n", w.Id)
}

func (w Worker) SolveByMercury() {
	fmt.Printf("Worker[%d] working on Mercury\n", w.Id)
}
