package rb

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
)

func calculateNodeHeight(n *treeNode) int {
	if n == nil {
		return 0
	}
	lh := calculateNodeHeight(n.left)
	rh := calculateNodeHeight(n.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func checkBalance(t *testing.T, n *treeNode) {
	if n == nil {
		return
	}

	checkBalance(t, n.left)
	checkBalance(t, n.right)

	if n.color == red {
		if n.parent != nil {
			require.Equal(t, n.parent.color, black)
		}
		if n.left != nil && n.right != nil {
			require.Equal(t, n.left.color, black)
			require.Equal(t, n.right.color, black)
		} else {
			require.Nil(t, n.left)
			require.Nil(t, n.right)
		}
	} else {
		require.Equal(t, n.color, black)
	}

	if n.left != nil {
		require.Equal(t, n.key.Compare(n.left.key), 1)
	}
	if n.right != nil {
		require.Equal(t, n.key.Compare(n.right.key), -1)
	}

	lh := calculateNodeHeight(n.left)
	rh := calculateNodeHeight(n.right)
	if lh > rh {
		require.LessOrEqual(t, lh-rh, lh)
	} else {
		require.LessOrEqual(t, rh-lh, rh)
	}
}

func TestNew(t *testing.T) {
	tr := New()
	require.NotNil(t, tr)
	require.Nil(t, tr.root)
	require.Equal(t, tr.Len(), 0)
}

func TestRBTree_createNode(t *testing.T) {
	tr := New()

	ele1 := container.Int64(1)
	n1 := tr.createNode(ele1, 1024, nil)
	require.NotNil(t, n1)
	require.Equal(t, n1.key, ele1)
	require.Equal(t, n1.value, 1024)
	require.Equal(t, n1.color, red)
	require.Nil(t, n1.left)
	require.Nil(t, n1.right)
	require.Nil(t, n1.parent)

	ele2 := container.Int64(2)
	n2 := tr.createNode(ele2, 1024, n1)
	require.NotNil(t, n2)
	require.Equal(t, n2.key, ele2)
	require.Equal(t, n2.value, 1024)
	require.Equal(t, n2.color, red)
	require.Nil(t, n2.left)
	require.Nil(t, n2.right)
	require.Equal(t, n2.parent, n1)
}

func TestRBTree(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tr := New()

	length := 259
	maxKey := length * 10
	keys := make([]container.Int64, length)

	// inserting
	for i := 0; i < length; i++ {
		for {
			k := container.Int64(r.Intn(maxKey) + 1)
			if ok := tr.Insert(k, int64(k*2+1)); ok {
				require.False(t, tr.Insert(k, int64(k*2+1)))
				keys[i] = k
				break
			}
		}

		require.Equal(t, tr.root.color, black)
		checkBalance(t, tr.root)
	}

	require.Equal(t, tr.Len(), length)

	// boundary
	for _, k := range []container.Int64{0, 0xfffffff} {
		require.True(t, tr.Insert(k, k))
		require.False(t, tr.Insert(k, k))
		require.NotNil(t, tr.Search(k))
		require.Equal(t, tr.Search(k), k)
		require.NotNil(t, tr.Delete(k))
		require.Nil(t, tr.Delete(k))
	}

	// search
	for i := 0; i < length; i++ {
		v := tr.Search(keys[i])
		require.NotNil(t, v)
		require.Equal(t, v, int64(keys[i]*2+1))
	}

	// delete
	for i := 0; i < length; i++ {
		require.NotNil(t, tr.Delete(keys[i]))
		require.Nil(t, tr.Delete(keys[i]))

		if tr.root != nil {
			require.Equal(t, tr.root.color, black)
		}
		checkBalance(t, tr.root)
	}

	require.Nil(t, tr.root)
	require.Equal(t, tr.Len(), 0)
}
