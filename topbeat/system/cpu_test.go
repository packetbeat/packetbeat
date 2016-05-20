// +build !integration

package system

import (
	"testing"

	"github.com/elastic/beats/libbeat/logp"
	sigar "github.com/elastic/gosigar"
	"github.com/stretchr/testify/assert"
)

func TestGetCpuTimes(t *testing.T) {

	cpu_stat, err := GetCpuTimes()

	assert.NotNil(t, cpu_stat)
	assert.Nil(t, err)

	assert.True(t, (cpu_stat.User > 0))
	assert.True(t, (cpu_stat.Sys > 0))

}

func TestCpuPercentage(t *testing.T) {

	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, []string{"*"})
	}
	cpu := CPU{}

	cpu1 := sigar.Cpu{
		User:    10855311,
		Nice:    0,
		Sys:     2021040,
		Idle:    17657874,
		Wait:    0,
		Irq:     0,
		SoftIrq: 0,
		Stolen:  0,
	}

	per := cpu.GetCpuStatEvent(&cpu1)
	assert.Equal(t, per["user_p"], 0.0)
	assert.Equal(t, per["system_p"], 0.0)

	cpu2 := sigar.Cpu{
		User:    10855693,
		Nice:    0,
		Sys:     2021058,
		Idle:    17657876,
		Wait:    0,
		Irq:     0,
		SoftIrq: 0,
		Stolen:  0,
	}

	per = cpu.GetCpuStatEvent(&cpu2)

	assert.Equal(t, per["user_p"], 0.9502)
	assert.Equal(t, per["system_p"], 0.0448)
}
