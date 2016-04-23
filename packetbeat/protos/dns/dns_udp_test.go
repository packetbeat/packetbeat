// +build !integration

// Unit tests and benchmarks for the dns package.
//
// The byte array test data was generated from pcap files using the gopacket
// test_creator.py script contained in the gopacket repository. The script was
// modified to drop the Ethernet, IP, and UDP headers from the byte arrays
// (skip the first 42 bytes).
//
// TODO:
//   * Add test validation for responsetime to make sure unit conversion
//     is being done correctly.
//   * Add test case to verify that Include_authorities and Include_additionals
//     are working.
//   * Add test case for Send_request and validate the stringified DNS message.
//   * Add test case for Send_response and validate the stringified DNS message.

package dns

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/elastic/beats/packetbeat/protos"
	"github.com/elastic/beats/packetbeat/publish"

	"github.com/elastic/beats/libbeat/common"

	mkdns "github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
)

// Verify that the interface for UDP has been satisfied.
var _ protos.UdpPlugin = &Dns{}

// DNS messages for testing. When adding a new test message, add it to the
// messages array and create a new benchmark test for the message.
var (
	// An array of all test messages.
	messages = []DnsTestMessage{
		elasticA,
		zoneIxfr,
		githubPtr,
		sophosTxt,
		ednsSecA,
	}

	elasticA = DnsTestMessage{
		id:      8529,
		opcode:  "QUERY",
		flags:   []string{"rd", "ra"},
		rcode:   "NOERROR",
		q_class: "IN",
		q_type:  "A",
		q_name:  "elastic.co.",
		q_etld:  "elastic.co.",
		answers: []string{"54.148.130.30", "54.69.104.66"},
		request: []byte{
			0x21, 0x51, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x65, 0x6c, 0x61,
			0x73, 0x74, 0x69, 0x63, 0x02, 0x63, 0x6f, 0x00, 0x00, 0x01, 0x00, 0x01,
		},
		response: []byte{
			0x21, 0x51, 0x81, 0x80, 0x00, 0x01, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x07, 0x65, 0x6c, 0x61,
			0x73, 0x74, 0x69, 0x63, 0x02, 0x63, 0x6f, 0x00, 0x00, 0x01, 0x00, 0x01, 0xc0, 0x0c, 0x00, 0x01,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x39, 0x00, 0x04, 0x36, 0x94, 0x82, 0x1e, 0xc0, 0x0c, 0x00, 0x01,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x39, 0x00, 0x04, 0x36, 0x45, 0x68, 0x42,
		},
	}

	zoneIxfr = DnsTestMessage{
		id:      16384,
		opcode:  "QUERY",
		flags:   []string{"ra"},
		rcode:   "NOERROR",
		q_class: "IN",
		q_type:  "IXFR",
		q_name:  "etas.com.",
		q_etld:  "etas.com.",
		answers: []string{"training2003p.", "training2003p.", "training2003p.",
			"training2003p.", "1.1.1.100"},
		request: []byte{
			0x40, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x65, 0x74, 0x61,
			0x73, 0x03, 0x63, 0x6f, 0x6d, 0x00, 0x00, 0xfb, 0x00, 0x01, 0xc0, 0x0c, 0x00, 0x06, 0x00, 0x01,
			0x00, 0x00, 0x0e, 0x10, 0x00, 0x2f, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x32,
			0x30, 0x30, 0x33, 0x70, 0x00, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
			0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x00, 0x02, 0x58, 0x00, 0x01, 0x51,
			0x80, 0x00, 0x00, 0x0e, 0x10, 0x4d, 0x53,
		},
		response: []byte{
			0x40, 0x00, 0x80, 0x80, 0x00, 0x01, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x04, 0x65, 0x74, 0x61,
			0x73, 0x03, 0x63, 0x6f, 0x6d, 0x00, 0x00, 0xfb, 0x00, 0x01, 0xc0, 0x0c, 0x00, 0x06, 0x00, 0x01,
			0x00, 0x00, 0x0e, 0x10, 0x00, 0x2f, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x32,
			0x30, 0x30, 0x33, 0x70, 0x00, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
			0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x00, 0x02, 0x58, 0x00, 0x01, 0x51,
			0x80, 0x00, 0x00, 0x0e, 0x10, 0xc0, 0x0c, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x0e, 0x10, 0x00,
			0x18, 0xc0, 0x26, 0xc0, 0x35, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x00, 0x02,
			0x58, 0x00, 0x01, 0x51, 0x80, 0x00, 0x00, 0x0e, 0x10, 0xc0, 0x0c, 0x00, 0x06, 0x00, 0x01, 0x00,
			0x00, 0x0e, 0x10, 0x00, 0x18, 0xc0, 0x26, 0xc0, 0x35, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00,
			0x3c, 0x00, 0x00, 0x02, 0x58, 0x00, 0x01, 0x51, 0x80, 0x00, 0x00, 0x0e, 0x10, 0x05, 0x69, 0x6e,
			0x64, 0x65, 0x78, 0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x0e, 0x10, 0x00, 0x04, 0x01,
			0x01, 0x01, 0x64, 0xc0, 0x0c, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x0e, 0x10, 0x00, 0x18, 0xc0,
			0x26, 0xc0, 0x35, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x00, 0x02, 0x58, 0x00,
			0x01, 0x51, 0x80, 0x00, 0x00, 0x0e, 0x10,
		},
	}

	githubPtr = DnsTestMessage{
		id:      344,
		opcode:  "QUERY",
		flags:   []string{"rd", "ra"},
		rcode:   "NOERROR",
		q_class: "IN",
		q_type:  "PTR",
		q_name:  "131.252.30.192.in-addr.arpa.",
		q_etld:  "192.in-addr.arpa.",
		answers: []string{"github.com."},
		authorities: []string{"a.root-servers.net.", "b.root-servers.net.", "c.root-servers.net.",
			"d.root-servers.net.", "e.root-servers.net.", "f.root-servers.net.", "g.root-servers.net.",
			"h.root-servers.net.", "i.root-servers.net.", "j.root-servers.net.", "k.root-servers.net.",
			"l.root-servers.net.", "m.root-servers.net."},
		request: []byte{
			0x01, 0x58, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x31, 0x33, 0x31,
			0x03, 0x32, 0x35, 0x32, 0x02, 0x33, 0x30, 0x03, 0x31, 0x39, 0x32, 0x07, 0x69, 0x6e, 0x2d, 0x61,
			0x64, 0x64, 0x72, 0x04, 0x61, 0x72, 0x70, 0x61, 0x00, 0x00, 0x0c, 0x00, 0x01,
		},
		response: []byte{
			0x01, 0x58, 0x81, 0x80, 0x00, 0x01, 0x00, 0x01, 0x00, 0x0d, 0x00, 0x00, 0x03, 0x31, 0x33, 0x31,
			0x03, 0x32, 0x35, 0x32, 0x02, 0x33, 0x30, 0x03, 0x31, 0x39, 0x32, 0x07, 0x69, 0x6e, 0x2d, 0x61,
			0x64, 0x64, 0x72, 0x04, 0x61, 0x72, 0x70, 0x61, 0x00, 0x00, 0x0c, 0x00, 0x01, 0xc0, 0x0c, 0x00,
			0x0c, 0x00, 0x01, 0x00, 0x00, 0x09, 0xe2, 0x00, 0x0c, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
			0x03, 0x63, 0x6f, 0x6d, 0x00, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x14,
			0x01, 0x6c, 0x0c, 0x72, 0x6f, 0x6f, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x03,
			0x6e, 0x65, 0x74, 0x00, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01,
			0x65, 0xc0, 0x52, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x63,
			0xc0, 0x52, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x62, 0xc0,
			0x52, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x61, 0xc0, 0x52,
			0x00, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x68, 0xc0, 0x52, 0x00,
			0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x66, 0xc0, 0x52, 0x00, 0x00,
			0x02, 0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x69, 0xc0, 0x52, 0x00, 0x00, 0x02,
			0x00, 0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x67, 0xc0, 0x52, 0x00, 0x00, 0x02, 0x00,
			0x01, 0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x6d, 0xc0, 0x52, 0x00, 0x00, 0x02, 0x00, 0x01,
			0x00, 0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x64, 0xc0, 0x52, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00,
			0x00, 0x07, 0xb8, 0x00, 0x04, 0x01, 0x6a, 0xc0, 0x52, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00,
			0x07, 0xb8, 0x00, 0x04, 0x01, 0x6b, 0xc0, 0x52,
		},
	}

	sophosTxt = DnsTestMessage{
		id:      8238,
		opcode:  "QUERY",
		flags:   []string{"rd", "ra"},
		rcode:   "NXDOMAIN",
		q_class: "IN",
		q_type:  "TXT",
		q_name: "3.1o19ss00s2s17s4qp375sp49r830n2n4n923s8839052s7p7768s53365226pp3.659p1r741os37393" +
			"648s2348o762q1066q53rq5p4614r1q4781qpr16n809qp4.879o3o734q9sns005o3pp76q83.2q65qns3spns" +
			"1081s5rn5sr74opqrqnpq6rn3ro5.i.00.mac.sophosxl.net.",
		q_etld: "sophosxl.net.",
		request: []byte{
			0x20, 0x2e, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x33, 0x3f, 0x31,
			0x6f, 0x31, 0x39, 0x73, 0x73, 0x30, 0x30, 0x73, 0x32, 0x73, 0x31, 0x37, 0x73, 0x34, 0x71, 0x70,
			0x33, 0x37, 0x35, 0x73, 0x70, 0x34, 0x39, 0x72, 0x38, 0x33, 0x30, 0x6e, 0x32, 0x6e, 0x34, 0x6e,
			0x39, 0x32, 0x33, 0x73, 0x38, 0x38, 0x33, 0x39, 0x30, 0x35, 0x32, 0x73, 0x37, 0x70, 0x37, 0x37,
			0x36, 0x38, 0x73, 0x35, 0x33, 0x33, 0x36, 0x35, 0x32, 0x32, 0x36, 0x70, 0x70, 0x33, 0x3f, 0x36,
			0x35, 0x39, 0x70, 0x31, 0x72, 0x37, 0x34, 0x31, 0x6f, 0x73, 0x33, 0x37, 0x33, 0x39, 0x33, 0x36,
			0x34, 0x38, 0x73, 0x32, 0x33, 0x34, 0x38, 0x6f, 0x37, 0x36, 0x32, 0x71, 0x31, 0x30, 0x36, 0x36,
			0x71, 0x35, 0x33, 0x72, 0x71, 0x35, 0x70, 0x34, 0x36, 0x31, 0x34, 0x72, 0x31, 0x71, 0x34, 0x37,
			0x38, 0x31, 0x71, 0x70, 0x72, 0x31, 0x36, 0x6e, 0x38, 0x30, 0x39, 0x71, 0x70, 0x34, 0x1a, 0x38,
			0x37, 0x39, 0x6f, 0x33, 0x6f, 0x37, 0x33, 0x34, 0x71, 0x39, 0x73, 0x6e, 0x73, 0x30, 0x30, 0x35,
			0x6f, 0x33, 0x70, 0x70, 0x37, 0x36, 0x71, 0x38, 0x33, 0x28, 0x32, 0x71, 0x36, 0x35, 0x71, 0x6e,
			0x73, 0x33, 0x73, 0x70, 0x6e, 0x73, 0x31, 0x30, 0x38, 0x31, 0x73, 0x35, 0x72, 0x6e, 0x35, 0x73,
			0x72, 0x37, 0x34, 0x6f, 0x70, 0x71, 0x72, 0x71, 0x6e, 0x70, 0x71, 0x36, 0x72, 0x6e, 0x33, 0x72,
			0x6f, 0x35, 0x01, 0x69, 0x02, 0x30, 0x30, 0x03, 0x6d, 0x61, 0x63, 0x08, 0x73, 0x6f, 0x70, 0x68,
			0x6f, 0x73, 0x78, 0x6c, 0x03, 0x6e, 0x65, 0x74, 0x00, 0x00, 0x10, 0x00, 0x01,
		},
		response: []byte{
			0x20, 0x2e, 0x81, 0x83, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x33, 0x3f, 0x31,
			0x6f, 0x31, 0x39, 0x73, 0x73, 0x30, 0x30, 0x73, 0x32, 0x73, 0x31, 0x37, 0x73, 0x34, 0x71, 0x70,
			0x33, 0x37, 0x35, 0x73, 0x70, 0x34, 0x39, 0x72, 0x38, 0x33, 0x30, 0x6e, 0x32, 0x6e, 0x34, 0x6e,
			0x39, 0x32, 0x33, 0x73, 0x38, 0x38, 0x33, 0x39, 0x30, 0x35, 0x32, 0x73, 0x37, 0x70, 0x37, 0x37,
			0x36, 0x38, 0x73, 0x35, 0x33, 0x33, 0x36, 0x35, 0x32, 0x32, 0x36, 0x70, 0x70, 0x33, 0x3f, 0x36,
			0x35, 0x39, 0x70, 0x31, 0x72, 0x37, 0x34, 0x31, 0x6f, 0x73, 0x33, 0x37, 0x33, 0x39, 0x33, 0x36,
			0x34, 0x38, 0x73, 0x32, 0x33, 0x34, 0x38, 0x6f, 0x37, 0x36, 0x32, 0x71, 0x31, 0x30, 0x36, 0x36,
			0x71, 0x35, 0x33, 0x72, 0x71, 0x35, 0x70, 0x34, 0x36, 0x31, 0x34, 0x72, 0x31, 0x71, 0x34, 0x37,
			0x38, 0x31, 0x71, 0x70, 0x72, 0x31, 0x36, 0x6e, 0x38, 0x30, 0x39, 0x71, 0x70, 0x34, 0x1a, 0x38,
			0x37, 0x39, 0x6f, 0x33, 0x6f, 0x37, 0x33, 0x34, 0x71, 0x39, 0x73, 0x6e, 0x73, 0x30, 0x30, 0x35,
			0x6f, 0x33, 0x70, 0x70, 0x37, 0x36, 0x71, 0x38, 0x33, 0x28, 0x32, 0x71, 0x36, 0x35, 0x71, 0x6e,
			0x73, 0x33, 0x73, 0x70, 0x6e, 0x73, 0x31, 0x30, 0x38, 0x31, 0x73, 0x35, 0x72, 0x6e, 0x35, 0x73,
			0x72, 0x37, 0x34, 0x6f, 0x70, 0x71, 0x72, 0x71, 0x6e, 0x70, 0x71, 0x36, 0x72, 0x6e, 0x33, 0x72,
			0x6f, 0x35, 0x01, 0x69, 0x02, 0x30, 0x30, 0x03, 0x6d, 0x61, 0x63, 0x08, 0x73, 0x6f, 0x70, 0x68,
			0x6f, 0x73, 0x78, 0x6c, 0x03, 0x6e, 0x65, 0x74, 0x00, 0x00, 0x10, 0x00, 0x01,
		},
	}

	ednsSecA = DnsTestMessage{
		id:      20498,
		opcode:  "QUERY",
		flags:   []string{"rd", "ad", "ra"},
		rcode:   "NOERROR",
		q_class: "IN",
		q_type:  "A",
		q_name:  "www.ietf.org.",
		q_etld:  "ietf.org.",
		answers: []string{"64.170.98.30", "iDA8bJnrAEz3jgYnyFRm567a76qlv1V0CqxOSd/o9nvnN0GlZLaVoDmuXpaIaoypbGxwzwgK/LY6CV2k6SWKwicBmpENL26hwyjkFzPDW8kX3ibFhtfsOb8pYe7nBj326actp/7iG+DRuDmPnkYBja+wDYk61doTtkqZg57fn3iS97tjNPCC9C9knRAuDYUG+dVxalazSwYrpvY97dUC1H2spD0g4UdDyCbGA46mouZ4GPzNMewgf948qxrnU8pWPk3nQW5TgLVkGoWgco2owfLElBqf6rJ4LOswuhaw8IpTtmw3FsixxTLQvKOE5nftd1nMhQQd9CaHjoKNAUEz5Q=="},
		request: []byte{
			0x50, 0x12, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x03, 0x77, 0x77, 0x77,
			0x04, 0x69, 0x65, 0x74, 0x66, 0x03, 0x6f, 0x72, 0x67, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00,
			0x29, 0x10, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00,
		},
		response: []byte{
			0x50, 0x12, 0x81, 0xa0, 0x00, 0x01, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x03, 0x77, 0x77, 0x77,
			0x04, 0x69, 0x65, 0x74, 0x66, 0x03, 0x6f, 0x72, 0x67, 0x00, 0x00, 0x01, 0x00, 0x01, 0xc0, 0x0c,
			0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x04, 0xd8, 0x00, 0x04, 0x40, 0xaa, 0x62, 0x1e, 0xc0, 0x0c,
			0x00, 0x2e, 0x00, 0x01, 0x00, 0x00, 0x04, 0xd8, 0x01, 0x1c, 0x00, 0x01, 0x05, 0x03, 0x00, 0x00,
			0x07, 0x08, 0x52, 0xd0, 0x71, 0xc7, 0x50, 0xef, 0x30, 0xaf, 0x9e, 0x04, 0x04, 0x69, 0x65, 0x74,
			0x66, 0x03, 0x6f, 0x72, 0x67, 0x00, 0x88, 0x30, 0x3c, 0x6c, 0x99, 0xeb, 0x00, 0x4c, 0xf7, 0x8e,
			0x06, 0x27, 0xc8, 0x54, 0x66, 0xe7, 0xae, 0xda, 0xef, 0xaa, 0xa5, 0xbf, 0x55, 0x74, 0x0a, 0xac,
			0x4e, 0x49, 0xdf, 0xe8, 0xf6, 0x7b, 0xe7, 0x37, 0x41, 0xa5, 0x64, 0xb6, 0x95, 0xa0, 0x39, 0xae,
			0x5e, 0x96, 0x88, 0x6a, 0x8c, 0xa9, 0x6c, 0x6c, 0x70, 0xcf, 0x08, 0x0a, 0xfc, 0xb6, 0x3a, 0x09,
			0x5d, 0xa4, 0xe9, 0x25, 0x8a, 0xc2, 0x27, 0x01, 0x9a, 0x91, 0x0d, 0x2f, 0x6e, 0xa1, 0xc3, 0x28,
			0xe4, 0x17, 0x33, 0xc3, 0x5b, 0xc9, 0x17, 0xde, 0x26, 0xc5, 0x86, 0xd7, 0xec, 0x39, 0xbf, 0x29,
			0x61, 0xee, 0xe7, 0x06, 0x3d, 0xf6, 0xe9, 0xa7, 0x2d, 0xa7, 0xfe, 0xe2, 0x1b, 0xe0, 0xd1, 0xb8,
			0x39, 0x8f, 0x9e, 0x46, 0x01, 0x8d, 0xaf, 0xb0, 0x0d, 0x89, 0x3a, 0xd5, 0xda, 0x13, 0xb6, 0x4a,
			0x99, 0x83, 0x9e, 0xdf, 0x9f, 0x78, 0x92, 0xf7, 0xbb, 0x63, 0x34, 0xf0, 0x82, 0xf4, 0x2f, 0x64,
			0x9d, 0x10, 0x2e, 0x0d, 0x85, 0x06, 0xf9, 0xd5, 0x71, 0x6a, 0x56, 0xb3, 0x4b, 0x06, 0x2b, 0xa6,
			0xf6, 0x3d, 0xed, 0xd5, 0x02, 0xd4, 0x7d, 0xac, 0xa4, 0x3d, 0x20, 0xe1, 0x47, 0x43, 0xc8, 0x26,
			0xc6, 0x03, 0x8e, 0xa6, 0xa2, 0xe6, 0x78, 0x18, 0xfc, 0xcd, 0x31, 0xec, 0x20, 0x7f, 0xde, 0x3c,
			0xab, 0x1a, 0xe7, 0x53, 0xca, 0x56, 0x3e, 0x4d, 0xe7, 0x41, 0x6e, 0x53, 0x80, 0xb5, 0x64, 0x1a,
			0x85, 0xa0, 0x72, 0x8d, 0xa8, 0xc1, 0xf2, 0xc4, 0x94, 0x1a, 0x9f, 0xea, 0xb2, 0x78, 0x2c, 0xeb,
			0x30, 0xba, 0x16, 0xb0, 0xf0, 0x8a, 0x53, 0xb6, 0x6c, 0x37, 0x16, 0xc8, 0xb1, 0xc5, 0x32, 0xd0,
			0xbc, 0xa3, 0x84, 0xe6, 0x77, 0xed, 0x77, 0x59, 0xcc, 0x85, 0x04, 0x1d, 0xf4, 0x26, 0x87, 0x8e,
			0x82, 0x8d, 0x01, 0x41, 0x33, 0xe5, 0x00, 0x00, 0x29, 0x0f, 0xa0, 0x00, 0x00, 0x80, 0x00, 0x00,
			0x00,
		},
	}
)

