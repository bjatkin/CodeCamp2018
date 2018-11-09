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
)

func main() {
	fmt.Println("test")

	w, h := 300.0, 300.0

	root := svg.NewFullscreen()
	center := root.NewG()
	center.Translate(w/2, h/2)

	mainBoard := Board{
		SVG: center.NewRect(10, 10),
	}
	mainBoard.SVG.SetAttribute("fill", Pcolor)
}
