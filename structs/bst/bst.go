package bst

import (
	"github.com/Yu-33/gohelper/structs/container"
)

type Element = container.Comparer

type Node interface {
	Element() Element
	Left() Node
	Right() Node
}