// Verify that an empty packet is safely handled (no panics).
func TestParseUdp_emptyPacket(t *testing.T) {
	dns := newDns(testing.Verbose())
	packet := newPacket(forward, []byte{})
	dns.ParseUdp(packet)
	assert.Empty(t, dns.transactions.Size(), "There should be no transactions.")
	client := dns.results.(*publish.ChanTransactions)
	close(client.Channel)
	assert.Nil(t, <-client.Channel, "No result should have been published.")
}

// Verify that a malformed packet is safely handled (no panics).
func TestParseUdp_malformedPacket(t *testing.T) {
	dns := newDns(testing.Verbose())
	garbage := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	packet := newPacket(forward, garbage)
	dns.ParseUdp(packet)
	assert.Empty(t, dns.transactions.Size(), "There should be no transactions.")

	// As a future addition, a malformed message should publish a result.
}

// Verify that the lone request packet is parsed.
func TestParseUdp_requestPacket(t *testing.T) {
	dns := newDns(testing.Verbose())
	packet := newPacket(forward, elasticA.request)
	dns.ParseUdp(packet)
	assert.Equal(t, 1, dns.transactions.Size(), "There should be one transaction.")
	client := dns.results.(*publish.ChanTransactions)
	close(client.Channel)
	assert.Nil(t, <-client.Channel, "No result should have been published.")
}

