package main

import (
	"fmt"
)

type Worker struct {
	Id int
	Board *Board
	MovesIn <-chan Move
	MovesOut chan<- Move
	Task Method
}

func (w Worker) Solve() {
	fmt.Printf("Worker[%d] working on '%d'\n", w.Id, w.Task)
}