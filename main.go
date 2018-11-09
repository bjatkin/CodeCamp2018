package main

import (
	"fmt"

	"github.com/dennwc/dom/svg"
)

//All the colors you'll need
const (
	Pcolor = "#041630"
	Scolor = "#5a8bbc"
	Black  = "#000020"
	White  = "#f4f4ff"
)

func main() {
	fmt.Println("test")

	w, h := 1409, 773

	doc := svg.NewFullscreen()
	root := doc.NewG()

	size := 500
	board := NewUIBoard(root, w/2, h/2, size, 5*(size/4))
	board.AddUIBox(Fill, 10, 10)
}
