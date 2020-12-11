package bst

import (
	"github.com/Yu-33/gohelper/structs/container"
)

// treeNode is used for Binary Search Tree.
//
// And it is also the implementation of interface container.KV and bst.Node
type treeNode struct {
	key   Key
	value Value
	left  *treeNode
	right *treeNode
}

// Key returns the key.
func (n *treeNode) Key() Key {
	return n.key
}

// Value returns the value.
func (n *treeNode) Value() Value {
	return n.value
}

// Left returns the left child of the Node.
func (n *treeNode) Left() Node {
	return n.left
}

// Right returns the right child of the Node.
func (n *treeNode) Right() Node {
	return n.right
}

// Tree implements the Binary Search Tree.
//
// And it is also the implementation of interface container.Container
type Tree struct {
	root *treeNode
	len  int
}

// New creates an Tree.
func New() *Tree {
	tr := &Tree{
		root: nil,
		len:  0,
	}
	return tr
}

// Len returns the number of elements.
func (tr *Tree) Len() int {
	return tr.len
}

// Insert inserts the giving key and value and returns the KV structure.
// Returns nil if key already exists.
func (tr *Tree) Insert(k Key, v Value) KV {
	var n *treeNode
	p := tr.root
	for p != nil {
		flag := k.Compare(p.key)
		if flag == -1 {
			if p.left == nil {
				n = tr.createNode(k, v)
				p.left = n
				break
			}
			p = p.left
		} else if flag == 1 {
			if p.right == nil {
				n = tr.createNode(k, v)
				p.right = n
				break
			}
			p = p.right
		} else {
			// The key already exists. Not allowed duplicates.
			return nil
		}
	}

	if p == nil {
		n = tr.createNode(k, v)
		tr.root = n
	}

	tr.len++
	return n
}

// Delete removes and returns the KV structure corresponding to the given key.
// Returns nil if not found.
func (tr *Tree) Delete(k Key) KV {
	var dd *treeNode
	d := tr.root

	for d != nil {
		flag := k.Compare(d.key)
		// Found the deletion key
		if flag == 0 {
			break
		}

		dd = d
		if flag == -1 {
			d = d.left
		} else {
			d = d.right
		}
	}

	// Not found.
	if d == nil {
		return nil
	}

	if d.left != nil && d.right != nil {
		xx := d
		x := d.left
		for x.right != nil {
			xx = x
			x = x.right
		}

		tr.swap(d, x)
		dd = xx
		d = x
	}

	var c *treeNode
	if d.left != nil {
		c = d.left
	} else {
		c = d.right
	}

	if dd == nil {
		tr.root = c
	} else if dd.left == d {
		dd.left = c
	} else {
		dd.right = c
	}

	d.left = nil
	d.right = nil

	tr.len--
	return d
}

// Search returns the KV structure corresponding to the given key.
// Returns nil if not found.
func (tr *Tree) Search(k Key) KV {
	p := tr.root
	for p != nil {
		flag := k.Compare(p.key)
		if flag == -1 {
			p = p.left
		} else if flag == 1 {
			p = p.right
		} else {
			return p
		}
	}
	return nil
}

// Iter return an Iterator, it's a wrap for bst.Iterator.
func (tr *Tree) Iter(start Key, boundary Key) container.Iterator {
	return NewIterator(tr.root, start, boundary)
}

func (tr *Tree) createNode(k Key, v Value) *treeNode {
	return &treeNode{
		key:   k,
		value: v,
		left:  nil,
		right: nil,
	}
}

func (tr *Tree) swap(n1, n2 *treeNode) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}
