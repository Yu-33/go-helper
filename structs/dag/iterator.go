package dag

import (
	"github.com/Yu-33/gohelper/structs/queue"
)

// IterTopo implements an iterator by Topological Sorting(Kahn Algorithm).
//
type IterTopo struct {
	q         *queue.Queue
	inDegrees map[*Vertex]int
}

// newIterTopo is an interval func helper creates an IterTopo.
func newIterTopo(g *DAG) *IterTopo {
	q := queue.Default()
	inDegrees := make(map[*Vertex]int, q.Len())

	// Init the iterator.
	vexIt := g.vertexes.Iter(nil, nil)
	for vexIt.Valid() {
		element := vexIt.Next()
		vex := element.Value().(*Vertex)

		degree := vex.in.Len()
		if degree == 0 {
			// Add the vertex that zero in-degree to queue.
			q.Push(vex)
		} else {
			inDegrees[vex] = degree
		}
	}

	it := &IterTopo{
		q:         q,
		inDegrees: inDegrees,
	}
	return it
}

// Valid represents whether has more vertex in iterator.
func (it *IterTopo) Valid() bool {
	return !it.q.Empty()
}

// Next returns a vertex that in-degree is zero.
// Returns nil if no more.
func (it *IterTopo) Next() *Vertex {
	if !it.Valid() {
		return nil
	}

	vex := it.q.Pop().(*Vertex)
	it.fillQueue(vex)

	return vex
}

// Batch returns all vertexes that in-degree is zero at once.
// Returns nil if no more.
func (it *IterTopo) Batch() []*Vertex {
	if !it.Valid() {
		return nil
	}

	vexes := make([]*Vertex, 0, it.q.Len())

	for !it.q.Empty() {
		vex := it.q.Pop().(*Vertex)
		vexes = append(vexes, vex)
	}

	for i := range vexes {
		it.fillQueue(vexes[i])
	}

	return vexes
}

func (it *IterTopo) fillQueue(vex *Vertex) {
	itOut := vex.out.Iter(nil, nil)
	for itOut.Valid() {
		element := itOut.Next()
		vex := element.Value().(*Vertex)

		it.inDegrees[vex]--
		if it.inDegrees[vex] == 0 {
			delete(it.inDegrees, vex)
			it.q.Push(vex)
		}
	}
}