// Verify that the lone response packet is parsed and that an error
// result is published.
func TestParseUdp_responseOnly(t *testing.T) {
	dns := newDns(testing.Verbose())
	q := elasticA
	packet := newPacket(reverse, q.response)
	dns.ParseUdp(packet)

	m := expectResult(t, dns)
	assert.Equal(t, "udp", mapValue(t, m, "transport"))
	assert.Nil(t, mapValue(t, m, "bytes_in"))
	assert.Equal(t, len(q.response), mapValue(t, m, "bytes_out"))
	assert.Nil(t, mapValue(t, m, "responsetime"))
	assert.Equal(t, common.ERROR_STATUS, mapValue(t, m, "status"))
	assert.Equal(t, OrphanedResponse.Error(), mapValue(t, m, "notes"))
	assertMapStrData(t, m, q)
}

// Verify that the first request is published without a response and that
// the status is error. This second packet will remain in the transaction
// map awaiting a response.
func TestParseUdp_duplicateRequests(t *testing.T) {
	dns := newDns(testing.Verbose())
	q := elasticA
	packet := newPacket(forward, q.request)
	dns.ParseUdp(packet)
	assert.Equal(t, 1, dns.transactions.Size(), "There should be one transaction.")
	packet = newPacket(forward, q.request)
	dns.ParseUdp(packet)
	assert.Equal(t, 1, dns.transactions.Size(), "There should be one transaction.")

	m := expectResult(t, dns)
	assert.Equal(t, "udp", mapValue(t, m, "transport"))
	assert.Equal(t, len(q.request), mapValue(t, m, "bytes_in"))
	assert.Nil(t, mapValue(t, m, "bytes_out"))
	assert.Nil(t, mapValue(t, m, "responsetime"))
	assert.Equal(t, common.ERROR_STATUS, mapValue(t, m, "status"))
	assert.Equal(t, DuplicateQueryMsg.Error(), mapValue(t, m, "notes"))
}

