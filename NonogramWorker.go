package CodeCamp2018

type Worker struct {
	Board Board
	MovesIn <-chan Move
	MovesOut chan<- Move
}

