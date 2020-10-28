package bst

import (
	"math/rand"
	"time"

	"github.com/Yu-33/gohelper/datastructs/container"
)

type MockNode struct {
	elements Elements
	left     *MockNode
	right    *MockNode
}

func (n *MockNode) Elements() Elements {
	return n.elements
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

func (tr *MockTree) Insert(elements Elements) bool {
	p := tr.root
	for p != nil {
		flag := elements.Compare(p.elements)
		if flag == -1 {
			if p.left == nil {
				p.left = &MockNode{elements: elements}
				break
			}
			p = p.left
		} else if flag == 1 {
			if p.right == nil {
				p.right = &MockNode{elements: elements}
				break
			}
			p = p.right
		} else {
			return false
		}
	}

	if p == nil {
		tr.root = &MockNode{elements: elements}
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
