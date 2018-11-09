package main

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/svg"
)

type UIBoard struct {
	root       *svg.G
	board      []*svg.Line
	squares    []UIBox
	colNumbers []svg.Element
	rowNumbers []svg.Element
}

func NewUIBoard(root *svg.G, x, y, height, width int) UIBoard {
	x = x - width/2
	y = y - height/2

	board := []*svg.Line{
		root.NewLine(),
		root.NewLine(),
		root.NewLine(),
		root.NewLine(),
		root.NewLine(),
		root.NewLine(),
	}

	for _, line := range board {
		line.SetStrokeWidth(1)
		line.SetAttribute("stroke", Scolor)
	}
	board[0].SetPos(dom.Point{x, y + height/2}, dom.Point{x + width/4, y})
	board[1].SetPos(dom.Point{x + width/4, y}, dom.Point{x + 3*(width/4), y})
	board[2].SetPos(dom.Point{x + 3*(width/4), y}, dom.Point{x + width, y + height/2})
	board[3].SetPos(dom.Point{x + width, y + height/2}, dom.Point{x + 3*(width/4), y + height})
	board[4].SetPos(dom.Point{x + 3*(width/4), y + height}, dom.Point{x + width/4, y + height})
	board[5].SetPos(dom.Point{x + width/4, y + height}, dom.Point{x, y + height/2})

	return UIBoard{
		board: board,
		root:  root,
	}
}

type UIBox struct {
	square *svg.Rect
	state  Mark
	x      int
	y      int
}

func (board *UIBoard) AddUIBox(state Mark, x int, y int) UIBox {
	ret := UIBox{
		square: board.root.NewRect(10, 10),
		state:  state,
		x:      x,
		y:      y,
	}
	ret.square.SetAttribute("fill", Scolor)
	return ret
}
