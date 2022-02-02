package resources

import (
	"context"
	"fmt"

	"github.com/elastic/beats/v7/libbeat/logp"
)

type FetchersRegistry interface {
	Register(key string, f Fetcher, c ...FetcherCondition) error
	Keys() []string
	ShouldRun(key string) bool
	Run(ctx context.Context, key string) ([]FetcherResult, error)
	Stop(ctx context.Context)
}

type fetchersRegistry struct {
	reg map[string]registeredFetcher
}

type registeredFetcher struct {
	f Fetcher
	c []FetcherCondition
}

func NewFetcherRegistry() FetchersRegistry {
	return &fetchersRegistry{
		reg: make(map[string]registeredFetcher),
	}
}

// Register registers a Fetcher implementation.
func (r *fetchersRegistry) Register(key string, f Fetcher, c ...FetcherCondition) error {
	if _, ok := r.reg[key]; ok {
		return fmt.Errorf("fetcher key collision: %q is already registered", key)
	}

	r.reg[key] = registeredFetcher{
		f: f,
		c: c,
	}

	return nil
}

func (r *fetchersRegistry) Keys() []string {
	keys := []string{}
	for k := range r.reg {
		keys = append(keys, k)
	}

	return keys
}

func (r *fetchersRegistry) ShouldRun(key string) bool {
	registered, ok := r.reg[key]
	if !ok {
		return false
	}

	for _, condition := range registered.c {
		if !condition.Condition() {
			logp.L().Infof("Conditional fetcher %q should not run because %q", key, condition.Name())
			return false
		}
	}

	return true
}

func (r *fetchersRegistry) Run(ctx context.Context, key string) ([]FetcherResult, error) {
	registered, ok := r.reg[key]
	if !ok {
		return nil, fmt.Errorf("fetcher %v not found", key)
	}

	return registered.f.Fetch(ctx)
}

func (r *fetchersRegistry) Stop(ctx context.Context) {
	for key, rfetcher := range r.reg {
		rfetcher.f.Stop()
		logp.L().Infof("Fetcher for key %q stopped", key)
	}
}
