package dag

import (
	"fmt"

	"github.com/Yu-33/gohelper/structs/container"
	"github.com/Yu-33/gohelper/structs/rb"
	"github.com/Yu-33/gohelper/structs/stack"
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

// InDegreeToString return strings by in-degree.
func (g *DAG) InDegreeToString() string {
	s := "{ "
	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		kv1 := it1.Next()
		n1 := kv1.Value().(*Node)
		s += fmt.Sprintf("%v: [ ", kv1.Key())

		it2 := n1.in.Iter(nil, nil)
		for it2.Valid() {
			kv2 := it2.Next()
			n2 := kv2.Value().(*Node)
			_ = n2
			s += fmt.Sprintf("%v ", kv2.Key())
		}
		s += "] "
	}
	s += "}"
	return s
}

// OutDegreeToString return strings by out-degree.
func (g *DAG) OutDegreeToString() string {
	s := "{ "
	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		kv1 := it1.Next()
		n1 := kv1.Value().(*Node)
		s += fmt.Sprintf("%v: [ ", kv1.Key())

		it2 := n1.out.Iter(nil, nil)
		for it2.Valid() {
			kv2 := it2.Next()
			n2 := kv2.Value().(*Node)
			_ = n2
			s += fmt.Sprintf("%v ", kv2.Key())
		}
		s += "] "
	}
	s += "}"
	return s
}

// AddVertex add new vertex into DAG, return false if vertex already exists.
func (g *DAG) AddVertex(vertex Vertex, v Value) bool {
	n := &Node{value: v}
	ok := g.vertexes.Insert(vertex, n)
	if !ok {
		return false
	}
	n.in = rb.New()
	n.out = rb.New()
	return true
}

// DelVertex delete a vertex from DAG, return false if vertex not exists.
func (g *DAG) DelVertex(vertex Vertex) bool {
	v1 := g.vertexes.Delete(vertex)
	if v1 == nil {
		return false
	}
	v1.(*Node).in = nil
	v1.(*Node).out = nil

	// Delete edge
	it := g.vertexes.Iter(nil, nil)
	for it.Valid() {
		kv2 := it.Next()
		_ = kv2.Value().(*Node).in.Delete(vertex)
		_ = kv2.Value().(*Node).out.Delete(vertex)
	}

	return true
}

// AddEdge add an edge from vertex to adjacency.
// return false when There is a loop between vertex and adjacency.
func (g *DAG) AddEdge(vertex, adjacency Vertex) bool {
	if vertex.Compare(adjacency) == 0 {
		return false
	}

	v1 := g.vertexes.Search(vertex)
	if v1 == nil {
		return false
	}
	v2 := g.vertexes.Search(adjacency)
	if v2 == nil {
		return false
	}

	vex := v1.(*Node)
	adj := v2.(*Node)

	// check loop
	s := stack.Default()
	s.Push([]interface{}{adjacency, adj})

	for !s.Empty() {
		x := s.Pop().([]interface{})
		k := x[0].(Vertex)
		n := x[1].(*Node)

		if k.Compare(vertex) == 0 {
			return false
		}

		it := n.out.Iter(nil, nil)
		for it.Valid() {
			kv3 := it.Next()
			s.Push([]interface{}{kv3.Key(), kv3.Value()})
		}
	}

	// add edge
	_ = vex.out.Insert(adjacency, adj)
	_ = adj.in.Insert(vertex, vex)

	return true
}

// DelEdge delete edges from vertex to adjacency.
func (g *DAG) DelEdge(vertex, adjacency Vertex) bool {
	v1 := g.vertexes.Search(vertex)
	if v1 == nil {
		return false
	}
	v2 := g.vertexes.Search(adjacency)
	if v2 == nil {
		return false
	}

	if v := v1.(*Node).out.Delete(adjacency); v == nil {
		return false
	}
	if v := v2.(*Node).in.Delete(vertex); v == nil {
		return false
	}

	return true
}

func (g *DAG) IterTopo() *IterTopo {
	return newIterTopo(g)
}
