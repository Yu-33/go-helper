package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func inOrder(root Node) []KV {
	var result []KV
	LDR(root, func(node Node) {
		result = append(result, node)
	})
	return result
}

func preOrder(root Node) []KV {
	var result []KV
	DLR(root, func(node Node) {
		result = append(result, node)
	})
	return result
}

func postOrder(root Node) []KV {
	var result []KV
	LRD(root, func(node Node) {
		result = append(result, node)
	})
	return result
}

func TestInOrder(t *testing.T) {
	tr := buildBSTree()

	r1 := inOrder(tr.root)

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
	r1 := preOrder(tr.root)

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

	r1 := postOrder(tr.root)

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
