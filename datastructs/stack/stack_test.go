package stack

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestStackNew(t *testing.T) {
	capacity := 17
	st := New(capacity)
	require.NotNil(t, st)
	require.NotNil(t, st.items)
	require.True(t, st.IsEmpty())
	require.Equal(t, st.Len(), 0)
	require.Equal(t, len(st.items), capacity)
	require.Equal(t, st.len, 0)
	require.Equal(t, st.cap, capacity)
}

func TestStack1(t *testing.T) {
	capacity := 17
	st := New(capacity)

	for x := 0; x < 2; x++ {
		// test push and make stack full
		for i := 0; i < capacity; i++ {
			st.Push(i)
		}
		require.Equal(t, st.Len(), capacity)
		require.False(t, st.IsEmpty())
		require.Equal(t, st.len, capacity)

		// test pop and make stack empty
		for i := capacity - 1; i >= 0; i-- {
			v := st.Pop()
			require.NotNil(t, v)
			require.Equal(t, v, i)
		}

		require.True(t, st.IsEmpty())
		require.Equal(t, st.len, 0)

		require.Nil(t, st.Pop())
	}
}

func TestStack2(t *testing.T) {
	capacity := 2
	st := New(capacity)

	p1 := unsafe.Pointer(&st.items[0])

	st.Push(1)
	st.Push(2)

	p2 := unsafe.Pointer(&st.items[0])
	require.Equal(t, p1, p2)

	st.Push(3)

	p3 := unsafe.Pointer(&st.items[0])
	require.NotEqual(t, p2, p3)

	require.Equal(t, st.cap, capacity*2)

	require.Equal(t, st.Pop(), 3)
	require.Equal(t, st.Pop(), 2)
	require.Equal(t, st.Pop(), 1)
}

func TestStack_Grow(t *testing.T) {
	capacity := 2
	st := New(capacity)
	require.Equal(t, st.Cap(), 2)
	st.Grow(4)
	require.Equal(t, st.Cap(), 4)
}
