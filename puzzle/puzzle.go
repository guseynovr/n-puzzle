/*
Package puzzle provides a type and functions for creating, manipulating,
and analyzing npuzzle.
asd
*/
package puzzle

import "fmt"

type Puzzle struct {
	Size     int
	Tiles    [][]int
	emptyX   int
	emptyY   int
	Target   [][]int
	TargetXY map[int]struct{ X, Y int }
	hash     string
}

func newPuzzle(size, x, y int, tiles [][]int) *Puzzle {
	p := Puzzle{
		Size:   size,
		Tiles:  tiles,
		emptyX: x,
		emptyY: y,
	}
	p.Target = p.targetState()
	p.updateHash()
	return &p
}

func (p *Puzzle) Hash() string {
	return p.hash
}

func (p *Puzzle) updateHash() {
	p.hash = fmt.Sprintf("%d%d%d%v", p.Size, p.emptyX, p.emptyY, p.Tiles)
}
