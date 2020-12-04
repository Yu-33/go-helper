package minheap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/datastructs/container"
)

func output(pq *MinHeap) {
	fmt.Printf("[")
	for i := 0; i < pq.Len(); i++ {
		fmt.Printf("%d, ", pq.items[i])
	}
	fmt.Printf("]\n")
}

// checkCorrect check every node is less than of equal to the left child and right child
func checkCorrect(t *testing.T, pq *MinHeap) {
	for i := 0; i < (pq.len-1)>>1; i++ {
		require.NotEqual(t, pq.items[i].Compare(pq.items[(i<<1)+1]), 1)
		require.NotEqual(t, pq.items[i].Compare(pq.items[(i<<1)+2]), 1)
	}
}

func TestNew(t *testing.T) {
	max := 17
	h := New(max)

	require.NotNil(t, h)
	require.True(t, h.Empty())
	require.Equal(t, h.len, 0)
	require.Equal(t, h.cap, max)
}

func TestMinHeap(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	max := 33
	maxKey := max * 10
	keys := make([]container.Int64, max)

	h := New(max)

	for x := 0; x < 2; x++ {
		// enqueue and make queue full
		for i := 0; i < max; i++ {
			k := container.Int64(r.Intn(maxKey) + 1)
			h.Push(k)
			keys[i] = k
			checkCorrect(t, h)
		}
		require.False(t, h.Empty())
		require.Equal(t, h.len, max)

		// output
		output(h)

		// dequeue and make queue emtpye
		last := h.Pop()
		require.NotNil(t, last)
		for i := 1; i < max; i++ {
			item := h.Pop()
			require.NotNil(t, item)
			require.NotEqual(t, item.Compare(last), -1)

			checkCorrect(t, h)

			last = item
		}
		require.True(t, h.Empty())
		require.Equal(t, h.len, 0)
		require.Nil(t, h.Pop())
	}

}
