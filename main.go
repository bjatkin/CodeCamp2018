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
	GuiMovesIn := make(chan Move, WorkerCount * 100)
	GuiMovesIn <- Move{0, 0, 0, Fill}

	rowHints := [][]int{
		{2, 2},
		{1, 1, 1},
		{0},
		{3},
		{1, 3},
	}
	columnHints := [][]int{
		{2, 1},
		{1},
		{1, 2},
		{1, 2},
		{2, 2},
	}

	board := drawUI(rowHints, columnHints)

	mainBoard := NewBoard(5, 5, rowHints, columnHints)
	nonogramMaster := NewMaster(mainBoard, GuiMovesIn)
	nonogramMaster.Solve()

	tick := time.Tick(500 * time.Millisecond)
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
