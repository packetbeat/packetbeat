package provider

import (
	"context"
	"errors"
	"testing"

	e "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/x-pack/beatless/core"
)

type simpleFunction struct {
	err error
}

func (s *simpleFunction) Run(ctx context.Context, client core.Client) error {
	return s.err
}

func (s *simpleFunction) Name() string {
	return "simpleFunction"
}

type mockClient struct{}

func (sc *mockClient) Publish(event beat.Event) error       { return nil }
func (sc *mockClient) PublishAll(events []beat.Event) error { return nil }
func (sc *mockClient) Close() error                         { return nil }
func (sc *mockClient) Wait()                                {}

func TestRunnable(t *testing.T) {
	t.Run("return an error when we cannot create the client", func(t *testing.T) {
		err := errors.New("oops")
		runnable := Runnable{
			config:     common.NewConfig(),
			makeClient: func(cfg *common.Config) (core.Client, error) { return nil, err },
			function:   &simpleFunction{err: nil},
		}

		errReceived := runnable.Run(context.Background())
		assert.Equal(t, err, e.Cause(errReceived))
	})

	t.Run("propagate functions errors to the coordinator", func(t *testing.T) {
		err := errors.New("function error")
		runnable := Runnable{
			config:     common.NewConfig(),
			makeClient: func(cfg *common.Config) (core.Client, error) { return &mockClient{}, nil },
			function:   &simpleFunction{err: err},
		}

		errReceived := runnable.Run(context.Background())
		assert.Equal(t, err, e.Cause(errReceived))
	})

	t.Run("when there is no error run and exit normaly", func(t *testing.T) {
		runnable := Runnable{
			config:     common.NewConfig(),
			makeClient: func(cfg *common.Config) (core.Client, error) { return &mockClient{}, nil },
			function:   &simpleFunction{err: nil},
		}

		errReceived := runnable.Run(context.Background())
		assert.NoError(t, errReceived)
	})
}
