package bst

import (
	"github.com/yu31/gohelper/structs/container"
)

// Type aliases for simplifying use in this package.
type Key = container.Key
type Value = container.Value
type Element = container.Element

// Node defines an interface of binary tree node type.
type Node interface {
	Element
	// Left returns the left child of the Node.
	Left() Node
	// Right returns the right child of the Node.
	Right() Node
}
