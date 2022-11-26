package puzzle

// IsSolvable
func (p *Puzzle) IsSolvable() bool {
	invCount := p.countInversions()

	if p.Size%2 != 0 {
		return invCount%2 == 0
	}
	emptyPosFromBottom := p.Size - p.Zero.Y
	res := false
	if emptyPosFromBottom%2 == 0 {
		res = invCount%2 != 0
	} else {
		res = invCount%2 == 0
	}
	if p.Size%8 == 0 || p.Size%8 == 6 {
		res = !res
	}
	return res
}

func (p *Puzzle) countInversions() int {
	var tiles []int
	if p.Size%2 == 0 {
		tiles = p.expand()
	} else {
		tiles = p.expandSnail()
	}
	// fmt.Println(tiles)

	count := 0
	for i, t := range tiles {
		if t == 0 {
			continue
		}
		for j := i + 1; j < len(tiles); j++ {
			if tiles[j] == 0 {
				continue
			}
			if t > tiles[j] {
				count++
			}
		}
	}
	return count
}

// expand returns p.tiles in form of a 1D array
func (p *Puzzle) expand() []int {
	result := make([]int, 0, p.Size*p.Size)
	for _, row := range p.Tiles {
		for _, t := range row {
			result = append(result, t.Value)
		}
	}
	return result
}

const (
	right = iota
	down
	left
	up
)

// expandSnail returns p.tiles for snail solution in form of a 1D array
func (p *Puzzle) expandSnail() []int {
	tileCnt := p.Size * p.Size
	result := make([]int, 0, tileCnt)
	p.iterateSnail(func(x, y int) {
		result = append(result, p.Tiles[y][x].Value)
	})
	return result
}

func (p *Puzzle) iterateSnail(f func(int, int)) {
	tileCnt := p.Size * p.Size
	i := 0
	x, y, xMax, yMax, xMin, yMin := 0, 0, p.Size, p.Size, 0, 0
	dir := right
	change := false
	for i < tileCnt {
		if !change {
			f(x, y)
			i++
		}
		change = false
		switch dir {
		case right:
			x++
			if x == xMax {
				xMax--
				x--
				dir = down
				change = true
			}
		case down:
			y++
			if y == yMax {
				yMax--
				y--
				dir = left
				change = true
			}
		case left:
			x--
			if x < xMin {
				xMin++
				x++
				dir = up
				change = true
			}
		case up:
			y--
			if y == yMin {
				yMin++
				y++
				dir = right
				change = true
			}
		}
	}
}
