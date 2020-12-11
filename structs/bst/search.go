package bst

import (
	"reflect"

	"github.com/Yu-33/gohelper/structs/stack"
)

// SearchRange to get the data for the specified range.
//
// Range interval: ( start <= x < boundary ).
// We will return data from the beginning if start is nil,
// And return data util the end if boundary is nil.
func SearchRange(root Node, start Key, boundary Key, f func(n Node)) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root
	for !s.Empty() || !reflect.ValueOf(p).IsNil() {
		if !reflect.ValueOf(p).IsNil() {
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
		} else {
			n := s.Pop().(Node)
			p = n.Right()

			f(n)
		}
	}
}

// SearchLastLT search for the last node that less than the key.
func SearchLastLT(root Node, key Key) Node {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Key())
		if flag == 1 {
			n = p
			p = p.Right()
		} else {
			p = p.Left()
		}
	}

	return n
}

// SearchLastLE search for the last node that less than or equal to the key.
func SearchLastLE(root Node, key Key) Node {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Key())
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

	return n
}

// SearchFirstGT search for the first node that greater than to the key.
func SearchFirstGT(root Node, key Key) Node {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Key())
		if flag == -1 {
			n = p
			p = p.Left()
		} else {
			p = p.Right()
		}
	}

	return n
}

// SearchFirstGE search for the first node that greater than or equal to the key.
func SearchFirstGE(root Node, key Key) Node {
	if root == nil || key == nil {
		return nil
	}

	var n Node

	p := root
	for !reflect.ValueOf(p).IsNil() {
		flag := key.Compare(p.Key())
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

	return n
}
