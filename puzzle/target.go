package puzzle

import (
	"fmt"
	"log"
)

func (p *Puzzle) IsSolved() bool {
	if len(p.Tiles) != len(p.Target) {
		log.Fatal(fmt.Errorf("target size(y) differs from the puzzle"))
		return false
	}
	for y, row := range p.Tiles {
		if len(row) != len(p.Target[y]) {
			log.Fatal(fmt.Errorf("target size(x) differs from the puzzle"))
			return false
		}
		for x, v := range row {
			if v != p.Target[y][x] {
				return false
			}
		}
	}
	return true
}

func (p *Puzzle) targetState() [][]int {
	target := make([][]int, 0, p.Size)
	for i := 0; i < p.Size; i++ {
		target = append(target, make([]int, p.Size))
	}
	i := 1
	lastX, lastY := 0, 0
	p.iterateSnail(func(x, y int) {
		target[y][x] = i
		i++
		lastX, lastY = x, y
	})
	target[lastY][lastX] = 0

	p.TargetXY = make(map[int]struct{ X, Y int })
	for y, row := range p.Target {
		for x, t := range row {
			p.TargetXY[t] = struct {
				X int
				Y int
			}{x, y}
		}
	}
	return target
}
