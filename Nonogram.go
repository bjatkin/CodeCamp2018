package CodeCamp2018

import "github.com/dennwc/dom/svg"

type Mark int

const (
	Fill  Mark = 0
	Cross Mark = 1
)

type Move struct {
	X    int
	Y    int
	Mark Mark
}

