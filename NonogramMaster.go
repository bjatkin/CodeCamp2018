package CodeCamp2018

type Master struct {
	Board *Board
	MovesIn <-chan Move
	MovesOut chan<- Move
	Workers []Worker
}

func newMaster(board *Board) (m *Master) {
	m.Board = board
	m.MovesIn = make(<-chan Move, WorkerCount)
	m.MovesOut = make(chan<- Move, WorkerCount)
	m.Workers = make([]Worker, 0)
	for i := 0; i < WorkerCount; i++ {
		m.newWorker()
	}
	return
}
