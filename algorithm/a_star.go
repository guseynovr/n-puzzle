package algorithm

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"npuzzle/puzzle"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func (s *Solver) AStar() (stats Stats) {
	open := &NodeQueue{byH: s.ByH}
	closed := make(map[string]*Node)

	heap.Push(open, newNode(*s.P, nil, s.H))
	heap.Init(open)
	stats.MaxStates, stats.TotalStates = 1, 1

	for open.Len() != 0 {
		if stats.MaxStates < open.Len()+len(closed) {
			stats.MaxStates = open.Len() + len(closed)
		}
		current := heap.Pop(open).(*Node)
		stats.TotalStates++
		closed[current.hash()] = current
		// s.debugAStar(current)
		if current.puzzle.IsSolved() {
			stats.Path = tracePath(current)
			stats.PathLen = len(stats.Path)
			*s.P = current.puzzle
			return
		}

		neighbours := getNeighbours(current, s.H)
		for _, n := range neighbours {
			if _, ok := closed[n.hash()]; ok {
				continue
			}
			index, ok := open.index(n)
			if ok && n.g < open.nodes[index].g {
				open.nodes[index] = n
				heap.Fix(open, index)
			} else if !ok {
				heap.Push(open, n)
			}
		}
		heap.Init(open)
	}
	log.Fatal("could not find the path, open set is empty")
	return
}

func (s *Solver) debugAStar(node *Node) {
	if !s.Debug {
		return
	}
	fmt.Printf("AStar: g=%d, h=%d\n%s\n", node.g, node.h, node.puzzle)
	if scanner.Scan() {
		scanner.Text()
	}
}

func nodeIndex(n *Node, slice []*Node) (int, bool) {
	for i, item := range slice {
		if item.hash() == n.hash() {
			return i, true
		}
	}
	return 0, false
}

func getNeighbours(n *Node, h func(puzzle.Puzzle) int) []*Node {
	res := []*Node{}
	if left, err := n.puzzle.FilledFromLeft(); err == nil {
		res = append(res, newNode(left, n, h))
	}
	if right, err := n.puzzle.FilledFromRight(); err == nil {
		res = append(res, newNode(right, n, h))
	}
	if top, err := n.puzzle.FilledFromAbove(); err == nil {
		res = append(res, newNode(top, n, h))
	}
	if bottom, err := n.puzzle.FilledFromBelow(); err == nil {
		res = append(res, newNode(bottom, n, h))
	}
	return res
}

func (s *Solver) popLowestF(open *[]*Node) *Node {
	if len(*open) == 0 {
		log.Fatal("empty open inside loop")
	}
	minF := int(^uint(0) >> 1)
	minH := int(^uint(0) >> 1)
	index := 0
	for i, n := range *open {
		if (!s.ByH && (n.f() < minF || (n.f() == minF && n.h < minH))) ||
			(s.ByH && (n.h < minH || (n.h == minH && n.f() < minF))) {
			minF = n.f()
			minH = n.h
			index = i
		}
	}
	result := (*open)[index]
	(*open)[index] = (*open)[len(*open)-1]
	(*open)[len(*open)-1] = nil
	*open = (*open)[:len(*open)-1]
	return result
}

func tracePath(node *Node) []puzzle.Puzzle {
	path := []puzzle.Puzzle{}
	for node.parent != nil {
		path = append(path, node.puzzle)
		node = node.parent
	}
	path = append(path, node.puzzle)
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}
	return path
}
