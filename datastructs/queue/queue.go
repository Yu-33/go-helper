package queue

const (
	defaultCapacity = 16
)

// Queue type implements a queue by dynamic array;
// Is not thread safe;
type Queue struct {
	items  []interface{}
	cap    int
	front  int
	behind int
}

// New return a new Queue Type, n is the initialization capacity;
// We will use the defaultCapacity if n <= 0
func New(n int) *Queue {
	if n <= 0 {
		n = defaultCapacity
	}

	n += 1
	q := &Queue{
		items:  make([]interface{}, n),
		cap:    n,
		front:  0,
		behind: 0,
	}
	return q
}

func (q *Queue) grow(n int) {
	if n > q.cap-1 {
		oldCap := q.cap
		oldLen := q.cap - 1
		items := q.items

		q.cap = n + 1
		q.items = make([]interface{}, q.cap)
		for i := 0; i < oldLen; i++ {
			q.items[i] = items[(i+q.front)%oldCap]
		}
		q.front = 0
		q.behind = oldLen
	}
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Len() int {
	return (q.behind - q.front + q.cap) % q.cap
}

func (q *Queue) Cap() int {
	return q.cap - 1
}

func (q *Queue) Grow(n int) {
	q.grow(n)
}

func (q *Queue) Enqueue(item interface{}) {
	if q.Len() == q.Cap() {
		q.grow((q.cap - 1) * 2)
	}
	q.items[q.behind] = item
	q.behind = (q.behind + 1) % q.cap
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := q.items[q.front]
	q.front = (q.front + 1) % q.cap
	return item
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[q.front]
}
