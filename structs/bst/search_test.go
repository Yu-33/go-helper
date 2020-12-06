package bst

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
)

func searchRangeByRecursion(root *MockNode, start Element, boundary Element) []Element {
	if root == nil {
		return nil
	}

	var result []Element

	var recursion func(node *MockNode, start Element, boundary Element)

	recursion = func(node *MockNode, start Element, boundary Element) {
		if node == nil {
			return
		}
		if start != nil && node.element.Compare(start) == -1 {
			recursion(node.right, start, boundary)
		} else if boundary != nil && node.element.Compare(boundary) == 1 {
			recursion(node.left, start, boundary)
		} else {
			// start <= node <= boundary
			recursion(node.left, start, boundary)
			result = append(result, node.element)
			recursion(node.right, start, boundary)
		}
	}

	recursion(root, start, boundary)

	return result
}

func searchRangeByIter(root *MockNode, start Element, boundary Element) []Element {
	if root == nil {
		return nil
	}
	var result []Element

	it := NewIterator(root, start, boundary)

	for it.Valid() {
		result = append(result, it.Next())
	}

	return result
}

func TestSearchRange(t *testing.T) {
	tr := &MockTree{}

	tr.Insert(container.Int64(24))
	tr.Insert(container.Int64(61))
	tr.Insert(container.Int64(67))
	tr.Insert(container.Int64(84))
	tr.Insert(container.Int64(91))
	tr.Insert(container.Int64(130))
	tr.Insert(container.Int64(133))
	tr.Insert(container.Int64(145))
	tr.Insert(container.Int64(150))
	tr.Insert(container.Int64(87))
	tr.Insert(container.Int64(97))
	tr.Insert(container.Int64(22))
	tr.Insert(container.Int64(35))
	tr.Insert(container.Int64(64))
	tr.Insert(container.Int64(76))

	// seeds := []container.Int64{22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150}

	var r1, r2, r3 []Element

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
	require.Equal(t, len(r1), 1)
	require.Equal(t, r1[0], container.Int64(22))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(77))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(77))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(77))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(147))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(147))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(147))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(150))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(150))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(150))
	require.Equal(t, r1[len(r1)-1], container.Int64(150))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, nil, container.Int64(156))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(156))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(156))
	require.Equal(t, r1[len(r1)-1], container.Int64(150))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	/* ---  test start != nil && boundary == nil --- */

	r1 = SearchRange(tr.root, container.Int64(21), container.Int64(13))
	r2 = searchRangeByRecursion(tr.root, container.Int64(21), container.Int64(13))
	r3 = searchRangeByIter(tr.root, container.Int64(21), container.Int64(13))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(65), container.Int64(27))
	r2 = searchRangeByRecursion(tr.root, container.Int64(65), container.Int64(27))
	r3 = searchRangeByIter(tr.root, container.Int64(65), container.Int64(27))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(68), container.Int64(132))
	r2 = searchRangeByRecursion(tr.root, container.Int64(68), container.Int64(132))
	r3 = searchRangeByIter(tr.root, container.Int64(68), container.Int64(132))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = SearchRange(tr.root, container.Int64(21), container.Int64(156))
	r2 = searchRangeByRecursion(tr.root, container.Int64(21), container.Int64(156))
	r3 = searchRangeByIter(tr.root, container.Int64(21), container.Int64(156))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)
}

func TestSearchLastLT(t *testing.T) {
	tr := &MockTree{}

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	tr.Insert(container.Int64(24))
	tr.Insert(container.Int64(61))
	tr.Insert(container.Int64(67))
	tr.Insert(container.Int64(84))
	tr.Insert(container.Int64(91))
	tr.Insert(container.Int64(130))
	tr.Insert(container.Int64(133))
	tr.Insert(container.Int64(145))
	tr.Insert(container.Int64(150))
	tr.Insert(container.Int64(87))
	tr.Insert(container.Int64(97))
	tr.Insert(container.Int64(22))
	tr.Insert(container.Int64(35))
	tr.Insert(container.Int64(64))
	tr.Insert(container.Int64(76))

	var element Element

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	element = SearchLastLT(tr.root, container.Int64(21))
	require.Nil(t, element)

	element = SearchLastLT(tr.root, container.Int64(22))
	require.Nil(t, element)

	element = SearchLastLT(tr.root, container.Int64(25))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(24))

	element = SearchLastLT(tr.root, container.Int64(63))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(61))

	element = SearchLastLT(tr.root, container.Int64(77))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(76))

	element = SearchLastLT(tr.root, container.Int64(84))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(76))

	element = SearchLastLT(tr.root, container.Int64(99))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(97))

	element = SearchLastLT(tr.root, container.Int64(132))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(130))

	element = SearchLastLT(tr.root, container.Int64(133))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(130))

	element = SearchLastLT(tr.root, container.Int64(146))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(145))

	element = SearchLastLT(tr.root, container.Int64(150))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(145))

	element = SearchLastLT(tr.root, container.Int64(156))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(150))
}

