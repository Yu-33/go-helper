package maxheap

import (
	"github.com/Yu-33/gohelper/structs/container"
)

const (
	defaultCapacity = 64
)

type Element = container.Comparator

// MaxHeap implements max heap and can used as priority queue.
type MaxHeap struct {
	items []Element
	cap   int
	len   int
}

// Default creates an MaxHeap with default parameters.
func Default() *MaxHeap {
	return New(defaultCapacity)
}

// New creates an MaxHeap with given initialization capacity.
func New(c int) *MaxHeap {
	h := &MaxHeap{
		items: make([]Element, c),
		cap:   c,
		len:   0,
	}
	return h
}

func (h *MaxHeap) grow(c int) {
	if c > h.cap {
		items := h.items
		h.items = make([]Element, c)
		h.cap = c
		copy(h.items, items)
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// Len return the number of element in the heap.
func (h *MaxHeap) Len() int {
	return h.len
}

// Cap return the current capacity of the heap.
func (h *MaxHeap) Cap() int {
	return h.cap
}

// Empty indicates whether the heap is empty.
func (h *MaxHeap) Empty() bool {
	return h.len == 0
}

// Push add element to the heap, Return the index number of the location.
func (h *MaxHeap) Push(item Element) int {
	if h.Len() == h.Cap() {
		h.grow(h.cap * 2)
	}

	h.items[h.len] = item

	var k, p int
	i := h.len

	for {
		k = i

		p = (i - 1) >> 1 // parent
		if p >= 0 && h.items[i].Compare(h.items[p]) == 1 {
			i = p
		}

		if k == i {
			break
		}

		h.swap(k, i)
	}

	h.len++
	return i
}

// Pop returns and removes the element that at the head.
func (h *MaxHeap) Pop() Element {
	if h.Empty() {
		return nil
	}

	item := h.items[0]
	h.len--

	h.items[0] = h.items[h.len]

	var i, left, right int
	k := 0

	for {
		i = k

		left = (i << 1) + 1
		if left < h.len && h.items[k].Compare(h.items[left]) == -1 {
			k = left
		}

		right = (i << 1) + 2
		if right < h.len && h.items[k].Compare(h.items[right]) == -1 {
			k = right
		}

		if i == k {
			break
		}

		h.swap(i, k)
	}

	return item
}

// Peek returns the element that at the head.
func (h *MaxHeap) Peek() Element {
	if h.Empty() {
		return nil
	}

	return h.items[0]
}
