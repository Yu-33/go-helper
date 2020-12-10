package bst

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
)

func searchRangeByRecursion(root *treeNode, start Key, boundary Key) []KV {
	if root == nil {
		return nil
	}

	var result []KV

	var recursion func(node *treeNode, start Key, boundary Key)

	recursion = func(node *treeNode, start Key, boundary Key) {
		if node == nil {
			return
		}
		if start != nil && node.key.Compare(start) == -1 {
			recursion(node.right, start, boundary)
		} else if boundary != nil && node.key.Compare(boundary) != -1 {
			recursion(node.left, start, boundary)
		} else {
			// start <= node <= boundary
			recursion(node.left, start, boundary)
			result = append(result, node)
			recursion(node.right, start, boundary)
		}
	}

	recursion(root, start, boundary)

	return result
}

func searchRangeByIter(root *treeNode, start Key, boundary Key) []KV {
	if root == nil {
		return nil
	}

	var result []KV

	it := NewIterator(root, start, boundary)

	for it.Valid() {
		n := it.Next()
		result = append(result, n)
	}

	return result
}

func TestSearchRange(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	var r1, r2, r3 []KV

	/* ------ test start == nil && boundary == nil */

	r1 = SearchRange(nil, nil, nil)
	r2 = searchRangeByRecursion(nil, nil, nil)
	r3 = searchRangeByIter(nil, nil, nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, nil)
	r2 = searchRangeByRecursion(tr.root, nil, nil)
	r3 = searchRangeByIter(tr.root, nil, nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	/* ---  test start != nil && boundary == nil --- */

	r1 = SearchRange(tr.root, container.Int64(21), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(21), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(21), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(22), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(22), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(22), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(27), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(27), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(27), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(62), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(62), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(62), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(132), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(132), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(132), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(144), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(144), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(144), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(150), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(150), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(150), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(156), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(156), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(156), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	/* ---  test start == nil && boundary != nil --- */

	r1 = SearchRange(tr.root, nil, container.Int64(21))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(21))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(21))
	require.Equal(t, len(r1), 0)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(22))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(22))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(22))
	require.Equal(t, len(r1), 0)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(77))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(77))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(77))
	require.Equal(t, len(r1), 7)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(147))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(147))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(147))
	require.Equal(t, len(r1), 14)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(150))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(150))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(150))
	require.Equal(t, len(r1), 14)
	require.Equal(t, r1[len(r1)-1].Key(), container.Int64(145))
	require.Equal(t, r1[len(r1)-1].Value(), int64(145*2+1))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(156))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(156))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(156))
	require.Equal(t, len(r1), 15)
	require.Equal(t, r1[len(r1)-1].Key(), container.Int64(150))
	require.Equal(t, r1[len(r1)-1].Value(), int64(150*2+1))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	/* ---  test start != nil && boundary == nil --- */

	r1 = SearchRange(tr.root, container.Int64(21), container.Int64(13))
	r2 = searchRangeByRecursion(tr.root, container.Int64(21), container.Int64(13))
	r3 = searchRangeByIter(tr.root, container.Int64(21), container.Int64(13))
	require.Equal(t, len(r1), 0)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(65), container.Int64(27))
	r2 = searchRangeByRecursion(tr.root, container.Int64(65), container.Int64(27))
	r3 = searchRangeByIter(tr.root, container.Int64(65), container.Int64(27))
	require.Equal(t, len(r1), 0)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(68), container.Int64(132))
	r2 = searchRangeByRecursion(tr.root, container.Int64(68), container.Int64(132))
	r3 = searchRangeByIter(tr.root, container.Int64(68), container.Int64(132))
	require.Equal(t, len(r1), 6)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(21), container.Int64(156))
	r2 = searchRangeByRecursion(tr.root, container.Int64(21), container.Int64(156))
	r3 = searchRangeByIter(tr.root, container.Int64(21), container.Int64(156))
	require.Equal(t, len(r1), 15)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)
}

