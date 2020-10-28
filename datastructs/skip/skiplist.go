package skip

import (
	"math/rand"
	"time"

	"github.com/Yu-33/gohelper/datastructs/container"
)

type Elements = container.Comparer

const (
	maxLevel = 0x1f
)

type Node struct {
	elements Elements
	next     []*Node
}

type List struct {
	head  *Node
	level int
	lens  []int
	r     *rand.Rand
}

func New() *List {
	sl := new(List)
	sl.head = sl.createNode(nil, maxLevel)
	sl.level = 0
	sl.lens = make([]int, maxLevel+1)
	sl.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return sl
}

// Len return number of elements
func (sl *List) Len() int {
	return sl.lens[0]
}

// Search for find the specified elements
func (sl *List) Search(elements Elements) Elements {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].elements.Compare(elements) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].elements.Compare(elements) == 0 {
			return p.next[i].elements
		}
	}
	return nil
}

// Insert for insert into specified elements, return false if duplicate;
func (sl *List) Insert(elements Elements) bool {
	var updates [maxLevel + 1]*Node

	level := sl.chooseLevel()
	if level > sl.level {
		sl.level = level
	}

	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].elements.Compare(elements) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].elements.Compare(elements) == 0 {
			return false
		}
		updates[i] = p
	}

	node := sl.createNode(elements, level)
	for i := 0; i <= level; i++ {
		node.next[i] = updates[i].next[i]
		updates[i].next[i] = node
		sl.lens[i]++
	}

	return true
}

// Delete for delete specified elements, return nil if not found
func (sl *List) Delete(elements Elements) Elements {
	var d *Node
	p := sl.head

	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].elements.Compare(elements) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].elements.Compare(elements) == 0 {
			if d == nil {
				d = p.next[i]
			}
			p.next[i] = p.next[i].next[i]
			sl.lens[i]--
		}
	}
	if d == nil {
		return nil
	}

	return d.elements
}

// Iter return a Iterator, include elements: start <= k <= boundary
// start == first node if start == nil and boundary == last node if boundary == nil
func (sl *List) Iter(start Elements, boundary Elements) container.Iterator {
	iter := newIterator(sl, start, boundary)
	return iter
}

func (sl *List) createNode(elements Elements, level int) *Node {
	n := new(Node)
	n.elements = elements
	n.next = make([]*Node, level+1)
	return n
}

func (sl *List) chooseLevel() int {
	level := 0
	for sl.r.Int63()&1 == 1 && level < maxLevel {
		level++
	}
	return level
}

// Search the last node that less than the 'key';
func (sl *List) searchLastLT(key Elements) *Node {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].elements.Compare(key) == -1 {
			p = p.next[i]
		}

		if i == 0 && p.elements != nil {
			return p
		}
	}

	return nil
}

// Search the last node that less than or equal to the 'key';
func (sl *List) searchLastLE(key Elements) *Node {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].elements.Compare(key) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil && p.next[i].elements.Compare(key) == 0 {
			return p.next[i]
		} else if i == 0 && p.elements != nil {
			return p
		}

	}

	return nil
}

// Search the first node that greater than to the 'key';
func (sl *List) searchFirstGT(key Elements) *Node {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].elements.Compare(key) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil {
			if p.next[i].elements.Compare(key) == 0 {
				return p.next[i].next[0]
			}
			if i == 0 {
				return p.next[i]
			}
		}

	}

	return nil
}

// Search the first node that greater than or equal to the 'key';
func (sl *List) searchFirstGE(key Elements) *Node {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].elements.Compare(key) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil {
			if p.next[i].elements.Compare(key) == 0 || i == 0 {
				return p.next[i]
			}
		}

	}

	return nil
}
