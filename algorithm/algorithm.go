package algorithm

import (
	"fmt"
	"math"
	"npuzzle/puzzle"
	"time"
)

func Solve(p *puzzle.Puzzle, h func(puzzle.Puzzle) int) (stats Stats) {
	start := time.Now()
	p.MakeAllIrrelevant()
	zeroTarget := p.GetPosition(0)
	border := p.Size - 1

	for i := 1; i < p.Size*p.Size; i++ {
		// find i pos
		iPos := p.GetPosition(i)
		iTarget := p.Tiles[iPos.Y][iPos.X].Target
		fmt.Printf("i: %d, pos %v\n", i, iPos)
		// find closest pos for 0 around i
		neighbour := ClosestNeighbour(iPos, iTarget,
			p.Zero, p.Size)
		// set that pos as target for 0
		// TODO:: if is in pos, continue to next i
		// TODO: same with 0
		p.Tiles[p.Zero.Y][p.Zero.X].Target = neighbour
		// lock zero
		p.Tiles[p.Zero.Y][p.Zero.X].Relevant = true
		//  mv zero to target
		p.Tiles[iPos.Y][iPos.X].Locked = true
		fmt.Printf("moving 0 to target %v\n", neighbour)
		if p.Zero.X != neighbour.X || p.Zero.Y != neighbour.Y {
			stats = stats.Append(AStar(p, h))
		}
		// unlock zero
		p.Tiles[p.Zero.Y][p.Zero.X].Relevant = false
		// lock i
		p.Tiles[iPos.Y][iPos.X].Relevant = true

		fmt.Printf("0 in place, moving %d to target %v\n",
			i, p.Tiles[iPos.Y][iPos.X].Target)
		// mv i to target
		p.Tiles[iPos.Y][iPos.X].Locked = false
		if iPos.X != iTarget.X || iPos.Y != iTarget.Y {
			stats = stats.Append(AStar(p, h))
		}
		fmt.Printf("%d in place(%v)\n", i, iTarget)
		if i%border != 0 {
			p.Tiles[iTarget.Y][iTarget.X].Locked = true
		}
		if i%border == 1 && i > 1 {
			border--
			prevPos := p.GetPosition(i - 1)
			p.Tiles[prevPos.Y][prevPos.X].Locked = true
			fmt.Printf("locked %d\n", p.Tiles[prevPos.Y][prevPos.X].Value)
		}
		if i == 6 || i == 7 {
			fmt.Println(p.Tiles)
		}
	}
	p.MakeAllRelevant()
	p.Tiles[p.Zero.Y][p.Zero.X].Target = zeroTarget
	stats = stats.Append(AStar(p, h))
	stats.t = time.Since(start)
	return
}

func ClosestNeighbour(iPos, iTarget, zero puzzle.Coordinates,
	size int) puzzle.Coordinates {

	neighbours := []puzzle.Coordinates{}
	if iPos.X > 0 {
		neighbours = append(neighbours, puzzle.Coordinates{iPos.X - 1, iPos.Y})
	}
	if iPos.X < size-1 {
		neighbours = append(neighbours, puzzle.Coordinates{iPos.X + 1, iPos.Y})
	}
	if iPos.Y > 0 {
		neighbours = append(neighbours, puzzle.Coordinates{iPos.X, iPos.Y - 1})
	}
	if iPos.Y < size-1 {
		neighbours = append(neighbours, puzzle.Coordinates{iPos.X, iPos.Y + 1})
	}
	fmt.Println("iPos", iPos, "neighbours", neighbours)
	minDist := int(^uint(0) >> 1)
	resIndex := 0
	for i, n := range neighbours {
		dist := int(
			math.Abs(float64(iTarget.X-n.X)) +
				math.Abs(float64(iTarget.Y-n.Y))*10)
		if dist < minDist {
			minDist = dist
			resIndex = i
		}
	}
	return neighbours[resIndex]
}
