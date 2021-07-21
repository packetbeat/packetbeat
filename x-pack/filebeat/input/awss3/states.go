// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package awss3

import (
	"strings"
	"sync"

	"github.com/elastic/beats/v7/libbeat/statestore"
	"github.com/elastic/beats/v7/libbeat/statestore/backend"

	"github.com/elastic/beats/v7/libbeat/logp"
)

// States handles list of s3 object state. One must use NewStates to instantiate a
// file states registry. Using the zero-value is not safe.
type States struct {
	sync.RWMutex

	// states store
	states []State

	// idx maps state IDs to state indexes for fast lookup and modifications.
	idx map[string]int
}

// NewStates generates a new states registry.
func NewStates() *States {
	return &States{
		states: nil,
		idx:    map[string]int{},
	}
}

// Update updates a state. If previous state didn't exist, new one is created
func (s *States) Update(newState State) {
	s.Lock()
	defer s.Unlock()

	id := newState.Bucket + newState.Key
	index := s.findPrevious(id)

	if index >= 0 {
		s.states[index] = newState
	} else {
		// No existing state found, add new one
		s.idx[id] = len(s.states)
		s.states = append(s.states, newState)
		logp.Debug("input", "New state added for %s", newState.Key)
	}
}

// FindPrevious lookups a registered state, that matching the new state.
// Returns a zero-state if no match is found.
func (s *States) FindPrevious(newState State) State {
	s.RLock()
	defer s.RUnlock()
	id := newState.Bucket + newState.Key
	i := s.findPrevious(id)
	if i < 0 {
		return State{}
	}
	return s.states[i]
}

func (s *States) IsNew(state State) bool {
	s.RLock()
	defer s.RUnlock()
	id := state.Bucket + state.Key
	i := s.findPrevious(id)

	if i < 0 {
		return true
	}

	return !s.states[i].IsEqual(&state)
}

// findPrevious returns the previous state for the file.
// In case no previous state exists, index -1 is returned
func (s *States) findPrevious(id string) int {
	if i, exists := s.idx[id]; exists {
		return i
	}
	return -1
}

// GetStates creates copy of the file states.
func (s *States) GetStates() []State {
	s.RLock()
	defer s.RUnlock()

	newStates := make([]State, len(s.states))
	copy(newStates, s.states)

	return newStates
}

func (s *States) readStatesFrom(store *statestore.Store) error {
	var states []State

	err := store.Each(func(key string, dec statestore.ValueDecoder) (bool, error) {
		if !strings.HasPrefix(key, awsS3StatePrefix) {
			return true, nil
		}

		// try to decode. Ingore faulty/incompatible values.
		var st State
		if err := dec.Decode(&st); err != nil {
			// XXX: Do we want to log here? In case we start to store other
			// state types in the registry, then this operation will likely fail
			// quite often, producing some false-positives in the logs...
			return true, nil
		}

		st.Id = key[len(awsS3StatePrefix):]
		states = append(states, st)
		return true, nil
	})

	if err != nil {
		return err
	}

	states = fixStates(states)

	for _, state := range states {
		s.Update(state)
	}

	return nil
}

// fixStates cleans up the registry states when updating from an older version
// of filebeat potentially writing invalid entries.
func fixStates(states []State) []State {
	if len(states) == 0 {
		return states
	}

	// we use a map of states here, so to identify and merge duplicate entries.
	idx := map[string]*State{}
	for i := range states {
		state := &states[i]

		old, exists := idx[state.Id]
		if !exists {
			idx[state.Id] = state
		} else {
			mergeStates(old, state) // overwrite the entry in 'old'
		}
	}

	if len(idx) == len(states) {
		return states
	}

	i := 0
	newStates := make([]State, len(idx))
	for _, state := range idx {
		newStates[i] = *state
		i++
	}
	return newStates
}

// mergeStates merges 2 states by trying to determine the 'newer' state.
// The st state is overwritten with the updated fields.
func mergeStates(st, other *State) {
	// update file meta-data. As these are updated concurrently by the
	// inputs, select the newer state based on the update timestamp.
	if st.LastModified.Before(other.LastModified) {
		st.Size = other.Size
		st.LastModified = other.LastModified
	}
}

func writeStates(store backend.Store, states []State) error {
	for i := range states {
		key := awsS3StatePrefix + states[i].Id
		if err := store.Set(key, states[i]); err != nil {
			return err
		}
	}
	return nil
}