// Verify that the request/response pair are parsed and that a result
// is published.
func TestParseUdp_requestResponse(t *testing.T) {
	parseUdpRequestResponse(t, newDns(testing.Verbose()), elasticA)
}

// Verify all DNS test messages are parsed correctly.
func TestParseUdp_allTestMessages(t *testing.T) {
	dns := newDns(testing.Verbose())
	for _, q := range messages {
		t.Logf("Testing with query for %s", q.q_name)
		parseUdpRequestResponse(t, dns, q)
	}
}

// Verify that expireTransaction publishes an event with an error status
// and note.
func TestExpireTransaction(t *testing.T) {
	dns := newDns(testing.Verbose())

	trans := newTransaction(time.Now(), DnsTuple{}, common.CmdlineTuple{})
	trans.Request = &DnsMessage{
		Data: &mkdns.Msg{
			Question: []mkdns.Question{{}},
		},
	}
	dns.expireTransaction(trans)

	m := expectResult(t, dns)
	assert.Nil(t, mapValue(t, m, "bytes_out"))
	assert.Nil(t, mapValue(t, m, "responsetime"))
	assert.Equal(t, common.ERROR_STATUS, mapValue(t, m, "status"))
	assert.Equal(t, NoResponse.Error(), mapValue(t, m, "notes"))
}

