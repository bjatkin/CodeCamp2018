package CodeCamp2018

import (
	"fmt"

	"github.com/dennwc/dom/svg"
)

const (
	Pcolor = "#0416300"
	Scolor = "#5a8bbc"
	Black  = "#000020"
	White  = "#f4f4ff"

	WorkerCount = 4
)

func main() {
	fmt.Println("test")

	w, h := 300.0, 300.0

	root := svg.NewFullscreen()
	center := root.NewG()
	center.Translate(w/2, h/2)

	rowHints := [][]int{
		{2, 2},
		{1, 1, 1},
		{0},
		{3},
		{1, 3},
	}
	columnHints := [][]int {
		{2, 1},
		{1},
		{1, 2},
		{1, 2},
		{2, 2},
	}
	mainBoard := newBoard(5, 5, rowHints, columnHints);
	nonogramMaster := newMaster(mainBoard)

}
