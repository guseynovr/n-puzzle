package algorithm

import "npuzzle/puzzle"

type Rectangle struct {
	topLeft  puzzle.Coordinates
	botRight puzzle.Coordinates
}

func (r Rectangle) contains(x, y int) bool {
	return x >= r.topLeft.X && x <= r.botRight.X &&
		y >= r.topLeft.Y && y <= r.botRight.Y
}
