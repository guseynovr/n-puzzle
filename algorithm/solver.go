package algorithm

import (
	"npuzzle/puzzle"
)

type Solver struct {
	P      *puzzle.Puzzle
	H      func(puzzle.Puzzle) int
	Stats  Stats
	Debug  bool
	ByH    bool
	Corner Rectangle
}
