package bst

import (
	"reflect"

	"github.com/Yu-33/gohelper/structs/stack"
)

func InOrder(node Node) []Element {
	if node == nil {
		return nil
	}

	result := make([]Element, 0)
	s := stack.Default()
	p := node

	var i int

	for !s.Empty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
			s.Push(p)
			p = p.Left()
		} else {
			p = s.Pop().(Node)
			result = append(result, p.Element())
			p = p.Right()
		}

		i++
	}
	//for !s.Empty() || !reflect.ValueOf(p).IsNil() {
	//	for !reflect.ValueOf(p).IsNil() {
	//		s.Push(p)
	//		p = p.Left()
	//	}
	//
	//	if !s.Empty() {
	//		p = s.Pop().(Node)
	//		result = append(result, p.Element())
	//		p = p.Right()
	//	}
	//}

	return result
}

func PreOrder(node Node) []Element {
	if node == nil {
		return nil
	}

	result := make([]Element, 0)
	s := stack.Default()
	p := node

	for !s.Empty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
			result = append(result, p.Element())
			s.Push(p)
			p = p.Left()
		} else {
			p = s.Pop().(Node)
			p = p.Right()
		}
	}

	return result
}

func PostOrder(node Node) []Element {
	if node == nil {
		return nil
	}

	result := make([]Element, 0)

	var lastVisit Node

	s := stack.Default()
	p := node

	for !reflect.ValueOf(p).IsNil() {
		s.Push(p)
		p = p.Left()
	}

	for !s.Empty() {
		p = s.Pop().(Node)
		if reflect.ValueOf(p.Right()).IsNil() || p.Right() == lastVisit {
			result = append(result, p.Element())
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

	return result
}
