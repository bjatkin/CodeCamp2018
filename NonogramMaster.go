package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Master struct {
	Board      Board
	MovesIn    chan Move
	MovesOut   chan Move
	GuiMovesIn chan Move
	Tasks      []Method
	Workers    []Worker

	WaitGroup *sync.WaitGroup
}

func NewMaster(board Board, GuiMovesIn chan Move) (m *Master) {
	m = &Master{}
	m.Board = board
	m.MovesIn = make(chan Move, WorkerCount)
	m.MovesOut = make(chan Move, WorkerCount)
	m.GuiMovesIn = GuiMovesIn
	m.Workers = make([]Worker, WorkerCount)
	m.WaitGroup = &sync.WaitGroup{}
	m.WaitGroup.Add(WorkerCount)
	for i := 0; i < WorkerCount; i++ {
		m.newWorker(i, m.WaitGroup)
	}
	m.Tasks = make([]Method, MethodCount)
	for i := 0; i < int(MethodCount); i++ {
		m.Tasks[i] = Method(i)
	}
	return
}

func (m Master) newWorker(id int, WaitGroup *sync.WaitGroup) {
	m.Workers[id] = Worker{
		Id:        id,
		Board:     m.Board,
		MovesIn:   make(chan Move, WorkerCount),
		MovesOut:  m.MovesIn,
		Done: make(chan bool),
		WaitGroup: WaitGroup,
	}
}

func (m Master) Solve() {
	go m.ProcessInbox()

	for i := 0; i < WorkerCount; i++ {
		m.ShuffleMethods()
		m.Workers[i].Tasks = m.Tasks
		go m.Workers[i].Solve()
	}

	go func() {
		m.WaitGroup.Wait()
		close(m.MovesIn)
	}()
}

func (m Master) ProcessInbox() {
	for {
		select {
		case move, ok := <- m.MovesIn:

			err := m.Board.MarkCell(move)
			if err != nil {
				fmt.Printf("%+v\n", err)
				fmt.Printf("%+v\n", move)
			}
			for _, worker := range m.Workers {
				worker.MovesIn <- move
			}

			m.GuiMovesIn <- move
			if !ok {
				for _, worker := range m.Workers {
					worker.Done <- true
				}
				close(m.GuiMovesIn)
			}
		}
	}
}

func (m Master) ShuffleMethods() {
	rand.Shuffle(len(m.Tasks), func(i, j int) {
		m.Tasks[i], m.Tasks[j] = m.Tasks[j], m.Tasks[i]
	})
}
