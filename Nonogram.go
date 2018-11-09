package main

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
