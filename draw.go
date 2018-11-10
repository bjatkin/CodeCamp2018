package main

import (
	"time"

	"github.com/dennwc/dom/svg"
)

func drawUI(size int, rowHints, columnHints [][]int) UIBoard {
	w, h := 1409, 773

	doc := svg.NewFullscreen()
	root := doc.NewG()
	defs := root.NewTag("defs")
	rG := defs.NewTag("radialGradient")
	rG.SetAttribute("id", "grad1")
	rG.SetAttribute("cx", "50%")
	rG.SetAttribute("cy", "50%")
	rG.SetAttribute("r", "50%")
	rG.SetAttribute("fx", "50%")
	rG.SetAttribute("fy", "50%")
	stop1 := rG.NewTag("stop")
	stop1.SetAttribute("offset", "0%")
	stop1.SetAttribute("style", "stop-color:#59b7ff;stop-opacity:0.5")
	stop2 := rG.NewTag("stop")
	stop2.SetAttribute("offset", "100%")
	stop2.SetAttribute("style", "stop-color:#041630;stop-opacity:0")
	bg := root.NewRect(w, h)
	bg.SetAttribute("fill", "url(#grad1)")

	cols := len(columnHints)
	rows := len(rowHints)
	board := NewUIBoard(root, w/2-(5*(size/4))/2, h/2-(size/2), size, 5*(size/4), cols, rows, columnHints, rowHints)

	//Fade the GUI in
	go func() {
		interval := time.Millisecond * 10
		for i := 0.0; i < 1; i += 0.01 {
			board.root.SetAttribute("opacity", i)
			time.Sleep(interval)
		}
	}()

	return board
}
