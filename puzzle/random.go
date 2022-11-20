package puzzle

import (
	"math/rand"
	"time"
)

// Random generates random puzzle with given size.
func Random(size int) *Puzzle {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	tiles1D := r.Perm(size * size)
	tiles := to2D(tiles1D, size)
	emptyX, emptyY := getEmptyXY(tiles)

	return newPuzzle(size, emptyX, emptyY, tiles)
}

func getEmptyXY(tiles [][]int) (int, int) {
	for y, row := range tiles {
		for x, cell := range row {
			if cell == 0 {
				return x, y
			}
		}
	}
	return 0, 0
}
