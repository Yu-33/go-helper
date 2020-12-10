package bst

import (
	"github.com/Yu-33/gohelper/structs/container"
)

// treeNode used in binary search tree.
// And also implements interface bst.Node and container.KV.
type treeNode struct {
	key   Key
	value Value
	left  *treeNode
	right *treeNode
}

// Implements interface container.KV.
// Implements interface bst.Node.
func (n *treeNode) Key() Key {
	return n.key
}

// Implements interface container.KV.
// Implements interface bst.Node.
func (n *treeNode) Value() Value {
	return n.value
}

// Implements interface bst.Node.
func (n *treeNode) Left() Node {
	return n.left
}

// Implements interface bst.Node.
func (n *treeNode) Right() Node {
	return n.right
}

// Tree implements the data struct binary search Tree.
//
// And also implements interface container.Container
type Tree struct {
	root *treeNode
	len  int
}

// New creates binary search Tree.
func New() *Tree {
	tr := &Tree{
		root: nil,
		len:  0,
	}
	return tr
}

// Len return number of elements.
func (tr *Tree) Len() int {
	return tr.len
}

// Insert inserts the key with value in the container.
// k and v must not be nil, otherwise it will crash.
// Returns false if key already exists.
func (tr *Tree) Insert(k Key, v Value) bool {
	p := tr.root
	for p != nil {
		flag := k.Compare(p.key)
		if flag == -1 {
			if p.left == nil {
				p.left = tr.createNode(k, v)
				break
			}
			p = p.left
		} else if flag == 1 {
			if p.right == nil {
				p.right = tr.createNode(k, v)
				break
			}
			p = p.right
		} else {
			// The key already exists. Not allowed duplicates.
			return false
		}
	}

	if p == nil {
		tr.root = tr.createNode(k, v)
	}

	tr.len++
	return true
}

// Delete remove and returns the value of the specified key.
// Returns nil if not found.
func (tr *Tree) Delete(k Key) Value {
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

	tr.len--
	return d.value
}

// Search get the value of specified key.
// Returns nil if not found.
func (tr *Tree) Search(k Key) Value {
	p := tr.root
	for p != nil {
		flag := k.Compare(p.key)
		if flag == -1 {
			p = p.left
		} else if flag == 1 {
			p = p.right
		} else {
			return p.value
		}
	}
	return nil
}

// Iter return a Iterator, it's a wraps for bst.Iterator.
func (tr *Tree) Iter(start Key, boundary Key) container.Iterator {
	it := NewIterator(tr.root, start, boundary)
	return it
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
