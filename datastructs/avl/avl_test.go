package avl

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/datastructs/container"
)

func recurseCalculateNodeHeight(n *Node) int {
	if n == nil {
		return 0
	}
	lh := recurseCalculateNodeHeight(n.left)
	rh := recurseCalculateNodeHeight(n.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func checkBalance(t *testing.T, tr *Tree, n *Node) {
	if n == nil {
		return
	}

	checkBalance(t, tr, n.left)
	checkBalance(t, tr, n.right)

	// check node height calculate
	require.Equal(t, tr.nodeHeight(n), recurseCalculateNodeHeight(n))

	if n.left != nil {
		require.Equal(t, n.element.Compare(n.left.element), 1)
	}
	if n.right != nil {
		require.Equal(t, n.element.Compare(n.right.element), -1)
	}

	lh := tr.nodeHeight(n.left)
	rh := tr.nodeHeight(n.right)
	if lh > rh {
		require.Equal(t, lh-rh, 1)
	} else if lh < rh {
		require.Equal(t, rh-lh, 1)
	}
}

func TestNew(t *testing.T) {
	tr := New()
	require.NotNil(t, tr)
	require.Nil(t, tr.root)
	require.Equal(t, tr.Len(), 0)
}

func TestAVLTree_createNode(t *testing.T) {
	tr := New()

	el1 := container.Int64(0xf)

	n1 := tr.createNode(el1)
	require.NotNil(t, n1)
	require.Equal(t, n1.element.Compare(el1), 0)
	require.Equal(t, n1.height, 1)
	require.Nil(t, n1.left)
	require.Nil(t, n1.right)
}

func TestAVLTree(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	tr := New()

	length := 259
	maxKey := length * 10
	keys := make([]container.Int64, length)

	for x := 0; x < 2; x++ {
		// insert
		for i := 0; i < length; i++ {
			for {
				k := container.Int64(r.Intn(maxKey) + 1)
				if ok := tr.Insert(k); ok {
					keys[i] = k
					break
				}
			}
			checkBalance(t, tr, tr.root)
		}

		require.Equal(t, tr.Len(), length)

		// boundary
		for _, k := range []container.Int64{0, 0xfffffff} {
			require.True(t, tr.Insert(k))
			require.False(t, tr.Insert(k))
			require.NotNil(t, tr.Search(k))
			require.Equal(t, tr.Search(k), k)
			require.NotNil(t, tr.Delete(k))
			require.Nil(t, tr.Delete(k))
		}

		// get
		for i := 0; i < length; i++ {
			n := tr.Search(keys[i])
			require.NotNil(t, n, fmt.Sprintf("key %d not found", keys[i]))
			require.Equal(t, n, keys[i])
		}

		// delete
		for i := 0; i < length; i++ {
			require.NotNil(t, tr.Delete(keys[i]))
			require.Nil(t, tr.Delete(keys[i]))
			require.Nil(t, tr.Search(keys[i]))
			checkBalance(t, tr, tr.root)
		}
		require.Nil(t, tr.root)
		require.Equal(t, tr.Len(), 0)
	}
}
