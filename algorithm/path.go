package algorithm

import (
	"fmt"
	"log"
	"npuzzle/puzzle"
)

func (s *Solver) makePath(src, dst puzzle.Coordinates) {
	s.Hor = Rectangle{}
	s.Ver = Rectangle{}

	dir := s.determineDirection(src, dst)
	fmt.Printf("dir=%s, src=%v, dst=%v\n", dir, src, dst)
	switch dir {
	case "left":
		s.leftPath(src, dst)
	case "right":
		s.rightPath(src, dst)
	case "top":
		s.topPath(src, dst)
	case "bot":
		s.botPath(src, dst)
	case "topleft":
		s.topLeftPath(src, dst)
	case "topright":
		s.topRightPath(src, dst)
	case "botleft":
		left := puzzle.Coordinates{dst.X, src.Y}
		right := src
		s.leftPath(right, left)
		top := left
		bot := dst
		s.botPath(top, bot)
		s.connectPaths(puzzle.Coordinates{dst.X, src.Y})
	case "botright":
		top := src
		bot := puzzle.Coordinates{src.X, dst.Y}
		s.botPath(top, bot)
		left := bot
		right := dst
		s.leftPath(left, right)
		s.connectPaths(puzzle.Coordinates{src.X, dst.Y})
	default:
		log.Fatalf("impossible direction: %s", dir)
	}
}

func (s *Solver) connectPaths(mid puzzle.Coordinates) {
	for y := -1; y < 2; y += 2 {
		for x := -1; x < 2; x += 2 {
			midNX := mid.X + x
			midNY := mid.Y + y
			fmt.Printf("midNX:%d, midnY:%d\n", midNX, midNY)
			if midNX < 0 || midNX > s.P.Size-1 ||
				midNY < 0 || midNY > s.P.Size-1 {
				continue
			}
			if !s.Ver.contains(midNX, midNY) && !s.Hor.contains(midNX, midNY) {
				s.Mid = puzzle.Coordinates{midNX, midNY}
				fmt.Println("mid:", s.Mid)
				return
			}
		}
	}
}

func (s *Solver) topLeftPath(src, dst puzzle.Coordinates) {
	top := puzzle.Coordinates{src.X, dst.Y}
	bot := src
	s.topPath(bot, top)
	left := dst
	right := top
	s.leftPath(right, left)
	s.connectPaths(puzzle.Coordinates{src.X, dst.Y})
	// if s.Hor.botRight.X == s.P.Size-1 ||
	// 	s.P.Tiles[s.Hor.topLeft.Y][s.Hor.botRight.X+1].Locked {
	// 	s.Ver.botRight.Y++
	// } else {
	// 	s.Hor.botRight.X++
	// }
}

func (s *Solver) topRightPath(src, dst puzzle.Coordinates) {
	left := src
	// if src.Y > 0 && !s.P.Tiles[src.Y-1][src.X].Locked {
	// 	left.Y--
	// }
	right := puzzle.Coordinates{dst.X, src.Y}
	s.rightPath(left, right)
	top := dst
	bot := right
	s.topPath(bot, top)
	s.connectPaths(puzzle.Coordinates{dst.X, src.Y})
	// if s.Hor.botRight.X == s.P.Size-1 ||
	// 	s.P.Tiles[s.Hor.botRight.Y][s.Hor.botRight.X].Locked {
	// 	s.Ver.botRight.Y++
	// } else {
	// 	s.Hor.botRight.X++
	// }
}

func (s *Solver) topPath(src, dst puzzle.Coordinates) {
	fmt.Printf("topPath: %v –> %v\n", src, dst)
	s.Ver.topLeft = dst
	s.Ver.botRight = src
	if dst.X == s.P.Size-1 || s.P.Tiles[dst.Y][dst.X+1].Locked {
		s.Ver.topLeft.X--
	} else {
		s.Ver.botRight.X++
	}
	if s.Ver.botRight.Y == 1 {
		s.Ver.botRight.Y++
	}
}

func (s *Solver) botPath(src, dst puzzle.Coordinates) {
	fmt.Printf("botPath: %v –> %v\n", src, dst)
	s.Ver.topLeft = src
	s.Ver.botRight = dst
	if dst.X == 0 || s.P.Tiles[dst.Y][dst.X-1].Locked {
		s.Ver.botRight.X++
	} else {
		s.Ver.topLeft.X--
	}
	if s.Ver.topLeft.Y == s.P.Size-2 {
		s.Ver.botRight.Y--
	}
}

func (s *Solver) leftPath(src, dst puzzle.Coordinates) {
	fmt.Printf("leftPath: %v –> %v\n", src, dst)
	s.Hor.topLeft = dst
	s.Hor.botRight = src
	if dst.Y == 0 || s.P.Tiles[dst.Y-1][dst.X].Locked {
		s.Hor.botRight.Y++
	} else {
		s.Hor.topLeft.Y--
	}
	if s.Hor.botRight.X == s.P.Size-2 {
		s.Hor.botRight.X++
	}
}

func (s *Solver) rightPath(src, dst puzzle.Coordinates) {
	fmt.Printf("rightPath: %v –> %v\n", src, dst)
	s.Hor.topLeft = src
	s.Hor.botRight = dst
	if dst.Y == s.P.Size-1 || s.P.Tiles[dst.Y+1][dst.X].Locked {
		s.Hor.topLeft.Y--
	} else {
		s.Hor.botRight.Y++
	}
	if s.Hor.topLeft.X == 1 {
		s.Hor.topLeft.X--
	}
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
	log.Fatalf("impossible direction: deltaX=%d, deltaY=%d", deltaX, deltaY)
	return ""
}
