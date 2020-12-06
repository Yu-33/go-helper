package stack

const (
	defaultCapacity = 32
)

// Stack implements a stack THAT LIFO.
type Stack struct {
	items []interface{}
	cap   int
	len   int
}

// Default return a Stack with default parameters.
func Default() *Stack {
	return New(defaultCapacity)
}

// New creates a Stack with the specified the initialization capacity.
func New(n int) *Stack {
	s := &Stack{
		items: make([]interface{}, n),
		cap:   n,
		len:   0,
	}
	return s
}

func (s *Stack) grow(n int) {
	if n > s.cap {
		items := s.items
		s.cap = n
		s.items = make([]interface{}, s.cap)
		copy(s.items, items)
	}
}

// Len return the number of elements in the stack.
func (s *Stack) Len() int {
	return s.len
}

// Cap return the current capacity of the stack.
func (s *Stack) Cap() int {
	return s.cap
}

// Empty indicates whether the stack is empty.
func (s *Stack) Empty() bool {
	return s.len == 0
}

// Push add element to the end of stack.
func (s *Stack) Push(item interface{}) {
	if s.cap == s.len {
		s.grow(s.cap << 1)
	}
	s.items[s.len] = item
	s.len++
}

// Pop returns and removes the element that at the end.
func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	s.len--
	item := s.items[s.len]
	return item
}

// Peek returns the element that at the end.
func (s *Stack) Peek() interface{} {
	if s.Empty() {
		return nil
	}
	item := s.items[s.len-1]
	return item
}
