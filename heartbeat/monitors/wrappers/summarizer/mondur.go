package summarizer

import (
	"time"

	"github.com/elastic/beats/v7/heartbeat/eventext"
	"github.com/elastic/beats/v7/heartbeat/look"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

type LightweightDurationSumPlugin struct {
	startedAt *time.Time
}

func (lwdsp *LightweightDurationSumPlugin) EachEvent(event *beat.Event) {
	// Effectively only runs once, on the first event
	if lwdsp.startedAt == nil {
		now := time.Now()
		lwdsp.startedAt = &now
	}
}

func (lwdsp *LightweightDurationSumPlugin) OnSummary(event *beat.Event) bool {
	eventext.MergeEventFields(event, mapstr.M{"monitor.duration": look.RTT(time.Since(*lwdsp.startedAt))})
	return false
}

type BrowserDurationSumPlugin struct {
	startedAt *time.Time
	endedAt   *time.Time
}

func (bwdsp *BrowserDurationSumPlugin) EachEvent(event *beat.Event) {
	// Effectively only runs once, on the first event
	et, _ := event.GetValue("synthetics.type")
	if et == "journey/start" {
		bwdsp.startedAt = &event.Timestamp
	} else if et == "journey/end" {
		bwdsp.endedAt = &event.Timestamp
	}
}

func (bwdsp *BrowserDurationSumPlugin) OnSummary(event *beat.Event) bool {
	if bwdsp.startedAt == nil {
		now := time.Now()
		bwdsp.startedAt = &now
	}
	if bwdsp.endedAt == nil {
		now := time.Now()
		bwdsp.endedAt = &now
	}
	eventext.MergeEventFields(event, mapstr.M{"monitor.duration": look.RTT(bwdsp.endedAt.Sub(*bwdsp.startedAt))})

	return false
}
