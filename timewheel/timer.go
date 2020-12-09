package timewheel

import (
	"container/list"
	"sync/atomic"
	"unsafe"
)

// Timer represents a single event. The given task will be executed when the timer expires.
type Timer struct {
	expiration int64 // in nanoseconds.
	task       func()

	// The bucket that holds the list to which this timer's element belongs.
	//
	// NOTICE: This field may be updated and read concurrently,
	// through Timer.Close() and Bucket.flush().
	b unsafe.Pointer // type: *bucket

	// The timer's Element in list.
	element *list.Element
}

func (t *Timer) getBucket() *bucket {
	return (*bucket)(atomic.LoadPointer(&t.b))
}

func (t *Timer) setBucket(b *bucket) {
	atomic.StorePointer(&t.b, unsafe.Pointer(b))
}

// Close prevents the Timer from firing.
//
// The func will be block until the timer has finally been removed from the TimeWheel.
// But, if the timer t has already expired and the t.task has been started in its own
// goroutine; Close does not wait for t.task to complete before returning. If the caller
// needs to know whether t.task is completed, it must coordinate with t.task explicitly.
func (t *Timer) Close() {
	for b := t.getBucket(); b != nil; b = t.getBucket() {
		// The b.delete may fail if t's bucket has changed due to TimeWheel call the b.flush.
		// Thus, we re-get t's possibly new bucket and retry until the bucket becomes nil or
		// delete successful.
		if ok := b.delete(t); ok {
			break
		}
	}
}
