// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build linux,386 linux,amd64

package socket

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/x-pack/auditbeat/module/system"
	"github.com/elastic/beats/v7/x-pack/auditbeat/module/system/socket/events"
	"github.com/elastic/beats/v7/x-pack/auditbeat/module/system/socket/guess"
	"github.com/elastic/beats/v7/x-pack/auditbeat/module/system/socket/helper"
	"github.com/elastic/beats/v7/x-pack/auditbeat/module/system/socket/state"
	"github.com/elastic/beats/v7/x-pack/auditbeat/tracing"
	"github.com/elastic/go-perf"
	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/providers/linux"

	"github.com/elastic/beats/v7/x-pack/auditbeat/module/system/socket/dns"
	// Register dns capture implementations
	_ "github.com/elastic/beats/v7/x-pack/auditbeat/module/system/socket/dns/afpacket"
)

const (
	moduleName     = "system"
	metricsetName  = "socket"
	fullName       = moduleName + "/" + metricsetName
	namespace      = "system.audit.socket"
	detailSelector = metricsetName + "detailed"
	auditbeatGroup = "auditbeat"
	// Magic value to detect clock-sync events generated by the metricset.
	clockSyncMagic uint64 = 0x42DEADBEEFABCDEF
)

var (
	kernelVersion string
	eventCount    uint64
)

var defaultMounts = []*mountPoint{
	{fsType: "tracefs", path: "/sys/kernel/tracing"},
	{fsType: "debugfs", path: "/sys/kernel/debug"},
}

// MetricSet for system/socket.
type MetricSet struct {
	system.SystemMetricSet
	templateVars common.MapStr
	config       Config
	log          *logp.Logger
	detailLog    *logp.Logger
	installer    helper.ProbeInstaller
	sniffer      dns.Sniffer
	perfChannel  *tracing.PerfChannel
	mountedFS    *mountPoint
	isDebug      bool
	isDetailed   bool
}

func init() {
	mb.Registry.MustAddMetricSet(moduleName, metricsetName, New,
		mb.DefaultMetricSet(),
		mb.WithNamespace(namespace),
	)
	var err error
	if kernelVersion, err = linux.KernelVersion(); err != nil {
		logp.Err("Failed fetching Linux kernel version: %v", err)
	}
}

// New constructs a new MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The %s dataset is beta.", fullName)

	config := defaultConfig
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, errors.Wrapf(err, "failed to unpack the %s config", fullName)
	}
	logger := logp.NewLogger(metricsetName)
	sniffer, err := dns.NewSniffer(base, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create DNS sniffer")
	}

	ms := &MetricSet{
		SystemMetricSet: system.NewSystemMetricSet(base),
		templateVars:    make(common.MapStr),
		config:          config,
		log:             logger,
		isDebug:         logp.IsDebug(metricsetName),
		detailLog:       logp.NewLogger(detailSelector),
		isDetailed:      logp.HasSelector(detailSelector),
		sniffer:         sniffer,
	}

	// Setup the metricset before Run() so that startup can be halted in case of
	// error.
	if err := ms.Setup(); err != nil {
		return nil, errors.Wrapf(err, "%s dataset setup failed", fullName)
	}
	return ms, nil
}

// Run the metricset. This will loop until the passed reporter is cancelled.
func (m *MetricSet) Run(r mb.PushReporterV2) {
	defer m.log.Infof("%s terminated.", fullName)
	defer m.Cleanup()

	st := state.NewState(r,
		m.log,
		m.config.FlowInactiveTimeout,
		m.config.SocketInactiveTimeout,
		m.config.FlowTerminationTimeout,
		m.config.ClockMaxDrift)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := m.sniffer.Monitor(ctx, func(tr dns.Transaction) {
		if err := st.OnDNSTransaction(tr); err != nil {
			m.log.Errorf("Unable to store DNS transaction %+v: %v", tr, err)
		}
	}); err != nil {
		err = errors.Wrap(err, "unable to start DNS sniffer")
		r.Error(err)
		m.log.Error(err)
		return
	}

	if err := m.perfChannel.Run(); err != nil {
		err = errors.Wrap(err, "unable to start perf channel")
		r.Error(err)
		m.log.Error(err)
		return
	}
	// Launch the clock-synchronization ticker.
	go m.clockSyncLoop(m.config.ClockSyncPeriod, r.Done())

	if procs, err := sysinfo.Processes(); err != nil {
		m.log.Error("Failed to bootstrap process table using /proc", err)
	} else {
		for _, p := range procs {
			if i, err := p.Info(); err == nil {
				process := state.CreateProcess(uint32(i.PID), i.Exe, i.Name, 0, i.Args).SetCreated(i.StartTime)

				if user, err := p.User(); err == nil {
					toUint32 := func(id string) uint32 {
						num, _ := strconv.Atoi(id)
						return uint32(num)
					}
					process.SetCreds(toUint32(user.UID), toUint32(user.GID), toUint32(user.EUID), toUint32(user.EGID))
				}

				st.ProcessStart(process)
			}
		}
		m.log.Info("Bootstrapped process table using /proc")
	}

	m.log.Infof("%s dataset is running.", fullName)
	// Dispatch loop.
	for running := true; running; {
		select {
		case <-r.Done():
			running = false

		case iface, ok := <-m.perfChannel.C():
			if !ok {
				running = false
				break
			}
			v, ok := iface.(events.Event)
			if !ok {
				m.log.Errorf("Received an event of wrong type: %T", iface)
				continue
			}
			if m.isDetailed {
				m.detailLog.Debug(v.String())
			}
			v.Update(st)
			atomic.AddUint64(&eventCount, 1)

		case err := <-m.perfChannel.ErrC():
			m.log.Errorf("Error received from perf channel: %v", err)
			running = false

		case numLost := <-m.perfChannel.LostC():
			if numLost != ^uint64(0) {
				m.log.Warnf("Lost %d events", numLost)
			} else {
				m.log.Warn("Lost the whole ringbuffer")
			}
		}
	}
}

