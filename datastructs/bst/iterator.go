package bst

import (
	"reflect"

	"github.com/Yu-33/gohelper/datastructs/stack"
)

type Iterator struct {
	s        *stack.Stack
	start    Element
	boundary Element
}

func NewIterator(root Node, start Element, boundary Element) *Iterator {
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
	return !it.s.Empty()
}

func (it *Iterator) Next() Element {
	if it.s.Empty() {
		return nil
	}

	p := it.s.Pop().(Node)
	elements := p.Element()

	fillStack(p.Right(), it.start, it.boundary, it.s)

	return elements
}

func fillStack(root Node, start Element, boundary Element, s *stack.Stack) {
	p := root
	for !reflect.ValueOf(p).IsNil() {
		if start != nil && p.Element().Compare(start) == -1 {
			p = p.Right()
			continue
		}
		if boundary != nil && p.Element().Compare(boundary) == 1 {
			p = p.Left()
			continue
		}

		s.Push(p)
		p = p.Left()
	}
}
