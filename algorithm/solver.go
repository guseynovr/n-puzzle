package algorithm

import "npuzzle/puzzle"

type Solver struct {
	P     *puzzle.Puzzle
	H     func(puzzle.Puzzle) int
	Ver   Rectangle
	Hor   Rectangle
	Stats Stats
	Debug bool
}