// Verify that an empty DNS request packet can be published.
func TestPublishTransaction_emptyDnsRequest(t *testing.T) {
	dns := newDns(testing.Verbose())

	trans := newTransaction(time.Now(), DnsTuple{}, common.CmdlineTuple{})
	trans.Request = &DnsMessage{
		Data: &mkdns.Msg{},
	}
	dns.publishTransaction(trans)

	m := expectResult(t, dns)
	assert.Equal(t, common.ERROR_STATUS, mapValue(t, m, "status"))
}

// Verify that an empty DNS response packet can be published.
func TestPublishTransaction_emptyDnsResponse(t *testing.T) {
	dns := newDns(testing.Verbose())

	trans := newTransaction(time.Now(), DnsTuple{}, common.CmdlineTuple{})
	trans.Response = &DnsMessage{
		Data: &mkdns.Msg{},
	}
	dns.publishTransaction(trans)

	m := expectResult(t, dns)
	assert.Equal(t, common.ERROR_STATUS, mapValue(t, m, "status"))
}

func TestPublishTransaction_edns(t *testing.T) {
	dns := newDns(testing.Verbose())
	q := ednsSecA
	packet := newPacket(forward, q.request)
	dns.ParseUdp(packet)
	assert.Equal(t, 1, dns.transactions.Size(), "There should be one transaction.")
	packet = newPacket(reverse, q.response)
	dns.ParseUdp(packet)
	assert.Empty(t, dns.transactions.Size(), "There should be no transactions.")

	m := expectResult(t, dns)
	assert.Equal(t, "udp", mapValue(t, m, "transport"))
	assert.Equal(t, len(q.request), mapValue(t, m, "bytes_in"))
	assert.Equal(t, len(q.response), mapValue(t, m, "bytes_out"))
	assert.NotNil(t, mapValue(t, m, "responsetime"))
	assert.Equal(t, common.OK_STATUS, mapValue(t, m, "status"))
	assert.Nil(t, mapValue(t, m, "notes"))
	assertMapStrData(t, m, q)
}

