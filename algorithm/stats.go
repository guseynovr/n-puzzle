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