// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package awss3

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	v2 "github.com/elastic/beats/v7/filebeat/input/v2"
	"github.com/elastic/beats/v7/libbeat/logp"
)

var inputCtx = v2.Context{
	Logger:      logp.NewLogger("test"),
	Cancelation: context.Background(),
}

func TestStatesDelete(t *testing.T) {
	type stateTestCase struct {
		states   func() *states
		deleteID string
		expected []state
	}

	lastModified := time.Date(2021, time.July, 22, 18, 38, 00, 0, time.UTC)
	statesFuncSingle := func() *states {
		states := newStates(inputCtx)
		states.Update(newState("bucket", "key", "etag", "listPrefix", lastModified), "")
		return states
	}

	statesFuncMultiple := func() *states {
		states := newStates(inputCtx)
		states.Update(newState("bucket", "key1", "etag1", "listPrefix", lastModified), "")
		states.Update(newState("bucket", "key2", "etag2", "listPrefix", lastModified), "")
		states.Update(newState("bucket", "key3", "etag3", "listPrefix", lastModified), "")
		return states
	}

	stateFirst := state{
		ID:           "bucketkey1etag1" + lastModified.String(),
		Bucket:       "bucket",
		Key:          "key1",
		Etag:         "etag1",
		ListPrefix:   "listPrefix",
		LastModified: lastModified,
	}

	stateSecond := state{
		ID:           "bucketkey2etag2" + lastModified.String(),
		Bucket:       "bucket",
		Key:          "key2",
		Etag:         "etag2",
		ListPrefix:   "listPrefix",
		LastModified: lastModified,
	}

	stateThird := state{
		ID:           "bucketkey3etag3" + lastModified.String(),
		Bucket:       "bucket",
		Key:          "key3",
		Etag:         "etag3",
		ListPrefix:   "listPrefix",
		LastModified: lastModified,
	}

	tests := map[string]stateTestCase{
		"delete empty states": {
			states: func() *states {
				return newStates(inputCtx)
			},
			deleteID: "an id",
			expected: []state{},
		},
		"delete not existing state": {
			states:   statesFuncSingle,
			deleteID: "an id",
			expected: []state{
				{
					ID:           "bucketkeyetag" + lastModified.String(),
					Bucket:       "bucket",
					Key:          "key",
					Etag:         "etag",
					ListPrefix:   "listPrefix",
					LastModified: lastModified,
				},
			},
		},
		"delete only one existing": {
			states:   statesFuncSingle,
			deleteID: "bucketkey",
			expected: []state{},
		},
		"delete first": {
			states:   statesFuncMultiple,
			deleteID: "bucketkey1",
			expected: []state{stateThird, stateSecond},
		},
		"delete last": {
			states:   statesFuncMultiple,
			deleteID: "bucketkey3",
			expected: []state{stateFirst, stateSecond},
		},
		"delete any": {
			states:   statesFuncMultiple,
			deleteID: "bucketkey2",
			expected: []state{stateFirst, stateThird},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			states := test.states()
			states.Delete(test.deleteID)
			assert.Equal(t, test.expected, states.GetStates())
		})
	}
}
