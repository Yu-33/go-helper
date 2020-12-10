package bst

import (
	"github.com/Yu-33/gohelper/structs/container"
)

type Key = container.Key
type Value = container.Value
type KV = container.KV

// Node defines an interface of binary tree node type.
type Node interface {
	KV
	Left() Node
	Right() Node
}
