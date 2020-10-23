package avl

import (
	"fmt"

	"github.com/Yu-33/helper/datastructs/bst"
	"github.com/Yu-33/helper/datastructs/container"
)

type Tree struct {
	root *Node
	len  int
}

func New() *Tree {
	tr := &Tree{
		root: nil,
		len:  0,
	}
	return tr
}

func (tr *Tree) Len() int {
	return tr.len
}

func (tr *Tree) Search(elements Elements) Elements {
	p := tr.root
	for p != nil {
		flag := elements.Compare(p.elements)
		if flag == -1 {
			p = p.left
		} else if flag == 1 {
			p = p.right
		} else {
			return p.elements
		}
	}
	return nil
}

func (tr *Tree) Insert(elements Elements) bool {
	var n *Node
	tr.root, n = tr.insert(tr.root, elements)
	if n == nil {
		return false
	}
	tr.len++
	return true
}

func (tr *Tree) Delete(elements Elements) Elements {
	var d *Node
	tr.root, d = tr.delete(tr.root, elements)
	if d == nil {
		return nil
	}

	tr.len--
	return d.elements
}

// Iter return a Iterator, include elements: start <= k <= boundary
// start == first node if start == nil and boundary == last node if boundary == nil
func (tr *Tree) Iter(start Elements, boundary Elements) container.Iterator {
	it := bst.NewIterator(tr.root, start, boundary)
	return it
}

// return (new root, new node)
func (tr *Tree) insert(root *Node, elements Elements) (*Node, *Node) {
	var n *Node

	if root == nil {
		n = tr.createNode(elements)
		root = n
	} else {
		flag := elements.Compare(root.elements)
		if flag == -1 {
			// insert into left subtree
			root.left, n = tr.insert(root.left, elements)
		} else if flag == 1 {
			// insert into right subtree
			root.right, n = tr.insert(root.right, elements)
		} else {
			// duplicate elements
			return root, nil
		}
		if n != nil {
			root = tr.reBalance(root)
		}
	}

	return root, n
}

// return (new root, delete node)
func (tr *Tree) delete(root *Node, elements Elements) (*Node, *Node) {
	var d *Node
	if root == nil {
		// not found
		return nil, nil
	} else {
		flag := elements.Compare(root.elements)
		if flag == -1 {
			// delete from left subtree
			root.left, d = tr.delete(root.left, elements)
		} else if flag == 1 {
			// delete from right subtree
			root.right, d = tr.delete(root.right, elements)
		} else {
			if root.left != nil && root.right != nil {
				if tr.nodeHeight(root.left) > tr.nodeHeight(root.right) {
					x := root.left
					for x.right != nil {
						x = x.right
					}
					root.elements, x.elements = x.elements, root.elements
					root.left, d = tr.delete(root.left, elements)
				} else {
					x := root.right
					for x.left != nil {
						x = x.left
					}
					root.elements, x.elements = x.elements, root.elements
					root.right, d = tr.delete(root.right, elements)
				}
			} else {
				d = root
				if d.left != nil {
					root = d.left
				} else {
					root = d.right
				}
			}
		}
		if root != nil {
			root = tr.reBalance(root)
		}
	}

	return root, d
}

func (tr *Tree) reBalance(n *Node) *Node {
	if n == nil {
		return nil
	}

	factor := tr.nodeHeight(n.left) - tr.nodeHeight(n.right)

	switch factor {
	case -1, 0, 1:
		n.height = tr.calculateHeight(n)
	case 2:
		// Left subtree higher than right subtree
		if tr.nodeHeight(n.left.right) > tr.nodeHeight(n.left.left) {
			n.left = tr.leftRotate(n.left)
		}
		n = tr.rightRotate(n)
	case -2:
		// Left subtree lower than right subtree
		if tr.nodeHeight(n.right.left) > tr.nodeHeight(n.right.right) {
			n.right = tr.rightRotate(n.right)
		}
		n = tr.leftRotate(n)
	default:
		panic(fmt.Sprintf("invalid factor <%d>", factor))
	}

	return n
}

func (tr *Tree) createNode(elements Elements) *Node {
	n := new(Node)
	n.elements = elements
	n.height = 1
	n.left = nil
	n.right = nil
	return n
}

func (tr *Tree) nodeHeight(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func (tr *Tree) calculateHeight(n *Node) int {
	lh := tr.nodeHeight(n.left)
	rh := tr.nodeHeight(n.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func (tr *Tree) leftRotate(n *Node) *Node {
	r := n.right

	n.right = r.left
	r.left = n

	n.height = tr.calculateHeight(n)
	r.height = tr.calculateHeight(r)

	return r
}

func (tr *Tree) rightRotate(n *Node) *Node {
	l := n.left

	n.left = l.right
	l.right = n

	n.height = tr.calculateHeight(n)
	l.height = tr.calculateHeight(l)

	return l
}