package dag

import (
	"fmt"

	"github.com/Yu-33/gohelper/structs/container"
	"github.com/Yu-33/gohelper/structs/rb"
	"github.com/Yu-33/gohelper/structs/stack"
)

// DAG type implements a Directed Acyclic Graph data structure.
type DAG struct {
	vertexes container.Container
}

func New() *DAG {
	g := new(DAG)
	g.vertexes = rb.New()
	return g
}

func (g *DAG) InDegreeToString() string {
	s := "{ "
	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		n1 := it1.Next().(*node)
		s += fmt.Sprintf("%v: [ ", n1.vex)

		it2 := n1.in.Iter(nil, nil)
		for it2.Valid() {
			n2 := it2.Next().(*node)
			s += fmt.Sprintf("%v ", n2.vex)
		}
		s += "] "
	}
	s += "}"
	return s
}

func (g *DAG) OutDegreeToString() string {
	s := "{ "
	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		n1 := it1.Next().(*node)
		s += fmt.Sprintf("%v: [ ", n1.vex)

		it2 := n1.out.Iter(nil, nil)
		for it2.Valid() {
			n2 := it2.Next().(*node)
			s += fmt.Sprintf("%v ", n2.vex)
		}
		s += "] "
	}
	s += "}"
	return s
}

// AddVertex add new vertex into DAG, return false if vertex already exists
func (g *DAG) AddVertex(vertex Vertex) bool {
	n := &node{vex: vertex}
	ok := g.vertexes.Insert(n)
	if !ok {
		return false
	}
	n.in = rb.New()
	n.out = rb.New()
	return true
}

// DelVertex delete a vertex from DAG, return false if vertex not exists
func (g *DAG) DelVertex(vertex Vertex) bool {
	n1 := g.vertexes.Delete(&node{vex: vertex})
	if n1 == nil {
		return false
	}
	n1.(*node).in = nil
	n1.(*node).out = nil

	// Delete edge
	it := g.vertexes.Iter(nil, nil)
	for it.Valid() {
		n2 := it.Next().(*node)
		_ = n2.in.Delete(n1)
		_ = n2.out.Delete(n1)
	}

	return true
}

// AddEdge add an edge from vertex to adjacency;
// return false when There is a loop between vertex and adjacency;
func (g *DAG) AddEdge(vertex, adjacency Vertex) bool {
	if vertex.Compare(adjacency) == 0 {
		return false
	}

	n1 := g.vertexes.Search(&node{vex: vertex})
	n2 := g.vertexes.Search(&node{vex: adjacency})
	if n1 == nil {
		return false
	}
	if n2 == nil {
		return false
	}

	vex := n1.(*node)
	adj := n2.(*node)

	// check loop
	s := stack.Default()
	s.Push(adj)

	for !s.Empty() {
		n := s.Pop().(*node)
		if n.Compare(vex) == 0 {
			return false
		}

		it := n.out.Iter(nil, nil)
		for it.Valid() {
			s.Push(it.Next().(*node))
		}
	}

	// add edge
	_ = vex.out.Insert(adj)
	_ = adj.in.Insert(vex)

	return true
}

func (g *DAG) DelEdge(vertex, adjacency Vertex) bool {
	n1 := g.vertexes.Search(&node{vex: vertex})
	if n1 == nil {
		return false
	}
	n2 := g.vertexes.Search(&node{vex: adjacency})
	if n2 == nil {
		return false
	}

	vex := n1.(*node)
	adj := n2.(*node)

	if v := vex.out.Delete(adj); v == nil {
		return false
	}
	if v := adj.in.Delete(vex); v == nil {
		return false
	}

	return true
}

func (g *DAG) IterTopo() *IterTopo {
	return newIterTopo(g)
}
