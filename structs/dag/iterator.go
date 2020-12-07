package dag

import (
	"github.com/Yu-33/gohelper/structs/queue"
)

// IterTopo implements Iterator by Topology algo.
type IterTopo struct {
	q         *queue.Queue
	inDegrees map[*node]int
}

func newIterTopo(g *DAG) *IterTopo {
	q := queue.Default()

	inDegrees := make(map[*node]int)

	vexIt := g.vertexes.Iter(nil, nil)
	for vexIt.Valid() {
		n := vexIt.Next().(*node)
		inDegrees[n] = n.in.Len()
	}

	for n, degree := range inDegrees {
		if degree == 0 {
			q.Push(n)
		}
	}

	it := &IterTopo{
		q:         q,
		inDegrees: inDegrees,
	}
	return it
}

func (it *IterTopo) Valid() bool {
	return !it.q.Empty()
}

func (it *IterTopo) Next() []Vertex {
	if !it.Valid() {
		return nil
	}

	vertexes := make([]Vertex, 0, it.q.Len())
	nodes := make([]*node, 0, it.q.Len())

	for !it.q.Empty() {
		n := it.q.Pop().(*node)
		vertexes = append(vertexes, n.vex)
		nodes = append(nodes, n)
	}

	for i := range nodes {
		itOut := nodes[i].out.Iter(nil, nil)
		for itOut.Valid() {
			n := itOut.Next().(*node)
			it.inDegrees[n]--
			if it.inDegrees[n] == 0 {
				it.q.Push(n)
			}
		}
	}

	return vertexes
}
