package avl

import (
	"github.com/Yu-33/helper/datastructs/bst"
	"github.com/Yu-33/helper/datastructs/container"
)

type Elements = container.Comparer

type Node struct {
	elements Elements
	left     *Node
	right    *Node
	height   int
}

func (n *Node) Elements() Elements {
	return n.elements
}

func (n *Node) Left() bst.Node {
	return n.left
}

func (n *Node) Right() bst.Node {
	return n.right
}