// Setup performs all the initialisations required for KProbes monitoring.
func (m *MetricSet) Setup() (err error) {
	m.log.Infof("Setting up %s for kernel %s", fullName, kernelVersion)

	//
	// Validate that tracefs / debugfs is present and kprobes are available
	//
	var traceFS *tracing.TraceFS
	if m.config.TraceFSPath == nil {

		if err := tracing.IsTraceFSAvailable(); err != nil {
			m.log.Debugf("tracefs/debugfs not found. Attempting to mount")
			for _, mount := range defaultMounts {
				if err = mount.mount(); err != nil {
					m.log.Debugf("Mount %s returned %v", mount, err)
					continue
				}
				if tracing.IsTraceFSAvailable() != nil {
					m.log.Warnf("Mounted %s but no kprobes available", mount, err)
					mount.unmount()
					continue
				}
				m.log.Debugf("Mounted %s", mount)
				m.mountedFS = mount
				break
			}
		}
		traceFS, err = tracing.NewTraceFS()
	} else {
		traceFS, err = tracing.NewTraceFSWithPath(*m.config.TraceFSPath)
	}
	if err != nil {
		return errors.Wrap(err, "tracefs/debugfs is not mounted or not writeable")
	}

	//
	// Setup initial template variables
	//
	m.templateVars.Update(baseTemplateVars)
	m.templateVars.Update(archVariables)

	//
	// Detect IPv6 support
	//

	hasIPv6, err := detectIPv6()
	if err != nil {
		m.log.Debugf("Error detecting IPv6 support: %v", err)
		hasIPv6 = false
	}
	m.log.Debugf("IPv6 supported: %v", hasIPv6)
	if m.config.EnableIPv6 != nil {
		if *m.config.EnableIPv6 && !hasIPv6 {
			return errors.New("requested IPv6 support but IPv6 is disabled in the system")
		}
		hasIPv6 = *m.config.EnableIPv6
	}
	m.log.Debugf("IPv6 enabled: %v", hasIPv6)
	m.templateVars["HAS_IPV6"] = hasIPv6

	//
	// Create probe installer
	//
	extra := WithNoOp()
	if m.config.DevelopmentMode {
		extra = WithFilterPort(22)
	}
	m.installer = newProbeInstaller(traceFS,
		WithGroup(auditbeatGroup),
		WithTemplates(m.templateVars),
		extra)
	defer func() {
		if err != nil {
			m.installer.UninstallInstalled()
		}
	}()

	//
	// remove existing KProbes from Auditbeat
	//
	if err = m.installer.UninstallIf(isOwnProbe); err != nil {
		return errors.Wrap(err, "unable to delete existing KProbes. Is Auditbeat already running?")
	}

	//
	// Load available kernel functions for tracing
	//
	functions, err := LoadTracingFunctions(traceFS)
	if err != nil {
		m.log.Debugf("Can't load available_tracing_functions. Using alternative. err=%v", err)
	}

	//
	// Resolve function names from alternatives
	//
	for varName, alternatives := range functionAlternatives {
		if exists, _ := m.templateVars.HasKey(varName); exists {
			return fmt.Errorf("variable %s overwrites existing key", varName)
		}
		found := false
		var selected string
		for _, selected = range alternatives {
			if found = m.isKernelFunctionAvailable(selected, functions); found {
				break
			}
		}
		if !found {
			return fmt.Errorf("none of the required functions for %s is found. One of %v is required", varName, alternatives)
		}
		if m.isDebug {
			m.log.Debugf("Selected kernel function %s for %s", selected, varName)
		}
		m.templateVars[varName] = selected
	}

	//
	// Make sure all the required kernel functions are available
	//
	for _, probeDef := range getKProbes(hasIPv6) {
		probeDef = probeDef.ApplyTemplate(m.templateVars)
		name := probeDef.Probe.Address
		if !m.isKernelFunctionAvailable(name, functions) {
			return fmt.Errorf("required function '%s' is not available for tracing in the current kernel (%s)", name, kernelVersion)
		}
	}

	//
	// Guess all the required parameters
	//
	if err = guess.GuessAll(m.installer,
		guess.Context{
			Log:     m.log,
			Vars:    m.templateVars,
			Timeout: m.config.GuessTimeout,
		}); err != nil {
		return errors.Wrap(err, "unable to guess one or more required parameters")
	}

	if m.isDebug {
		names := make([]string, 0, len(m.templateVars))
		for name := range m.templateVars {
			names = append(names, name)
		}
		sort.Strings(names)
		m.log.Debugf("%d template variables in use:", len(m.templateVars))
		for _, key := range names {
			m.log.Debugf("  %s = %v", key, m.templateVars[key])
		}
	}

	//
	// Create perf channel
	//
	m.perfChannel, err = tracing.NewPerfChannel(
		tracing.WithBufferSize(m.config.PerfQueueSize),
		tracing.WithErrBufferSize(m.config.ErrQueueSize),
		tracing.WithLostBufferSize(m.config.LostQueueSize),
		tracing.WithRingSizeExponent(m.config.RingSizeExp),
		tracing.WithTID(perf.AllThreads),
		tracing.WithTimestamp())
	if err != nil {
		return errors.Wrapf(err, "unable to create perf channel")
	}

	//
	// Register Kprobes
	//
	for _, probeDef := range getKProbes(hasIPv6) {
		format, decoder, err := m.installer.Install(probeDef)
		if err != nil {
			return errors.Wrapf(err, "unable to register probe %s", probeDef.Probe.String())
		}
		if err = m.perfChannel.MonitorProbe(format, decoder); err != nil {
			return errors.Wrapf(err, "unable to monitor probe %s", probeDef.Probe.String())
		}
	}
	return nil
}

