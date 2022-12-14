package algorithm

import (
	"fmt"
	"log"
	"math"
	"npuzzle/puzzle"
	"sort"
	"time"
)

func (s Solver) Solve() Stats {
	start := time.Now()
	s.P.MakeAllIrrelevant()
	var prev puzzle.Coordinates

	for i := 1; i < s.P.Size*s.P.Size-8; i++ {
		s.nextTile(i, &prev)
	}
	s.ByH = false
	s.P.Next = 0
	s.P.MakeAllRelevant()
	s.debug("MakeAllRelevant")
	s.Stats = s.Stats.Append(s.AStar())
	s.Stats.PathLen++
	s.debug("at the end")
	s.Stats.t = time.Since(start)
	s.Stats.PathLen--
	return s.Stats
}

func (s *Solver) nextTile(i int, prev *puzzle.Coordinates) {
	s.P.Next = i
	src := s.P.GetPosition(i)
	dst := s.P.Tiles[src.Y][src.X].Target

	if src == dst {
		s.P.Tiles[src.Y][src.X].Locked = true
		s.P.Tiles[src.Y][src.X].Relevant = true
		s.debug("src=dst, continue")
		*prev = dst
		return
	}

	s.moveZeroToNext(src)
	s.debug("moveZeroToNext")

	inCorner := s.inCorner(dst)
	if inCorner {
		s.cornerTile(&src, dst, *prev)
	}

	s.moveItem(src, dst)
	s.debug("moveItem")

	if inCorner {
		s.P.Tiles[prev.Y][prev.X].Locked = true
	}
	*prev = dst

	s.unlockPath()
	s.debug("unlockPath")
}

func (s *Solver) cornerTile(src *puzzle.Coordinates, dst, prev puzzle.Coordinates) {
	s.P.Tiles[prev.Y][prev.X].Locked = false

	s.makeCorner(dst)
	s.debug(fmt.Sprintln("makeCorner:", s.Corner))

	dst2 := s.zeroInCorner(*src, dst)
	s.P.Tiles[src.Y][src.X].Target = dst2
	s.debug(fmt.Sprintf("dst2:%v", dst2))

	s.moveItem(*src, dst2)
	s.P.Tiles[dst2.Y][dst2.X].Target = dst
	s.P.Tiles[dst2.Y][dst2.X].Locked = false
	s.debug("moveItem to corner")

	s.moveZeroToCorner(dst2)
	s.debug("moveZeroToNext")

	s.lockCorner()
	s.debug("lockCorner")

	*src = dst2
}

func (s *Solver) makeCorner(dst puzzle.Coordinates) {
	s.Corner.topLeft.X = dst.X - 1
	s.Corner.topLeft.Y = dst.Y - 2
	s.Corner.botRight.X = dst.X + 1
	s.Corner.botRight.Y = dst.Y + 2
}

func (s *Solver) inCorner(dst puzzle.Coordinates) bool {
	return ((dst.X == 0 || s.P.Tiles[dst.Y][dst.X-1].Locked) &&
		(dst.X == s.P.Size-1 || s.P.Tiles[dst.Y][dst.X+1].Locked)) ||
		((dst.Y == 0 || s.P.Tiles[dst.Y-1][dst.X].Locked) &&
			(dst.Y == s.P.Size-1 || s.P.Tiles[dst.Y+1][dst.X].Locked))
}

func (s *Solver) moveItem(src, dst puzzle.Coordinates) {
	s.P.Tiles[src.Y][src.X].Relevant = true
	s.Stats = s.Stats.Append(s.AStar())
	s.Stats.Path = s.Stats.Path[:len(s.Stats.Path)-1]
	s.P.Tiles[dst.Y][dst.X].Locked = true
}

func (s *Solver) moveZeroToNext(next puzzle.Coordinates) {
	s.P.Tiles[next.Y][next.X].Locked = true
	zeroPos := s.zeroNearNext(s.P.Zero, next)
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = zeroPos
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = true
	s.Stats = s.Stats.Append(s.AStar())
	s.Stats.Path = s.Stats.Path[:len(s.Stats.Path)-1]
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = puzzle.Coordinates{-1, -1}
	s.P.Tiles[next.Y][next.X].Locked = false
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = false
}

func (s *Solver) lockCorner() {
	for y, row := range s.P.Tiles {
		for x := range row {
			if !s.Corner.contains(x, y) /*  && !s.P.Tiles[y][x].Relevant */ {
				s.P.Tiles[y][x].Locked = true
			} else {
				s.P.Tiles[y][x].Locked = false
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

func (s *Solver) moveZeroToCorner(next puzzle.Coordinates) {
	s.P.Tiles[next.Y][next.X].Locked = true
	s.P.Tiles[next.Y][next.X].Relevant = false
	zeroPos := s.zeroInCorner(s.P.Zero, next)
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = zeroPos
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = true
	stats := s.AStar()
	s.Stats.Path = s.Stats.Path[:len(s.Stats.Path)-1]
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Target = puzzle.Coordinates{-1, -1}
	s.P.Tiles[next.Y][next.X].Locked = false
	s.P.Tiles[next.Y][next.X].Relevant = true
	s.P.Tiles[s.P.Zero.Y][s.P.Zero.X].Relevant = false

	s.Stats = s.Stats.Append(stats)
}

func (s *Solver) zeroInCorner(zero puzzle.Coordinates,
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
	sort.Slice(neighbours, func(i, j int) bool {
		d1 := int(math.Sqrt(
			math.Pow(float64(neighbours[i].X-zero.X), 2)+
				math.Pow(float64(neighbours[i].Y-zero.Y), 2)) * 10)
		d2 := int(math.Sqrt(
			math.Pow(float64(neighbours[j].X-zero.X), 2)+
				math.Pow(float64(neighbours[j].Y-zero.Y), 2)) * 10)
		return d1 < d2
	})
	for _, n := range neighbours {
		if s.Corner.contains(n.X, n.Y) && !s.P.Tiles[n.Y][n.X].Locked &&
			!s.P.Tiles[n.Y][n.X].Relevant {
			return n
		}
	}

	log.Fatalf("zero pos inside path not found: zero=%v, next=%v\n", zero, next)
	return puzzle.Coordinates{}
}

func (s *Solver) zeroNearNext(zero puzzle.Coordinates,
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
	sort.Slice(neighbours, func(i, j int) bool {
		d1 := int(math.Sqrt(
			math.Pow(float64(neighbours[i].X-zero.X), 2)+
				math.Pow(float64(neighbours[i].Y-zero.Y), 2)) * 10)
		d2 := int(math.Sqrt(
			math.Pow(float64(neighbours[j].X-zero.X), 2)+
				math.Pow(float64(neighbours[j].Y-zero.Y), 2)) * 10)
		return d1 < d2
	})
	for _, n := range neighbours {
		if !s.P.Tiles[n.Y][n.X].Locked {
			return n
		}
	}
	log.Fatal("zero pos inside path not found")
	return puzzle.Coordinates{}
}
