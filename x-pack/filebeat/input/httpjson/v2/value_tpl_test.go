package v2

import (
	"net/http"
	"testing"
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/stretchr/testify/assert"
)

func TestValueTpl(t *testing.T) {
	cases := []struct {
		name        string
		value       string
		paramCtx    transformContext
		paramTr     *transformable
		paramDefVal string
		expected    string
		setup       func()
		teardown    func()
	}{
		{
			name:  "can render values from ctx",
			value: "{{.last_response.body.param}}",
			paramCtx: transformContext{
				lastResponse: newTestTransformable(common.MapStr{"param": 25}, nil, ""),
			},
			paramTr:     emptyTransformable(),
			paramDefVal: "",
			expected:    "25",
		},
		{
			name:  "can render default value if execute fails",
			value: "{{.last_response.body.does_not_exist}}",
			paramCtx: transformContext{
				lastResponse: emptyTransformable(),
			},
			paramTr:     emptyTransformable(),
			paramDefVal: "25",
			expected:    "25",
		},
		{
			name:        "can render default value if template is empty",
			value:       "",
			paramCtx:    emptyTransformContext(),
			paramTr:     emptyTransformable(),
			paramDefVal: "25",
			expected:    "25",
		},
		{
			name:        "can render default value if execute panics",
			value:       "{{.last_response.panic}}",
			paramDefVal: "25",
			expected:    "25",
		},
		{
			name:     "func hour",
			value:    `{{ hour -1 }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "-1h0m0s",
		},
		{
			name:     "func now",
			setup:    func() { timeNow = func() time.Time { return time.Unix(1604582732, 0).UTC() } },
			teardown: func() { timeNow = time.Now },
			value:    `{{ now }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05 13:25:32 +0000 UTC",
		},
		{
			name:     "func now with duration",
			setup:    func() { timeNow = func() time.Time { return time.Unix(1604582732, 0).UTC() } },
			teardown: func() { timeNow = time.Now },
			value:    `{{ now (-1|hour) }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05 12:25:32 +0000 UTC",
		},
		{
			name:     "func parseDate",
			value:    `{{ parseDate "2020-11-05T12:25:32.1234567Z" "RFC3339Nano" }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05 12:25:32.1234567 +0000 UTC",
		},
		{
			name:     "func parseDate defaults to RFC3339",
			value:    `{{ parseDate "2020-11-05T12:25:32Z" }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05 12:25:32 +0000 UTC",
		},
		{
			name:     "func parseDate with custom layout",
			value:    `{{ (parseDate "Thu Nov  5 12:25:32 +0000 2020" "Mon Jan _2 15:04:05 -0700 2006").UTC }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05 12:25:32 +0000 UTC",
		},
		{
			name:     "func formatDate",
			setup:    func() { timeNow = func() time.Time { return time.Unix(1604582732, 0).UTC() } },
			teardown: func() { timeNow = time.Now },
			value:    `{{ formatDate (now) "UnixDate" "America/New_York" }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "Thu Nov  5 08:25:32 EST 2020",
		},
		{
			name:     "func formatDate defaults to UTC",
			setup:    func() { timeNow = func() time.Time { return time.Unix(1604582732, 0).UTC() } },
			teardown: func() { timeNow = time.Now },
			value:    `{{ formatDate (now) "UnixDate" }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "Thu Nov  5 13:25:32 UTC 2020",
		},
		{
			name:     "func formatDate falls back to UTC",
			setup:    func() { timeNow = func() time.Time { return time.Unix(1604582732, 0).UTC() } },
			teardown: func() { timeNow = time.Now },
			value:    `{{ formatDate (now) "UnixDate" "wrong/tz"}}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "Thu Nov  5 13:25:32 UTC 2020",
		},
		{
			name:     "func parseTimestamp",
			value:    `{{ (parseTimestamp 1604582732).UTC }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05 13:25:32 +0000 UTC",
		},
		{
			name:     "func parseTimestampMilli",
			value:    `{{ (parseTimestampMilli 1604582732000).UTC }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05 13:25:32 +0000 UTC",
		},
		{
			name:     "func parseTimestampNano",
			value:    `{{ (parseTimestampNano 1604582732000000000).UTC }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05 13:25:32 +0000 UTC",
		},
		{
			name:     "can execute functions pipeline",
			setup:    func() { timeNow = func() time.Time { return time.Unix(1604582732, 0).UTC() } },
			teardown: func() { timeNow = time.Now },
			value:    `{{ -1 | hour | now | formatDate }}`,
			paramCtx: emptyTransformContext(),
			paramTr:  emptyTransformable(),
			expected: "2020-11-05T12:25:32Z",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup()
			}
			if tc.teardown != nil {
				t.Cleanup(tc.teardown)
			}
			tpl := &valueTpl{}
			assert.NoError(t, tpl.Unpack(tc.value))
			got := tpl.Execute(tc.paramCtx, tc.paramTr, tc.paramDefVal)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func newTestTransformable(body common.MapStr, header http.Header, url string) *transformable {
	tr := emptyTransformable()
	if len(body) > 0 {
		tr.body = body
	}
	if len(header) > 0 {
		tr.header = header
	}
	if url != "" {
		tr.url = newURL(url)
	}
	return tr
}
