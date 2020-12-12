package dag

import (
	"github.com/yu31/gohelper/structs/container"
)

// Type aliases for simplifying use in this package.
type Key = container.Comparator
type Value = container.Value

// Vertex represents a vertex in DAG.
type Vertex struct {
	key   Key
	value Value
	out   container.Container // out-edge
	in    container.Container // in-edge
}

// Returns the key of vertex.
func (vex *Vertex) Key() Key {
	return vex.key
}

// Returns the value of vertex.
func (vex *Vertex) Value() Value {
	return vex.value
}
