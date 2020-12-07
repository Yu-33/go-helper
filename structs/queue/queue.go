package queue

const (
	defaultCapacity = 64
)

// Queue type implements a queue by dynamic array.
// Is not thread safe.
type Queue struct {
	items  []interface{}
	cap    int
	front  int
	behind int
}

// Default creates an Queue with default parameters.
func Default() *Queue {
	return New(defaultCapacity)
}

// New creates an Queue with given initialization capacity.
func New(c int) *Queue {
	c += 1
	q := &Queue{
		items:  make([]interface{}, c),
		cap:    c,
		front:  0,
		behind: 0,
	}
	return q
}

func (q *Queue) grow(c int) {
	if c > q.cap-1 {
		oldCap := q.cap
		oldLen := q.cap - 1
		items := q.items

		q.cap = c + 1
		q.items = make([]interface{}, q.cap)
		for i := 0; i < oldLen; i++ {
			q.items[i] = items[(i+q.front)%oldCap]
		}
		q.front = 0
		q.behind = oldLen
	}
}

// Len return the number of elements in the queue.
func (q *Queue) Len() int {
	return (q.behind - q.front + q.cap) % q.cap
}

// Cap return the current capacity of the queue.
func (q *Queue) Cap() int {
	return q.cap - 1
}

// Empty indicates whether the queue is empty.
func (q *Queue) Empty() bool {
	return q.Len() == 0
}

// Push add element to the end of queue.
func (q *Queue) Push(item interface{}) {
	if q.Len() == q.Cap() {
		q.grow((q.cap - 1) * 2)
	}
	q.items[q.behind] = item
	q.behind = (q.behind + 1) % q.cap
}

// Pop returns and removes the element that at the head.
func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	item := q.items[q.front]
	q.front = (q.front + 1) % q.cap
	return item
}

// Peek returns the element that at the head.
func (q *Queue) Peek() interface{} {
	if q.Empty() {
		return nil
	}
	return q.items[q.front]
}
