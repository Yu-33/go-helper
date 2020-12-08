package dqueue

import (
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
	require.Equal(t, dq.sleeping, int32(0))
	require.NotNil(t, dq.wakeupC)
	require.NotNil(t, dq.exitC)
	require.NotNil(t, dq.wg)

	go dq.polling()
	time.Sleep(time.Millisecond * 10)

	require.Equal(t, dq.sleeping, int32(1))
}

func TestDefault(t *testing.T) {
	dq := Default()
	require.NotNil(t, dq)
	require.Equal(t, dq.pq.Cap(), defaultCapacity)
}

func TestDQueue_After(t *testing.T) {
	dq := Default()
	defer dq.Close()

	var x int
	exitC := make(chan struct{})

	go func() {
		dq.Receive(func(value Value) {
			x = value.(int)
			close(exitC)
		})
	}()

	start := time.Now()
	dq.After(time.Second, 1024)

	<-exitC

	require.Equal(t, x, 1024)
	require.Equal(t, int(time.Since(start).Seconds()), 1)
}

func TestDQueue_Expire(t *testing.T) {
	dq := Default()
	defer dq.Close()

	var x int
	exitC := make(chan struct{})

	go func() {
		dq.Receive(func(value Value) {
			x = value.(int)
			close(exitC)
		})
	}()

	start := time.Now()
	dq.Expire(time.Now().Add(time.Second).UnixNano(), 1024)

	<-exitC

	require.Equal(t, x, 1024)
	require.Equal(t, int(time.Since(start).Seconds()), 1)
}
