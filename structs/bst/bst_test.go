package bst

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
)

func checkCorrect(t *testing.T, n *treeNode) {
	if n == nil {
		return
	}
	checkCorrect(t, n.left)
	checkCorrect(t, n.right)

	if n.left != nil {
		require.Equal(t, n.key.Compare(n.left.key), 1)
	}
	if n.right != nil {
		require.Equal(t, n.key.Compare(n.right.key), -1)
	}
}

func buildBSTree() (tr *Tree) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 25
	maxKey := length * 10

	tr = New()

	for i := 0; i < length; i++ {
		for {
			k := container.Int64(r.Intn(maxKey) + 1)
			if ok := tr.Insert(k, k*2+1); ok {
				break
			}
		}
	}

	return
}

func Test_Interface(t *testing.T) {
	// Ensure the interface is implemented.
	var node Node
	var kv container.KV
	var ct container.Container
	var it container.Iterator

	node = &treeNode{}
	_ = node
	kv = &treeNode{}
	_ = kv
	ct = New()
	_ = ct
	it = NewIterator(node, nil, nil)
	_ = it
}

func TestNew(t *testing.T) {
	tr := New()
	require.NotNil(t, tr)
	require.Nil(t, tr.root)
	require.Equal(t, tr.Len(), 0)
}

func TestTree(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tr := New()

	length := 65
	maxKey := length * 10
	keys := make([]container.Int, length)

	for x := 0; x < 2; x++ {
		// insert
		for i := 0; i < length; i++ {
			for {
				k := container.Int(r.Intn(maxKey) + 1)
				if ok := tr.Insert(k, int64(k*2+1)); ok {
					require.False(t, tr.Insert(k, int64(k*2+1)))
					keys[i] = k
					break
				}
			}
			checkCorrect(t, tr.root)
		}

		require.Equal(t, length, tr.len)

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
			checkCorrect(t, tr.root)
		}

		require.Nil(t, tr.root)
		require.Equal(t, 0, tr.Len())
	}
}

func TestTree_Len(t *testing.T) {
	tr := New()

	require.True(t, tr.Insert(container.Int(12), 1))
	require.True(t, tr.Insert(container.Int(18), 1))
	require.True(t, tr.Insert(container.Int(33), 1))

	// Insert duplicate key.
	require.False(t, tr.Insert(container.Int(12), 1))
	require.False(t, tr.Insert(container.Int(18), 1))
	require.False(t, tr.Insert(container.Int(33), 1))

	require.Equal(t, tr.Len(), 3)

	require.NotNil(t, tr.Delete(container.Int(18)))
	// Delete key not exists.
	require.Nil(t, tr.Delete(container.Int(18)))

	require.Equal(t, tr.Len(), 2)
}

func TestAVLTree_createNode(t *testing.T) {
	tr := New()

	el1 := container.Int64(0xf)

	n1 := tr.createNode(el1, 1024)
	require.NotNil(t, n1)
	require.Equal(t, n1.key.Compare(el1), 0)
	require.Nil(t, n1.left)
	require.Nil(t, n1.right)
}
