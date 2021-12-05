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

package scheduler

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"golang.org/x/sync/semaphore"

	"github.com/elastic/beats/v7/heartbeat/config"
	"github.com/elastic/beats/v7/heartbeat/scheduler/timerqueue"
	"github.com/elastic/beats/v7/libbeat/common/atomic"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/monitoring"
)

const (
	statePreRunning int = iota + 1
	stateRunning
	stateStopped
)

var debugf = logp.MakeDebug("scheduler")

// ErrInvalidTransition is returned from start/stop when making an invalid state transition, say from preRunning to stopped
var ErrInvalidTransition = fmt.Errorf("invalid state transition")

// Scheduler represents our async timer based scheduler.
type Scheduler struct {
	limit       int64
	limitSem    *semaphore.Weighted
	state       atomic.Int
	location    *time.Location
	timerQueue  *timerqueue.TimerQueue
	ctx         context.Context
	cancelCtx   context.CancelFunc
	stats       schedulerStats
	jobLimitSem map[string]*semaphore.Weighted
	runOnceMode bool
}

type schedulerStats struct {
	activeJobs         *monitoring.Uint // gauge showing number of active jobs
	activeTasks        *monitoring.Uint // gauge showing number of active tasks
	waitingTasks       *monitoring.Uint // number of tasks waiting to run, but constrained by scheduler limit
	jobsPerSecond      *monitoring.Uint // rate of job processing computed over the past hour
	jobsMissedDeadline *monitoring.Uint // counter for number of jobs that missed start deadline
}

// TaskFunc represents a single task in a job. Optionally returns continuation of tasks to
// be executed within current job.
type TaskFunc func(ctx context.Context) []TaskFunc

// Schedule defines an interface for getting the next scheduled runtime for a job
type Schedule interface {
	// Next returns the next runAt a scheduled event occurs after the given runAt
	Next(now time.Time) (next time.Time)
	// Returns true if this schedule type should run once immediately before checking Next.
	// Cron tasks run at exact times so should set this to false.
	RunOnInit() bool
}

func getJobLimitSem(jobLimitByType map[string]config.JobLimit) map[string]*semaphore.Weighted {
	jobLimitSem := map[string]*semaphore.Weighted{}
	for jobType, jobLimit := range jobLimitByType {
		if jobLimit.Limit > 0 {
			jobLimitSem[jobType] = semaphore.NewWeighted(jobLimit.Limit)
		}
	}
	return jobLimitSem
}

// New creates a new Scheduler
func New(limit int64, registry *monitoring.Registry) *Scheduler {
	return NewWithLocation(limit, registry, time.Local, nil, false)
}

// NewWithLocation creates a new Scheduler using the given runAt zone.
func NewWithLocation(limit int64, registry *monitoring.Registry, location *time.Location, jobLimitByType map[string]config.JobLimit, runOnceMode bool) *Scheduler {
	ctx, cancelCtx := context.WithCancel(context.Background())

	if limit < 1 {
		limit = math.MaxInt64
	}

	jobsMissedDeadlineCounter := monitoring.NewUint(registry, "jobs.missed_deadline")
	activeJobsGauge := monitoring.NewUint(registry, "jobs.active")
	activeTasksGauge := monitoring.NewUint(registry, "tasks.active")
	waitingTasksGauge := monitoring.NewUint(registry, "tasks.waiting")

	sched := &Scheduler{
		limit:       limit,
		location:    location,
		state:       atomic.MakeInt(statePreRunning),
		ctx:         ctx,
		cancelCtx:   cancelCtx,
		limitSem:    semaphore.NewWeighted(limit),
		jobLimitSem: getJobLimitSem(jobLimitByType),
		timerQueue:  timerqueue.NewTimerQueue(ctx),
		runOnceMode: runOnceMode,

		stats: schedulerStats{
			activeJobs:         activeJobsGauge,
			activeTasks:        activeTasksGauge,
			waitingTasks:       waitingTasksGauge,
			jobsMissedDeadline: jobsMissedDeadlineCounter,
		},
	}

	return sched
}

