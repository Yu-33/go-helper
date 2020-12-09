package timewheel

import (
	"container/list"
	"sync"
	"sync/atomic"
)

// Each tick(time interval) have a bucket to store all timers(tasks) that belonging to this tick.
type bucket struct {
	expiration int64
	timers     *list.List
	mu         *sync.Mutex
	flushMu    *sync.Mutex // represents whether the bucket is performing flush.
}

func (b *bucket) getExpiration() int64 {
	return atomic.LoadInt64(&b.expiration)
}

func (b *bucket) setExpiration(expiration int64) bool {
	return atomic.SwapInt64(&b.expiration, expiration) != expiration
}

// insert add t to the b.timers, it only called by tw.add.
func (b *bucket) insert(t *Timer) {
	b.mu.Lock()

	e := b.timers.PushBack(t)
	t.setBucket(b)
	t.element = e

	b.mu.Unlock()
}

// delete remove t from the TimeWheel, it only called by t.Close().
//
// return true indicates t has been removed from TimeWheel.
// return false indicates the bucket has changed and can't perform delete operations.
func (b *bucket) delete(t *Timer) bool {
	b.flushMu.Lock()
	b.mu.Lock()

	ok := true
	if t.getBucket() != b {
		// If delete is called just after the TimeWheel's goroutine has under cases:
		//   - moved t from the b to another non-nil bucket "ab" (through: tw.process -> b.flush -> tw.submit -> tw.add -> ab.insert)
		//     and in the case, t.getBucket will return "ab".
		//
		// In either cases, the return value maybe does not equal to b.
		// return false to make the caller to try again.
		ok = false
	} else if t.element == nil {
		// If delete is called after following cases happens:
		//   1. the timer t add by tw.AfterFunc.
		//   2. the next time is zero in tw.Schedule.
		// In either cases, the timer t not in TimeWheel and is nil (set by b.flush),
		// and it can be considered as a successful deletion.
	} else {
		b.timers.Remove(t.element)
		t.setBucket(nil)
		t.element = nil
	}

	b.mu.Unlock()
	b.flushMu.Unlock()

	return ok
}

func (b *bucket) flush(submit func(*Timer)) {
	b.flushMu.Lock()
	b.mu.Lock()

	timers := b.timers
	// Reset the times in bucket.
	b.timers = list.New()
	b.setExpiration(-1)

	b.mu.Unlock()

	// Re submit and remove the Timer from list.
	for e := timers.Front(); e != nil; {
		next := e.Next()

		v := timers.Remove(e)
		t := v.(*Timer)

		// The timer t may not re-enqueue in the following cases:
		//   1. the timer add by tw.AfterFunc.
		//   2. the next time is zero in tw.Schedule.
		// Thus, set the t.element to nil before submit to prevents unexpected when call t.Close.
		t.element = nil

		submit(t)

		e = next
	}

	b.flushMu.Unlock()
}

func newBucket() *bucket {
	return &bucket{
		expiration: -1,
		timers:     list.New(),
		mu:         new(sync.Mutex),
		flushMu:    new(sync.Mutex),
	}
}

func createBuckets(n int) []*bucket {
	buckets := make([]*bucket, n)
	for i := 0; i < n; i++ {
		buckets[i] = newBucket()
	}
	return buckets
}
