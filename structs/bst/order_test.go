package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInOrder(t *testing.T) {
	tr := buildBSTree()

	r1 := InOrder(tr.root)

	var f func(node *treeNode)

	var r2 []KV
	f = func(node *treeNode) {
		if node == nil {
			return
		}
		f(node.left)
		r2 = append(r2, node)
		f(node.right)
	}

	f(tr.root)

	require.Equal(t, r1, r2)
}

func TestPreOrder(t *testing.T) {
	tr := buildBSTree()
	r1 := PreOrder(tr.root)

	var f func(node *treeNode)

	var r2 []KV
	f = func(node *treeNode) {
		if node == nil {
			return
		}
		r2 = append(r2, node)
		f(node.left)
		f(node.right)
	}

	f(tr.root)

	require.Equal(t, r1, r2)
}

func TestPostOrder(t *testing.T) {
	tr := buildBSTree()

	r1 := PostOrder(tr.root)

	var f func(node *treeNode)

	var r2 []KV
	f = func(node *treeNode) {
		if node == nil {
			return
		}
		f(node.left)
		f(node.right)
		r2 = append(r2, node)
	}

	f(tr.root)

	require.Equal(t, r1, r2)
}
