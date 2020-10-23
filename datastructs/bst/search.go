package bst

import (
	"reflect"

	"github.com/Yu-33/helper/datastructs/stack"
)

// SearchRange search for elements in a specified range from and and boundary (start <= k <= boundary)
func SearchRange(root Node, start Elements, boundary Elements) []Elements {
	if root == nil {
		return nil
	}

	var result []Elements

	s := stack.New(-1)
	p := root
	for !s.IsEmpty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
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
		} else {
			p = s.Pop().(Node)
			result = append(result, p.Elements())
			p = p.Right()
		}
	}

	return result
}

// SearchLastLT search for the last node that less than the 'key';
func SearchLastLT(root Node, key Elements) Elements {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Elements())
		if flag == 1 {
			n = p
			p = p.Right()
		} else {
			p = p.Left()
		}
	}

	if n != nil {
		return n.Elements()
	}

	return nil
}

// SearchLastLE search for the last node that less than or equal to the 'key';
func SearchLastLE(root Node, key Elements) Elements {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Elements())
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
		return n.Elements()
	}

	return nil
}

// SearchFirstGT search for the first node that greater than to the 'key';
func SearchFirstGT(root Node, key Elements) Elements {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Elements())
		if flag == -1 {
			n = p
			p = p.Left()
		} else {
			p = p.Right()
		}
	}

	if n != nil {
		return n.Elements()
	}

	return nil
}

// SearchFirstGE search for the first node that greater than or equal to the 'key';
func SearchFirstGE(root Node, key Elements) Elements {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Elements())
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
		return n.Elements()
	}

	return nil
}
