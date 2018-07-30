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

// +build integration

package oplog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/elastic/beats/metricbeat/module/mongodb"
)

func TestFetch(t *testing.T) {
	compose.EnsureUp(t, "mongodb")
	
	err := initiateReplicaSet()
	if err != nil {
		t.FailNow()
	}

	f := mbtest.NewEventFetcher(t, getConfig())
	event, err := f.Fetch()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)

	// Check event fields
	allocated := event["size"].(common.MapStr)["allocated"].(int64)
	assert.True(t, allocated >= 0)

	used := event["size"].(common.MapStr)["used"].(int64)
	assert.True(t, used > 0)

	firstTs := event["first"].(common.MapStr)["timestamp"].(int64)
	assert.True(t, firstTs >= 0)

	window := event["window"].(int64)
	assert.True(t, window >= 0)
}

func TestData(t *testing.T) {
	compose.EnsureUp(t, "mongodb")

	f := mbtest.NewEventFetcher(t, getConfig())
	err := mbtest.WriteEvent(f, t)
	if err != nil {
		t.Fatal("write", err)
	}
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "mongodb",
		"metricsets": []string{"oplog"},
		"hosts":      []string{mongodb.GetEnvHost() + ":" + mongodb.GetEnvPort()},
	}
}

func initiateReplicaSet() error {
        url := mongodb.GetEnvHost() + ":" + mongodb.GetEnvPort()
	mongoSession, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	defer mongoSession.Close()

	// get oplog.rs collection
	db := mongoSession.DB("admin")
	config := ReplicaConfig{"beats", []Host{{0, url}}}
	var initiateResult map[string]interface{}
	if err := db.Run(bson.M{"replSetInitiate": config}, &initiateResult); err != nil {
		if err.Error() != "already initialized" {
			return err
		}
	}

	return nil
}

type ReplicaConfig struct {
	id      string `bson:_id`
	members []Host `bson:hosts`
}

type Host struct {
	id   int    `bson:_id`
	host string `bson:host`
}

