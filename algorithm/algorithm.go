package algorithm

import "npuzzle/puzzle"

func AStar(p puzzle.Puzzle, h func(puzzle.Puzzle) int) (stats Stats) {
	open := []*Node{}
	closed := make(map[string]*Node)

	open = append(open, newNode(p, nil, h))
	for {
		current := popLowestF(&open)
		closed[current.hash()] = current
		if current.puzzle.IsSolved() {
			//TODO: update stats?
			return
		}
	}
	return
}

func popLowestF(open *[]*Node) *Node {
	min := int(^uint(0) >> 1)
	index := 0
	for i, n := range *open {
		if n.f() < min {
			min = n.f()
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
	path = append(path, node.parent.puzzle)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
