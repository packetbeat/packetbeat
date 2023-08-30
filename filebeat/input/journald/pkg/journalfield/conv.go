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

package journalfield

import (
	"fmt"
	"math/bits"
	"regexp"
	"strconv"
	"strings"

	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// FieldConversion provides the mappings and conversion rules for raw fields of journald entries.
type FieldConversion map[string]Conversion

// Conversion configures the conversion rules for a field.
type Conversion struct {
	Names     []string
	IsInteger bool
	Dropped   bool
}

// Converter applis configured conversion rules to journald entries, producing
// a new mapstr.M.
type Converter struct {
	log         *logp.Logger
	conversions FieldConversion
}

// NewConverter creates a new Converter from the given conversion rules. If
// conversions is nil, internal default conversion rules will be applied.
func NewConverter(log *logp.Logger, conversions FieldConversion) *Converter {
	if conversions == nil {
		conversions = journaldEventFields
	}

	return &Converter{log: log, conversions: conversions}
}

// Convert creates a mapstr.M from the raw fields by applying the
// configured conversion rules.
// Field type conversion errors are logged to at debug level and the original
// value is added to the map.
func (c *Converter) Convert(entryFields map[string]string) mapstr.M {
	fields := mapstr.M{}
	var custom mapstr.M

	for entryKey, v := range entryFields {
		if fieldConversionInfo, ok := c.conversions[entryKey]; !ok {
			if custom == nil {
				custom = mapstr.M{}
			}
			normalized := strings.ToLower(strings.TrimLeft(entryKey, "_"))
			_, _ = custom.Put(normalized, v)
		} else if !fieldConversionInfo.Dropped {
			value, err := convertValue(fieldConversionInfo, v)
			if err != nil {
				value = v
				c.log.Debugf("Journald mapping error: %v", err)
			}
			for _, name := range fieldConversionInfo.Names {
				_, _ = fields.Put(name, value)
			}
		}
	}

	if len(custom) != 0 {
		_, _ = fields.Put("journald.custom", custom)
	}

	return withECSEnrichment(fields)
}

func convertValue(fc Conversion, value string) (interface{}, error) {
	if fc.IsInteger {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			// On some versions of systemd the 'syslog.pid' can contain the username
			// appended to the end of the pid. In most cases this does not occur
			// but in the cases that it does, this tries to strip ',\w*' from the
			// value and then perform the conversion.
			s := strings.Split(value, ",")
			v, err = strconv.ParseInt(s[0], 10, 64)
			if err != nil {
				return value, fmt.Errorf("failed to convert field %s \"%v\" to int: %w", fc.Names[0], value, err)
			}
		}
		return v, nil
	}
	return value, nil
}

func withECSEnrichment(fields mapstr.M) mapstr.M {
	// from https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html
	// we see journald.object fields are populated by systemd on behalf of a different program
	// so we want them to favor their use in root fields as they are the from the effective program
	// performing the action.
	setGidUidFields("journald", fields)
	setGidUidFields("journald.object", fields)
	setProcessFields("journald", fields)
	setProcessFields("journald.object", fields)
	expandCapabilities(fields)
	return fields
}

func setGidUidFields(prefix string, fields mapstr.M) {
	var auditLoginUid string
	if found, _ := fields.HasKey(prefix + ".audit.login_uid"); found {
		auditLoginUid = fmt.Sprint(getIntegerFromFields(prefix+".audit.login_uid", fields))
		_, _ = fields.Put("user.id", auditLoginUid)
	}

	if found, _ := fields.HasKey(prefix + ".uid"); !found {
		return
	}

	uid := fmt.Sprint(getIntegerFromFields(prefix+".uid", fields))
	gid := fmt.Sprint(getIntegerFromFields(prefix+".gid", fields))
	if auditLoginUid != "" && auditLoginUid != uid {
		putStringIfNotEmtpy("user.effective.id", uid, fields)
		putStringIfNotEmtpy("user.effective.group.id", gid, fields)
	} else {
		putStringIfNotEmtpy("user.id", uid, fields)
		putStringIfNotEmtpy("user.group.id", gid, fields)
	}
}

var cmdlineRegexp = regexp.MustCompile(`"(\\"|[^"])*?"|[^\s]+`)