func TestSearchLastLT(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	var kv KV

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	kv = SearchLastLT(tr.root, container.Int64(21))
	require.Nil(t, kv)

	kv = SearchLastLT(tr.root, container.Int64(22))
	require.Nil(t, kv)

	kv = SearchLastLT(tr.root, container.Int64(25))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(24))
	require.Equal(t, kv.Value(), int64(24*2+1))

	kv = SearchLastLT(tr.root, container.Int64(63))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(61))
	require.Equal(t, kv.Value(), int64(61*2+1))

	kv = SearchLastLT(tr.root, container.Int64(77))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(76))
	require.Equal(t, kv.Value(), int64(76*2+1))

	kv = SearchLastLT(tr.root, container.Int64(84))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(76))
	require.Equal(t, kv.Value(), int64(76*2+1))

	kv = SearchLastLT(tr.root, container.Int64(99))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(97))
	require.Equal(t, kv.Value(), int64(97*2+1))

	kv = SearchLastLT(tr.root, container.Int64(132))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(130))
	require.Equal(t, kv.Value(), int64(130*2+1))

	kv = SearchLastLT(tr.root, container.Int64(133))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(130))
	require.Equal(t, kv.Value(), int64(130*2+1))

	kv = SearchLastLT(tr.root, container.Int64(146))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(145))
	require.Equal(t, kv.Value(), int64(145*2+1))

	kv = SearchLastLT(tr.root, container.Int64(150))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(145))
	require.Equal(t, kv.Value(), int64(145*2+1))

	kv = SearchLastLT(tr.root, container.Int64(156))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(150))
	require.Equal(t, kv.Value(), int64(150*2+1))
}

func TestSearchLastLE(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	var kv KV

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	kv = SearchLastLE(tr.root, container.Int64(21))
	require.Nil(t, kv)

	kv = SearchLastLE(tr.root, container.Int64(22))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(22))
	require.Equal(t, kv.Value(), int64(22*2+1))

	kv = SearchLastLE(tr.root, container.Int64(25))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(24))
	require.Equal(t, kv.Value(), int64(24*2+1))

	kv = SearchLastLE(tr.root, container.Int64(63))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(61))
	require.Equal(t, kv.Value(), int64(61*2+1))

	kv = SearchLastLE(tr.root, container.Int64(77))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(76))
	require.Equal(t, kv.Value(), int64(76*2+1))

	kv = SearchLastLE(tr.root, container.Int64(76))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(76))
	require.Equal(t, kv.Value(), int64(76*2+1))

	kv = SearchLastLE(tr.root, container.Int64(99))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(97))
	require.Equal(t, kv.Value(), int64(97*2+1))

	kv = SearchLastLE(tr.root, container.Int64(132))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(130))
	require.Equal(t, kv.Value(), int64(130*2+1))

	kv = SearchLastLE(tr.root, container.Int64(133))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(133))
	require.Equal(t, kv.Value(), int64(133*2+1))

	kv = SearchLastLE(tr.root, container.Int64(146))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(145))
	require.Equal(t, kv.Value(), int64(145*2+1))

	kv = SearchLastLE(tr.root, container.Int64(150))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(150))
	require.Equal(t, kv.Value(), int64(150*2+1))

	kv = SearchLastLE(tr.root, container.Int64(156))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(150))
	require.Equal(t, kv.Value(), int64(150*2+1))
}

func TestSearchFirstGT(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	var kv KV

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	kv = SearchFirstGT(tr.root, container.Int64(21))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(22))
	require.Equal(t, kv.Value(), int64(22*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(24))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(35))
	require.Equal(t, kv.Value(), int64(35*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(25))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(35))
	require.Equal(t, kv.Value(), int64(35*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(63))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(64))
	require.Equal(t, kv.Value(), int64(64*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(77))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(84))
	require.Equal(t, kv.Value(), int64(84*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(99))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(130))
	require.Equal(t, kv.Value(), int64(130*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(132))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(133))
	require.Equal(t, kv.Value(), int64(133*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(133))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(145))
	require.Equal(t, kv.Value(), int64(145*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(147))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(150))
	require.Equal(t, kv.Value(), int64(150*2+1))

	kv = SearchFirstGT(tr.root, container.Int64(150))
	require.Nil(t, kv)
	kv = SearchFirstGT(tr.root, container.Int64(151))
	require.Nil(t, kv)
}

func TestSearchFirstGE(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	var kv KV

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	kv = SearchFirstGE(tr.root, container.Int64(21))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(22))
	require.Equal(t, kv.Value(), int64(22*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(24))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(24))
	require.Equal(t, kv.Value(), int64(24*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(25))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(35))
	require.Equal(t, kv.Value(), int64(35*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(63))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(64))
	require.Equal(t, kv.Value(), int64(64*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(77))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(84))
	require.Equal(t, kv.Value(), int64(84*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(99))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(130))
	require.Equal(t, kv.Value(), int64(130*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(132))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(133))
	require.Equal(t, kv.Value(), int64(133*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(133))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(133))
	require.Equal(t, kv.Value(), int64(133*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(146))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(150))
	require.Equal(t, kv.Value(), int64(150*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(150))
	require.NotNil(t, kv)
	require.Equal(t, kv.Key(), container.Int64(150))
	require.Equal(t, kv.Value(), int64(150*2+1))

	kv = SearchFirstGE(tr.root, container.Int64(151))
	require.Nil(t, kv)
}
