package throttler

import (
	"sync"
	"testing"
	"time"

	"github.com/elastic/beats/libbeat/common/atomic"
	"github.com/stretchr/testify/require"
)

func TestThrottling(t *testing.T) {
	throttler := NewThrottler(5)

	// Use a wait group to block the first 5 jobs, thus occupying the throttler.
	stopProcesses := sync.WaitGroup{}
	stopProcesses.Add(1)

	// Wait group that will complete when all release statements are processed
	releaseWg := sync.WaitGroup{}

	// We should be able to acquire slots without blocking before
	// starting the throttler
	acquiredCount := atomic.NewUint(0)
	for i := 0; i < 10; i++ {
		releaseWg.Add(1)
		go func() {
			acquired, release := throttler.AcquireSlot()
			require.True(t, acquired)
			acquiredCount.Inc()
			stopProcesses.Wait()
			release()
			releaseWg.Done()
		}()
	}

	throttler.Start()

	// Now that we've started the throttler 5 jobs should run and be stuck due to the wait group
	// We wait 2 seconds to ensure that it's run as much as it's ever going to run
	time.Sleep(time.Second * 2)
	// Test that our throttler hasn't let any more jobs through
	require.Equal(t, acquiredCount.Load(), uint(5))
	//Now unblock the jobs so the remaining ones can finish
	stopProcesses.Done()

	started := time.Now()
	elapsed := time.Duration(0)
	// Wait until all 10 jobs have run
	for acquiredCount.Load() < 10 && elapsed < time.Second*2 {
		time.Sleep(time.Millisecond)
		elapsed = time.Now().Sub(started)
	}
	require.Equal(t, acquiredCount.Load(), uint(10))

	// Now we can stop the throttler
	throttler.Stop()

	// Acquiring after the throttler is stopped should not work
	// Let's wait till we're done releasing for this
	releaseWg.Wait()
	acquired, _ := throttler.AcquireSlot()
	require.False(t, acquired)
}
