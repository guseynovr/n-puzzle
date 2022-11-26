package algorithm

type NodeQueue []*Node

func (q NodeQueue) Len() int {
	return len(q)
}

func (q NodeQueue) Less(i, j int) bool {
	return q[i].f() < q[j].f()
}

func (q NodeQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *NodeQueue) Push(n any) {
	*q = append(*q, n.(*Node))
}

func (q *NodeQueue) Pop() any {
	if len(*q) == 0 {
		return nil
	}
	last := (*q)[len(*q)-1]
	*q = (*q)[0 : len(*q)-1]
	return last
}

func (q *NodeQueue) Index(node *Node) (int, bool) {
	for i, n := range *q {
		if n.hash() == node.hash() {
			return i, true
		}
	}
	return 0, false
}
