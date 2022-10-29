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
	tiles := p.expand()

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
