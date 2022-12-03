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

	for i := 1; i < s.P.Size*s.P.Size-8; i++ {
		s.P.Next = i
		s.Mid = puzzle.Coordinates{-1, -1}
		src := s.P.GetPosition(i)
		dst := s.P.Tiles[src.Y][src.X].Target

		if src == dst {
			s.P.Tiles[src.Y][src.X].Locked = true
			s.P.Tiles[src.Y][src.X].Relevant = true
			s.debug("src=dst, continue")
			continue
		}

		s.makePath(src, dst)
		s.debug(fmt.Sprintf("makePath: ver %v, hor %v\n", s.Ver, s.Hor))

		if s.Mid.X == -1 {
			s.moveZeroToPath(src)
			s.debug("moveZeroToPath")

			s.lockPath(true)
			s.debug("lockPath")

			s.moveItem(src, dst)
			s.debug("moveItem")

			s.unlockPath()
			s.debug("unlockPath")
			continue
		}

		s.moveZeroToPathMid(src, true)
		s.debug("moveZeroToPath")

		s.lockPath(true)
		s.debug("lockPath")

		s.P.Tiles[src.Y][src.X].Target = s.Mid
		s.moveItem(src, s.Mid)
		s.debug("moveItem")

		s.unlockPath()
		s.debug("unlockPath")

		s.moveZeroToPathMid(src, false)
		s.debug("moveZeroToPath")

		s.lockPath(false)
		s.debug("lockPath")

		s.P.Tiles[s.Mid.Y][s.Mid.X].Target = src
		s.moveItem(s.Mid, dst)
		s.debug("moveItem")

		s.unlockPath()
		s.debug("unlockPath")
	}
	s.P.Next = 0
	s.P.MakeAllRelevant()
	s.debug("MakeAllRelevant")
	s.Stats = s.Stats.Append(s.AStar())
	s.debug("at the end")
	s.Stats.t = time.Since(start)
	return s.Stats
}

// func (s *Solver) LockLastPart() {
// 	for y := s.P.Size - 2; y < s.P.Size; y++ {
// 		for x := s.P.Size - 3; x < s.P.Size; x++ {
// 			s.P.Tiles[y][x].Locked = true
// 		}
// 	}
// }

func (s *Solver) moveItem(src, dst puzzle.Coordinates) {
	s.P.Tiles[src.Y][src.X].Relevant = true
	fmt.Printf("moveItem: src=%v, dst%v\n", src, dst)
	stats := s.AStar()
	s.P.Tiles[dst.Y][dst.X].Locked = true
	s.Stats = stats.Append(stats)
}

func (s *Solver) moveZeroToPathMid(next puzzle.Coordinates, ver bool) {
	s.P.Tiles[next.Y][next.X].Locked = true
	zeroPos := s.zeroInPathPosMid(s.P.Zero, next, ver)
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = zeroPos
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = true
	stats := s.AStar()
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = puzzle.Coordinates{-1, -1}
	s.P.Tiles[next.Y][next.X].Locked = false
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = false

	s.Stats = s.Stats.Append(stats)
}

func (s *Solver) moveZeroToPath(next puzzle.Coordinates) {
	s.P.Tiles[next.Y][next.X].Locked = true
	zeroPos := s.zeroInPathPos(s.P.Zero, next)
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = zeroPos
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = true
	stats := s.AStar()
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = puzzle.Coordinates{-1, -1}
	s.P.Tiles[next.Y][next.X].Locked = false
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = false

	s.Stats = s.Stats.Append(stats)
}

func (s *Solver) lockPath(ver bool) {
	// fmt.Printf("lockPath tiles 0,0:%#v\n", s.P.Tiles[0][0])
	for y, row := range s.P.Tiles {
		for x := range row {
			if (ver && s.Ver.contains(x, y)) || (!ver && s.Hor.contains(x, y)) ||
				(s.Mid.X == x && s.Mid.Y == y) {
				s.P.Tiles[y][x].Locked = false
			} else {
				s.P.Tiles[y][x].Locked = true
			}
		}
	}
	// fmt.Printf("lockPath tiles 0,0:%#v\n", s.P.Tiles[0][0])
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

func (s *Solver) zeroInPathPosMid(zero puzzle.Coordinates,
	next puzzle.Coordinates, ver bool) puzzle.Coordinates {

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
		if ((ver && s.Ver.contains(n.X, n.Y)) ||
			(!ver && s.Hor.contains(n.X, n.Y))) &&
			!s.P.Tiles[n.Y][n.X].Locked {
			return n
		}
	}
	log.Fatal("zero pos inside path not found")
	return puzzle.Coordinates{}
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
