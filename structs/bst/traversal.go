package bst

import (
	"reflect"

	"github.com/Yu-33/gohelper/structs/stack"
)

// LDR return node by in-order traversal.
// The order: Left -> Mid -> Right.
func LDR(root Node, f func(n Node)) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root

	for !s.Empty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
			s.Push(p)
			p = p.Left()
		} else {
			n := s.Pop().(Node)
			p = n.Right()

			f(n)
		}
	}

	//for !s.Empty() || !reflect.ValueOf(p).IsNil() {
	//	for !reflect.ValueOf(p).IsNil() {
	//		s.Push(p)
	//		p = p.Left()
	//	}
	//
	//	if !s.Empty() {
	//		n := s.Pop().(Node)
	//		f(n)
	//		p = n.Right()
	//	}
	//}
}

// DLR return node by pre-order traversal.
// The order: Mid -> Left -> Right.
func DLR(root Node, f func(n Node)) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root

	for !s.Empty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
			f(p)

			s.Push(p)
			p = p.Left()
		} else {
			p = s.Pop().(Node)
			p = p.Right()
		}
	}
}

// LRD return node by post-order traversal.
// The order: Left -> Right -> Mid.
func LRD(root Node, f func(n Node)) {
	if root == nil {
		return
	}

	var lastVisit Node

	s := stack.Default()
	p := root

	for !reflect.ValueOf(p).IsNil() {
		s.Push(p)
		p = p.Left()
	}

	for !s.Empty() {
		p = s.Pop().(Node)
		if reflect.ValueOf(p.Right()).IsNil() || p.Right() == lastVisit {
			f(p)

			lastVisit = p
		} else {
			s.Push(p)
			p = p.Right()
			for !reflect.ValueOf(p).IsNil() {
				s.Push(p)
				p = p.Left()
			}
		}
	}
}
