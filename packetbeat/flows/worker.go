package flows

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"net"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/publisher"
)

type flowsProcessor struct {
	spool    spool
	table    *flowMetaTable
	counters *counterReg
	timeout  time.Duration
}

var (
	ErrInvalidTimeout = errors.New("timeout must not >= 1s")
	ErrInvalidPeriod  = errors.New("report period must be -1 or >= 1s")
)

func newFlowsWorker(
	pub publisher.Client,
	table *flowMetaTable,
	counters *counterReg,
	timeout, period time.Duration,
) (*worker, error) {
	oneSecond := 1 * time.Second

	if timeout < oneSecond {
		return nil, ErrInvalidTimeout
	}

	if 0 < period && period < oneSecond {
		return nil, ErrInvalidPeriod
	}

	tickDuration := timeout
	ticksTimeout := 1
	ticksPeriod := -1
	if period > 0 {
		tickDuration = time.Duration(gcd(int64(timeout), int64(period)))
		if tickDuration < oneSecond {
			tickDuration = oneSecond
		}

		ticksTimeout = int(timeout / tickDuration)
		if ticksTimeout == 0 {
			ticksTimeout = 1
		}

		ticksPeriod = int(period / tickDuration)
		if ticksPeriod == 0 {
			ticksPeriod = 1
		}
	}

	debugf("new flows worker. timeout=%v, period=%v, tick=%v, ticksTO=%v, ticksP=%v",
		timeout, period, tickDuration, ticksTimeout, ticksPeriod)

	defaultBatchSize := 1024
	processor := &flowsProcessor{
		table:    table,
		counters: counters,
		timeout:  timeout,
	}
	processor.spool.init(pub, defaultBatchSize)

	return newWorker(func(w *worker) {
		defer w.finished()

		// round time to nearest 10 seconds for alignment
		aligned := time.Unix(((time.Now().Unix()+9)/10)*10, 0)
		waitStart := aligned.Sub(time.Now())

		debugf("worker wait start(%v): %v", aligned, waitStart)
		if cont := w.sleep(waitStart); !cont {
			return
		}

		nTimeout := ticksTimeout
		nPeriod := ticksPeriod
		reportPeriodically := ticksPeriod > 0
		w.periodicaly(tickDuration, func() error {
			nTimeout--
			nPeriod--
			debugf("worker tick, nTimeout=%v, nPeriod=%v", nTimeout, nPeriod)

			handleTimeout := nTimeout == 0
			handleReports := reportPeriodically && nPeriod == 0
			if handleTimeout {
				nTimeout = ticksTimeout
			}
			if handleReports {
				nPeriod = ticksPeriod
			}

			processor.execute(w, handleTimeout, handleReports)
			return nil
		})
	}), nil
}

func (fw *flowsProcessor) execute(w *worker, checkTimeout, handleReports bool) {
	if !checkTimeout && !handleReports {
		return
	}

	debugf("exec tick, timeout=%v, report=%v", checkTimeout, handleReports)

	// get counter names snapshot if reports must be generated
	fw.counters.mutex.Lock()
	intNames := fw.counters.ints.getNames()
	uintNames := fw.counters.uints.getNames()
	floatNames := fw.counters.floats.getNames()
	fw.counters.mutex.Unlock()

	fw.table.Lock()
	defer fw.table.Unlock()
	ts := time.Now()

	// TODO: create snapshot inside flows/tables, so deletion of timedout flows
	//       and reporting flows stats can be done more concurrent to packet
	//       processing.

	for table := fw.table.tables.head; table != nil; table = table.next {
		var next *biFlow
		for flow := table.flows.head; flow != nil; flow = next {
			next = flow.next

			debugf("handle flow: %v, %v", flow.id.flowIDMeta, flow.id.flowID)

			reportFlow := handleReports
			if checkTimeout {
				if ts.Sub(flow.ts) > fw.timeout {
					debugf("kill flow")

					reportFlow = true
					flow.kill() // mark flow as killed
					table.remove(flow)
				}
			}

			if reportFlow {
				debugf("report flow")
				fw.report(w, ts, flow, intNames, uintNames, floatNames)
			}
		}
	}

	fw.spool.flush()
}

func (fw *flowsProcessor) report(
	w *worker,
	ts time.Time,
	flow *biFlow,
	intNames, uintNames, floatNames []string,
) {
	if event := createEvent(ts, flow, intNames, uintNames, floatNames); event != nil {
		debugf("add event: %v", event)
		fw.spool.publish(event)
	}
}

func createEvent(
	ts time.Time, f *biFlow,
	intNames, uintNames, floatNames []string,
) common.MapStr {
	event := common.MapStr{
		"@timestamp": common.Time(ts),
		"type":       "flow",
	}

	// add ethernet layer meta data
	if src, dst, ok := f.id.EthAddr(); ok {
		event["mac_source"] = net.HardwareAddr(src).String()
		event["mac_dest"] = net.HardwareAddr(dst).String()
	}

	// add vlan
	if vlan := f.id.OutterVLan(); vlan != nil {
		event["outter_vlan"] = binary.LittleEndian.Uint16(vlan)
	}
	if vlan := f.id.VLan(); vlan != nil {
		event["vlan"] = binary.LittleEndian.Uint16(vlan)
	}

	// ipv4 layer meta data
	if src, dst, ok := f.id.OutterIPv4Addr(); ok {
		event["outter_ip4_source"] = net.IP(src).String()
		event["outter_ip4_dest"] = net.IP(dst).String()
	}
	if src, dst, ok := f.id.IPv4Addr(); ok {
		event["ip4_source"] = net.IP(src).String()
		event["ip4_dest"] = net.IP(dst).String()
	}

	// ipv6 layer meta data
	if src, dst, ok := f.id.OutterIPv6Addr(); ok {
		event["outter_ip6_source"] = net.IP(src).String()
		event["outter_ip6_dest"] = net.IP(dst).String()
	}
	if src, dst, ok := f.id.IPv6Addr(); ok {
		event["ip6_source"] = net.IP(src).String()
		event["ip6_dest"] = net.IP(dst).String()
	}

	// udp layer meta data
	if src, dst, ok := f.id.UDPAddr(); ok {
		event["port_source"] = binary.LittleEndian.Uint16(src)
		event["port_dest"] = binary.LittleEndian.Uint16(dst)
		event["transp"] = "udp"
	}

	// tcp layer meta data
	if src, dst, ok := f.id.TCPAddr(); ok {
		event["port_source"] = binary.LittleEndian.Uint16(src)
		event["port_dest"] = binary.LittleEndian.Uint16(dst)
		event["transp"] = "tcp"
	}

	if id := f.id.ConnectionID(); id != nil {
		event["connection_id"] = base64.StdEncoding.EncodeToString(id)
	}

	if f.stats[0] != nil {
		event["stats_source"] = encodeStats(f.stats[0], intNames, uintNames, floatNames)
	}
	if f.stats[1] != nil {
		event["stats_dest"] = encodeStats(f.stats[1], intNames, uintNames, floatNames)
	}

	return event
}

func encodeStats(
	stats *flowStats,
	ints, uints, floats []string,
) map[string]interface{} {
	report := make(map[string]interface{})

	for i, name := range ints {
		report[name] = stats.ints[i]
	}

	for i, name := range uints {
		report[name] = stats.uints[i]
	}

	for i, name := range floats {
		report[name] = stats.floats[i]
	}

	return report
}
