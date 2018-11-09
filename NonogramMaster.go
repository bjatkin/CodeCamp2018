package CodeCamp2018

type Master struct {
	Board Board
	MovesIn <-chan Move
	MovesOut chan<- Move
}
