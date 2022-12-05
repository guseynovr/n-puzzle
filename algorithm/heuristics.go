package algorithm

import (
	"math"
	"npuzzle/puzzle"
)

func OutOfPlace(p puzzle.Puzzle) int {
	result := 0
	var nextPos puzzle.Coordinates
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Relevant && (t.Target.X != x || t.Target.Y != y) {
				result += 10
			}
			if t.Value == p.Next {
				nextPos = puzzle.Coordinates{x, y}
			}
		}
	}
	if p.Next > 0 {
		result += penalty(p.Zero, nextPos)
	}
	return result
}

func Manhattan(p puzzle.Puzzle) int {
	result := 0
	var nextPos puzzle.Coordinates
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Relevant {
				result += int(math.Round(
					math.Abs(float64(x-t.Target.X))+
						math.Abs(float64(y-t.Target.Y))) * 10)
			}
			if t.Value == p.Next {
				nextPos = puzzle.Coordinates{x, y}
			}
		}
	}
	if p.Next > 0 {
		result += penalty(p.Zero, nextPos)
	}
	return result
}

func Euclidean(p puzzle.Puzzle) int {
	result := 0
	var nextPos puzzle.Coordinates
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Relevant {
				result += int(math.Sqrt(
					math.Pow(float64(x-t.Target.X), 2)+
						math.Pow(float64(y-t.Target.Y), 2)) * 10)
			}
			if t.Value == p.Next {
				nextPos = puzzle.Coordinates{x, y}
			}
		}
	}
	if p.Next > 0 {
		result += penalty(p.Zero, nextPos)
	}
	return result
}

func Diagonal(p puzzle.Puzzle) int {
	result := 0
	var nextPos puzzle.Coordinates
	diagDist := math.Sqrt(200)
	//dx = abs(current_cell.x – goal.x)
	// dy = abs(current_cell.y – goal.y)
	// h = D * (dx + dy) + (D2 - 2 * D) * min(dx, dy)
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Relevant {
				dx := math.Abs(float64(x - t.Target.X))
				dy := math.Abs(float64(y - t.Target.Y))
				min := dx
				if dy < dx {
					min = dy
				}
				result += int(10*(dx+dy) + (diagDist-20)*min)
			}
			if t.Value == p.Next {
				nextPos = puzzle.Coordinates{x, y}
			}
		}
	}
	if p.Next > 0 {
		result += penalty(p.Zero, nextPos)
	}
	return result
}

func penalty(zeroPos, nextPos puzzle.Coordinates) int {
	// for y := -1; y < 2; y++ {
	// 	for x := -1; x < 2; x++ {
	// 		if nextPos.X+x == zeroPos.X && nextPos.Y+y == zeroPos.Y {
	// 			return 0
	// 		}
	// 	}
	// }
	zeroDist := int(math.Round(
		math.Abs(float64(zeroPos.X-nextPos.X))+
			math.Abs(float64(zeroPos.Y-nextPos.Y))) * 10)
	return zeroDist
}