// Verify that a non-edns answer to a edns query publishes Notes.
func TestPublishTransaction_respEdnsNoSupport(t *testing.T) {
	dns := newDns(testing.Verbose())
	q := ednsSecA
	q.response = q.response[:len(q.response)-11] // Remove OPT RR

	packet := newPacket(forward, q.request)
	dns.ParseUdp(packet)
	assert.Equal(t, 1, dns.transactions.Size(), "There should be one transaction.")
	packet = newPacket(reverse, q.response)
	dns.ParseUdp(packet)
	assert.Empty(t, dns.transactions.Size(), "There should be no transactions.")

	m := expectResult(t, dns)
	assert.Equal(t, "udp", mapValue(t, m, "transport"))
	assert.Equal(t, len(q.request), mapValue(t, m, "bytes_in"))
	assert.Equal(t, len(q.response), mapValue(t, m, "bytes_out"))
	assert.NotNil(t, mapValue(t, m, "responsetime"))
	assert.Equal(t, common.OK_STATUS, mapValue(t, m, "status"))
	assert.Equal(t, RespEdnsNoSupport.Error(), mapValue(t, m, "notes"))
	assertRequest(t, m, q)
}

// Verify that a edns response to a non-edns query publishes Notes.
func TestPublishTransaction_respEdnsUnexpected(t *testing.T) {
	dns := newDns(testing.Verbose())
	q := ednsSecA
	q.request = q.request[:len(q.request)-11] // Remove OPT RR

	packet := newPacket(forward, q.request)
	dns.ParseUdp(packet)
	assert.Equal(t, 1, dns.transactions.Size(), "There should be one transaction.")
	packet = newPacket(reverse, q.response)
	dns.ParseUdp(packet)
	assert.Empty(t, dns.transactions.Size(), "There should be no transactions.")

	m := expectResult(t, dns)
	assert.Equal(t, "udp", mapValue(t, m, "transport"))
	assert.Equal(t, len(q.request), mapValue(t, m, "bytes_in"))
	assert.Equal(t, len(q.response), mapValue(t, m, "bytes_out"))
	assert.NotNil(t, mapValue(t, m, "responsetime"))
	assert.Equal(t, common.OK_STATUS, mapValue(t, m, "status"))
	assert.Equal(t, RespEdnsUnexpected.Error(), mapValue(t, m, "notes"))
	assertMapStrData(t, m, q)
}

