package bst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInOrder(t *testing.T) {
	tr := buildBSTree()

	r1 := InOrder(tr.root)
	fmt.Println(r1)

	var f func(node *MockNode)

	r2 := make([]Element, 0)
	f = func(node *MockNode) {
		if node == nil {
			return
		}
		f(node.left)
		r2 = append(r2, node.element)
		f(node.right)
	}

	f(tr.root)

	fmt.Println(r2)

	require.Equal(t, r1, r2)
}

func TestPreOrder(t *testing.T) {
	tr := buildBSTree()
	r1 := PreOrder(tr.root)
	fmt.Println(r1)

	var f func(node *MockNode)

	r2 := make([]Element, 0)
	f = func(node *MockNode) {
		if node == nil {
			return
		}
		r2 = append(r2, node.element)
		f(node.left)
		f(node.right)
	}

	f(tr.root)

	fmt.Println(r2)

	require.Equal(t, r1, r2)
}

func TestPostOrder(t *testing.T) {
	tr := buildBSTree()

	r1 := PostOrder(tr.root)
	fmt.Println(r1)

	var f func(node *MockNode)

	r2 := make([]Element, 0)
	f = func(node *MockNode) {
		if node == nil {
			return
		}
		f(node.left)
		f(node.right)
		r2 = append(r2, node.element)
	}

	f(tr.root)

	fmt.Println(r2)

	require.Equal(t, r1, r2)
}
