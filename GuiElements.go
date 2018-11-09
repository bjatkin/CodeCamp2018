package main

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/svg"
)

type UIBoard struct {
	root       *svg.G
	board      []*svg.Line
	squares    []UIBox
	colNumbers []*svg.Text
	rowNumbers []*svg.Text
	x          int
	y          int
	width      int
	height     int
	cols       int
	rows       int
}

func NewUIBoard(root *svg.G, x, y, height, width, rows, cols int) UIBoard {
	board := []*svg.Line{
		root.NewLine(),
		root.NewLine(),
		root.NewLine(),
		root.NewLine(),
		root.NewLine(),
		root.NewLine(),
	}

	cNum := []*svg.Text{
		root.NewText("test"),
	}
	rNum := []*svg.Text{
		root.NewText("test row"),
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
		board:      board,
		root:       root,
		x:          x,
		y:          y,
		width:      width,
		height:     height,
		rows:       rows,
		cols:       cols,
		colNumbers: cNum,
		rowNumbers: rNum,
	}
}

type UIBox struct {
	square *svg.Rect
	state  Mark
	x      int
	y      int
}

func (board *UIBoard) AddUIBox(state Mark, x, y, width, height int) UIBox {
	ret := UIBox{
		square: board.root.NewRect(width-1, height-1),
		state:  state,
		x:      x,
		y:      y,
	}
	ret.square.SetAttribute("fill", Scolor)
	ret.square.Translate(float64(x), float64(y))
	return ret
}

func (board *UIBoard) UpdateCoord(state Mark, x, y int) {
	if x+1 > board.cols || y+1 > board.rows {
		return
	}
	minX := board.x + board.width/4
	maxX := board.x + 3*(board.width/4)
	minY := board.y + board.height/5

	width := (maxX - minX) / board.cols
	height := width
	board.AddUIBox(state, minX+width*x, minY+height*y, width, height)
}