// Benchmarks UDP parsing for the given test message.
func benchmarkUdp(b *testing.B, q DnsTestMessage) {
	dns := newDns(false)
	for i := 0; i < b.N; i++ {
		packet := newPacket(forward, q.request)
		dns.ParseUdp(packet)
		packet = newPacket(reverse, q.response)
		dns.ParseUdp(packet)

		client := dns.results.(*publish.ChanTransactions)
		<-client.Channel
	}
}

// Benchmark UDP parsing against each test message.
func BenchmarkUdpElasticA(b *testing.B)  { benchmarkUdp(b, elasticA) }
func BenchmarkUdpZoneIxfr(b *testing.B)  { benchmarkUdp(b, zoneIxfr) }
func BenchmarkUdpGithubPtr(b *testing.B) { benchmarkUdp(b, githubPtr) }
func BenchmarkUdpSophosTxt(b *testing.B) { benchmarkUdp(b, sophosTxt) }

// Benchmark that runs with parallelism to help find concurrency related
// issues. To run with parallelism, the 'go test' cpu flag must be set
// greater than 1, otherwise it just runs concurrently but not in parallel.
func BenchmarkParallelUdpParse(b *testing.B) {
	rand.Seed(22)
	numMessages := len(messages)
	dns := newDns(false)
	client := dns.results.(*publish.ChanTransactions)

	// Drain the results channal while the test is running.
	go func() {
		totalMessages := 0
		for r := range client.Channel {
			_ = r
			totalMessages++
		}
		fmt.Printf("Parsed %d messages.\n", totalMessages)
	}()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		// Each iteration parses one message, either a request or a response.
		// The request and response could be parsed on different goroutines.
		for pb.Next() {
			q := messages[rand.Intn(numMessages)]
			var packet *protos.Packet
			if rand.Intn(2) == 0 {
				packet = newPacket(forward, q.request)
			} else {
				packet = newPacket(reverse, q.response)
			}
			dns.ParseUdp(packet)
		}
	})

	defer close(client.Channel)
}

// parseUdpRequestResponse parses a request then a response packet and validates
// the published result.
func parseUdpRequestResponse(t testing.TB, dns *Dns, q DnsTestMessage) {
	packet := newPacket(forward, q.request)
	dns.ParseUdp(packet)
	packet = newPacket(reverse, q.response)
	dns.ParseUdp(packet)
	assert.Empty(t, dns.transactions.Size(), "There should be no transactions.")

	m := expectResult(t, dns)
	assert.Equal(t, "udp", mapValue(t, m, "transport"))
	assert.Equal(t, len(q.request), mapValue(t, m, "bytes_in"))
	assert.Equal(t, len(q.response), mapValue(t, m, "bytes_out"))
	assert.NotNil(t, mapValue(t, m, "responsetime"))

	if assert.ObjectsAreEqual("NOERROR", mapValue(t, m, "dns.response_code")) {
		assert.Equal(t, common.OK_STATUS, mapValue(t, m, "status"))
	} else {
		assert.Equal(t, common.ERROR_STATUS, mapValue(t, m, "status"))
	}

	assert.Nil(t, mapValue(t, m, "notes"))
	assertMapStrData(t, m, q)
}
