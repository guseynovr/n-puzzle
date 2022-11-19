package puzzle

import (
	"math/rand"
	"time"
)

func Random(size int) *Puzzle {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	tiles := r.Perm(size * size)
	cells := to2D(tiles, size)
	emptyX, emptyY := getEmptyXY(cells)

	result := Puzzle{
		size:   size,
		cells:  cells,
		emptyX: emptyX,
		emptyY: emptyY,
	}
	result.target = result.targetState()
	return &result
}

func getEmptyXY(cells [][]int) (int, int) {
	for y, row := range cells {
		for x, cell := range row {
			if cell == 0 {
				return x, y
			}
		}
	}
	return 0, 0
}
