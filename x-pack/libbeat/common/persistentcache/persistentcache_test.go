// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package persistentcache

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	logp.DevelopmentSetup()
}

func TestPutGet(t *testing.T) {
	t.Parallel()

	registry := newTestRegistry(t)

	cache, err := newPersistentCache(registry, "test", 0, PersistentCacheOptions{})
	require.NoError(t, err)
	defer cache.Close()

	type valueType struct {
		Something string
	}

	var key = "somekey"
	var value = valueType{Something: "foo"}

	err = cache.Put(key, value)
	assert.NoError(t, err)

	var result valueType
	err = cache.Get(key, &result)
	assert.NoError(t, err)
	assert.Equal(t, value, result)

	err = cache.Get("notexist", &result)
	assert.Error(t, err)
}

func TestPersist(t *testing.T) {
	t.Parallel()

	registry := newTestRegistry(t)

	cache, err := newPersistentCache(registry, "test", 0, PersistentCacheOptions{})
	require.NoError(t, err)

	type valueType struct {
		Something string
	}

	var key = "somekey"
	var value = valueType{Something: "foo"}

	err = cache.Put(key, value)
	assert.NoError(t, err)

	cache.Close()

	cache, err = newPersistentCache(registry, "test", 0, PersistentCacheOptions{})
	require.NoError(t, err)

	var result valueType
	err = cache.Get(key, &result)
	assert.NoError(t, err)
	assert.Equal(t, value, result)
}

func TestCleanup(t *testing.T) {
	t.Parallel()

	registry := newTestRegistry(t)

	removeChan := make(chan common.Value)
	options := PersistentCacheOptions{
		RemovalListener: func(k string, v common.Value) {
			removeChan <- v
		},
	}

	cache, err := newPersistentCache(registry, "test", 0, options)
	require.NoError(t, err)

	now := time.Now()
	cache.clock = func() time.Time { return now }

	type valueType struct {
		Something string
	}

	var key = "somekey"
	var value = valueType{Something: "foo"}

	err = cache.PutWithTimeout(key, value, 5*time.Minute)
	assert.NoError(t, err)

	var result valueType
	err = cache.Get(key, &result)
	assert.NoError(t, err)
	assert.Equal(t, value, result)

	now = now.Add(10 * time.Minute)
	removedCount := cache.CleanUp()
	assert.Equal(t, 1, removedCount)

	err = cache.Get(key, &result)
	assert.Error(t, err)
}

func newTestRegistry(t *testing.T) *PersistentCacheRegistry {
	t.Helper()

	tempDir, err := ioutil.TempDir("", "beat-data-dir-")
	require.NoError(t, err)

	t.Cleanup(func() { os.RemoveAll(tempDir) })

	return &PersistentCacheRegistry{
		path: filepath.Join(tempDir, cacheFile),
	}
}
