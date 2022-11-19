package puzzle

import (
	"fmt"
	"log"
)

func (p *Puzzle) IsSolved() bool {
	if len(p.cells) != len(p.target) {
		log.Fatal(fmt.Errorf("target size(y) differs from the puzzle"))
		return false
	}
	for y, row := range p.cells {
		if len(row) != len(p.target[y]) {
			log.Fatal(fmt.Errorf("target size(x) differs from the puzzle"))
			return false
		}
		for x, v := range row {
			if v != p.target[y][x] {
				return false
			}
		}
	}
	return true
}

func (p *Puzzle) targetState() [][]int {
	target := make([][]int, 0, p.size)
	for i := 0; i < p.size; i++ {
		target = append(target, make([]int, p.size))
	}
	i := 1
	lastX, lastY := 0, 0
	p.iterateSnail(func(x, y int) {
		target[y][x] = i
		i++
		lastX, lastY = x, y
	})
	target[lastY][lastX] = 0
	return target
}
