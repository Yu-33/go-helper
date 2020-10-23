package stack

const (
	defaultCapacity = 16
)

// Stack implements a stack THAT LIFO
type Stack struct {
	items []interface{}
	cap   int
	len   int
}

// New return a new Stack Type, n is the initialization capacity;
// We will use the defaultCapacity if n <= 0
func New(n int) *Stack {
	if n <= 0 {
		n = defaultCapacity
	}
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

func (s *Stack) Cap() int {
	return s.cap
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Grow(n int) {
	s.grow(n)
}

func (s *Stack) Push(item interface{}) {
	if s.cap == s.len {
		s.grow(s.cap << 1)
	}
	s.items[s.len] = item
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	s.len--
	item := s.items[s.len]
	return item
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	item := s.items[s.len-1]
	return item
}
