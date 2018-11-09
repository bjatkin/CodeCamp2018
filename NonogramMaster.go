package main

type Master struct {
	Board    Board
	MovesIn  <-chan Move
	MovesOut chan<- Move
}