func setProcessFields(prefix string, fields mapstr.M) {
	if found, _ := fields.HasKey(prefix + ".pid"); found {
		pid := getIntegerFromFields(prefix+".pid", fields)
		_, _ = fields.Put("process.pid", pid)
	}

	name := getStringFromFields(prefix+".name", fields)
	if name != "" {
		_, _ = fields.Put("process.name", name)
	}

	executable := getStringFromFields(prefix+".executable", fields)
	if executable != "" {
		_, _ = fields.Put("process.executable", executable)
	}

	cmdline := getStringFromFields(prefix+".process.command_line", fields)
	if cmdline == "" {
		return
	}

	_, _ = fields.Put("process.command_line", cmdline)

	args := cmdlineRegexp.FindAllString(cmdline, -1)
	if len(args) > 0 {
		_, _ = fields.Put("process.args", args)
		_, _ = fields.Put("process.args_count", len(args))
	}
}

// expandCapabilites expands the hex string of capabilities bits in the
// journald.process.capabilities field in-place into an array of conventional
// capabilities names in process.thread.capabilities.effective. If a
// capability is unknown it is rendered as the numeric value of the cap.
// The original capabilities string is not altered. If any error is
// encountered no modification is made to the fields.
func expandCapabilities(fields mapstr.M) {
	cs, err := fields.GetValue("journald.process.capabilities")
	if err != nil {
		return
	}
	c, ok := cs.(string)
	if !ok {
		return
	}
	w, err := strconv.ParseUint(c, 16, 64)
	if err != nil {
		return
	}
	if w == 0 {
		return
	}
	caps := make([]string, 0, bits.OnesCount64(w))
	for i := 0; w != 0; i++ {
		if w&1 != 0 {
			if i < len(capTable) {
				caps = append(caps, capTable[i])
			} else {
				caps = append(caps, strconv.Itoa(i))
			}
		}
		w >>= 1
	}
	fields.Put("process.thread.capabilities.effective", caps)
}

// include/uapi/linux/capability.h
var capTable = [...]string{
	0:  "cap_chown",
	1:  "cap_dac_override",
	2:  "cap_dac_read_search",
	3:  "cap_fowner",
	4:  "cap_fsetid",
	5:  "cap_kill",
	6:  "cap_setgid",
	7:  "cap_setuid",
	8:  "cap_setpcap",
	9:  "cap_linux_immutable",
	10: "cap_net_bind_service",
	11: "cap_net_broadcast",
	12: "cap_net_admin",
	13: "cap_net_raw",
	14: "cap_ipc_lock",
	15: "cap_ipc_owner",
	16: "cap_sys_module",
	17: "cap_sys_rawio",
	18: "cap_sys_chroot",
	19: "cap_sys_ptrace",
	20: "cap_sys_pacct",
	21: "cap_sys_admin",
	22: "cap_sys_boot",
	23: "cap_sys_nice",
	24: "cap_sys_resource",
	25: "cap_sys_time",
	26: "cap_sys_tty_config",
	27: "cap_mknod",
	28: "cap_lease",
	29: "cap_audit_write",
	30: "cap_audit_control",
	31: "cap_setfcap",
	32: "cap_mac_override",
	33: "cap_mac_admin",
	34: "cap_syslog",
	35: "cap_wake_alarm",
	36: "cap_block_suspend",
	37: "cap_audit_read",
	38: "cap_perfmon",
	39: "cap_bpf",
	40: "cap_checkpoint_restore",
}

func getStringFromFields(key string, fields mapstr.M) string {
	value, _ := fields.GetValue(key)
	str, _ := value.(string)
	return str
}

func getIntegerFromFields(key string, fields mapstr.M) int64 {
	value, _ := fields.GetValue(key)
	i, _ := value.(int64)
	return i
}

func putStringIfNotEmtpy(k, v string, fields mapstr.M) {
	if v == "" {
		return
	}
	_, _ = fields.Put(k, v)
}

// helpers for creating a field conversion table.
var ignoredField = Conversion{Dropped: true}

func text(names ...string) Conversion {
	return Conversion{Names: names, IsInteger: false, Dropped: false}
}

func integer(names ...string) Conversion {
	return Conversion{Names: names, IsInteger: true, Dropped: false}
}
