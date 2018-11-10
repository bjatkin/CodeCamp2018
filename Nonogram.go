package main

type Mark int

const (
	Empty Mark = 0
	Fill  Mark = 1
	Cross Mark = 2
	MarkCount Mark = 3
)

type Move struct {
	WorkerId int
	MethodId Method
	X        int
	Y        int
	Mark     Mark
}

type Method int

const (
	Boxes       Method = 0
	Spaces      Method = 1
	Forcing     Method = 2
	Glue        Method = 3
	Joining     Method = 4
	Splitting   Method = 5
	Punctuating Method = 6
	Mercury     Method = 7
	MethodCount Method = 8
)
