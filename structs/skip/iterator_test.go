package skip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
)

func TestList_Iter(t *testing.T) {
	sl := New()

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	seeds := []container.Int64{22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150}

	for _, k := range seeds {
		sl.Insert(k, int64(k*2+1))
	}

	var iter container.Iterator
	var kv KV

	/* ------ test start == nil && boundary == nil */

	iter = sl.Iter(nil, nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	/* ---  test start != nil && boundary == nil --- */
	//
	// start < first node
	iter = sl.Iter(container.Int64(21), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start == first node
	iter = sl.Iter(container.Int64(22), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start > first node && start < last node
	iter = sl.Iter(container.Int64(27), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 2; i < len(seeds); i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start > first node && start < last node
	iter = sl.Iter(container.Int64(62), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 4; i < len(seeds); i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start > root node && start < last node
	iter = sl.Iter(container.Int64(132), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 12; i < len(seeds); i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start == last node
	iter = sl.Iter(container.Int64(150), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	kv = iter.Next()
	require.Equal(t, kv.Key(), container.Int64(150))
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start > last node
	iter = sl.Iter(container.Int64(156), nil)
	require.NotNil(t, iter)
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	/* ---  test start == nil && boundary != nil --- */
	//
	// boundary < first node
	iter = sl.Iter(nil, container.Int64(21))
	require.NotNil(t, iter)
	require.False(t, iter.Valid())
	kv = iter.Next()
	require.Nil(t, kv)

	// boundary == first node
	iter = sl.Iter(nil, container.Int64(22))
	require.NotNil(t, iter)
	require.False(t, iter.Valid())
	kv = iter.Next()
	require.Nil(t, kv)

	// boundary > first node
	iter = sl.Iter(nil, container.Int64(24))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	kv = iter.Next()
	require.Equal(t, kv.Key(), container.Int64(22))
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// boundary < last node && bound > first node
	iter = sl.Iter(nil, container.Int64(147))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds)-1; i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// boundary == last node
	iter = sl.Iter(nil, container.Int64(150))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds)-1; i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// boundary > last node
	iter = sl.Iter(nil, container.Int64(156))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := range seeds {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	/* ---  test start != nil && boundary != nil --- */
	//

	// start < boundary && start > first node && bound < last node
	iter = sl.Iter(container.Int64(68), container.Int64(132))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 6; i < len(seeds)-3; i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start < boundary && start < first node && bound > last node
	iter = sl.Iter(container.Int64(21), container.Int64(153))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		kv := iter.Next()
		require.NotNil(t, kv, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, kv.Key(), seeds[i])
		require.Equal(t, kv.Value(), int64(seeds[i]*2+1))
	}
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start == boundary, start and boundary exists.
	iter = sl.Iter(container.Int64(24), container.Int64(24))
	require.NotNil(t, iter)
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start == boundary, start and boundary not exists.
	iter = sl.Iter(container.Int64(25), container.Int64(25))
	require.NotNil(t, iter)
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start < boundary && start < first node && bound < first node
	iter = sl.Iter(container.Int64(21), container.Int64(13))
	require.NotNil(t, iter)
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())

	// start > boundary && start > first node
	iter = sl.Iter(container.Int64(65), container.Int64(27))
	require.NotNil(t, iter)
	kv = iter.Next()
	require.Nil(t, kv)
	require.False(t, iter.Valid())
}
