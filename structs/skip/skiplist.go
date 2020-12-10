package skip

import (
	"math/rand"
	"time"

	"github.com/Yu-33/gohelper/structs/container"
)

type Key = container.Key
type Value = container.Value
type KV = container.KV

const (
	maxLevel = 0x1f
)

// listNode used in skip list and implements container.KV.
//
// And also implements interface container.KV.
type listNode struct {
	key   Key
	value Value
	next  []*listNode
}

// Implements interface container.KV.
func (n *listNode) Key() Key {
	return n.key
}

// Implements interface container.KV.
func (n *listNode) Value() Value {
	return n.value
}

// List implements data struct of skip list.
//
// And also implements interface container.Container
type List struct {
	head  *listNode
	level int
	lens  []int
	r     *rand.Rand
}

// New creates an skip list.
func New() *List {
	sl := new(List)
	sl.head = sl.createNode(nil, nil, maxLevel)
	sl.level = 0
	sl.lens = make([]int, maxLevel+1)
	sl.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return sl
}

// Len return number of elements.
func (sl *List) Len() int {
	return sl.lens[0]
}

// Insert inserts the key with value in the container.
// k and v must not be nil, otherwise it will crash.
// Returns false if key already exists.
func (sl *List) Insert(k Key, v Value) bool {
	var updates [maxLevel + 1]*listNode

	level := sl.chooseLevel()
	if level > sl.level {
		sl.level = level
	}

	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			// The key already exists. Not allowed duplicates.
			return false
		}
		updates[i] = p
	}

	node := sl.createNode(k, v, level)
	for i := 0; i <= level; i++ {
		node.next[i] = updates[i].next[i]
		updates[i].next[i] = node
		sl.lens[i]++
	}

	return true
}

// Delete remove and returns the value of the specified key.
// Returns nil if not found.
func (sl *List) Delete(k Key) Value {
	var d *listNode
	p := sl.head

	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
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

	return d.value
}

// Search get the value of specified key.
// Returns nil if not found.
func (sl *List) Search(k Key) Value {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			return p.next[i].value
		}
	}
	return nil
}

// Iter return a Iterator, it's a wraps for skip.Iterator
func (sl *List) Iter(start Key, boundary Key) container.Iterator {
	iter := newIterator(sl, start, boundary)
	return iter
}

func (sl *List) createNode(k Key, v Value, level int) *listNode {
	return &listNode{
		key:   k,
		value: v,
		next:  make([]*listNode, level+1),
	}
}

func (sl *List) chooseLevel() int {
	level := 0
	for sl.r.Int63()&1 == 1 && level < maxLevel {
		level++
	}
	return level
}

// Search the last node that less than the 'key'.
func (sl *List) searchLastLT(k Key) *listNode {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}

		if i == 0 && p.key != nil {
			return p
		}
	}
	return nil
}

// Search the last node that less than or equal to the 'key'.
func (sl *List) searchLastLE(k Key) *listNode {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			return p.next[i]
		} else if i == 0 && p.key != nil {
			return p
		}

	}
	return nil
}

// Search the first node that greater than to the 'key'.
func (sl *List) searchFirstGT(k Key) *listNode {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil {
			if p.next[i].key.Compare(k) == 0 {
				return p.next[i].next[0]
			}
			if i == 0 {
				return p.next[i]
			}
		}

	}
	return nil
}

// Search the first node that greater than or equal to the 'key'.
func (sl *List) searchFirstGE(k Key) *listNode {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil {
			if p.next[i].key.Compare(k) == 0 || i == 0 {
				return p.next[i]
			}
		}

	}
	return nil
}
