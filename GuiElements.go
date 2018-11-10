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

func flipArray(numbers []int) []int {
	ret := make([]int, len(numbers))
	for i, n := range numbers {
		ret[i] = n
	}
	for i := 0; i < len(ret)/2; i++ {
		j := len(ret) - i - 1
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

func (board *UIBoard) AddText(horizontal bool, nums []int, index int) UIText {
	nums = flipArray(nums)
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
	cross  []*svg.Line
	state  Mark
	x      int
	y      int
}

func (board *UIBoard) AddUIBox(move Move, width, height int) UIBox {
	// Check for a box at these coordinates
	var ret UIBox
	found := false
	if len(board.squares) > 0 {
		for _, box := range board.squares {
			if box.x == move.X && box.y == move.Y {
				ret = box
				found = true
			}
		}
	}

	if !found {
		cross := []*svg.Line{
			board.root.NewLine(),
			board.root.NewLine(),
		}

		ret = UIBox{
			square: board.root.NewRect(width-2, height-2),
			cross:  cross,
			state:  move.Mark,
			x:      move.X,
			y:      move.Y,
		}
		ret.square.Translate(float64(move.X), float64(move.Y))
		ret.cross[0].SetPos(dom.Point{move.X + 5, move.Y + 5}, dom.Point{move.X + width - 7, move.Y + height - 7})
		ret.cross[1].SetPos(dom.Point{move.X + 5, move.Y + height - 7}, dom.Point{move.X + width - 7, move.Y + 5})
	}

	drawColor := Scolor
	switch move.MethodId {
	case Boxes:
		drawColor = "#FF0000"
	case Forcing:
		drawColor = "#00FF00"
	case Glue:
		drawColor = "#00FFFF"
	}

	switch move.Mark {
	case Fill:
		ret.square.SetAttribute("fill", drawColor)
		ret.square.SetAttribute("fill-opacity", "1.0")
		ret.square.SetAttribute("style", "")
		ret.cross[0].SetStrokeWidth(0)
		ret.cross[1].SetStrokeWidth(0)
	case Cross:
		ret.square.SetAttribute("fill-opacity", "0.0")
		ret.square.SetAttribute("style", "stroke-width: 1px; stroke: "+drawColor)
		ret.cross[0].SetStrokeWidth(1)
		ret.cross[1].SetStrokeWidth(1)
		ret.cross[0].SetAttribute("stroke", drawColor)
		ret.cross[1].SetAttribute("stroke", drawColor)
	case Empty:
		ret.square.SetAttribute("fill-opacity", 0.0)
		ret.square.SetAttribute("style", "")
		ret.cross[0].SetStrokeWidth(0)
		ret.cross[1].SetStrokeWidth(0)
	}

	return ret
}

func (board *UIBoard) UpdateCoord(move Move) {
	if move.X+1 > board.cols || move.Y+1 > board.rows {
		return
	}
	minX := board.x + board.width/4
	maxX := board.x + 3*(board.width/4)
	minY := board.y + board.height/5

	width := (maxX - minX) / board.cols
	height := width

	move.X = minX+width*move.X
	move.Y = minY+height*move.Y
	board.squares = append(board.squares, board.AddUIBox(move, width, height))
}
