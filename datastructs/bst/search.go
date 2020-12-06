package bst

import (
	"reflect"

	"github.com/Yu-33/gohelper/datastructs/stack"
)

// SearchRange search for elements in a specified range from and and boundary (start <= k <= boundary)
func SearchRange(root Node, start Element, boundary Element) []Element {
	if root == nil {
		return nil
	}

	var result []Element

	s := stack.Default()
	p := root
	for !s.Empty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
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
		} else {
			p = s.Pop().(Node)
			result = append(result, p.Element())
			p = p.Right()
		}
	}

	return result
}

// SearchLastLT search for the last node that less than the 'key';
func SearchLastLT(root Node, key Element) Element {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Element())
		if flag == 1 {
			n = p
			p = p.Right()
		} else {
			p = p.Left()
		}
	}

	if n != nil {
		return n.Element()
	}

	return nil
}

// SearchLastLE search for the last node that less than or equal to the 'key';
func SearchLastLE(root Node, key Element) Element {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Element())
		if flag == 1 {
			n = p
			p = p.Right()
		} else if flag == -1 {
			p = p.Left()
		} else {
			n = p
			break
		}
	}

	if n != nil {
		return n.Element()
	}

	return nil
}

// SearchFirstGT search for the first node that greater than to the 'key';
func SearchFirstGT(root Node, key Element) Element {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Element())
		if flag == -1 {
			n = p
			p = p.Left()
		} else {
			p = p.Right()
		}
	}

	if n != nil {
		return n.Element()
	}

	return nil
}

// SearchFirstGE search for the first node that greater than or equal to the 'key';
func SearchFirstGE(root Node, key Element) Element {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Element())
		if flag == -1 {
			n = p
			p = p.Left()
		} else if flag == 1 {
			p = p.Right()
		} else {
			n = p
			break
		}
	}

	if n != nil {
		return n.Element()
	}

	return nil
}
