// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package journalctl

import (
	"sync"

	input "github.com/elastic/beats/v7/filebeat/input/v2"
)

// Ensure, that JctlMock does implement Jctl.
// If this is not the case, regenerate this file with moq.
var _ Jctl = &JctlMock{}

// JctlMock is a mock implementation of Jctl.
//
//	func TestSomethingThatUsesJctl(t *testing.T) {
//
//		// make and configure a mocked Jctl
//		mockedJctl := &JctlMock{
//			KillFunc: func() error {
//				panic("mock out the Kill method")
//			},
//			NextFunc: func(canceler input.Canceler) ([]byte, error) {
//				panic("mock out the Next method")
//			},
//		}
//
//		// use mockedJctl in code that requires Jctl
//		// and then make assertions.
//
//	}
type JctlMock struct {
	// KillFunc mocks the Kill method.
	KillFunc func() error

	// NextFunc mocks the Next method.
	NextFunc func(canceler input.Canceler) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// Kill holds details about calls to the Kill method.
		Kill []struct {
		}
		// Next holds details about calls to the Next method.
		Next []struct {
			// Canceler is the canceler argument value.
			Canceler input.Canceler
		}
	}
	lockKill sync.RWMutex
	lockNext sync.RWMutex
}

// Kill calls KillFunc.
func (mock *JctlMock) Kill() error {
	if mock.KillFunc == nil {
		panic("JctlMock.KillFunc: method is nil but Jctl.Kill was just called")
	}
	callInfo := struct {
	}{}
	mock.lockKill.Lock()
	mock.calls.Kill = append(mock.calls.Kill, callInfo)
	mock.lockKill.Unlock()
	return mock.KillFunc()
}

// KillCalls gets all the calls that were made to Kill.
// Check the length with:
//
//	len(mockedJctl.KillCalls())
func (mock *JctlMock) KillCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockKill.RLock()
	calls = mock.calls.Kill
	mock.lockKill.RUnlock()
	return calls
}

// Next calls NextFunc.
func (mock *JctlMock) Next(canceler input.Canceler) ([]byte, error) {
	if mock.NextFunc == nil {
		panic("JctlMock.NextFunc: method is nil but Jctl.Next was just called")
	}
	callInfo := struct {
		Canceler input.Canceler
	}{
		Canceler: canceler,
	}
	mock.lockNext.Lock()
	mock.calls.Next = append(mock.calls.Next, callInfo)
	mock.lockNext.Unlock()
	return mock.NextFunc(canceler)
}

// NextCalls gets all the calls that were made to Next.
// Check the length with:
//
//	len(mockedJctl.NextCalls())
func (mock *JctlMock) NextCalls() []struct {
	Canceler input.Canceler
} {
	var calls []struct {
		Canceler input.Canceler
	}
	mock.lockNext.RLock()
	calls = mock.calls.Next
	mock.lockNext.RUnlock()
	return calls
}
