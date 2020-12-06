package dqueue

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/Yu-33/gohelper/structs/container"
	"github.com/Yu-33/gohelper/structs/minheap"
)

const (
	defaultCapacity = 64
	defaultChanSize = 8
)

type Value interface{}
type Receiver func(value Value)

// Item declares a data-type used in priority queue.
type Item struct {
	Expiration int64 // Expiration time of nano timestamp
	Value      Value
}

// Implements container.Comparer.
func (k1 *Item) Compare(target container.Comparer) int {
	k2 := target.(*Item)
	if k1.Expiration < k2.Expiration {
		return -1
	}
	if k1.Expiration > k2.Expiration {
		return 1
	}
	return 0
}

// DQueue implements a delay queue base on priority queue (min heap).
// Inspired by https://github.com/RussellLuo/timingwheel/blob/master/delayqueue/delayqueue.go.
type DQueue struct {
	C chan Value // Notify channel

	mu *sync.Mutex
	pq *minheap.MinHeap // priority queue implemented by min heap.

	sleeping int32         // Similar to the sleeping state of runtime.timers. 1 => true, 0 => false.
	wakeupC  chan struct{} // Used to wakeup poll goroutine when item add to queue head.

	exitC chan struct{} // Used to make poll goroutine exit.
}

// Default return a DQueue with default parameters.
func Default() *DQueue {
	return New(defaultCapacity, defaultChanSize)
}

// New create a new DQueue with specified qCap(queue capacity) and chanSize(channel buffer capacity).
func New(qCap int, chanSize int) *DQueue {
	dq := &DQueue{
		C:        make(chan Value, chanSize),
		pq:       minheap.New(qCap),
		mu:       new(sync.Mutex),
		sleeping: 0,
		wakeupC:  make(chan struct{}),
		exitC:    make(chan struct{}),
	}

	go dq.polling()
	return dq
}

// Offer add a new value to queue with specified delay time.
func (dq *DQueue) Offer(delay time.Duration, value Value) {
	dq.mu.Lock()
	item := &Item{Expiration: time.Now().Add(delay).UnixNano(), Value: value}
	index := dq.pq.Push(item)
	dq.mu.Unlock()

	// A new item with the earliest expiration is added.
	if index == 0 && atomic.CompareAndSwapInt32(&dq.sleeping, 1, 0) {
		dq.wakeupC <- struct{}{}
	}
}

// Receive register a func to be called if some item expires.
func (dq *DQueue) Receive(f Receiver) {
	go func() {
		for {
			select {
			case <-dq.exitC:
				return
			case value := <-dq.C:
				f(value)
			}
		}
	}()
}

// Close to notify the polling exit. can't be called repeatedly.
func (dq *DQueue) Close() {
	dq.mu.Lock()
	close(dq.exitC)
	dq.mu.Unlock()
}

func (dq *DQueue) peekAndShift() (*Item, int64) {
	element := dq.pq.Peek()
	if element == nil {
		// queue is empty
		return nil, 0
	}

	item := element.(*Item)
	delay := item.Expiration - time.Now().UnixNano()
	if delay > 0 {
		return nil, delay
	}

	// Removed from queue top
	_ = dq.pq.Pop()
	return item, 0
}

func (dq *DQueue) polling() {
	defer func() {
		// Reset the sleeping states
		atomic.StoreInt32(&dq.sleeping, 0)
	}()

LOOP:
	for {
		dq.mu.Lock()
		item, delay := dq.peekAndShift()
		if item == nil {
			// No items left or at least one item is pending.

			// We must ensure the atomicity of the whole operation, which is
			// composed of the above PeekAndShift and the following StoreInt32,
			// to avoid possible race conditions between Offer and Poll.
			atomic.StoreInt32(&dq.sleeping, 1)
		}
		dq.mu.Unlock()

		// No items in queue. Waiting to be wakeup.
		if item == nil && delay == 0 {
			select {
			case <-dq.exitC:
				break LOOP
			case <-dq.wakeupC:
			}

			continue LOOP
		}

		// At least one item is pending. Go to sleep.
		if delay > 0 {
			select {
			case <-dq.exitC:
				break LOOP
			case <-dq.wakeupC:
				// A new item with an "earlier" expiration than the current "earliest" one is added.
			case <-time.After(time.Duration(delay)):
				// The current "earliest" item expires.

				// Reset the sleeping state since there's no need to receive from wakeupC.
				if atomic.SwapInt32(&dq.sleeping, 0) == 0 {
					// A caller of Offer() is being blocked on sending to wakeupC,
					// drain wakeupC to unblock the caller.
					<-dq.wakeupC
				}
			}

			continue LOOP
		}

		// Send expired element to channel.
		select {
		case <-dq.exitC:
			break LOOP
		case dq.C <- item.Value:
			// The expired element has been sent out successfully.
		}
	}
}
