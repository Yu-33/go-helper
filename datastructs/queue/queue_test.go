package queue

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestQueueNew(t *testing.T) {
	capacity := 17
	q := New(capacity)

	require.NotNil(t, q)
	require.True(t, q.IsEmpty())
	require.Equal(t, q.Len(), 0)
	require.Equal(t, q.Cap(), capacity)

	require.NotNil(t, q.items)
	require.Equal(t, len(q.items), capacity+1)
	require.Equal(t, q.cap, capacity+1)
	require.Equal(t, q.front, 0)
	require.Equal(t, q.behind, 0)

}

func TestQueue1(t *testing.T) {
	capacity := 17
	q := New(capacity)

	// test enqueue and make length of queue equal of capacity
	for i := 0; i < capacity; i++ {
		q.Enqueue(i)
	}

	// queue already full
	require.Equal(t, q.front, 0)
	require.Equal(t, q.behind, capacity)

	require.Equal(t, q.Len(), capacity)
	require.Equal(t, q.Cap(), capacity)

	// test dequeue and make queue empty
	for i := 0; i < capacity; i++ {
		item := q.Dequeue()
		require.NotNil(t, item)
		require.Equal(t, item, i)
	}

	// queue is empty
	require.Equal(t, q.front, capacity)
	require.Equal(t, q.behind, capacity)

	require.True(t, q.IsEmpty())
	require.Equal(t, q.Len(), 0)
	require.Equal(t, q.Cap(), capacity)
	require.Nil(t, q.Dequeue())
}

func TestQueue2(t *testing.T) {
	capacity := 2
	q := New(capacity)

	p1 := unsafe.Pointer(&q.items[0])

	q.Enqueue(1)
	q.Enqueue(2)

	p2 := unsafe.Pointer(&q.items[0])
	require.Equal(t, p1, p2)

	q.Enqueue(3)

	p3 := unsafe.Pointer(&q.items[0])
	require.NotEqual(t, p2, p3)

	require.Equal(t, q.Cap(), capacity*2)

	require.Equal(t, q.Dequeue(), 1)
	require.Equal(t, q.Dequeue(), 2)
	require.Equal(t, q.Dequeue(), 3)
}

func TestQueue3(t *testing.T) {
	capacity := 2
	q := New(capacity)

	q.Enqueue(1)
	q.Enqueue(2)
	require.Equal(t, q.Dequeue(), 1)
	q.Enqueue(3)

	require.Greater(t, q.front, q.behind)

	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)

	require.Equal(t, q.Dequeue(), 2)
	require.Equal(t, q.Dequeue(), 3)
	require.Equal(t, q.Dequeue(), 4)
	require.Equal(t, q.Dequeue(), 5)
	require.Equal(t, q.Dequeue(), 6)
}

func TestQueue_Grow(t *testing.T) {
	capacity := 2
	q := New(capacity)

	require.Equal(t, q.Cap(), 2)
	q.Grow(4)
	require.Equal(t, q.Cap(), 4)
}
