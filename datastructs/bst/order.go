package bst

import (
	"reflect"

	"github.com/Yu-33/helper/datastructs/stack"
)

func InOrder(node Node) []Elements {
	if node == nil {
		return nil
	}

	result := make([]Elements, 0)
	s := stack.New(-1)
	p := node

	var i int

	for !s.IsEmpty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
			s.Push(p)
			p = p.Left()
		} else {
			p = s.Pop().(Node)
			result = append(result, p.Elements())
			p = p.Right()
		}

		i++
	}
	//for !s.IsEmpty() || !reflect.ValueOf(p).IsNil() {
	//	for !reflect.ValueOf(p).IsNil() {
	//		s.Push(p)
	//		p = p.Left()
	//	}
	//
	//	if !s.IsEmpty() {
	//		p = s.Pop().(Node)
	//		result = append(result, p.Elements())
	//		p = p.Right()
	//	}
	//}

	return result
}

func PreOrder(node Node) []Elements {
	if node == nil {
		return nil
	}

	result := make([]Elements, 0)
	s := stack.New(-1)
	p := node

	for !s.IsEmpty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
			result = append(result, p.Elements())
			s.Push(p)
			p = p.Left()
		} else {
			p = s.Pop().(Node)
			p = p.Right()
		}
	}

	return result
}

func PostOrder(node Node) []Elements {
	if node == nil {
		return nil
	}

	result := make([]Elements, 0)

	var lastVisit Node

	s := stack.New(-1)
	p := node

	for !reflect.ValueOf(p).IsNil() {
		s.Push(p)
		p = p.Left()
	}

	for !s.IsEmpty() {
		p = s.Pop().(Node)
		if reflect.ValueOf(p.Right()).IsNil() || p.Right() == lastVisit {
			result = append(result, p.Elements())
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
