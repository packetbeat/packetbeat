package elb

import (
	"sync"
	"testing"
	"time"

	"github.com/elastic/beats/libbeat/common"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/libbeat/common/bus"
)

type testEventAccumulator struct {
	events []bus.Event
	lock   sync.Mutex
}

func (tea *testEventAccumulator) add(e bus.Event) {
	tea.lock.Lock()
	defer tea.lock.Unlock()

	tea.events = append(tea.events, e)
}

func (tea *testEventAccumulator) len() int {
	tea.lock.Lock()
	defer tea.lock.Unlock()

	return len(tea.events)
}

func (tea *testEventAccumulator) get() []bus.Event {
	tea.lock.Lock()
	defer tea.lock.Unlock()

	res := make([]bus.Event, len(tea.events))
	copy(res, tea.events)
	return res
}

func (tea *testEventAccumulator) waitForNumEvents(t *testing.T, targetLen int, timeout time.Duration) {
	start := time.Now()

	for time.Now().Sub(start) < timeout {
		if tea.len() >= targetLen {
			return
		}
		time.Sleep(time.Millisecond)
	}

	t.Fatalf("Timed out waiting for num events to be %d", targetLen)
}

func Test_internalBuilder(t *testing.T) {
	lbl := fakeLbl()
	lbls := []*lbListener{lbl}
	fetcher := newMockFetcher(lbls, nil)
	pBus := bus.New("test")

	cfg := &Config{
		Region: "us-east-1",
	}

	provider, err := internalBuilder(pBus, cfg, fetcher)
	require.NoError(t, err)

	provider.Start()

	startListener := pBus.Subscribe("start")
	stopListener := pBus.Subscribe("stop")
	listenerDone := make(chan struct{})
	defer close(listenerDone)

	var events testEventAccumulator
	go func() {
		for {
			select {
			case e := <-startListener.Events():
				events.add(e)
			case e := <-stopListener.Events():
				events.add(e)
			case <-listenerDone:
				return
			}
		}
	}()

	// Run twice to ensure that duplicates don't create two start events
	provider.watcher.once()
	provider.watcher.once()
	events.waitForNumEvents(t, 1, time.Second)

	assert.Equal(t, 1, events.len())

	expectedStartEvent := bus.Event{
		"start":   true,
		"hashKey": lbl.uuid(),
		"host":    *lbl.lb.DNSName,
		"port":    *lbl.listener.Port,
		"meta":    common.MapStr{"elb": lbl.toMap()},
	}

	require.Equal(t, expectedStartEvent, events.get()[0])

	fetcher.setLbls([]*lbListener{})

	// Run twice to ensure that duplicates don't create two stop events
	provider.watcher.once()
	provider.watcher.once()
	events.waitForNumEvents(t, 2, time.Second)

	require.Equal(t, 2, events.len())

	expectedStopEvent := bus.Event{
		"stop":    true,
		"hashKey": lbl.uuid(),
	}

	require.Equal(t, expectedStopEvent, events.get()[1])

}
