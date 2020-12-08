package timewheel

import (
	"time"
)

// Scheduler determines the execution plan of a task.
type Scheduler interface {
	// Next returns the next execution time after the given (previous) time.
	// It will return a zero time if no next time is scheduled.
	Next(prev time.Time) (next time.Time)
	// Run will be called when scheduler expired.
	Run()
}

// Schedule calls f (in its own goroutine) according to the execution
// plan scheduled by s. It returns a Timer that can be used to cancel the
// call using its Stop method.
//
// If the caller want to terminate the execution plan halfway, it must
// stop the timer and ensure that the timer is stopped actually, since in
// the current implementation, there is a gap between the expiring and the
// restarting of the timer. The wait time for ensuring is short since the
// gap is very small.
//
// Internally, ScheduleFunc will ask the first execution time (by calling
// s.Next()) initially, and create a timer if the execution time is non-zero.
// Afterwards, it will ask the next execution time each time f is about to
// be executed, and f will be called at the next execution time if the time
// is non-zero.
func (tw *TimeWheel) Schedule(sh Scheduler) (t *Timer) {
	next := sh.Next(time.Now())
	if next.IsZero() {
		// No time is scheduled, return nil.
		return
	}

	t = &Timer{
		expiration: next.UnixNano(),
		task: func() {
			// Schedule the task to execute at the next time if possible.
			next := sh.Next(time.Unix(0, t.expiration))
			if !next.IsZero() {
				// Resubmit the timer to next cycle.
				t.expiration = next.UnixNano()
				tw.submit(t)
			}

			// Actually execute the task.
			sh.Run()
		},
	}

	tw.submit(t)
	return
}

// AfterFunc waits for the duration to elapse and then calls f in its own goroutine.
// It returns a Timer that can be used to cancel the call using its Stop method.
func (tw *TimeWheel) AfterFunc(d time.Duration, f func()) *Timer {
	t := &Timer{
		expiration: time.Now().Add(d).UnixNano(),
		task:       f,
	}
	tw.submit(t)
	return t
}
