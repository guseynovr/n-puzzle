/*
Package puzzle provides a type and functions for creating, manipulating,
and analyzing npuzzle.
asd
*/
package puzzle

import "fmt"

type Tile struct {
	Value    int
	Relevant bool
	Locked   bool
	Target   Coordinates
}

type Coordinates struct {
	X, Y int
}

type Puzzle struct {
	Size  int
	Tiles [][]Tile
	Zero  Coordinates
	Next  int
	// Target [][]int
	hash string
	// TargetXY map[int]Coordinates
}

func newPuzzle(size, x, y int, tiles [][]int) *Puzzle {
	p := Puzzle{
		Size: size,
		Zero: Coordinates{x, y},
	}
	p.initTiles(tiles)
	p.initTargetState()
	p.updateHash()
	return &p
}

func (p *Puzzle) initTiles(tiles [][]int) {
	p.Tiles = make([][]Tile, 0, p.Size)
	for _, row := range tiles {
		tileRow := make([]Tile, 0, p.Size)
		for _, v := range row {
			tile := Tile{Value: v, Relevant: true}
			tileRow = append(tileRow, tile)
		}
		p.Tiles = append(p.Tiles, tileRow)
	}
}

func (p *Puzzle) Hash() string {
	return p.hash
}

func (p *Puzzle) updateHash() {
	p.hash = fmt.Sprintf("%d%v%v", p.Size, p.Zero, p.Tiles)
}

func (p *Puzzle) DeepCopy() Puzzle {
	pCopy := *p
	pCopy.Tiles = make([][]Tile, p.Size)
	for i := range pCopy.Tiles {
		pCopy.Tiles[i] = make([]Tile, p.Size)
		copy(pCopy.Tiles[i], p.Tiles[i])
	}
	return pCopy
}