// Cleanup must be called so that kprobes are not left around after exit.
func (m *MetricSet) Cleanup() {
	if m.perfChannel != nil {
		if err := m.perfChannel.Close(); err != nil {
			m.log.Warnf("Failed to close perf channel on exit: %v", err)
		}
	}
	if m.installer != nil {
		if err := m.installer.UninstallIf(isOwnProbe); err != nil {
			m.log.Warnf("Failed to remove KProbes on exit: %v", err)
		}
	}
	if m.mountedFS != nil {
		if err := m.mountedFS.unmount(); err != nil {
			m.log.Errorf("Failed to umount %s: %v", m.mountedFS, err)
		} else {
			m.log.Debugf("Unmounted %s", m.mountedFS)
		}
	}
}

func (m *MetricSet) clockSyncLoop(interval time.Duration, done <-chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	triggerClockSync()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			triggerClockSync()
		}
	}
}

func (m *MetricSet) isKernelFunctionAvailable(name string, tracingFns common.StringSet) bool {
	if tracingFns.Count() != 0 {
		return tracingFns.Has(name)
	}
	defer m.installer.UninstallInstalled()
	checkProbe := helper.ProbeDef{
		Probe: tracing.Probe{
			Name:      "check_" + name,
			Address:   name,
			Fetchargs: "%ax:u64", // dump decoder needs it.
		},
		Decoder: tracing.NewDumpDecoder,
	}
	_, _, err := m.installer.Install(checkProbe)
	return err == nil
}

func triggerClockSync() {
	// This generates a uname (SYS_UNAME) syscall event that contains
	// clockSyncMagic at the first 8 bytes of the passed buffer and
	// the current UNIX nano timestamp at the following 8 bytes.
	//
	// The magic bytes are used to filter-out legitimate uname() calls
	// from this process and the timestamp is used as a reference point for
	// synchronization with the internal clock that the kernel uses for stamping
	// the tracing events it produces.
	var buf unix.Utsname
	tracing.MachineEndian.PutUint64(buf.Sysname[:], clockSyncMagic)
	tracing.MachineEndian.PutUint64(buf.Sysname[8:], uint64(time.Now().UnixNano()))
	unix.Uname(&buf)
}

func isOwnProbe(probe tracing.Probe) bool {
	return probe.Group == auditbeatGroup
}

type mountPoint struct {
	fsType string
	path   string
}

func (m mountPoint) mount() error {
	return unix.Mount(m.fsType, m.path, m.fsType, 0, "")
}

func (m mountPoint) unmount() error {
	return syscall.Unmount(m.path, 0)
}

func (m *mountPoint) String() string {
	return m.fsType + " at " + m.path
}

func detectIPv6() (bool, error) {
	// Check that AF_INET6 is available.
	// This fails when the kernel is booted with ipv6.disable=1
	fd, err := unix.Socket(unix.AF_INET6, unix.SOCK_DGRAM, 0)
	if err != nil {
		return false, nil
	}
	unix.Close(fd)
	loopback, err := helper.NewIPv6Loopback()
	if err != nil {
		return false, err
	}
	defer loopback.Cleanup()
	_, err = loopback.AddRandomAddress()
	// Assume that all failures for Add..() are caused by missing IPv6 support.
	return err == nil, nil
}
