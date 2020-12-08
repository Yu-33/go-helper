package timewheel

import (
	"container/list"
	"sync"
	"sync/atomic"
)

// Each tick(time interval) has a bucket and the bucket store all timers belonging to this tick.
type bucket struct {
	expiration int64
	mu         *sync.Mutex
	timers     *list.List
}

func (b *bucket) getExpiration() int64 {
	return atomic.LoadInt64(&b.expiration)
}

func (b *bucket) setExpiration(expiration int64) bool {
	return atomic.SwapInt64(&b.expiration, expiration) != expiration
}

func (b *bucket) insert(t *Timer) {
	b.mu.Lock()
	defer b.mu.Unlock()

	e := b.timers.PushBack(t)
	t.element = e
	t.setBucket(b)
}

func (b *bucket) delete(t *Timer) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.remove(t)
}

func (b *bucket) remove(t *Timer) bool {
	if t.getBucket() != b {
		// If remove is called from t.Stop, and this happens just after the timing wheel's goroutine has:
		//     1. removed t from b (through b.Flush -> b.remove)
		//     2. moved t from b to another bucket ab (through b.Flush -> b.remove and ab.Add)
		// then t.getBucket will return nil for case 1, or ab (non-nil) for case 2.
		// In either case, the returned value does not equal to b.
		return false
	}
	b.timers.Remove(t.element)
	t.setBucket(nil)
	t.element = nil
	return true
}

func (b *bucket) flush(reinsert func(*Timer)) {
	b.mu.Lock()

	timers := make([]*Timer, 0, b.timers.Len())
	for e := b.timers.Front(); e != nil; {
		next := e.Next()

		t := e.Value.(*Timer)
		b.remove(t)
		timers = append(timers, t)

		e = next
	}
	b.mu.Unlock()

	b.setExpiration(-1)

	for i := 0; i < len(timers); i++ {
		reinsert(timers[i])
	}
}

func createBuckets(size int) []*bucket {
	buckets := make([]*bucket, size)
	for i := 0; i < size; i++ {
		buckets[i] = &bucket{
			expiration: -1,
			mu:         new(sync.Mutex),
			timers:     list.New(),
		}
	}
	return buckets
}
