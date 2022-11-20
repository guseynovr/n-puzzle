package puzzle

import "fmt"

func (p *Puzzle) FillFromLeft() error {
	if p.emptyX == 0 {
		return fmt.Errorf("can't fill empty tile from left")
	}
	p.Tiles[p.emptyY][p.emptyX], p.Tiles[p.emptyY][p.emptyX-1] =
		p.Tiles[p.emptyY][p.emptyX-1], p.Tiles[p.emptyY][p.emptyX]
	p.emptyX -= 1
	p.updateHash()
	return nil
}

func (p *Puzzle) FillFromRight() error {
	if p.emptyX == p.Size-1 {
		return fmt.Errorf("can't fill empty tile from right")
	}
	p.Tiles[p.emptyY][p.emptyX], p.Tiles[p.emptyY][p.emptyX+1] =
		p.Tiles[p.emptyY][p.emptyX+1], p.Tiles[p.emptyY][p.emptyX]
	p.emptyX += 1
	p.updateHash()
	return nil
}

func (p *Puzzle) FillFromAbove() error {
	if p.emptyY == 0 {
		return fmt.Errorf("can't fill empty tile from above")
	}
	p.Tiles[p.emptyY][p.emptyX], p.Tiles[p.emptyY-1][p.emptyX] =
		p.Tiles[p.emptyY-1][p.emptyX], p.Tiles[p.emptyY][p.emptyX]
	p.emptyY -= 1
	p.updateHash()
	return nil
}

func (p *Puzzle) FillFromBelow() error {
	if p.emptyY == p.Size-1 {
		return fmt.Errorf("can't fill empty tile from below")
	}
	p.Tiles[p.emptyY][p.emptyX], p.Tiles[p.emptyY+1][p.emptyX] =
		p.Tiles[p.emptyY+1][p.emptyX], p.Tiles[p.emptyY][p.emptyX]
	p.emptyY += 1
	p.updateHash()
	return nil
}

func (p Puzzle) FilledFromLeft() (Puzzle, error) {
	if p.emptyX == 0 {
		return p, fmt.Errorf("can't fill empty tile from left")
	}
	p = p.deepCopy()
	err := p.FillFromLeft()
	return p, err
}

func (p Puzzle) FilledFromRight() (Puzzle, error) {
	if p.emptyX == p.Size-1 {
		return p, fmt.Errorf("can't fill empty tile from right")
	}
	p = p.deepCopy()
	err := p.FillFromRight()
	return p, err
}

func (p Puzzle) FilledFromAbove() (Puzzle, error) {
	if p.emptyY == 0 {
		return p, fmt.Errorf("can't fill empty tile from above")
	}
	p = p.deepCopy()
	err := p.FillFromAbove()
	return p, err
}

func (p Puzzle) FilledFromBelow() (Puzzle, error) {
	if p.emptyY == p.Size-1 {
		return p, fmt.Errorf("can't fill empty tile from below")
	}
	p = p.deepCopy()
	err := p.FillFromBelow()
	return p, err
}

func (p *Puzzle) deepCopy() Puzzle {
	pCopy := *p
	pCopy.Tiles = make([][]int, p.Size)
	for i := range pCopy.Tiles {
		pCopy.Tiles[i] = make([]int, p.Size)
		copy(pCopy.Tiles[i], p.Tiles[i])
	}
	return pCopy
}
