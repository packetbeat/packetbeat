package protos

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/packetbeat/publish"
)

const (
	DefaultTransactionHashSize                 = 2 ^ 16
	DefaultTransactionExpiration time.Duration = 10 * time.Second
)

// ProtocolData interface to represent an upper
// protocol private data. Used with types like
// HttpStream, MysqlStream, etc.
type ProtocolData interface{}

type Packet struct {
	Ts      time.Time
	Tuple   common.IPPortTuple
	Payload []byte
}

var ErrInvalidPort = errors.New("port number out of range")

// Protocol Plugin Port configuration with validation on init
type PortsConfig struct {
	Ports []int
}

type enhTransactions struct {
	t    publish.Transactions
	meta common.EventMetadata
}

func (p *PortsConfig) Init(ports ...int) error {
	return p.Set(ports)
}

func (p *PortsConfig) Set(ports []int) error {
	if err := validatePorts(ports); err != nil {
		return err
	}
	p.Ports = ports
	return nil
}

func validatePorts(ports []int) error {
	for port := range ports {
		if port < 0 || port > 65535 {
			return ErrInvalidPort
		}
	}
	return nil
}

type Protocols interface {
	BpfFilter(withVlans bool, withICMP bool) string
	GetTCP(proto Protocol) TCPPlugin
	GetUDP(proto Protocol) UDPPlugin
	GetAll() map[Protocol]Plugin
	GetAllTCP() map[Protocol]TCPPlugin
	GetAllUDP() map[Protocol]UDPPlugin
	// Register(proto Protocol, plugin ProtocolPlugin)
}

// list of protocol plugins
type ProtocolsStruct struct {
	all map[Protocol]Plugin
	tcp map[Protocol]TCPPlugin
	udp map[Protocol]UDPPlugin
}

// Singleton of Protocols type.
var Protos = ProtocolsStruct{
	all: map[Protocol]Plugin{},
	tcp: map[Protocol]TCPPlugin{},
	udp: map[Protocol]UDPPlugin{},
}

func (s ProtocolsStruct) Init(
	testMode bool,
	results publish.Transactions,
	configs map[string]*common.Config,
	listConfigs []*common.Config,
) error {
	if len(configs) > 0 {
		logp.Warn("Deprecated: dictionary style protocols configuration has been deprecated. Please use list-style protocols configuration.")
	}

	for proto := range protocolSyms {
		logp.Info("registered protocol plugin: %v", proto)
	}

	for name, config := range configs {
		if err := s.configureProtocol(testMode, results, name, config); err != nil {
			return err
		}
	}

	for _, config := range listConfigs {
		module := struct {
			Name string `config:"type" validate:"required"`
		}{}
		if err := config.Unpack(&module); err != nil {
			return err
		}

		if err := s.configureProtocol(testMode, results, module.Name, config); err != nil {
			return err
		}
	}

	return nil
}

func (s ProtocolsStruct) configureProtocol(
	testMode bool,
	results publish.Transactions,
	name string,
	config *common.Config,
) error {
	// XXX: icmp is special, ignore here :/
	if name == "icmp" {
		return nil
	}

	proto, exists := protocolSyms[name]
	if !exists {
		logp.Err("Unknown protocol plugin: %v", name)
		return nil
	}

	plugin, exists := protocolPlugins[proto]
	if !exists {
		logp.Err("Protocol plugin '%v' not registered (%v).", name, proto.String())
		return nil
	}

	if !config.Enabled() {
		logp.Info("Protocol plugin '%v' disabled by config", name)
		return nil
	}

	meta := struct {
		Config common.EventMetadata `config:",inline"` // Fields and tags to add to events.
	}{}
	if err := config.Unpack(&meta); err != nil {
		return err
	}

	inst, err := plugin(testMode, resultsChannel(results, meta.Config), config)
	if err != nil {
		logp.Err("Failed to register protocol plugin: %v", err)
		return err
	}

	s.register(proto, inst)
	return nil
}

func resultsChannel(
	results publish.Transactions,
	meta common.EventMetadata,
) publish.Transactions {
	if len(meta.Fields) == 0 && len(meta.Tags) == 0 {
		return results
	}
	return &enhTransactions{results, meta}
}

func (s ProtocolsStruct) GetTCP(proto Protocol) TCPPlugin {
	plugin, exists := s.tcp[proto]
	if !exists {
		return nil
	}

	return plugin
}

func (s ProtocolsStruct) GetUDP(proto Protocol) UDPPlugin {
	plugin, exists := s.udp[proto]
	if !exists {
		return nil
	}

	return plugin
}

func (s ProtocolsStruct) GetAll() map[Protocol]Plugin {
	return s.all
}

func (s ProtocolsStruct) GetAllTCP() map[Protocol]TCPPlugin {
	return s.tcp
}

func (s ProtocolsStruct) GetAllUDP() map[Protocol]UDPPlugin {
	return s.udp
}

// BpfFilter returns a Berkeley Packer Filter (BFP) expression that
// will match against packets for the registered protocols. If with_vlans is
// true the filter will match against both IEEE 802.1Q VLAN encapsulated
// and unencapsulated packets
func (s ProtocolsStruct) BpfFilter(withVlans bool, withICMP bool) string {
	// Sort the protocol IDs so that the return value is consistent.
	var protos []int
	for proto := range s.all {
		protos = append(protos, int(proto))
	}
	sort.Ints(protos)

	var expressions []string
	for _, key := range protos {
		proto := Protocol(key)
		plugin := s.all[proto]
		for _, port := range plugin.GetPorts() {
			hasTCP := false
			hasUDP := false

			if _, present := s.tcp[proto]; present {
				hasTCP = true
			}
			if _, present := s.udp[proto]; present {
				hasUDP = true
			}

			var expr string
			if hasTCP && !hasUDP {
				expr = "tcp port %d"
			} else if !hasTCP && hasUDP {
				expr = "udp port %d"
			} else {
				expr = "port %d"
			}

			expressions = append(expressions, fmt.Sprintf(expr, port))
		}
	}

	if withICMP {
		expressions = append(expressions, "icmp", "icmp6")
	}

	filter := strings.Join(expressions, " or ")
	if withVlans {
		filter = fmt.Sprintf("%s or (vlan and (%s))", filter, filter)
	}
	return filter
}

func (s ProtocolsStruct) register(proto Protocol, plugin Plugin) {
	if _, exists := s.all[proto]; exists {
		logp.Warn("Protocol (%s) plugin will overwritten by another plugin", proto.String())
	}

	s.all[proto] = plugin

	success := false
	if tcp, ok := plugin.(TCPPlugin); ok {
		s.tcp[proto] = tcp
		success = true
	}
	if udp, ok := plugin.(UDPPlugin); ok {
		s.udp[proto] = udp
		success = true
	}
	if !success {
		logp.Warn("Protocol (%s) register failed, port: %v", proto.String(), plugin.GetPorts())
	}
}

func (e *enhTransactions) PublishTransaction(event common.MapStr) bool {
	err := common.MergeFields(event, e.meta.Fields, e.meta.FieldsUnderRoot)
	if err != nil {
		logp.Err("Failed to merge fields: %v", err)
		return false
	}

	err = common.AddTags(event, e.meta.Tags)
	if err != nil {
		logp.Err("Failed to add tags: %v", err)
		return false
	}

	return e.t.PublishTransaction(event)
}
