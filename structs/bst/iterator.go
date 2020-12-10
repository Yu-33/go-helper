package bst

import (
	"reflect"

	"github.com/Yu-33/gohelper/structs/stack"
)

// Iterator implements universal iterators are used for all-type binary trees.
//
// Iterator used to get the specified range of data.
// The range is start <= x < boundary, and we allowed the start or boundary is nil.
// And also implements interface container.Iterator.
type Iterator struct {
	s        *stack.Stack
	start    Key
	boundary Key
}

// New creates an Iterator with given parameters.
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

// Valid represents whether have more elements.
//
// And also implements interface container.Iterator.
func (it *Iterator) Valid() bool {
	return !it.s.Empty()
}

// Next returns a k/v pair and moved the iterator to the next pair.
// Returns nil if no more elements.
//
// And also implements interface container.Iterator.
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
