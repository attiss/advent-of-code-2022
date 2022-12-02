package main

const (
	win = iota
	loose
	draw
)

const (
	rock = iota
	paper
	scissor
)

type Battle struct {
	OpponentShape int
	MyShape       int
}

func (b Battle) Outcome() int {
	if b.OpponentShape == b.MyShape {
		return draw
	}
	if b.OpponentShape == rock && b.MyShape == paper ||
		b.OpponentShape == paper && b.MyShape == scissor ||
		b.OpponentShape == scissor && b.MyShape == rock {
		return win
	}

	return loose
}

func (b Battle) Score() int {
	var score int

	switch b.Outcome() {
	case win:
		score += 6
	case loose:
		score += 0
	case draw:
		score += 3
	}

	switch b.MyShape {
	case rock:
		score += 1
	case paper:
		score += 2
	case scissor:
		score += 3
	}

	return score
}
