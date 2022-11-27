package algorithm

import (
	"log"
	"npuzzle/puzzle"
)

func AStar(p *puzzle.Puzzle, h func(puzzle.Puzzle) int) (stats Stats) {
	open := []*Node{}
	closed := make(map[string]*Node)

	open = append(open, newNode(*p, nil, h))
	stats.MaxStates, stats.TotalStates = 1, 1

	for len(open) != 0 {
		if stats.MaxStates < len(open)+len(closed) {
			stats.MaxStates = len(open) + len(closed)
		}
		current := popLowestF(&open)
		closed[current.hash()] = current
		if current.puzzle.IsSolved() {
			stats.Path = tracePath(current)
			stats.PathLen = len(stats.Path)
			*p = current.puzzle.DeepCopy()
			return
		}

		neighbours := allNeighbours(current, h)
		for _, n := range neighbours {
			if _, ok := closed[n.hash()]; ok {
				continue
			}
			index, ok := nodeIndex(n, open)
			if ok && (n.g < open[index].g ||
				n.f() == open[index].f() && n.h < open[index].h) {
				open[index] = n
			} else if !ok {
				open = append(open, n)
				stats.TotalStates++
			}
		}
	}
	log.Fatal("could not find the path, open set is empty")
	return
}

func nodeIndex(n *Node, slice []*Node) (int, bool) {
	for i, item := range slice {
		if item.hash() == n.hash() {
			return i, true
		}
	}
	return 0, false
}

func allNeighbours(n *Node, h func(puzzle.Puzzle) int) []*Node {
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

func popLowestF(open *[]*Node) *Node {
	if len(*open) == 0 {
		log.Fatal("empty open inside loop")
	}
	minF := int(^uint(0) >> 1)
	minH := int(^uint(0) >> 1)
	index := 0
	for i, n := range *open {
		if n.f() < minF || n.f() == minF && n.h < minH {
			minF = n.f()
			minH = n.h
			index = i
		}
	}
	result := (*open)[index]
	(*open)[index] = (*open)[len(*open)-1]
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