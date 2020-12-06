package bst

import (
	"github.com/Yu-33/gohelper/structs/container"
)

type Element = container.Comparer

// Node declare an interface of binary tree node type.
type Node interface {
	Element() Element
	Left() Node
	Right() Node
}
