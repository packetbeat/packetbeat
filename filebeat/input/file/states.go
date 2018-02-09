package file

import (
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/logp"
)

// States handles list of FileState. One must use NewStates to instantiate a
// file states regisry. Using the zero-value is not safe.
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
	s.UpdateWithTs(newState, time.Now())
}

// UpdateWithTs updates a state, assigning the given timestamp.
// If previous state didn't exist, new one is created
func (s *States) UpdateWithTs(newState State, ts time.Time) {
	s.Lock()
	defer s.Unlock()

	id := newState.ID()
	index := s.findPrevious(id)
	newState.Timestamp = ts

	if index >= 0 {
		s.states[index] = newState
	} else {
		// No existing state found, add new one
		s.idx[id] = len(s.states)
		s.states = append(s.states, newState)
		logp.Debug("input", "New state added for %s", newState.Source)
	}
}

func (s *States) FindPrevious(newState State) State {
	s.RLock()
	defer s.RUnlock()
	i := s.findPrevious(newState.ID())
	if i < 0 {
		return State{}
	}
	return s.states[i]
}

// findPreviousState returns the previous state fo the file
// In case no previous state exists, index -1 is returned
func (s *States) findPrevious(id string) int {
	if i, exists := s.idx[id]; exists {
		return i
	}
	return -1
}

// Cleanup cleans up the state array. All states which are older then `older` are removed
// The number of states that were cleaned up and number of states that can be
// cleaned up in the future is returned.
func (s *States) Cleanup() (int, int) {
	s.Lock()
	defer s.Unlock()

	statesBefore := len(s.states)
	numCanExpire := 0

	currentTime := time.Now()
	states := s.states[:0]

	for i := range s.states {
		state := &s.states[i]
		expired := (state.TTL > 0 && currentTime.Sub(state.Timestamp) > state.TTL)

		if state.TTL == 0 || expired {
			if state.Finished {
				logp.Debug("state", "State removed for %v because of older: %v", state.Source, state.TTL)
				delete(s.idx, state.ID())
				continue // drop state
			} else {
				logp.Err("State for %s should have been dropped, but couldn't as state is not finished.", state.Source)
			}
		}

		if state.TTL >= 0 {
			numCanExpire++
		}
		states = append(states, *state) // in-place copy old state
	}
	s.states = states

	return statesBefore - len(s.states), numCanExpire
}

// Count returns number of states
func (s *States) Count() int {
	s.RLock()
	defer s.RUnlock()

	return len(s.states)
}

// Returns a copy of the file states
func (s *States) GetStates() []State {
	s.RLock()
	defer s.RUnlock()

	newStates := make([]State, len(s.states))
	copy(newStates, s.states)

	return newStates
}

// SetStates overwrites all internal states with the given states array
func (s *States) SetStates(states []State) {
	s.Lock()
	defer s.Unlock()
	s.states = states

	// create new index
	s.idx = map[string]int{}
	for i := range states {
		s.idx[states[i].ID()] = i
	}
}

// Copy create a new copy of the states object
func (s *States) Copy() *States {
	new := NewStates()
	new.SetStates(s.GetStates())
	return new
}
