package algorithm

import "npuzzle/puzzle"

type Stats struct {
	TotalStates int
	MaxStates   int
	PathLen     int
	Path        []puzzle.Puzzle
}
