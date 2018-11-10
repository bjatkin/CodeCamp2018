package main

import (
	"math/rand"
	"time"
)

//All the colors you'll need
const (
	BGcolor = "#041630"
	Scolor  = "#5a8bbc"
	Black   = "#000020"
	White   = "#f4f4ff"

	WorkerCount = 4
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	GuiMovesIn := make(chan Move, WorkerCount*1000)

	rowHints := [][]int{
		{13, 3},
		{11, 2},
		{10, 2},
		{9, 1},
		{4, 3},
		{2, 3, 2, 1},
		{2, 3, 10},
		{1, 2, 4, 4},
		{3, 1, 2},
		{1, 1, 2, 3},
		{2, 1, 3},
		{2, 3, 2},
		{4, 3, 1},
		{1, 2},
		{3, 1, 3},
		{8, 1, 1, 3},
		{6, 3, 2},
		{6, 7, 2},
		{11, 2},
		{10, 2},
	}
	columnHints := [][]int{
		{5, 1, 1, 1},
		{7, 1, 1},
		{8, 3},
		{5, 2, 3},
		{4, 2, 3},
		{4, 1, 5, 5},
		{9, 1, 1, 6},
		{10, 7},
		{6, 2, 2},
		{3, 1, 5, 3},
		{2, 2, 3, 3},
		{1, 3, 1, 1, 3},
		{1, 6, 3},
		{1, 1, 2, 6},
		{1, 1, 2, 6},
		{1, 6},
		{1, 3},
		{1, 3},
		{1, 1, 2, 1, 4},
		{2, 3, 4},
	}

	board := drawUI(500, columnHints, rowHints)

	mainBoard := NewBoard(len(rowHints), len(columnHints), rowHints, columnHints)
	nonogramMaster := NewMaster(mainBoard, GuiMovesIn)
	nonogramMaster.Solve()

	tick := time.Tick(50 * time.Millisecond)
	for {
		select {
		case <-tick:
			select {
			case move := <-GuiMovesIn:
				board.UpdateCoord(move.Mark, move.X, move.Y)
			}
		}
	}
}
