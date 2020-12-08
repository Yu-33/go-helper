package timewheel

import (
	"container/list"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBucket_Flush(t *testing.T) {
	b := &bucket{
		expiration: -1,
		mu:         new(sync.Mutex),
		timers:     list.New(),
	}

	b.insert(&Timer{})
	b.insert(&Timer{})
	l1 := b.timers.Len()
	require.Equal(t, l1, 2)

	b.flush(func(*Timer) {})
	l2 := b.timers.Len()
	require.Equal(t, l2, 0)
}