// Start the scheduler. Starting a stopped scheduler returns an error.
func (s *Scheduler) Start() error {
	if s.state.Load() == stateStopped {
		return ErrInvalidTransition
	}
	if !s.state.CAS(statePreRunning, stateRunning) {
		return nil // we already running, just exit
	}

	s.timerQueue.Start()

	// Missed deadline reporter
	go s.missedDeadlineReporter()

	return nil
}

func (s *Scheduler) missedDeadlineReporter() {
	interval := time.Second * 15

	t := time.NewTicker(interval)

	// Counter used to check if we're missing more checks now than before
	missedAtLastCheck := uint64(0)
	for {
		select {
		case <-s.ctx.Done():
			t.Stop()
			return
		case <-t.C:
			missingNow := s.stats.jobsMissedDeadline.Get()
			missedDelta := missingNow - missedAtLastCheck
			if missedDelta > 0 {
				logp.Warn("%d tasks have missed their schedule deadlines in the last %s.", missedDelta, interval)
			}
			missedAtLastCheck = missingNow
		}
	}
}

// Stop all executing tasks in the scheduler. Cannot be restarted after Stop.
func (s *Scheduler) Stop() error {
	if s.state.CAS(stateRunning, stateStopped) {
		s.cancelCtx()
		return nil
	} else if s.state.Load() == stateStopped {
		return nil
	}

	return ErrInvalidTransition
}

// ErrAlreadyStopped is returned when an Add operation is attempted after the scheduler
// has already stopped.
var ErrAlreadyStopped = errors.New("attempted to add job to already stopped scheduler")

// Add adds the given TaskFunc to the current scheduler. Will return an error if the scheduler
// is done.
func (s *Scheduler) Add(sched Schedule, id string, entrypoint TaskFunc, jobType string) (removeFn context.CancelFunc, err error) {
	if s.state.Load() == stateStopped {
		return nil, ErrAlreadyStopped
	}

	jobCtx, jobCtxCancel := context.WithCancel(s.ctx)

	// lastRanAt stores the last runAt the task was invoked
	// The initial value is runAt.Now() because we use it to get the next runAt a job is scheduled to run
	lastRanAt := time.Now().In(s.location)

	var taskFn timerqueue.TimerTaskFn

	taskFn = func(_ time.Time) {
		select {
		case <-jobCtx.Done():
			debugf("Job '%v' canceled", id)
			return
		default:
		}
		s.stats.activeJobs.Inc()
		debugf("Job '%s' started", id)
		lastRanAt := newSchedJob(jobCtx, s, id, jobType, entrypoint).run()
		s.stats.activeJobs.Dec()
		if !s.runOnceMode {
			s.runOnce(sched.Next(lastRanAt), taskFn)
			debugf("Job '%v' returned at %v", id, time.Now())
		}

	}

	// We skip using the scheduler to execute the initial tasks for jobs that have RunOnInit returning true.
	// You might think it'd be simpler to just invoke runOnce in either case with 0 as a lastRanAt value,
	// however, that would caused the missed deadline stats to be incremented. Given that, it's easier
	// and slightly more efficient to simply run these tasks immediately in a goroutine.
	if sched.RunOnInit() {
		go taskFn(time.Now())
	} else {
		s.runOnce(sched.Next(lastRanAt), taskFn)
	}

	return func() {
		debugf("Remove scheduler job '%v'", id)
		jobCtxCancel()
	}, nil
}

func (s *Scheduler) runOnce(runAt time.Time, taskFn timerqueue.TimerTaskFn) {
	now := time.Now().In(s.location)
	if runAt.Before(now) {
		// Our last invocation went long!
		s.stats.jobsMissedDeadline.Inc()
	}

	// Schedule task to run sometime in the future. Wrap the task in a go-routine so it doesn't
	// block the timer thread.
	asyncTask := func(now time.Time) { go taskFn(now) }
	s.timerQueue.Push(runAt, asyncTask)
}

func (s *Scheduler) RunOnceCompleted() bool {
	return s.stats.activeJobs.Get() == 0 && s.stats.activeTasks.Get() == 0 && s.stats.waitingTasks.Get() == 0
}
