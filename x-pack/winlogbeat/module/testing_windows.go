// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package module

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/processors/script/javascript"
	"github.com/elastic/beats/winlogbeat/checkpoint"
	"github.com/elastic/beats/winlogbeat/eventlog"

	// Register javascript modules.
	_ "github.com/elastic/beats/libbeat/processors/script/javascript/module"
	_ "github.com/elastic/beats/winlogbeat/processors/script/javascript/module/winlogbeat"
)

var update = flag.Bool("update", false, "update golden files")

// Option configures the test behavior.
type Option func(*params)

type params struct {
	ignoreFields []string
}

// WithFieldFilter filters the specified fields from the event prior to
// creating the golden file.
func WithFieldFilter(filter []string) Option {
	return func(p *params) {
		p.ignoreFields = filter
	}
}

// TestPipeline tests the given pipeline by reading events from the .evtx files
// and processing them with the pipeline. Then it compares the results against
// a saved golden file. Use -update to regenerate the golden files.
func TestPipeline(t *testing.T, evtx string, pipeline string, opts ...Option) {
	files, err := filepath.Glob(evtx)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("glob", evtx, "didn't match any files")
	}

	var p params
	for _, o := range opts {
		o(&p)
	}

	for _, f := range files {
		t.Run(filepath.Base(f), func(t *testing.T) {
			testPipeline(t, f, pipeline, &p)
		})
	}
}

func testPipeline(t testing.TB, evtx string, pipeline string, p *params) {
	t.Helper()

	path, err := filepath.Abs(evtx)
	if err != nil {
		t.Fatal(err)
	}

	// Open evtx file.
	log, err := eventlog.New(common.MustNewConfigFrom(common.MapStr{
		"name":           path,
		"api":            "wineventlog",
		"no_more_events": "stop",
	}))
	if err != nil {
		t.Fatal(err)
	}
	defer log.Close()

	if err = log.Open(checkpoint.EventLogState{}); err != nil {
		t.Fatal(err)
	}

	// Load javascript processor.
	processor, err := javascript.New(common.MustNewConfigFrom(common.MapStr{
		"file": pipeline,
	}))
	if err != nil {
		t.Fatal(err)
	}

	// Read and process events.
	var events []common.MapStr
	for stop := false; !stop; {
		records, err := log.Read()
		if err == io.EOF {
			stop = true
		} else if err != nil {
			t.Fatal(err)
		}

		for _, r := range records {
			record := r.ToEvent()
			record.Fields.Delete("event.created")
			record.Fields.Delete("log.file")

			evt, err := processor.Run(&record)
			if err != nil {
				t.Fatalf("%v while processing event:\n%v", err, record.Fields.StringToPrint())
			}

			// Ensure timezone is UTC. In the normal Beats output this is handled
			// by the encoder (go-structform).
			evt.PutValue("@timestamp", evt.Timestamp.UTC())

			events = append(events, filterEvent(evt.Fields, p.ignoreFields))
		}
	}

	if *update {
		writeGolden(t, path, events)
		return
	}

	expected := readGolden(t, path)
	if !assert.Len(t, events, len(expected)) {
		return
	}
	for i, e := range events {
		assert.EqualValues(t, expected[i], normalize(t, e))
	}
}

func writeGolden(t testing.TB, source string, events []common.MapStr) {
	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll("testdata", 0755); err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join("testdata", filepath.Base(source)+".golden.json")
	if err := ioutil.WriteFile(outPath, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func readGolden(t testing.TB, source string) []common.MapStr {
	inPath := filepath.Join("testdata", filepath.Base(source)+".golden.json")

	data, err := ioutil.ReadFile(inPath)
	if err != nil {
		t.Fatal(err)
	}

	var events []common.MapStr
	if err = json.Unmarshal(data, &events); err != nil {
		t.Fatal(err)
	}

	return events
}

func normalize(t testing.TB, m common.MapStr) common.MapStr {
	data, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}

	var out common.MapStr
	if err = json.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}

	return out
}

func filterEvent(m common.MapStr, ignores []string) common.MapStr {
	for _, f := range ignores {
		m.Delete(f)
	}
	return m
}
