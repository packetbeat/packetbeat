package aws

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/x-pack/beatless/provider"
)

type arrayBackedClient struct {
	Events []beat.Event
	err    error
}

func (a *arrayBackedClient) Publish(event beat.Event) error {
	if a.err != nil {
		return a.err
	}
	a.Events = append(a.Events, event)
	return nil
}

func (a *arrayBackedClient) PublishAll(events []beat.Event) error {
	if a.err != nil {
		return a.err
	}
	a.Events = append(a.Events, events...)
	return nil
}

func (a *arrayBackedClient) Wait()        { return }
func (a *arrayBackedClient) Close() error { return nil }

func TestCloudwatchLogs(t *testing.T) {
	cfg := common.MustNewConfigFrom(map[string]interface{}{
		"name":        "ph",
		"description": "my long description",
		"role":        "arn:aws:iam::00000000:role/beatless",
		"triggers": []map[string]interface{}{
			map[string]interface{}{
				"log_group_name": "foo",
				"filter_name":    "bar",
			},
		},
	})

	t.Run("when publish is succesful", func(t *testing.T) {
		client := &arrayBackedClient{}
		cwl, err := NewCloudwatchLogs(&provider.DefaultProvider{}, cfg)
		if !assert.NoError(t, err) {
			return
		}

		c, _ := cwl.(*CloudwatchLogs)
		handler := c.createHandler(client)

		err = handler(generateCloudwatchLogRawEvent())

		assert.NoError(t, err)
	})

	t.Run("when publish is not succesful", func(t *testing.T) {
		e := errors.New("something bad")
		client := &arrayBackedClient{err: e}
		cwl, err := NewCloudwatchLogs(&provider.DefaultProvider{}, cfg)
		if !assert.NoError(t, err) {
			return
		}

		c, _ := cwl.(*CloudwatchLogs)
		handler := c.createHandler(client)

		err = handler(generateCloudwatchLogRawEvent())

		assert.Equal(t, e, err)
	})

}

func generateCloudwatchLogRawEvent() events.CloudwatchLogsEvent {
	rawEvent := events.CloudwatchLogsData{
		Owner:     "foobar",
		LogGroup:  "foo",
		LogStream: "/var/foobar",
		LogEvents: []events.CloudwatchLogsLogEvent{
			events.CloudwatchLogsLogEvent{
				ID:        "1234",
				Timestamp: time.Now().Unix(),
				Message:   "hello world",
			},
		},
	}

	b, _ := json.Marshal(&rawEvent)

	data := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, data)
	zw := gzip.NewWriter(encoder)
	zw.Write(b)
	zw.Close()
	encoder.Close()

	return events.CloudwatchLogsEvent{
		AWSLogs: events.CloudwatchLogsRawData{
			Data: data.String(),
		},
	}
}
