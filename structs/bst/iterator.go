package bst

import (
	"reflect"

	"github.com/Yu-33/gohelper/structs/stack"
)

// Iterator to get the data for the specified range.
//
// Range interval: ( start <= x < boundary ).
// We will return data from the beginning if start is nil,
// And return data util the end if boundary is nil.
//
// The Iterator return data with in-order traversal,
// And its can used for all-type binary search trees.
//
// And it is also the implementation of interface container.Iterator
type Iterator struct {
	s        *stack.Stack
	start    Key
	boundary Key
}

// NewIterator creates an Iterator with given parameters.
func NewIterator(root Node, start Key, boundary Key) *Iterator {
	s := stack.Default()

	fillStack(root, start, boundary, s)

	it := &Iterator{
		s:        s,
		start:    start,
		boundary: boundary,
	}

	return it
}

// Valid represents whether to have more elements in the Iterator.
func (it *Iterator) Valid() bool {
	return !it.s.Empty()
}

// Next returns a k/v pair and moved the iterator to the next pair.
// Returns nil if no more elements.
func (it *Iterator) Next() KV {
	if it.s.Empty() {
		return nil
	}

	p := it.s.Pop().(Node)
	n := p

	fillStack(p.Right(), it.start, it.boundary, it.s)

	return n
}

func fillStack(root Node, start Key, boundary Key, s *stack.Stack) {
	p := root
	for !reflect.ValueOf(p).IsNil() {
		if start != nil && p.Key().Compare(start) == -1 {
			p = p.Right()
			continue
		}
		if boundary != nil && p.Key().Compare(boundary) != -1 {
			p = p.Left()
			continue
		}

		s.Push(p)
		p = p.Left()
	}
}
