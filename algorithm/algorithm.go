package algorithm

import (
	"bufio"
	"fmt"
	"log"
	"npuzzle/puzzle"
	"os"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

func (s Solver) Solve() Stats {
	start := time.Now()
	s.P.MakeAllIrrelevant()

	y := 0
	for ; y < s.P.Size-2; y++ {
		for x := 0; x < s.P.Size; x++ {
			src := s.P.GetTilePosWithTarget(puzzle.Coordinates{x, y})
			dst := s.P.Tiles[src.Y][src.X].Target

			s.makePath(src, dst)

			s.moveZeroToPath(src)
			s.debug("moveZeroToPath")
			s.lockPath()
			s.debug("lockPath")

			s.moveItem(src, dst)
			s.debug("moveItem")

			s.unlockPath()
			s.debug("unlockPath")
		}
	}
	for x := 0; x < s.P.Size-3; x++ {
		for y := y; y < s.P.Size; y++ {
			src := s.P.GetTilePosWithTarget(puzzle.Coordinates{x, y})
			dst := s.P.Tiles[src.Y][src.X].Target

			s.makeFinalPath(src, dst)

			s.moveZeroToPath(src)
			s.debug("moveZeroToPath")

			s.lockPath()
			s.debug("lockPath")

			s.moveItem(src, dst)
			s.debug("moveItem")

			s.unlockPath()
			s.debug("unlockPath")
		}
	}
	s.P.MakeAllRelevant()
	s.debug("last section")
	s.Stats = s.Stats.Append(AStar(s.P, s.H))
	s.debug("at the end")
	s.Stats.t = time.Since(start)
	return s.Stats
}

func (s *Solver) LockLastPart() {
	for y := s.P.Size - 2; y < s.P.Size; y++ {
		for x := s.P.Size - 3; x < s.P.Size; x++ {
			s.P.Tiles[y][x].Locked = true
		}
	}
}

func (s *Solver) moveItem(src, dst puzzle.Coordinates) {
	s.P.Tiles[src.Y][src.X].Relevant = true
	if dst.X == s.P.Size-1 && s.P.Tiles[dst.Y][dst.X-1].Locked {
		s.P.Tiles[dst.Y][dst.X-1].Locked = false
	}
	if dst.Y == s.P.Size-1 {
		s.P.Tiles[dst.Y-1][dst.X].Locked = false
	}
	stats := AStar(s.P, s.H)
	s.P.Tiles[dst.Y][dst.X].Locked = true
	if dst.X == s.P.Size-1 {
		s.P.Tiles[dst.Y][dst.X-1].Locked = true
	}
	if dst.Y == s.P.Size-1 {
		s.P.Tiles[dst.Y-1][dst.X].Locked = true
	}
	s.Stats = stats.Append(stats)
}

func (s *Solver) moveZeroToPath(next puzzle.Coordinates) {

	s.P.Tiles[next.Y][next.X].Locked = true
	zeroPos := s.zeroInPathPos(s.P.Zero, next)
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = zeroPos
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = true
	stats := AStar(s.P, s.H)
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = puzzle.Coordinates{-1, -1}
	s.P.Tiles[next.Y][next.X].Locked = false
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = false

	s.Stats = s.Stats.Append(stats)
}

func (s *Solver) makePath(src, dst puzzle.Coordinates) {
	s.makeVerticalPath(src, dst)
	s.makeHorizontalPath(src, dst)

	fmt.Printf("ver %v, hor %v\n", s.Ver, s.Hor)
}

func (s *Solver) makeFinalPath(src, dst puzzle.Coordinates) {

	s.Hor.topLeft.X = dst.X
	s.Hor.topLeft.Y = s.P.Size - 2

	s.Hor.botRight.X = src.X
	if src.X < 2 {
		s.Hor.botRight.X = 2
	}
	s.Hor.botRight.Y = s.P.Size - 1
	s.Ver = s.Hor
}

func (s *Solver) lockPath() {
	for y, row := range s.P.Tiles {
		for x := range row {
			if s.Ver.contains(x, y) || s.Hor.contains(x, y) {
				continue
			}
			s.P.Tiles[y][x].Locked = true
		}
	}
}

func (s *Solver) unlockPath() {
	for y, row := range s.P.Tiles {
		for x, t := range row {
			if t.Relevant {
				continue
			}
			s.P.Tiles[y][x].Locked = false
		}
	}
}

func (s *Solver) makeVerticalPath(src, dst puzzle.Coordinates) {

	s.Ver.topLeft.X = dst.X
	if dst.X == s.P.Size-1 {
		s.Ver.topLeft.X = dst.X - 1
	}
	s.Ver.topLeft.Y = dst.Y

	s.Ver.botRight.X = s.Ver.topLeft.X + 1
	s.Ver.botRight.Y = src.Y
	if src.Y-dst.Y < 2 {
		s.Ver.botRight.Y = dst.Y + 2
	}
}

func (s *Solver) makeHorizontalPath(src, dst puzzle.Coordinates) {

	s.Hor.topLeft.X = dst.X
	if dst.X > src.X {
		s.Hor.topLeft.X = src.X
	}
	if s.Hor.topLeft.X == s.P.Size-1 {
		s.Hor.topLeft.X -= 1
	}
	s.Hor.topLeft.Y = src.Y - 1
	if src.Y > 0 && s.P.Tiles[src.Y-1][src.X].Locked {
		s.Hor.topLeft.Y = src.Y
	}
	if src.Y-dst.Y < 1 {
		s.Hor.topLeft.Y = dst.Y + 1
	}
	if src.Y == dst.Y {
		s.Hor.topLeft.Y = src.Y
	}

	s.Hor.botRight.X = src.X
	if dst.X > src.X {
		s.Hor.botRight.X = dst.X
	}
	s.Hor.botRight.Y = s.Hor.topLeft.Y + 1
}

func (s *Solver) debug(msg string) {
	if !s.Debug {
		return
	}
	fmt.Println(msg)
	fmt.Print(s.P)
	if scanner.Scan() {
		scanner.Text()
	}
}

func (s *Solver) zeroInPathPos(zero puzzle.Coordinates,
	next puzzle.Coordinates) puzzle.Coordinates {

	neighbours := []puzzle.Coordinates{}
	if next.X > 0 {
		neighbours = append(neighbours, puzzle.Coordinates{next.X - 1, next.Y})
	}
	if next.X < s.P.Size-1 {
		neighbours = append(neighbours, puzzle.Coordinates{next.X + 1, next.Y})
	}
	if next.Y > 0 {
		neighbours = append(neighbours, puzzle.Coordinates{next.X, next.Y - 1})
	}
	if next.Y < s.P.Size-1 {
		neighbours = append(neighbours, puzzle.Coordinates{next.X, next.Y + 1})
	}
	for _, n := range neighbours {
		if (s.Ver.contains(n.X, n.Y) || s.Hor.contains(n.X, n.Y)) &&
			!s.P.Tiles[n.Y][n.X].Locked {
			return n
		}
	}
	log.Fatal("zero pos inside path not found")
	return puzzle.Coordinates{}
}
