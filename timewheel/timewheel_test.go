package timewheel

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	size := int64(3)
	tw := New(time.Millisecond, size)
	require.NotNil(t, tw)
	require.Equal(t, tw.tick, int64(time.Millisecond))
	require.Equal(t, tw.size, size)
	require.Equal(t, tw.interval, int64(time.Millisecond)*size)
	require.Greater(t, tw.current, int64(0))
	require.Equal(t, len(tw.buckets), int(size))
	require.NotNil(t, tw.queue)
	require.True(t, tw.overflow == nil)
}

func TestTimeWheel_AfterFunc(t *testing.T) {
	tw := New(time.Millisecond, 3)
	go tw.Start()
	defer tw.Stop()

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

	for _, d := range seeds {
		t.Run(d.String(), func(t *testing.T) {
			retC := make(chan time.Time)

			start := time.Now()

			min := start
			max := start.Add(d + time.Millisecond*5)

			_ = tw.AfterFunc(d, func() { retC <- time.Now() })

			got := <-retC

			require.Greater(t, got.UnixNano(), min.UnixNano(), fmt.Sprintf("%s: got: %s, want: %s", d.String(), got.String(), min.String()))
			require.Less(t, got.UnixNano(), max.UnixNano(), fmt.Sprintf("%s: got: %s, want: %s", d.String(), got.String(), max.String()))
		})
	}
}

type Task struct {
	mu    *sync.Mutex
	seeds []time.Duration
	index int
	retC  chan time.Time
}

func (s *Task) Next(prev time.Time) time.Time {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.index >= len(s.seeds) {
		return time.Time{}
	}
	next := prev.Add(s.seeds[s.index])
	s.index += 1
	return next
}

func (s *Task) Run() {
	s.retC <- time.Now()
}

func TestTimeWheel_Schedule(t *testing.T) {
	tw := New(time.Millisecond, 20)
	go tw.Start()
	defer tw.Stop()

	seeds := []time.Duration{
		1 * time.Millisecond,   // start + 1ms
		4 * time.Millisecond,   // start + 5ms
		5 * time.Millisecond,   // start + 10ms
		40 * time.Millisecond,  // start + 50ms
		50 * time.Millisecond,  // start + 100ms
		400 * time.Millisecond, // start + 400ms
		500 * time.Millisecond, // start + 500ms
		501 * time.Millisecond, // start + 501ms
	}

	retC := make(chan time.Time)

	s := &Task{
		mu:    new(sync.Mutex),
		seeds: seeds,
		retC:  retC,
	}

	lapse := time.Duration(0)
	start := time.Now()

	_ = tw.Schedule(s)

	for _, d := range seeds {
		lapse += d
		min := start
		max := start.Add(lapse + time.Millisecond*5)

		got := <-retC

		require.Greater(t, got.UnixNano(), min.UnixNano(), fmt.Sprintf("%s: got: %s, want: %s", d.String(), got.String(), min.String()))
		require.Less(t, got.UnixNano(), max.UnixNano(), fmt.Sprintf("%s: got: %s, want: %s", d.String(), got.String(), max.String()))
	}
}
