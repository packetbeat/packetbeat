// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package azureeventhub

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Azure/azure-event-hubs-go/v3"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/filebeat/channel"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
)

var (
	config = azureInputConfig{
		SAKey:            "",
		SAName:           "",
		SAContainer:      ephContainerName,
		ConnectionString: "",
		ConsumerGroup:    "",
	}
)

func TestProcessEvents(t *testing.T) {
	// Stub outlet for receiving events generated by the input.
	o := &stubOutleter{}
	out, err := newStubOutlet(o)
	if err != nil {
		t.Fatal(err)
	}
	input := azureInput{
		config: config,
		outlet: out,
	}
	var sn int64 = 12
	now := time.Now()
	var off int64 = 1234
	var pID int16 = 1

	properties := eventhub.SystemProperties{
		SequenceNumber: &sn,
		EnqueuedTime:   &now,
		Offset:         &off,
		PartitionID:    &pID,
		PartitionKey:   nil,
	}
	single := "{\"test\":\"this is some message\",\"time\":\"2019-12-17T13:43:44.4946995Z\"}"
	msg := fmt.Sprintf("{\"records\":[%s]}", single)
	ev := eventhub.Event{
		Data:             []byte(msg),
		SystemProperties: &properties,
	}
	err = input.processEvents(&ev, "0")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(o.Events), 1)
	parse, err := time.Parse(time.RFC3339, "2019-12-17T13:43:44.4946995Z")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, o.Events[0].Timestamp, parse)
	message, err := o.Events[0].Fields.GetValue("message")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, message, single)
}

func TestParseMultipleMessages(t *testing.T) {
	date1 := "2019-12-17T13:43:44.4946995Z"
	date2 := "2019-12-17T14:43:44.4946995Z"
	date3 := "2019-12-17T15:43:44.4946995Z"
	msg := fmt.Sprintf("{\"records\":[{\"test\":\"this is some message\",\"time\":\"%s\"},"+
		"{\"test\":\"this is 2nd message\",\"time\":\"%s\"},"+
		"{\"test\":\"this is 3rd message\",\"time\":\"%s\"}]}", date1, date2, date3)
	input := azureInput{}
	messages := input.parseMultipleMessages([]byte(msg))
	assert.NotNil(t, messages)

	assert.Equal(t, len(messages), 3)

	pdate1, err := time.Parse(time.RFC3339, date1)
	if err != nil {
		t.Fatal(err)
	}
	pdate2, err := time.Parse(time.RFC3339, date2)
	if err != nil {
		t.Fatal(err)
	}
	pdate3, err := time.Parse(time.RFC3339, date3)
	if err != nil {
		t.Fatal(err)
	}
	dates := []time.Time{pdate1, pdate2, pdate3}
	msgs := []string{
		fmt.Sprintf("{\"test\":\"this is some message\",\"time\":\"%s\"}", date1),
		fmt.Sprintf("{\"test\":\"this is 2nd message\",\"time\":\"%s\"}", date2),
		fmt.Sprintf("{\"test\":\"this is 3rd message\",\"time\":\"%s\"}", date3)}
	for _, ms := range messages {
		for key, val := range ms {
			assert.Contains(t, dates, key)
			assert.Contains(t, msgs, val)
		}
	}
}

type stubOutleter struct {
	sync.Mutex
	cond   *sync.Cond
	done   bool
	Events []beat.Event
}

func newStubOutlet(stub *stubOutleter) (channel.Outleter, error) {
	stub.cond = sync.NewCond(stub)
	defer stub.Close()

	connector := channel.ConnectorFunc(func(_ *common.Config, _ beat.ClientConfig) (channel.Outleter, error) {
		return stub, nil
	})
	return connector.ConnectWith(nil, beat.ClientConfig{
		Processing: beat.ProcessingConfig{},
	})
}
func (o *stubOutleter) Close() error {
	o.Lock()
	defer o.Unlock()
	o.done = true
	return nil
}
func (o *stubOutleter) Done() <-chan struct{} { return nil }
func (o *stubOutleter) OnEvent(event beat.Event) bool {
	o.Lock()
	defer o.Unlock()
	o.Events = append(o.Events, event)
	o.cond.Broadcast()
	return o.done
}
