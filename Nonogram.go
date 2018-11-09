package CodeCamp2018

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

