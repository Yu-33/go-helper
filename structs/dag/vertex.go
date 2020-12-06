package dag

import (
	"github.com/Yu-33/gohelper/structs/container"
)

type Vertex = container.Comparer

type node struct {
	vex Vertex
	in  container.Container
	out container.Container
}

func (k1 *node) Compare(target Vertex) int {
	k2 := target.(*node).vex
	return k1.vex.Compare(k2)
}
