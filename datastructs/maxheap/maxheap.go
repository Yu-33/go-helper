package maxheap

import (
	"github.com/Yu-33/gohelper/datastructs/container"
)

const (
	defaultCapacity = 16
)

type Element = container.Comparer

// MaxHeap implements max heap and can used as priority queue.
type MaxHeap struct {
	items []Element
	cap   int
	len   int
}

// New return a MaxHeap with the specifies initialization cap.
// We will use the defaultCapacity if n <= 0.
func New(c int) *MaxHeap {
	if c <= 0 {
		c = defaultCapacity
	}
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

// Len return the number of elements in the heap.
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

// Push add element to the heap.
func (h *MaxHeap) Push(item Element) {
	if h.Len() == h.Cap() {
		h.grow(h.cap * 2)
	}

	h.items[h.len] = item

	var i, p int
	k := h.len

	for {
		i = k

		p = (k - 1) >> 1 // parent
		if p >= 0 && h.items[k].Compare(h.items[p]) == 1 {
			k = p
		}

		if i == k {
			break
		}

		h.swap(i, k)
	}

	h.len++
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
