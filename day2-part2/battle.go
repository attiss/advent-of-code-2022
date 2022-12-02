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
	Outcome       int
}

func (b Battle) MyShape() int {
	var myShape int

	switch b.Outcome {
	case draw:
		myShape = b.OpponentShape
	case win:
		switch b.OpponentShape {
		case rock:
			myShape = paper
		case paper:
			myShape = scissor
		case scissor:
			myShape = rock
		}
	case loose:
		switch b.OpponentShape {
		case rock:
			myShape = scissor
		case paper:
			myShape = rock
		case scissor:
			myShape = paper
		}
	}

	return myShape
}

func (b Battle) Score() int {
	var score int

	switch b.Outcome {
	case win:
		score += 6
	case loose:
		score += 0
	case draw:
		score += 3
	}

	switch b.MyShape() {
	case rock:
		score += 1
	case paper:
		score += 2
	case scissor:
		score += 3
	}

	return score
}
