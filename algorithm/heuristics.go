package algorithm

import (
	"math"
	"npuzzle/puzzle"
)

func OutOfPlace(p puzzle.Puzzle) int {
	result := 0
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Relevant && (t.Target.X != x || t.Target.Y != y) {
				result += 10
			}
		}
	}
	return result
}

func Manhattan(p puzzle.Puzzle) int {
	result := 0
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Relevant {
				result += int(
					math.Abs(float64(x-t.Target.X)) +
						math.Abs(float64(y-t.Target.Y))*10)
			}
		}
	}
	return result
}

func Euclidean(p puzzle.Puzzle) int {
	result := 0
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Relevant {
				result += int(
					math.Sqrt(
						math.Pow(float64(x-t.Target.X), 2)+
							math.Pow(float64(y-t.Target.Y), 2)) * 10)
			}
		}
	}
	return result
}
