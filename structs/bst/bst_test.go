package bst

import (
	"math/rand"
	"time"

	"github.com/Yu-33/gohelper/structs/container"
)

type MockNode struct {
	element Element
	left    *MockNode
	right   *MockNode
}

func (n *MockNode) Element() Element {
	return n.element
}

func (n *MockNode) Left() Node {
	return n.left
}

func (n *MockNode) Right() Node {
	return n.right
}

type MockTree struct {
	root *MockNode
}

func (tr *MockTree) Insert(element Element) bool {
	p := tr.root
	for p != nil {
		flag := element.Compare(p.element)
		if flag == -1 {
			if p.left == nil {
				p.left = &MockNode{element: element}
				break
			}
			p = p.left
		} else if flag == 1 {
			if p.right == nil {
				p.right = &MockNode{element: element}
				break
			}
			p = p.right
		} else {
			return false
		}
	}

	if p == nil {
		tr.root = &MockNode{element: element}
	}

	return true
}

func buildBSTree() (tr *MockTree) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 25
	maxKey := length * 10

	tr = &MockTree{}

	for i := 0; i < length; i++ {
		for {
			k := container.Int64(r.Intn(maxKey) + 1)
			if ok := tr.Insert(k); ok {
				break
			}
		}
	}

	return
}
