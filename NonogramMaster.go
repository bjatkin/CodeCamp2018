package main

type Master struct {
	Board *Board
	MovesIn <-chan Move
	MovesOut chan<- Move
	Tasks chan Method
	Workers []Worker
}

func NewMaster(board *Board) (m *Master) {
	m = &Master{}
	m.Board = board
	m.MovesIn = make(<-chan Move, WorkerCount)
	m.MovesOut = make(chan<- Move, WorkerCount)
	m.Workers = make([]Worker, WorkerCount)
	for i := 0; i < WorkerCount; i++ {
		m.newWorker(i)
	}
	m.Tasks = make(chan Method, MethodCount)
	for i := 0; i < int(MethodCount); i++ {
		m.Tasks <- Method(i)
	}
	return
}

func (m Master) newWorker(id int) {
	m.Workers[id] = Worker {
		Id: id,
		Board: m.Board,
		MovesIn: make(<-chan Move, WorkerCount),
		MovesOut: make(chan<- Move, WorkerCount),
	}
}

func (m Master) Solve() {
	for i:= 0; i < WorkerCount; i++ {
		m.Workers[i].Task = <-m.Tasks
		m.Workers[i].Solve()
	}
}