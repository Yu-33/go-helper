package maxheap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
)

// checkCorrect check every node is greater than of equal to the left child and right child.
func checkCorrect(t *testing.T, h *MaxHeap) {
	// Check the index.
	for i := 0; i < h.len; i++ {
		require.Equal(t, h.items[i].index, i)
	}
	for i := 0; i < (h.len-1)>>1; i++ {
		require.NotEqual(t, h.items[i].key.Compare(h.items[(i<<1)+1].key), -1)
		require.NotEqual(t, h.items[i].key.Compare(h.items[(i<<1)+2].key), -1)
	}
}

func TestNew(t *testing.T) {
	_ = Default()
	max := 17
	h := New(max)

	require.NotNil(t, h)
	require.True(t, h.Empty())
	require.Equal(t, h.len, 0)
	require.Equal(t, h.cap, max)
}

func TestMaxHeap(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	max := 1025
	maxKey := max * 100

	h := New(max)

	for x := 0; x < 2; x++ {
		// enqueue
		for i := 0; i < max; i++ {
			k := container.Int64(r.Intn(maxKey) + 1)

			item := h.Push(k, int(k*2+1))
			require.Equal(t, item.key.(container.Int64), k)
		}

		checkCorrect(t, h)
		require.False(t, h.Empty())
		require.Equal(t, h.Len(), max)

		// dequeue and make queue empty
		p1 := h.Peek()
		last := h.Pop()
		require.NotNil(t, last)
		require.Equal(t, last, p1)
		for i := 1; i < max; i++ {
			p1 := h.Peek()
			item := h.Pop()
			require.NotNil(t, item)
			require.Equal(t, item, p1)

			require.NotEqual(t, item.key.Compare(last.key), 1)
			require.Equal(t, item.value, int(item.key.(container.Int64))*2+1)

			last = item
		}

		checkCorrect(t, h)

		require.True(t, h.Empty())
		require.Equal(t, h.Len(), 0)
		require.Nil(t, h.Pop())
	}

}
