package bst

import (
	"github.com/Yu-33/helper/datastructs/container"
)

type Elements = container.Comparer

type Node interface {
	Elements() Elements
	Left() Node
	Right() Node
}
