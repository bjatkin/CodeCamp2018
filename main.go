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
		{0},
		{2},
		{0},
		{0},
		{1, 1, 1},
	}
	columnHints := [][]int{
		{1, 2},
		{0},
		{0},
		{0},
		{2},
	}

	board := drawUI(500, columnHints, rowHints)

	mainBoard := NewBoard(5, 5, rowHints, columnHints)
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
