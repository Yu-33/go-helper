package avl

import (
	"fmt"

	"github.com/Yu-33/gohelper/structs/bst"
	"github.com/Yu-33/gohelper/structs/container"
)

type Key = container.Key
type Value = container.Value

// treeNode used in avl tree.
// And also implements interface bst.Node and container.KV.
type treeNode struct {
	key    Key
	value  Value
	left   *treeNode
	right  *treeNode
	height int
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
func (n *treeNode) Left() bst.Node {
	return n.left
}

// Implements interface bst.Node.
func (n *treeNode) Right() bst.Node {
	return n.right
}

// Tree implements the data struct of AVL tree.
//
// And also implements interface container.Container
type Tree struct {
	root *treeNode
	len  int
}

// New creates an AVL Tree.
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

// Insert inserts the key with value in the container.
// k and v must not be nil, otherwise it will crash.
// Returns false if key already exists.
func (tr *Tree) Insert(k Key, v Value) bool {
	var n *treeNode
	tr.root, n = tr.insert(tr.root, k, v)
	if n == nil {
		return false
	}
	tr.len++
	return true
}

// Delete remove and returns the value of the specified key.
// Returns nil if not found.
func (tr *Tree) Delete(k Key) Value {
	var d *treeNode
	tr.root, d = tr.delete(tr.root, k)
	if d == nil {
		return nil
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
	it := bst.NewIterator(tr.root, start, boundary)
	return it
}

func (tr *Tree) swap(n1, n2 *treeNode) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}

// return (new root, new node)
func (tr *Tree) insert(root *treeNode, k Key, v Value) (*treeNode, *treeNode) {
	var n *treeNode

	if root == nil {
		n = tr.createNode(k, v)
		root = n
	} else {
		flag := k.Compare(root.key)
		if flag == -1 {
			// Insert into the left subtree.
			root.left, n = tr.insert(root.left, k, v)
		} else if flag == 1 {
			// Insert into the right subtree
			root.right, n = tr.insert(root.right, k, v)
		} else {
			// The key already exists. Not allow duplicates.
			return root, nil
		}
		if n != nil {
			root = tr.reBalance(root)
		}
	}

	return root, n
}

// return (new root, delete node).
func (tr *Tree) delete(root *treeNode, k Key) (*treeNode, *treeNode) {
	var d *treeNode
	if root == nil {
		// not found
		return nil, nil
	} else {
		flag := k.Compare(root.key)
		if flag == -1 {
			// delete from the left subtree
			root.left, d = tr.delete(root.left, k)
		} else if flag == 1 {
			// delete from the right subtree
			root.right, d = tr.delete(root.right, k)
		} else {
			if root.left != nil && root.right != nil {
				if tr.nodeHeight(root.left) > tr.nodeHeight(root.right) {
					x := root.left
					for x.right != nil {
						x = x.right
					}
					tr.swap(root, x)
					root.left, d = tr.delete(root.left, k)
				} else {
					x := root.right
					for x.left != nil {
						x = x.left
					}
					tr.swap(root, x)
					root.right, d = tr.delete(root.right, k)
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

func (tr *Tree) reBalance(n *treeNode) *treeNode {
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
		panic(fmt.Errorf("avl: unexpected cases with invalid factor <%d>", factor))
	}

	return n
}

func (tr *Tree) createNode(k Key, v Value) *treeNode {
	return &treeNode{
		key:    k,
		value:  v,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func (tr *Tree) nodeHeight(n *treeNode) int {
	if n == nil {
		return 0
	}
	return n.height
}

func (tr *Tree) calculateHeight(n *treeNode) int {
	lh := tr.nodeHeight(n.left)
	rh := tr.nodeHeight(n.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func (tr *Tree) leftRotate(n *treeNode) *treeNode {
	r := n.right

	n.right = r.left
	r.left = n

	n.height = tr.calculateHeight(n)
	r.height = tr.calculateHeight(r)

	return r
}

func (tr *Tree) rightRotate(n *treeNode) *treeNode {
	l := n.left

	n.left = l.right
	l.right = n

	n.height = tr.calculateHeight(n)
	l.height = tr.calculateHeight(l)

	return l
}