func TestSearchLastLE(t *testing.T) {
	tr := &MockTree{}

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	tr.Insert(container.Int64(24))
	tr.Insert(container.Int64(61))
	tr.Insert(container.Int64(67))
	tr.Insert(container.Int64(84))
	tr.Insert(container.Int64(91))
	tr.Insert(container.Int64(130))
	tr.Insert(container.Int64(133))
	tr.Insert(container.Int64(145))
	tr.Insert(container.Int64(150))
	tr.Insert(container.Int64(87))
	tr.Insert(container.Int64(97))
	tr.Insert(container.Int64(22))
	tr.Insert(container.Int64(35))
	tr.Insert(container.Int64(64))
	tr.Insert(container.Int64(76))

	var element Element

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	element = SearchLastLE(tr.root, container.Int64(21))
	require.Nil(t, element)

	element = SearchLastLE(tr.root, container.Int64(22))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(22))

	element = SearchLastLE(tr.root, container.Int64(25))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(24))

	element = SearchLastLE(tr.root, container.Int64(63))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(61))

	element = SearchLastLE(tr.root, container.Int64(77))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(76))

	element = SearchLastLE(tr.root, container.Int64(76))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(76))

	element = SearchLastLE(tr.root, container.Int64(99))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(97))

	element = SearchLastLE(tr.root, container.Int64(132))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(130))

	element = SearchLastLE(tr.root, container.Int64(133))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(133))

	element = SearchLastLE(tr.root, container.Int64(146))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(145))

	element = SearchLastLE(tr.root, container.Int64(150))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(150))

	element = SearchLastLE(tr.root, container.Int64(156))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(150))
}

func TestSearchFirstGT(t *testing.T) {
	tr := &MockTree{}

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	tr.Insert(container.Int64(24))
	tr.Insert(container.Int64(61))
	tr.Insert(container.Int64(67))
	tr.Insert(container.Int64(84))
	tr.Insert(container.Int64(91))
	tr.Insert(container.Int64(130))
	tr.Insert(container.Int64(133))
	tr.Insert(container.Int64(145))
	tr.Insert(container.Int64(150))
	tr.Insert(container.Int64(87))
	tr.Insert(container.Int64(97))
	tr.Insert(container.Int64(22))
	tr.Insert(container.Int64(35))
	tr.Insert(container.Int64(64))
	tr.Insert(container.Int64(76))

	var element Element

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	element = SearchFirstGT(tr.root, container.Int64(21))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(22))

	element = SearchFirstGT(tr.root, container.Int64(24))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(35))

	element = SearchFirstGT(tr.root, container.Int64(25))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(35))

	element = SearchFirstGT(tr.root, container.Int64(63))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(64))

	element = SearchFirstGT(tr.root, container.Int64(77))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(84))

	element = SearchFirstGT(tr.root, container.Int64(99))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(130))

	element = SearchFirstGT(tr.root, container.Int64(132))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(133))

	element = SearchFirstGT(tr.root, container.Int64(133))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(145))

	element = SearchFirstGT(tr.root, container.Int64(147))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(150))

	element = SearchFirstGT(tr.root, container.Int64(150))
	require.Nil(t, element)
	element = SearchFirstGT(tr.root, container.Int64(151))
	require.Nil(t, element)
}

func TestSearchFirstGE(t *testing.T) {
	tr := &MockTree{}

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	tr.Insert(container.Int64(24))
	tr.Insert(container.Int64(61))
	tr.Insert(container.Int64(67))
	tr.Insert(container.Int64(84))
	tr.Insert(container.Int64(91))
	tr.Insert(container.Int64(130))
	tr.Insert(container.Int64(133))
	tr.Insert(container.Int64(145))
	tr.Insert(container.Int64(150))
	tr.Insert(container.Int64(87))
	tr.Insert(container.Int64(97))
	tr.Insert(container.Int64(22))
	tr.Insert(container.Int64(35))
	tr.Insert(container.Int64(64))
	tr.Insert(container.Int64(76))

	var element Element

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	element = SearchFirstGE(tr.root, container.Int64(21))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(22))

	element = SearchFirstGE(tr.root, container.Int64(24))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(24))

	element = SearchFirstGE(tr.root, container.Int64(25))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(35))

	element = SearchFirstGE(tr.root, container.Int64(63))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(64))

	element = SearchFirstGE(tr.root, container.Int64(77))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(84))

	element = SearchFirstGE(tr.root, container.Int64(99))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(130))

	element = SearchFirstGE(tr.root, container.Int64(132))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(133))

	element = SearchFirstGE(tr.root, container.Int64(133))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(133))

	element = SearchFirstGE(tr.root, container.Int64(146))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(150))

	element = SearchFirstGE(tr.root, container.Int64(150))
	require.NotNil(t, element)
	require.Equal(t, element, container.Int64(150))

	element = SearchFirstGE(tr.root, container.Int64(151))
	require.Nil(t, element)
}
