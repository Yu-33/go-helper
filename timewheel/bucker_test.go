package timewheel

import (
	"fmt"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func Test_bucket_new(t *testing.T) {
	b := newBucket()
	require.NotNil(t, b)
	require.Equal(t, b.expiration, int64(-1))
	require.NotNil(t, b.flushMu)
	require.NotNil(t, b.mu)
}

func Test_bucket_insert(t *testing.T) {
	b := newBucket()

	n := 129
	for i := 0; i < n; i++ {
		b.insert(&Timer{})
	}

	timers := b.timers
	require.Equal(t, timers.Len(), n)

	for e := timers.Front(); e != nil; e = e.Next() {
		timer := e.Value.(*Timer)

		require.Equal(t, timer.element, e)
		require.Equal(t, timer.getBucket(), b)
	}
}

func Test_bucket_flush1(t *testing.T) {
	b := newBucket()

	b.insert(&Timer{})
	b.insert(&Timer{})
	p1 := unsafe.Pointer(b.timers)

	require.Equal(t, b.timers.Len(), 2)

	b.flush(func(*Timer) {})
	p2 := unsafe.Pointer(b.timers)

	require.Equal(t, b.timers.Len(), 0)
	require.NotEqual(t, p1, p2)
}

func Test_bucket_flush2(t *testing.T) {
	b := newBucket()

	n := 17
	for i := 0; i < n; i++ {
		b.insert(&Timer{})
	}

	require.Equal(t, b.timers.Len(), n)

	b.flush(b.insert)

	require.Equal(t, b.timers.Len(), n)
}

func Test_bucket_flush3(t *testing.T) {
	b := newBucket()

	cases := []struct {
		name string
		N    int // the data size (i.e. number of existing timers)
	}{
		{"N-1w", 10000},
		{"N-10w", 100000},
		{"N-100w", 1000000},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for i := 0; i < c.N; i++ {
				b.insert(&Timer{})
			}
			start := time.Now()
			b.flush(func(timer *Timer) {
				_ = timer
			})
			elapse := time.Since(start)

			fmt.Printf("flush %s use: %s\n", c.name, elapse.String())
		})
	}
}
