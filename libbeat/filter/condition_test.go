package filter

import (
	"testing"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/stretchr/testify/assert"
)

func TestBadCondition(t *testing.T) {

	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, []string{"*"})
	}

	configs := []ConditionConfig{
		ConditionConfig{
			Equals: &ConditionFilter{fields: map[string]interface{}{
				"proc.pid": 0.08,
			}},
		},

		ConditionConfig{
			Range: &ConditionFilter{fields: map[string]interface{}{
				"gtr": 0.3,
			}},
		},

		ConditionConfig{
			Range: &ConditionFilter{fields: map[string]interface{}{
				"gt": "fdfdd",
			}},
		},
		ConditionConfig{
			Regexp: &ConditionFilter{fields: map[string]interface{}{
				"proc.name": "58gdhsga-=kw++w00",
			}},
		},
	}

	for _, config := range configs {
		_, err := NewCondition(config)
		assert.NotNil(t, err)
	}
}

func GetConditions(t *testing.T, configs []ConditionConfig) []Condition {
	conds := []Condition{}

	for _, config := range configs {

		cond, err := NewCondition(config)
		assert.Nil(t, err)
		conds = append(conds, *cond)
	}
	assert.True(t, len(conds) == len(configs))

	return conds
}

func TestEqualsCondition(t *testing.T) {

	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, []string{"*"})
	}

	configs := []ConditionConfig{
		ConditionConfig{
			Equals: &ConditionFilter{fields: map[string]interface{}{
				"type": "process",
			}},
		},

		ConditionConfig{
			Equals: &ConditionFilter{fields: map[string]interface{}{
				"type":     "process",
				"proc.pid": 305,
			}},
		},

		ConditionConfig{
			Range: &ConditionFilter{fields: map[string]interface{}{
				"proc.cpu.total_p.gt": 0.5,
			}},
		},
	}

	conds := GetConditions(t, configs)

	event := common.MapStr{
		"@timestamp": "2016-04-14T20:41:06.258Z",
		"proc": common.MapStr{
			"cmdline": "/usr/libexec/secd",
			"cpu": common.MapStr{
				"start_time": "Apr10",
				"system":     1988,
				"total":      6029,
				"total_p":    0.08,
				"user":       4041,
			},
			"name":     "secd",
			"pid":      305,
			"ppid":     1,
			"state":    "running",
			"username": "monica",
		},
		"type": "process",
	}

	assert.True(t, conds[0].Check(event))
	assert.True(t, conds[1].Check(event))
	assert.False(t, conds[2].Check(event))
}

func TestContainsCondition(t *testing.T) {

	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, []string{"*"})
	}

	configs := []ConditionConfig{
		ConditionConfig{
			Contains: &ConditionFilter{fields: map[string]interface{}{
				"proc.name":     "sec",
				"proc.username": "monica",
			}},
		},

		ConditionConfig{
			Contains: &ConditionFilter{fields: map[string]interface{}{
				"type":      "process",
				"proc.name": "secddd",
			}},
		},
	}

	conds := GetConditions(t, configs)

	event := common.MapStr{
		"@timestamp": "2016-04-14T20:41:06.258Z",
		"proc": common.MapStr{
			"cmdline": "/usr/libexec/secd",
			"cpu": common.MapStr{
				"start_time": "Apr10",
				"system":     1988,
				"total":      6029,
				"total_p":    0.08,
				"user":       4041,
			},
			"name":     "secd",
			"pid":      305,
			"ppid":     1,
			"state":    "running",
			"username": "monica",
		},
		"type": "process",
	}

	assert.True(t, conds[0].Check(event))
	assert.False(t, conds[1].Check(event))
}

