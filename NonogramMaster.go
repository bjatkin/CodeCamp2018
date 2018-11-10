package main

import "math/rand"

type Master struct {
	Board    *Board
	MovesIn  chan Move
	MovesOut chan Move
	Tasks    []Method
	Workers  []Worker
}

func NewMaster(board *Board) (m *Master) {
	m = &Master{}
	m.Board = board
	m.MovesIn = make(chan Move, WorkerCount)
	m.MovesOut = make(chan Move, WorkerCount)
	m.Workers = make([]Worker, WorkerCount)
	for i := 0; i < WorkerCount; i++ {
		m.newWorker(i)
	}
	m.Tasks = make([]Method, MethodCount)
	for i := 0; i < int(MethodCount); i++ {
		m.Tasks[i] = Method(i)
	}
	return
}

func (m Master) newWorker(id int) {
	m.Workers[id] = Worker{
		Id:       id,
		Board:    m.Board,
		MovesIn:  make(<-chan Move, WorkerCount),
		MovesOut: m.MovesIn,
	}
}

func (m Master) Solve() {
	for i := 0; i < WorkerCount; i++ {
		m.ShuffleMethods()
		m.Workers[i].Tasks = m.Tasks
		m.Workers[i].Solve()
	}
}

func (m Master) ShuffleMethods() {
	rand.Shuffle(len(m.Tasks), func(i, j int) {
		m.Tasks[i], m.Tasks[j] = m.Tasks[j], m.Tasks[i]
	})
}
