package puzzle

func (p *Puzzle) IsSolvable() bool {
	invCount := p.countInversions()

	if p.size%2 != 0 {
		return invCount%2 == 0
	}
	emptyPosFromBottom := p.size - p.emptyY
	if emptyPosFromBottom%2 == 0 {
		return invCount%2 != 0
	} else {
		return invCount%2 == 0
	}
}

func (p *Puzzle) countInversions() int {
	var tiles []int
	if p.size%2 == 0 {
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

// expand returns p.cells in form of a 1D array
func (p *Puzzle) expand() []int {
	result := make([]int, 0, p.size*p.size)
	for _, row := range p.cells {
		result = append(result, row...)
	}
	return result
}

const (
	right = iota
	down
	left
	up
)

// expandSnail returns p.cells for snail solution in form of a 1D array
func (p *Puzzle) expandSnail() []int {
	tileCnt := p.size * p.size
	result := make([]int, 0, tileCnt)
	p.iterateSnail(func(i int) {
		result = append(result, i)
	})
	return result
}

func (p *Puzzle) iterateSnail(f func(int)) {
	tileCnt := p.size * p.size
	i := 0
	x, y, xMax, yMax, xMin, yMin := 0, 0, p.size, p.size, 0, 0
	dir := right
	change := false
	for i < tileCnt {
		if !change {
			f(p.cells[y][x])
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
