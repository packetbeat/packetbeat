// Need for unit and integration tests
package elasticsearch

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs/outil"
)

func GetValidQueryResult() QueryResult {
	result := QueryResult{
		Ok:    true,
		Index: "testIndex",
		Type:  "testType",
		ID:    "12",
		Source: []byte(`{
			"ok": true,
			"_type":"testType",
			"_index":"testIndex",
			"_id":"12",
			"_version": 2,
			"found": true,
			"exists": true,
			"created": true,
			"lastname":"ruflin",
			"firstname": "nicolas"}`,
		),
		Version: 2,
		Found:   true,
		Exists:  true,
		Created: true,
		Matches: []string{"abc", "def"},
	}

	return result
}

func GetValidSearchResults() SearchResults {
	hits := Hits{
		Total: 0,
		Hits:  nil,
	}

	results := SearchResults{
		Took: 19,
		Shards: []byte(`{
    		"total" : 3,
    		"successful" : 2,
    		"failed" : 1
  		}`),
		Hits: hits,
		Aggs: nil,
	}

	return results
}

func TestReadQueryResult(t *testing.T) {
	queryResult := GetValidQueryResult()

	json := queryResult.Source
	result, err := readQueryResult(json)

	assert.Nil(t, err)
	assert.Equal(t, queryResult.Ok, result.Ok)
	assert.Equal(t, queryResult.Index, result.Index)
	assert.Equal(t, queryResult.Type, result.Type)
	assert.Equal(t, queryResult.ID, result.ID)
	assert.Equal(t, queryResult.Version, result.Version)
	assert.Equal(t, queryResult.Found, result.Found)
	assert.Equal(t, queryResult.Exists, result.Exists)
	assert.Equal(t, queryResult.Created, result.Created)
}

// Check empty query result object
func TestReadQueryResult_empty(t *testing.T) {
	result, err := readQueryResult(nil)
	assert.Nil(t, result)
	assert.Nil(t, err)
}

// Check invalid query result object
func TestReadQueryResult_invalid(t *testing.T) {
	// Invalid json string
	json := []byte(`{"name":"ruflin","234"}`)

	result, err := readQueryResult(json)
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestReadSearchResult(t *testing.T) {
	resultsObject := GetValidSearchResults()

	json := []byte(`{
  		"took" : 19,
  		"_shards" : {
    		"total" : 3,
    		"successful" : 2,
    		"failed" : 1
  		},
  		"hits" : {},
  		"aggs" : {}
  	}`)

	results, err := readSearchResult(json)

	assert.Nil(t, err)
	assert.Equal(t, resultsObject.Took, results.Took)
	assert.Equal(t, resultsObject.Hits, results.Hits)
	assert.Equal(t, resultsObject.Shards, results.Shards)
	assert.Equal(t, resultsObject.Aggs, results.Aggs)
}

func TestReadSearchResult_empty(t *testing.T) {
	results, err := readSearchResult(nil)
	assert.Nil(t, results)
	assert.Nil(t, err)
}

func TestReadSearchResult_invalid(t *testing.T) {
	// Invalid json string
	json := []byte(`{"took":"19","234"}`)

	results, err := readSearchResult(json)
	assert.Nil(t, results)
	assert.Error(t, err)
}

func newTestClient(url string) *Client {
	client, err := NewClient(ClientSettings{
		URL:              url,
		Index:            outil.MakeSelector(),
		Timeout:          60 * time.Second,
		CompressionLevel: 3,
	}, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func (r QueryResult) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		logp.Warn("failed to marshal QueryResult (%v): %#v", err, r)
		return "ERROR"
	}
	return string(out)
}
