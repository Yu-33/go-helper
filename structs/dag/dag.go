package dag

import (
	"github.com/yu31/gohelper/structs/container"
	"github.com/yu31/gohelper/structs/rb"
	"github.com/yu31/gohelper/structs/stack"
)

// DAG implements data struct of Directed Acyclic Graph.
type DAG struct {
	vertexes container.Container
}

// New creates an DAG.
func New() *DAG {
	g := new(DAG)
	g.vertexes = rb.New()
	return g
}

// Len returns number of vertex in DAG.
func (g *DAG) Len() int {
	return g.vertexes.Len()
}

// AddVertex adds new vertex with k/v to DAG.
// Returns false if vertex already exists.
func (g *DAG) AddVertex(k Key, v Value) bool {
	vex := &Vertex{key: k, value: v}
	if g.vertexes.Insert(k, vex) == nil {
		return false
	}

	vex.in = rb.New()
	vex.out = rb.New()
	return true
}

// DelVertex removes the vertex by giving key and returns its value.
// Returns nil if vertex not exists.
func (g *DAG) DelVertex(k Key) Value {
	element := g.vertexes.Delete(k)
	if element == nil {
		return nil
	}

	vex := element.Value().(*Vertex)
	vex.in = nil
	vex.out = nil

	// Delete edges form other vertices that attach to this vertex.
	it := g.vertexes.Iter(nil, nil)
	for it.Valid() {
		element := it.Next()
		vt := element.Value().(*Vertex)
		_ = vt.in.Delete(k)
		_ = vt.out.Delete(k)
	}

	return vex.value
}

// GetVertex get the value of a given key.
func (g *DAG) GetVertex(k Key) Value {
	element := g.vertexes.Search(k)
	if element == nil {
		return nil
	}
	return element.Value().(*Vertex).value
}

// AddEdge attaches an edge from vertex to adjacency.
//
// Returns false if there a ring between vertex and adjacency after attaching.
//
// And will be crashing in following cases:
//   - vertex equal to adjacency.
//   - vertex or adjacency does not exist.
// FIXME: refactor error process.
func (g *DAG) AddEdge(vex, adj Key) bool {
	if vex.Compare(adj) == 0 {
		panic("dag:AddEdge: vertex can not equal to adjacency")
	}

	e1 := g.vertexes.Search(vex)
	if e1 == nil {
		panic("dag:AddEdge: vertex not exists")
	}
	e2 := g.vertexes.Search(adj)
	if e2 == nil {
		panic("dag:AddEdge: adjacency not exists")
	}

	vex1 := e1.Value().(*Vertex)
	vex2 := e2.Value().(*Vertex)

	// FIXME: refactor it.
	// Check whether has ring.
	s := stack.Default()
	s.Push([]interface{}{adj, vex2})

	for !s.Empty() {
		x := s.Pop().([]interface{})
		k := x[0].(Key)
		n := x[1].(*Vertex)

		if k.Compare(vex) == 0 {
			return false
		}

		it := n.out.Iter(nil, nil)
		for it.Valid() {
			element := it.Next()
			s.Push([]interface{}{element.Key(), element.Value()})
		}
	}

	// Attach edges
	// FIXME: We needs check edges exists first ?
	_ = vex1.out.Insert(adj, vex2)
	_ = vex2.in.Insert(vex, vex1)

	return true
}

// DelEdge detaches edges from vertex to adjacency.
//
// And will be crashing in following cases:
//   - vertex equal to adjacency.
//   - vertex or adjacency does not exist.
// FIXME: refactor error process.
func (g *DAG) DelEdge(vex, adj Key) bool {
	if vex.Compare(adj) == 0 {
		panic("dag:DelEdge: vertex can not equal to adjacency")
	}

	e1 := g.vertexes.Search(vex)
	if e1 == nil {
		panic("dag:DelEdge: vertex not exists")
	}
	e2 := g.vertexes.Search(adj)
	if e2 == nil {
		panic("dag:DelEdge: adjacency not exists")
	}

	if element := e1.Value().(*Vertex).out.Delete(adj); element == nil {
		return false
	}
	if element := e2.Value().(*Vertex).in.Delete(vex); element == nil {
		return false
	}

	return true
}

func (g *DAG) IterTopo() *IterTopo {
	return newIterTopo(g)
}
