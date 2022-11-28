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

	for i := 0; i < s.P.Size*s.P.Size-8; i++ {
		src := s.P.GetPosition(i)
		dst := s.P.Tiles[src.Y][src.X].Target

		s.makePath(src, dst)
		s.debug(fmt.Sprintf("makePath: ver %v, hor %v\n", s.Ver, s.Hor))

		s.moveZeroToPath(src)
		s.debug("moveZeroToPath")

		s.lockPath()
		s.debug("lockPath")

		s.moveItem(src, dst)
		s.debug("moveItem")

		s.unlockPath()
		s.debug("unlockPath")
	}
	s.P.MakeAllRelevant()
	s.debug("MakeAllRelevant")
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
	// if dst.X == s.P.Size-1 && s.P.Tiles[dst.Y][dst.X-1].Locked {
	// 	s.P.Tiles[dst.Y][dst.X-1].Locked = false
	// }
	// if dst.Y == s.P.Size-1 {
	// 	s.P.Tiles[dst.Y-1][dst.X].Locked = false
	// }
	stats := AStar(s.P, s.H)
	s.P.Tiles[dst.Y][dst.X].Locked = true
	// if dst.X == s.P.Size-1 {
	// 	s.P.Tiles[dst.Y][dst.X-1].Locked = true
	// }
	// if dst.Y == s.P.Size-1 {
	// 	s.P.Tiles[dst.Y-1][dst.X].Locked = true
	// }
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
	dir := s.determineDirection(src, dst)
	s.makeVerticalPath(src, dst, dir)
	s.makeHorizontalPath(src, dst, dir)

	fmt.Printf("ver %v, hor %v\n", s.Ver, s.Hor)
}

func (s *Solver) determineDirection(src, dst puzzle.Coordinates) string {
	deltaX, deltaY := dst.X-src.X, dst.Y-src.Y
	if deltaX < 0 && deltaY == 0 {
		return "left"
	} else if deltaX > 0 && deltaY == 0 {
		return "right"
	} else if deltaX == 0 && deltaY < 0 {
		return "top"
	} else if deltaX == 0 && deltaY > 0 {
		return "bot"
	} else if deltaX < 0 && deltaY < 0 {
		return "topleft"
	} else if deltaX > 0 && deltaY < 0 {
		return "topright"
	} else if deltaX > 0 && deltaY > 0 {
		return "botright"
	} else if deltaX < 0 && deltaY > 0 {
		return "botleft"
	}
	log.Fatal("impossible direction")
	return ""
}

// func (s *Solver) makeFinalPath(src, dst puzzle.Coordinates) {

// 	s.Hor.topLeft.X = dst.X
// 	s.Hor.topLeft.Y = s.P.Size - 2

// 	s.Hor.botRight.X = src.X
// 	if src.X < 2 {
// 		s.Hor.botRight.X = 2
// 	}
// 	s.Hor.botRight.Y = s.P.Size - 1
// 	s.Ver = s.Hor
// }

func (s *Solver) lockPath() {
	for y, row := range s.P.Tiles {
		for x := range row {
			if s.Ver.contains(x, y) || s.Hor.contains(x, y) {
				s.P.Tiles[y][x].Locked = false
			} else {
				s.P.Tiles[y][x].Locked = true
			}
		}
	}
}

func (s *Solver) unlockPath() {
	for y, row := range s.P.Tiles {
		for x, t := range row {
			if t.Relevant {
				s.P.Tiles[y][x].Locked = true
			} else {
				s.P.Tiles[y][x].Locked = false
			}
		}
	}
}

func (s *Solver) makeVerticalPath(src, dst puzzle.Coordinates, dir string) {
	switch dir {
	case "left":
	case "right":
	case "top":
	case "bot":
	case "topleft":
	case "topright":
	case "botright":
	case "botleft":
	}
}

func (s *Solver) makeHorizontalPath(src, dst puzzle.Coordinates, dir string) {
	switch dir {
	case "left":
		s.Hor.topLeft = dst
		s.Hor.botRight = src
		if dst.Y == 0 || s.P.Tiles[dst.Y-1][dst.X].Locked {
			s.Hor.botRight.Y++
		} else {
			s.Hor.topLeft.Y--
		}
	case "right":
		s.Hor.topLeft = src
		s.Hor.botRight = dst
		if dst.Y == s.P.Size-1 || s.P.Tiles[dst.Y+1][dst.X].Locked {
			s.Hor.topLeft.Y--
		} else {
			s.Hor.botRight.Y++
		}
	case "top":
	case "bot":
	case "topleft":
	case "topright":
	case "botright":
	case "botleft":
	}
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
