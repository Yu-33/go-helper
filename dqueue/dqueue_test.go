package dqueue

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	capacity := 8
	dq := newDQueue(capacity)
	require.NotNil(t, dq)
	defer dq.Close()

	require.NotNil(t, dq.C)
	require.NotNil(t, dq.mu)
	require.NotNil(t, dq.pq)
	require.Equal(t, dq.pq.Cap(), capacity)
	require.Equal(t, atomic.LoadInt32(&dq.sleeping), int32(0))
	require.NotNil(t, dq.wakeupC)
	require.NotNil(t, dq.exitC)
	require.NotNil(t, dq.wg)

	go dq.polling()
	time.Sleep(time.Millisecond * 10)

	require.Equal(t, atomic.LoadInt32(&dq.sleeping), int32(1))
}

func TestDefault(t *testing.T) {
	dq := Default()
	require.NotNil(t, dq)
	require.Equal(t, dq.pq.Cap(), defaultCapacity)
}

type Result struct {
	T time.Time
	V time.Duration
}

func TestDQueue_After(t *testing.T) {
	dq := Default()
	defer dq.Close()

	retC := make(chan *Result)

	go func() {
		dq.Receive(func(value Value) {
			retC <- &Result{
				T: time.Now(),
				V: value.(time.Duration),
			}
		})
	}()

	seeds := []time.Duration{
		time.Millisecond * 1,
		time.Millisecond * 5,
		time.Millisecond * 10,
		time.Millisecond * 50,
		time.Millisecond * 100,
		time.Millisecond * 400,
		time.Millisecond * 500,
		time.Second * 1,
	}

	lapse := time.Duration(0)
	start := time.Now()

	for _, d := range seeds {
		dq.After(d, d)
	}

	for _, d := range seeds {
		lapse += d
		min := start.Add(d)
		max := start.Add(lapse + time.Millisecond*5)

		got := <-retC

		require.Equal(t, d, got.V)
		require.Greater(t, got.T.UnixNano(), min.UnixNano(), fmt.Sprintf("%s: got: %s, want: %s", d.String(), got.T.String(), min.String()))
		require.Less(t, got.T.UnixNano(), max.UnixNano(), fmt.Sprintf("%s: got: %s, want: %s", d.String(), got.T.String(), max.String()))
	}
}

func TestDQueue_Expire(t *testing.T) {
	dq := Default()
	defer dq.Close()

	retC := make(chan *Result)

	go func() {
		dq.Receive(func(value Value) {
			retC <- &Result{
				T: time.Now(),
				V: value.(time.Duration),
			}
		})
	}()

	seeds := []time.Duration{
		time.Millisecond * 1,
		time.Millisecond * 5,
		time.Millisecond * 10,
		time.Millisecond * 50,
		time.Millisecond * 100,
		time.Millisecond * 400,
		time.Millisecond * 500,
		time.Second * 1,
	}

	lapse := time.Duration(0)
	start := time.Now()

	for _, d := range seeds {
		dq.Expire(time.Now().Add(d).UnixNano(), d)
	}

	for _, d := range seeds {
		lapse += d
		min := start.Add(d)
		max := start.Add(lapse + time.Millisecond*5)

		got := <-retC

		require.Equal(t, d, got.V)
		require.Greater(t, got.T.UnixNano(), min.UnixNano(), fmt.Sprintf("%s: got: %s, want: %s", d.String(), got.T.String(), min.String()))
		require.Less(t, got.T.UnixNano(), max.UnixNano(), fmt.Sprintf("%s: got: %s, want: %s", d.String(), got.T.String(), max.String()))
	}
}
