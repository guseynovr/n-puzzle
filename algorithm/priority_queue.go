package algorithm

type NodeQueue struct {
	nodes []*Node
	byH   bool
}

func (q *NodeQueue) Len() int {
	return len(q.nodes)
}

func (q *NodeQueue) Less(i, j int) bool {
	return (!q.byH &&
		(q.nodes[i].f() < q.nodes[j].f() ||
			(q.nodes[i].f() == q.nodes[j].f() &&
				q.nodes[i].h < q.nodes[j].h))) ||
		(q.byH &&
			(q.nodes[i].h < q.nodes[j].h ||
				(q.nodes[i].h == q.nodes[j].h &&
					q.nodes[i].f() < q.nodes[j].f())))
}

func (q *NodeQueue) Swap(i, j int) {
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
}

func (q *NodeQueue) Push(n any) {
	q.nodes = append(q.nodes, n.(*Node))
}

func (q *NodeQueue) Pop() any {
	if len(q.nodes) == 0 {
		return nil
	}
	last := q.nodes[len(q.nodes)-1]
	q.nodes[len(q.nodes)-1] = nil
	q.nodes = q.nodes[0 : len(q.nodes)-1]
	return last
}

func (q *NodeQueue) index(node *Node) (int, bool) {
	for i, n := range q.nodes {
		if n.hash() == node.hash() {
			return i, true
		}
	}
	return 0, false
}
