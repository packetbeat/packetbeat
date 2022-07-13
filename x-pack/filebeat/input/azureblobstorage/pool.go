// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !aix
// +build !aix

package azureblobstorage

import (
	"context"
	"sync"

	"github.com/elastic/elastic-agent-libs/logp"
)

type Pool interface {
	Start()
	Stop()
	Submit(job Job)
	AvailableWorkers() int32
}

type pool struct {
	ctx           context.Context
	workerErrChan chan error
	jobQueue      chan Job
	readyPool     chan chan Job
	workers       []Worker
	dispatcherWg  sync.WaitGroup
	workersWg     *sync.WaitGroup
	quit          chan bool
	log           *logp.Logger
}

var (
	poolError string = "worker pool error : %w"
)

// NewWorkerPool returns an instance of a worker pool with 'maxWorkers' ready to accept work
func NewWorkerPool(ctx context.Context, maxWorkers int, log *logp.Logger) Pool {

	workersWg := sync.WaitGroup{}

	readyPool := make(chan chan Job, maxWorkers)
	workers := make([]Worker, maxWorkers)
	errChan := make(chan error)

	// creates workers
	for i := 0; i < maxWorkers; i++ {
		workers[i] = NewWorker(ctx, i+1, readyPool, &workersWg, errChan)
	}

	return &pool{
		ctx:           ctx,
		workerErrChan: errChan,
		jobQueue:      make(chan Job),
		readyPool:     readyPool,
		workers:       workers,
		dispatcherWg:  sync.WaitGroup{},
		workersWg:     &workersWg,
		quit:          make(chan bool),
		log:           log,
	}
}

func (q *pool) Start() {
	// puts workers in ready state
	for i := 0; i < len(q.workers); i++ {
		q.workers[i].Start()
	}

	// starts dispatcher
	go q.dispatch()
}

// This is a blocking if all workers are busy
func (q *pool) Submit(job Job) {
	q.jobQueue <- job
}

// AvailableWorkers returns the number of free workers at any point of time
func (q *pool) AvailableWorkers() int32 {
	return int32(len(q.readyPool))
}

func (q *pool) Stop() {
	q.quit <- true
	q.dispatcherWg.Wait()
}

func (q *pool) dispatch() {
	// starts dispatching jobs
	q.dispatcherWg.Add(1)
	for {
		select {

		case err := <-q.workerErrChan:
			q.log.Errorf(poolError, err)

		case <-q.ctx.Done():
			q.log.Errorf(poolError, q.ctx.Err())
			q.Stop()

		case job := <-q.jobQueue:
			workerXChannel := <-q.readyPool // free worker 'x' found
			workerXChannel <- job           // assigns job to worker 'x'

		case <-q.quit:
			// frees all workers & gracefully close worker pool
			for i := 0; i < len(q.workers); i++ {
				q.workers[i].Stop()
			}
			// waits for all workers to finish their jobs
			q.workersWg.Wait()
			// stops dispatcher
			q.dispatcherWg.Done()
			return
		}
	}
}
