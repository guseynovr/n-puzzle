package puzzle

import "fmt"

func (p *Puzzle) FillFromLeft() error {
	if p.Zero.X == 0 || p.Tiles[p.Zero.Y][p.Zero.X-1].Locked {
		return fmt.Errorf("can't fill empty tile from left")
	}
	p.Tiles[p.Zero.Y][p.Zero.X], p.Tiles[p.Zero.Y][p.Zero.X-1] =
		p.Tiles[p.Zero.Y][p.Zero.X-1], p.Tiles[p.Zero.Y][p.Zero.X]
	p.Zero.X -= 1
	p.updateHash()
	return nil
}

func (p *Puzzle) FillFromRight() error {
	if p.Zero.X == p.Size-1 || p.Tiles[p.Zero.Y][p.Zero.X+1].Locked {
		return fmt.Errorf("can't fill empty tile from right")
	}
	p.Tiles[p.Zero.Y][p.Zero.X], p.Tiles[p.Zero.Y][p.Zero.X+1] =
		p.Tiles[p.Zero.Y][p.Zero.X+1], p.Tiles[p.Zero.Y][p.Zero.X]
	p.Zero.X += 1
	p.updateHash()
	return nil
}

func (p *Puzzle) FillFromAbove() error {
	if p.Zero.Y == 0 || p.Tiles[p.Zero.Y-1][p.Zero.X].Locked {
		return fmt.Errorf("can't fill empty tile from above")
	}
	p.Tiles[p.Zero.Y][p.Zero.X], p.Tiles[p.Zero.Y-1][p.Zero.X] =
		p.Tiles[p.Zero.Y-1][p.Zero.X], p.Tiles[p.Zero.Y][p.Zero.X]
	p.Zero.Y -= 1
	p.updateHash()
	return nil
}

func (p *Puzzle) FillFromBelow() error {
	if p.Zero.Y == p.Size-1 || p.Tiles[p.Zero.Y+1][p.Zero.X].Locked {
		return fmt.Errorf("can't fill empty tile from below")
	}
	p.Tiles[p.Zero.Y][p.Zero.X], p.Tiles[p.Zero.Y+1][p.Zero.X] =
		p.Tiles[p.Zero.Y+1][p.Zero.X], p.Tiles[p.Zero.Y][p.Zero.X]
	p.Zero.Y += 1
	p.updateHash()
	return nil
}

func (p Puzzle) FilledFromLeft() (Puzzle, error) {
	if p.Zero.X == 0 || p.Tiles[p.Zero.Y][p.Zero.X-1].Locked {
		return p, fmt.Errorf("can't fill empty tile from left")
	}
	p = p.DeepCopy()
	err := p.FillFromLeft()
	return p, err
}

func (p Puzzle) FilledFromRight() (Puzzle, error) {
	if p.Zero.X == p.Size-1 || p.Tiles[p.Zero.Y][p.Zero.X+1].Locked {
		return p, fmt.Errorf("can't fill empty tile from right")
	}
	p = p.DeepCopy()
	err := p.FillFromRight()
	return p, err
}

func (p Puzzle) FilledFromAbove() (Puzzle, error) {
	if p.Zero.Y == 0 || p.Tiles[p.Zero.Y-1][p.Zero.X].Locked {
		return p, fmt.Errorf("can't fill empty tile from above")
	}
	p = p.DeepCopy()
	err := p.FillFromAbove()
	return p, err
}

func (p Puzzle) FilledFromBelow() (Puzzle, error) {
	if p.Zero.Y == p.Size-1 || p.Tiles[p.Zero.Y+1][p.Zero.X].Locked {
		return p, fmt.Errorf("can't fill empty tile from below")
	}
	p = p.DeepCopy()
	err := p.FillFromBelow()
	return p, err
}
