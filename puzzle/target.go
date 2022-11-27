package puzzle

import (
	"log"
)

func (p *Puzzle) IsSolved() bool {
	// if len(p.Tiles) != len(p.Target) {
	// 	log.Fatal(fmt.Errorf("target size(y) differs from the puzzle"))
	// 	return false
	// }
	for y, row := range p.Tiles {
		// if len(row) != len(p.Target[y]) {
		// 	log.Fatal(fmt.Errorf("target size(x) differs from the puzzle"))
		// 	return false
		// }
		for x, v := range row {
			if v.Relevant && (x != v.Target.X || y != v.Target.Y) {
				return false
			}
		}
	}
	return true
}

func (p *Puzzle) initTargetState() {
	target := make(map[int]Coordinates)
	i := 1
	lastX, lastY := 0, 0
	p.iterateSnail(func(x, y int) {
		target[i] = Coordinates{x, y}
		i++
		lastX, lastY = x, y
	})
	target[0] = Coordinates{lastX, lastY}

	for y, row := range p.Tiles {
		for x, t := range row {
			p.Tiles[y][x].Target = target[t.Value]
		}
	}
}

func (p *Puzzle) MakeAllIrrelevant() {
	for y, row := range p.Tiles {
		for x := range row {
			p.Tiles[y][x].Relevant = false
		}
	}
}

func (p *Puzzle) MakeAllRelevant() {
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Value == 0 {
				continue
			}
			p.Tiles[y][x].Relevant = true
		}
	}
}

func (p *Puzzle) GetPosition(i int) Coordinates {
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Value == i {
				return Coordinates{x, y}
			}
		}
	}
	log.Fatal("Requested unexistent tile: ", i)
	return Coordinates{}
}

func (p *Puzzle) GetTile(pos Coordinates) Tile {
	return p.Tiles[pos.Y][pos.X]
}

func (p *Puzzle) UnlockIrrelevant() {
	for y, row := range p.Tiles {
		for x, t := range row {
			if !t.Relevant {
				p.Tiles[y][x].Locked = false
			}
		}
	}
}

func (p *Puzzle) GetTilePosWithTarget(target Coordinates) Coordinates {
	for y, row := range p.Tiles {
		for x, t := range row {
			if t.Target == target {
				return Coordinates{x, y}
			}
		}
	}
	log.Fatal("no such target")
	return Coordinates{}
}
