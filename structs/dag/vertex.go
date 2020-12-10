package dag

import (
	"github.com/Yu-33/gohelper/structs/container"
)

type Vertex container.Comparator
type Value container.Value
type KV = container.KV

type Node struct {
	value Value
	in    container.Container
	out   container.Container
}
