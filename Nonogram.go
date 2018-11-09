package main

type Mark int

const (
	Fill  Mark = 0
	Cross Mark = 1
	Empty Mark = 2
)

type Move struct {
	X    int
	Y    int
	Mark Mark
}

type Method int

const (
	Boxes Method = 0
	Spaces Method = 1
	Forcing Method = 2
	Glue Method = 3
	Joining Method = 4
	Splitting Method = 5
	Punctuating Method = 6
	Mercury Method = 7
	MethodCount Method = 8
)