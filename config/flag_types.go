package config

import (
	"fmt"
	"npuzzle/algorithm"
	"npuzzle/puzzle"
	"strconv"
	"strings"
)

type size int

func (s *size) UnmarshalText(text []byte) error {
	n, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}
	if n < 3 {
		return fmt.Errorf("size cannot be less than 3")
	}
	*s = size(n)
	return nil
}

func (s size) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprint(int(s))), nil
}

type heuristic struct {
	F    func(puzzle.Puzzle) int
	Desc string
}

func (h *heuristic) UnmarshalText(text []byte) error {
	str := strings.ToLower(string(text))
	switch str {
	case "manhattan":
		*h = heuristic{
			F:    algorithm.Manhattan,
			Desc: "manhattan",
		}
	case "tiles":
		*h = heuristic{
			F:    algorithm.OutOfPlace,
			Desc: "tiles",
		}
	case "euclidean":
		*h = heuristic{
			F:    algorithm.Euclidean,
			Desc: "euclidean",
		}
	case "diagonal":
		*h = heuristic{
			F:    algorithm.Diagonal,
			Desc: "diagonal",
		}
	default:
		return fmt.Errorf("unsupported heuristic: %s", text)
	}
	return nil
}

func (h heuristic) MarshalText() ([]byte, error) {
	return []byte(h.Desc), nil
}
