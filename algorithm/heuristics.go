package algorithm

import (
	"math"
	"npuzzle/puzzle"
)

func OutOfPlace(p puzzle.Puzzle) int {
	result := 0
	for y, row := range p.Tiles {
		for x, t := range row {
			if p.Target[y][x] != t {
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
			result += int(
				math.Abs(float64(x-p.TargetXY[t].X)) +
					math.Abs(float64(y-p.TargetXY[t].Y))*10)
		}
	}
	return result
}

func Euclidean(p puzzle.Puzzle) int {
	result := 0
	for y, row := range p.Tiles {
		for x, t := range row {
			result += int(
				math.Sqrt(
					math.Pow(float64(x-p.TargetXY[t].X), 2)+
						math.Pow(float64(y-p.TargetXY[t].Y), 2)) * 10)
		}
	}
	return result
}
