package http_endpoint

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/elastic/beats/v7/libbeat/common"
)

func Test_httpReadJSON(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantObjs   []common.MapStr
		wantStatus int
		wantErr    bool
	}{
		{
			name:     "single object",
			body:     `{"a": 42, "b": "c"}`,
			wantObjs: []common.MapStr{{"a": float64(42), "b": "c"}},
		},
		{
			name:     "array accepted",
			body:     `[{"a":"b"},{"c":"d"}]`,
			wantObjs: []common.MapStr{{"a": "b"}, {"c": "d"}},
		},
		{
			name:       "not an object not accepted",
			body:       `42`,
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name:       "sequence of objects not accepted (CRLF)",
			body:       "{\"a\":1}\r\n{\"a\":2}",
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name:       "sequence of objects not accepted (no separator)",
			body:       "{\"a\":1}{\"a\":2}",
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name:       "not an object mixed",
			body:       "[{\"a\":1},\n42,\n{\"a\":2}]",
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObjs, gotStatus, err := httpReadJSON(strings.NewReader(tt.body))
			if (err != nil) != tt.wantErr {
				t.Errorf("httpReadNDJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObjs, tt.wantObjs) {
				t.Errorf("httpReadNDJSON() gotObjs = %v, want %v", gotObjs, tt.wantObjs)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("httpReadNDJSON() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func Test_httpReadNDJSON(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantObjs   []common.MapStr
		wantStatus int
		wantErr    bool
	}{
		{
			name:     "single object",
			body:     `{"a": 42, "b": "c"}`,
			wantObjs: []common.MapStr{{"a": float64(42), "b": "c"}},
		},
		{
			name:       "array not accepted",
			body:       `[{},{}]`,
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name:       "not an object not accepted",
			body:       `42`,
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name:     "sequence of objects accepted (CRLF)",
			body:     "{\"a\":1}\r\n{\"a\":2}",
			wantObjs: []common.MapStr{{"a": float64(1)}, {"a": float64(2)}},
		},
		{
			name:     "sequence of objects accepted (LF)",
			body:     "{\"a\":1}\n{\"a\":2}",
			wantObjs: []common.MapStr{{"a": float64(1)}, {"a": float64(2)}},
		},
		{
			name:     "sequence of objects accepted (SP)",
			body:     "{\"a\":1} {\"a\":2}",
			wantObjs: []common.MapStr{{"a": float64(1)}, {"a": float64(2)}},
		},
		{
			name:     "sequence of objects accepted (no separator)",
			body:     "{\"a\":1}{\"a\":2}",
			wantObjs: []common.MapStr{{"a": float64(1)}, {"a": float64(2)}},
		},
		{
			name:       "not an object mixed",
			body:       "{\"a\":1}\n42\n{\"a\":2}",
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObjs, gotStatus, err := httpReadNDJSON(strings.NewReader(tt.body))
			if (err != nil) != tt.wantErr {
				t.Errorf("httpReadNDJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObjs, tt.wantObjs) {
				t.Errorf("httpReadNDJSON() gotObjs = %v, want %v", gotObjs, tt.wantObjs)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("httpReadNDJSON() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
