package algorithm

import (
	"fmt"
	"npuzzle/puzzle"
)

type Node struct {
	puzzle puzzle.Puzzle
	parent *Node
	g      int // pathLen from start to current node
	h      int // pathlen from current node to target node estimated with h(n)
}

func newNode(p puzzle.Puzzle, parent *Node, h func(puzzle.Puzzle) int) *Node {
	n := Node{
		puzzle: p,
		parent: parent,
		h:      h(p),
	}
	if parent != nil {
		n.g = parent.g + 10
	}
	return &n
}

func (n Node) f() int {
	return n.g + n.h
}

func (n Node) hash() string {
	return n.puzzle.Hash()
}

func (n Node) String() string {
	return fmt.Sprintf("(g=%d,h=%d)", n.g, n.h)
}
