package puzzle

import (
	"fmt"
)

type Puzzle struct {
	size   int
	cells  [][]int
	emptyX int
	emptyY int
}

func (p *Puzzle) FillFromLeft() error {
	if p.emptyX == 0 {
		return fmt.Errorf("can't fill empty cell from left")
	}
	p.cells[p.emptyY][p.emptyX], p.cells[p.emptyY][p.emptyX-1] =
		p.cells[p.emptyY][p.emptyX-1], p.cells[p.emptyY][p.emptyX]
	p.emptyX -= 1
	return nil
}

func (p *Puzzle) FillFromRight() error {
	if p.emptyX == p.size-1 {
		return fmt.Errorf("can't fill empty cell from right")
	}
	p.cells[p.emptyY][p.emptyX], p.cells[p.emptyY][p.emptyX+1] =
		p.cells[p.emptyY][p.emptyX+1], p.cells[p.emptyY][p.emptyX]
	p.emptyX += 1
	return nil
}

func (p *Puzzle) FillFromAbove() error {
	if p.emptyY == 0 {
		return fmt.Errorf("can't fill empty cell from above")
	}
	p.cells[p.emptyY][p.emptyX], p.cells[p.emptyY-1][p.emptyX] =
		p.cells[p.emptyY-1][p.emptyX], p.cells[p.emptyY][p.emptyX]
	p.emptyY -= 1
	return nil
}

func (p *Puzzle) FillFromBelow() error {
	if p.emptyY == p.size-1 {
		return fmt.Errorf("can't fill empty cell from below")
	}
	p.cells[p.emptyY][p.emptyX], p.cells[p.emptyY+1][p.emptyX] =
		p.cells[p.emptyY+1][p.emptyX], p.cells[p.emptyY][p.emptyX]
	p.emptyY += 1
	return nil
}
