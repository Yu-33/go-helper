package dag

import (
	"github.com/Yu-33/gohelper/structs/queue"
)

// IterTopo implements Iterator by Topology algo.
type IterTopo struct {
	q         *queue.Queue
	inDegrees map[Vertex]int
}

func newIterTopo(g *DAG) *IterTopo {
	q := queue.Default()

	inDegrees := make(map[Vertex]int)

	vexIt := g.vertexes.Iter(nil, nil)
	for vexIt.Valid() {
		kv := vexIt.Next()
		degree := kv.Value().(*Node).in.Len()

		inDegrees[kv.Key()] = degree

		// Add the vertex with zero degree into the queue.
		if degree == 0 {
			q.Push(kv)
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

// FIXME:
func (it *IterTopo) Next() []KV {
	if !it.Valid() {
		return nil
	}
	vertexes := make([]KV, 0, it.q.Len())

	for !it.q.Empty() {
		kv := it.q.Pop().(KV)
		vertexes = append(vertexes, kv)
	}

	for i := range vertexes {
		itOut := vertexes[i].Value().(*Node).out.Iter(nil, nil)
		for itOut.Valid() {
			kv := itOut.Next()
			it.inDegrees[kv.Key()]--
			if it.inDegrees[kv.Key()] == 0 {
				it.q.Push(kv)
			}
		}
	}

	return vertexes
}