func TestRegexpCondition(t *testing.T) {

	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, []string{"*"})
	}

	configs := []ConditionConfig{
		ConditionConfig{
			Regexp: &ConditionFilter{fields: map[string]interface{}{
				"source": "apache2/error.*",
			}},
		},

		ConditionConfig{
			Regexp: &ConditionFilter{fields: map[string]interface{}{
				"source": "apache2/access.*",
			}},
		},

		ConditionConfig{
			Regexp: &ConditionFilter{fields: map[string]interface{}{
				"source":  "apache2/error.*",
				"message": "[client 1.2.3.4]",
			}},
		},
	}

	conds := GetConditions(t, configs)

	event := common.MapStr{
		"@timestamp": "2016-04-14T20:41:06.258Z",
		"message":    `[Fri Dec 16 01:46:23 2005] [error] [client 1.2.3.4] Directory index forbidden by rule: /home/test/`,
		"source":     "/var/log/apache2/error.log",
		"type":       "log",
		"input_type": "log",
		"offset":     30,
	}

	event1 := common.MapStr{
		"@timestamp": "2016-04-14T20:41:06.258Z",
		"message":    `127.0.0.1 - - [28/Jul/2006:10:27:32 -0300] "GET /hidden/ HTTP/1.0" 404 7218`,
		"source":     "/var/log/apache2/access.log",
		"type":       "log",
		"input_type": "log",
		"offset":     30,
	}

	assert.True(t, conds[0].Check(event))
	assert.False(t, conds[1].Check(event))
	assert.True(t, conds[2].Check(event))

	assert.True(t, conds[1].Check(event1))
	assert.False(t, conds[2].Check(event1))
}

func TestRangeCondition(t *testing.T) {

	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, []string{"*"})
	}

	configs := []ConditionConfig{
		ConditionConfig{
			Range: &ConditionFilter{fields: map[string]interface{}{
				"http.code.gte": 400,
				"http.code.lt":  500,
			}},
		},

		ConditionConfig{
			Range: &ConditionFilter{fields: map[string]interface{}{
				"bytes_out.gte": 2800,
			}},
		},

		ConditionConfig{
			Range: &ConditionFilter{fields: map[string]interface{}{
				"bytes_out.gte":   2800,
				"responsetime.gt": 30,
			}},
		},

		ConditionConfig{
			Range: &ConditionFilter{fields: map[string]interface{}{
				"proc.cpu.total_p.gte": 0.5,
			}},
		},
	}

	conds := GetConditions(t, configs)

	event := common.MapStr{
		"@timestamp":    "2015-06-11T09:51:23.642Z",
		"bytes_in":      126,
		"bytes_out":     28033,
		"client_ip":     "127.0.0.1",
		"client_port":   42840,
		"client_proc":   "",
		"client_server": "mar.local",
		"http": common.MapStr{
			"code":           404,
			"content_length": 76985,
			"phrase":         "Not found",
		},
		"ip":           "127.0.0.1",
		"method":       "GET",
		"params":       "",
		"path":         "/jszip.min.js",
		"port":         8000,
		"proc":         "",
		"query":        "GET /jszip.min.js",
		"responsetime": 30,
		"server":       "mar.local",
		"status":       "OK",
		"type":         "http",
	}

	event1 := common.MapStr{
		"@timestamp": "2016-04-20T07:46:44.633Z",
		"proc": common.MapStr{
			"cmdline": "/System/Library/Frameworks/CoreServices.framework/Frameworks/Metadata.framework/Versions/A/Support/mdworker -s mdworker -c MDSImporterWorker -m com.apple.mdworker.single",
			"cpu": common.MapStr{
				"start_time": "09:19",
				"system":     22,
				"total":      66,
				"total_p":    0.6,
				"user":       44,
			},
			"name":     "mdworker",
			"pid":      44978,
			"ppid":     1,
			"state":    "running",
			"username": "test",
		},
		"type": "process",
	}

	assert.True(t, conds[0].Check(event))
	assert.True(t, conds[1].Check(event))
	assert.False(t, conds[2].Check(event))
	assert.True(t, conds[3].Check(event1))
	assert.False(t, conds[3].Check(event))
}
