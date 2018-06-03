// +build windows

package procs

import (
	"encoding/hex"
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestParseTableRaw(t *testing.T) {
	IPv4 := extractTCPRowOwnerPID
	IPv6 := extractTCP6RowOwnerPID

	for idx, testCase := range []struct {
		name     string
		factory  extractorFactory
		raw      string
		expected []portProcMapping
		mustErr  bool
	}{
		{"Empty table IPv4", IPv4,
			"00000000", nil, false},
		{"Empty table IPv6", IPv6,
			"00000000", nil, false},
		{"Short table (no length)", IPv4,
			"000000", nil, true},
		{"Short table (partial entry)", IPv6,
			"01000000AAAAAAAAAAAAAAAAAAAA", nil, true},
		{"One entry (IPv4)", IPv4,
			"01000000" +
				"77777777AAAAAAAA12340000BBBBBBBBFFFF0000CCCCCCCC",
			[]portProcMapping{
				{port: 0x1234, pid: 0xCCCCCCCC},
			}, false},
		{"Two entries (IPv6)", IPv6,
			"02000000" +
				// First entry
				"11112222333344445555666677778888F0F0F0F0" +
				"ABCDEFFF" + // local port
				"FFFFEEEEDDDDCCCCBBBBAAAA999988880A0A0A0A" +
				"33333333" + // remote port
				"77777777" +
				"01000000" + // pid
				// second entry
				"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBBBBBBB" +
				"0000FFFF" + // local port
				"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBCCCCCCCC" +
				"44444444" + // remote port
				"77777777" +
				"FFFF0000" + // pid
				"",
			[]portProcMapping{
				{port: 0xABCD, pid: 1},
				{port: 0, pid: 0xffff},
			}, false},
	} {
		msg := fmt.Sprintf("Test case #%d: %s", idx+1, testCase.name)
		table, err := hex.DecodeString(testCase.raw)
		assert.NoError(t, err, msg)
		var result []portProcMapping
		callback := func(port uint16, pid int) {
			result = append(result, portProcMapping{port: port, pid: pid})
		}
		err = parseTable(table, testCase.factory(callback))
		if testCase.mustErr {
			assert.Error(t, err, msg)
		} else {
			assert.NoError(t, err, msg)
			assert.Len(t, result, len(testCase.expected), msg)
			assert.Equal(t, testCase.expected, result, msg)
		}
	}
}

func TestParseTableSizes(t *testing.T) {
	// Make sure the structs in Golang have the expected size
	assert.Equal(t, uintptr(sizeOfTCPRowOwnerPID), unsafe.Sizeof(TCPRowOwnerPID{}))
	assert.Equal(t, uintptr(sizeOfTCP6RowOwnerPID), unsafe.Sizeof(TCP6RowOwnerPID{}))
}
