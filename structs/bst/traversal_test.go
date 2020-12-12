package bst

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gohelper/structs/container"
)

func buildBSTree() (tr *Tree) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 25
	maxKey := length * 100

	tr = New()

	for i := 0; i < length; i++ {
		for {
			k := container.Int64(r.Intn(maxKey) + 1)
			if tr.Insert(k, k*2+1) != nil {
				break
			}
		}
	}
	return
}

func TestLDR(t *testing.T) {
	tr := buildBSTree()

	var r1 []Element
	LDR(tr.root, func(n Node) {
		r1 = append(r1, n)
	})

	var f func(node *treeNode)
	var r2 []Element
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

func TestDLR(t *testing.T) {
	tr := buildBSTree()

	var r1 []Element
	DLR(tr.root, func(n Node) {
		r1 = append(r1, n)
	})

	var f func(node *treeNode)
	var r2 []Element
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

func TestLRD(t *testing.T) {
	tr := buildBSTree()

	var r1 []Element
	LRD(tr.root, func(n Node) {
		r1 = append(r1, n)
	})

	var f func(node *treeNode)
	var r2 []Element
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
