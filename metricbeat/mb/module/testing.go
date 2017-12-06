package module

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/testing"
)

// receiveOneEvent receives one event from the events channel then closes the
// returned done channel. If no events are received it will close the returned
// done channel after the timeout period elapses.
func receiveOneEvent(d testing.Driver, events <-chan beat.Event, timeout time.Duration) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		select {
		case <-time.Tick(timeout):
			d.Error("error", errors.New("timeout waiting for an event"))
		case event, ok := <-events:
			if !ok {
				return
			}
			outputJSON(d, &event)
		}
	}()

	return done
}

func outputJSON(d testing.Driver, event *beat.Event) {
	out := event.Fields.Clone()
	out.Put("@timestamp", common.Time(event.Timestamp))
	jsonData, err := json.MarshalIndent(out, "", " ")
	if err != nil {
		d.Error("convert error", err)
		return
	}

	d.Result(string(jsonData))
}
