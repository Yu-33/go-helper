package bst

import (
	"reflect"

	"github.com/Yu-33/helper/datastructs/stack"
)

type Iterator struct {
	s        *stack.Stack
	start    Elements
	boundary Elements
}

func NewIterator(root Node, start Elements, boundary Elements) *Iterator {
	s := stack.New(-1)

	fillStack(root, start, boundary, s)

	it := &Iterator{
		s:        s,
		start:    start,
		boundary: boundary,
	}

	return it
}

func (it *Iterator) Valid() bool {
	return !it.s.IsEmpty()
}

func (it *Iterator) Next() Elements {
	if it.s.IsEmpty() {
		return nil
	}

	p := it.s.Pop().(Node)
	elements := p.Elements()

	fillStack(p.Right(), it.start, it.boundary, it.s)

	return elements
}

func fillStack(root Node, start Elements, boundary Elements, s *stack.Stack) {
	p := root
	for !reflect.ValueOf(p).IsNil() {
		if start != nil && p.Elements().Compare(start) == -1 {
			p = p.Right()
			continue
		}
		if boundary != nil && p.Elements().Compare(boundary) == 1 {
			p = p.Left()
			continue
		}

		s.Push(p)
		p = p.Left()
	}
}
