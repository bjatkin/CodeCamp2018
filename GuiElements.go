package main

import (
	"fmt"

	"github.com/dennwc/dom"
	"github.com/dennwc/dom/svg"
)

type UIBoard struct {
	root       *svg.G
	board      []*svg.Line
	squares    []UIBox
	colNumbers []UIText
	rowNumbers []UIText
	x          int
	y          int
	width      int
	height     int
	cols       int
	rows       int
}

func NewUIBoard(root *svg.G, x, y, height, width, rows, cols int, rowNum, colNum [][]int) UIBoard {
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

	ret := UIBoard{
		board:  board,
		root:   root,
		x:      x,
		y:      y,
		width:  width,
		height: height,
		rows:   rows,
		cols:   cols,
	}

	for i, r := range rowNum {
		ret.rowNumbers = append(ret.rowNumbers, ret.AddText(false, r, i))
	}

	for i, c := range colNum {
		ret.colNumbers = append(ret.colNumbers, ret.AddText(true, c, i))
	}

	return ret
}

type UIText struct {
	txt []*svg.Text
}

func (board *UIBoard) AddText(horizontal bool, nums []int, index int) UIText {
	ret := UIText{}
	minX := board.x + board.width/4
	maxX := board.x + 3*(board.width/4)
	minY := board.y + board.height/5

	size := ((maxX - minX) / board.cols)
	fontSize := size - 15

	if fontSize > 25 {
		fontSize = 25
	}

	if fontSize < 6 {
		fontSize = 6
	}
	for i, num := range nums {
		n := board.root.NewText(fmt.Sprintf("%d", num))
		n.SetAttribute("style", fmt.Sprintf("fill: %s; font: %dpx sans-serif;", Scolor, fontSize))

		if horizontal {
			n.Translate(float64((minX+index*size)+size/3), float64(minY-fontSize*i)-2)
		} else {
			n.Translate(float64(minX-fontSize-(fontSize*i)), float64(((minY+size)+index*size)-size/2))
		}
		ret.txt = append(ret.txt, n)
	}
	return ret
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
