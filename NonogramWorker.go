package main

type Worker struct {
	Board *Board
	MovesIn <-chan Move
	MovesOut chan<- Move
}
