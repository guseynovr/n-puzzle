package algorithm

import (
	"fmt"
	"npuzzle/puzzle"
	"strings"
	"time"
)

type Stats struct {
	Heuristics  string
	TotalStates int // complexity in time
	MaxStates   int // complexity in size
	PathLen     int
	Path        []puzzle.Puzzle
	t           time.Duration
}

func (s Stats) String() string {
	sb := strings.Builder{}
	// sb.WriteString("Stats:\n")
	sb.WriteString(fmt.Sprintf("Complexity in time: %d\n", s.TotalStates))
	sb.WriteString(fmt.Sprintf("Complexity in size: %d\n", s.MaxStates))
	sb.WriteString(fmt.Sprintf("Path len: %d\n", s.PathLen))
	// sb.WriteString("Path sequence:\n")
	// for _, st := range s.Path {
	// 	sb.WriteString(st.String() + "\n")
	// }
	sb.WriteString("Time required: " + s.t.String())
	return sb.String()
}

func (s Stats) Append(s2 Stats) Stats {
	fmt.Printf("append: Total=%d, Max=%d, PathLen=%d\n",
		s2.TotalStates, s2.MaxStates, s2.PathLen)
	return Stats{
		TotalStates: s.TotalStates + s2.TotalStates,
		MaxStates:   s.MaxStates + s2.MaxStates,
		PathLen:     s.PathLen + s2.PathLen,
		Path:        append(s.Path, s2.Path...),
	}
}
