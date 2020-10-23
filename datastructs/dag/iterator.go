package dag

import (
	"github.com/Yu-33/helper/datastructs/queue"
)

// Iterator Topology
type IteratorTopo struct {
	q         *queue.Queue
	inDegrees map[*node]int
}

func newIteratorTopo(g *DAG) *IteratorTopo {
	q := queue.New(-1)

	inDegrees := make(map[*node]int)

	vexIt := g.vertexes.Iter(nil, nil)
	for vexIt.Valid() {
		n := vexIt.Next().(*node)
		inDegrees[n] = n.in.Len()
	}

	for n, degree := range inDegrees {
		if degree == 0 {
			q.Enqueue(n)
		}
	}

	it := &IteratorTopo{
		q:         q,
		inDegrees: inDegrees,
	}
	return it
}

func (it *IteratorTopo) Valid() bool {
	return !it.q.IsEmpty()
}

func (it *IteratorTopo) Next() []Vertex {
	if !it.Valid() {
		return nil
	}

	vertexes := make([]Vertex, 0, it.q.Len())
	nodes := make([]*node, 0, it.q.Len())

	for !it.q.IsEmpty() {
		n := it.q.Dequeue().(*node)
		vertexes = append(vertexes, n.vex)
		nodes = append(nodes, n)
	}

	for i := range nodes {
		itOut := nodes[i].out.Iter(nil, nil)
		for itOut.Valid() {
			n := itOut.Next().(*node)
			it.inDegrees[n]--
			if it.inDegrees[n] == 0 {
				it.q.Enqueue(n)
			}
		}
	}

	return vertexes
}
