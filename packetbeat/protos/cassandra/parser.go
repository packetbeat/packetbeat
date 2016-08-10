package cassandra

import (
	"errors"
	"time"

	"github.com/elastic/beats/libbeat/common/streambuf"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/packetbeat/protos/applayer"
	. "github.com/elastic/beats/packetbeat/protos/cassandra/internal/gocql"
)

type parser struct {
	buf       streambuf.Buffer
	config    *parserConfig
	framer    *Framer
	message   *message
	onMessage func(m *message) error
}

type parserConfig struct {
	maxBytes   int
	compressor Compressor
	ignoredOps map[FrameOp]bool
}

// check whether this ops is enabled or not
func (p *parser) FrameOpsIgnored() bool {

	if p.config.ignoredOps != nil && len(p.config.ignoredOps) > 0 {
		//default map value is false
		v := p.config.ignoredOps[p.framer.Header.Op]
		if v {
			return true
		}
	}
	return false
}

type message struct {
	applayer.Message

	// indicator for parsed message being complete or requires more messages
	// (if false) to be merged to generate full message.
	isComplete bool

	failed  bool
	ignored bool

	data   map[string]interface{}
	header map[string]interface{}
	// list element use by 'transactions' for correlation
	next *message

	transactionTimeout time.Duration

	results transactions
}

// Error code if stream exceeds max allowed size on append.
var (
	ErrStreamTooLarge = errors.New("Stream data too large")
	isDebug           = false
)

func (p *parser) init(
	cfg *parserConfig,
	onMessage func(*message) error,
) {
	*p = parser{
		buf:       streambuf.Buffer{},
		config:    cfg,
		onMessage: onMessage,
	}

	isDebug = logp.IsDebug("cassandra")

}

func (p *parser) append(data []byte) error {
	_, err := p.buf.Write(data)
	if err != nil {
		return err
	}

	if p.config.maxBytes > 0 && p.buf.Total() > p.config.maxBytes {
		return ErrStreamTooLarge
	}
	return nil
}

func (p *parser) feed(ts time.Time, data []byte) error {
	if err := p.append(data); err != nil {
		return err
	}

	for p.buf.Total() > 0 {
		if p.message == nil {
			// allocate new message object to be used by parser with current timestamp
			p.message = p.newMessage(ts)
			p.message.data = map[string]interface{}{}
		}

		msg, err := p.parse()
		if err != nil {
			return err
		}
		if msg == nil {
			break // wait for more data
		}

		// reset buffer and message -> handle next message in buffer
		p.buf.Reset()
		p.message = nil

		// call message handler callback
		if err := p.onMessage(msg); err != nil {
			return err
		}
	}

	return nil
}

func (p *parser) newMessage(ts time.Time) *message {
	return &message{
		Message: applayer.Message{
			Ts: ts,
		},
	}
}

func (p *parser) parse() (*message, error) {

	// if p.frame is nil then create a new framer, or continue to process the last message
	if p.framer == nil {
		if isDebug {
			debugf("start new framer")
		}
		p.framer = NewFramer(&p.buf, p.config.compressor)
	}

	// check if the frame header were parsed or not
	if p.framer.Header == nil {
		if isDebug {
			debugf("start to parse header")
		}
		if !p.buf.Avail(9) {
			debugf("not enough head bytes, ignore")
			return nil, nil
		}

		_, err := p.framer.ReadHeader()
		if err != nil {
			logp.Err("%v", err)
			p.framer = nil
			return nil, err
		}
	}

	//check if the ops need to be ignored
	if p.FrameOpsIgnored() {
		// as we already ignore the content, we now mark the result is ignored
		p.message.ignored = true
		if isDebug {
			debugf("Ops: %s was marked to be ignored, ignoring, request:%v", p.framer.Header.Op.String(), p.framer.Header.Version.IsRequest())
		}
	}

	msg := p.message

	if p.framer.Header.BodyLength > 0 {

		//let's wait for enough buf
		if !p.buf.Avail(p.framer.Header.BodyLength) {
			if isDebug {
				debugf("buf not enough for body, waiting for more, return")
			}
			return nil, nil
		}

		//check if the ops already ignored
		if p.message.ignored {
			p.buf.Collect(p.framer.Header.BodyLength)
			finalCollectedFrameLength := p.buf.BufferConsumed()
			if finalCollectedFrameLength-p.framer.Header.HeadLength != p.framer.Header.BodyLength {
				return nil, errors.New("data messed while parse frame body")
			}

		}

		// start to parse body
		if !p.message.ignored {
			data, err := p.framer.ReadFrame()
			if err != nil {
				// if the frame parsed failed, should ignore the whole message
				p.framer = nil
				return nil, err
			}

			msg.data = data

			// dealing with un-parsed content
			frameParsedLength := p.buf.BufferConsumed()

			// collect leftover
			unParsedSize := p.framer.Header.BodyLength + p.framer.Header.HeadLength - frameParsedLength
			if unParsedSize > 0 {

				// double check the buf size
				if p.buf.Avail(unParsedSize) {
					p.buf.Collect(unParsedSize)
				} else {
					logp.Err("should be enough bytes for cleanup,but not enough")
					return nil, errors.New("should be enough bytes,but not enough")
				}
			}

		}

		finalCollectedFrameLength := p.buf.BufferConsumed()
		if finalCollectedFrameLength-p.framer.Header.HeadLength != p.framer.Header.BodyLength {
			logp.Err("body_length:%d, head_length:%d, all_consumed:%d", p.framer.Header.BodyLength, p.framer.Header.HeadLength, finalCollectedFrameLength)
			return nil, errors.New("data messed while parse frame body")
		}

	}

	dir := applayer.NetOriginalDirection

	isRequest := true
	if p.framer.Header.Version.IsResponse() {
		dir = applayer.NetReverseDirection
		isRequest = false
	}

	msg.Size = uint64(p.buf.BufferConsumed())
	msg.IsRequest = isRequest
	msg.Direction = dir

	msg.header = p.framer.Header.ToMap()

	if msg.IsRequest {
		p.message.results.requests.append(msg)
	} else {
		p.message.results.responses.append(msg)
	}
	p.framer = nil
	return msg, nil
}
