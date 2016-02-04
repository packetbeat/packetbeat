package logstash

// TODO:
//  - test with connection timeout

import (
	"io"
	"net"
	"testing"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/streambuf"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/stretchr/testify/assert"
)

type mockAddr string

const (
	driverCmdQuit = iota
	driverCmdPublish
)

type testClientDriver interface {
	Stop()
	Publish(events []common.MapStr)
	Returns() []testClientReturn
}

type clientFactory func(TransportClient) testClientDriver

type testClientReturn struct {
	n   int
	err error
}

type testDriverCommand struct {
	code   int
	events []common.MapStr
}

const (
	cmdError = iota
	cmdOK
	cmdMessage
)

type mockTransportCommand struct {
	code    uint8
	message []byte
	err     error
}

type mockTransport struct {
	buf     streambuf.Buffer
	ch      chan []byte
	control chan mockTransportCommand
}

func newLumberjackTestClient(conn TransportClient) *client {
	c, err := newLumberjackClient(conn, 3, testMaxWindowSize, 250*time.Millisecond)
	if err != nil {
		panic(err)
	}
	return c
}

func (a mockAddr) Network() string { return "fake" }
func (a mockAddr) String() string  { return string(a) }

func newMockTransport() *mockTransport {
	return &mockTransport{
		ch:      make(chan []byte),
		control: make(chan mockTransportCommand),
	}
}

func (m *mockTransport) Connect(timeout time.Duration) error {
	return nil
}

func (m *mockTransport) IsConnected() bool {
	return true
}

func (m *mockTransport) Close() error {
	// close(m.ch)
	// close(m.control)
	return nil
}

func (m *mockTransport) Read(b []byte) (n int, err error) {
	cmd := <-m.control
	switch cmd.code {
	case cmdError:
		return 0, cmd.err
	case cmdOK:
		return 0, nil
	case cmdMessage:
		m.buf.Write(cmd.message)
		return m.buf.Read(b)
	}
	return 0, nil
}

func (m *mockTransport) Write(b []byte) (int, error) {
	m.ch <- b
	cmd := <-m.control
	switch cmd.code {
	case cmdError:
		return 0, cmd.err
	case cmdOK:
		return len(b), nil
	case cmdMessage:
		m.buf.Write(cmd.message)
		return len(b), nil
	}
	return 0, nil
}

func (m *mockTransport) recv(into io.Writer) {
	bytes, ok := <-m.ch
	if ok && len(bytes) > 0 {
		into.Write(bytes)
	}
}

func (m *mockTransport) sendError(e error) {
	m.control <- mockTransportCommand{code: cmdError, err: e}
}

func (m *mockTransport) sendOK() {
	m.control <- mockTransportCommand{code: cmdOK}
}

func (m *mockTransport) sendBytes(b []byte) {
	m.control <- mockTransportCommand{code: cmdMessage, message: b}
}

func (m *mockTransport) LocalAddr() net.Addr  { return mockAddr("client") }
func (m *mockTransport) RemoteAddr() net.Addr { return mockAddr("server") }

func (m *mockTransport) SetDeadline(t time.Time) error      { return nil }
func (m *mockTransport) SetReadDeadline(t time.Time) error  { return nil }
func (m *mockTransport) SetWriteDeadline(t time.Time) error { return nil }

func recvMessage(buf *streambuf.Buffer, transp *mockTransport) (*message, error) {
	for {
		transp.recv(buf)
		resp, err := readMessage(buf)
		transp.sendOK()
		if resp != nil || err != nil {
			return resp, err
		}
	}
}

func sendAck(transp *mockTransport, seq uint32) {
	buf := streambuf.New(nil)
	buf.WriteByte('2')
	buf.WriteByte('A')
	buf.WriteNetUint32(seq)
	transp.sendBytes(buf.Bytes())
}

func enableLogging(selectors []string) {
	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, selectors)
	}
}

const testMaxWindowSize = 64

func testSendZero(t *testing.T, factory clientFactory) {
	enableLogging([]string{"*"})

	server := newMockServerTCP(t, 1*time.Second, "")
	defer server.Close()

	sock, transp, err := server.connectPair(1 * time.Second)
	if err != nil {
		t.Fatalf("Failed to connect server and client: %v", err)
	}

	client := factory(transp)
	defer sock.Close()
	defer transp.Close()

	client.Publish(make([]common.MapStr, 0))

	client.Stop()
	returns := client.Returns()

	assert.Equal(t, 1, len(returns))
	if len(returns) == 1 {
		assert.Equal(t, 0, returns[0].n)
		assert.Nil(t, returns[0].err)
	}
}

func testSimpleEvent(t *testing.T, factory clientFactory) {
	enableLogging([]string{"*"})
	server := newMockServerTCP(t, 1*time.Second, "")

	sock, transp, err := server.connectPair(1 * time.Second)
	if err != nil {
		t.Fatalf("Failed to connect server and client: %v", err)
	}
	client := factory(transp)
	conn := &mockConn{sock, streambuf.New(nil)}
	defer transp.Close()
	defer sock.Close()

	event := common.MapStr{"name": "me", "line": 10}
	client.Publish([]common.MapStr{event})

	// receive window message
	err = sock.SetReadDeadline(time.Now().Add(1 * time.Second))
	win, err := conn.recvMessage()
	assert.Nil(t, err)

	// receive data message
	msg, err := conn.recvMessage()
	assert.Nil(t, err)

	// send ack
	conn.sendACK(1)

	client.Stop()

	// validate
	assert.NotNil(t, win)
	assert.NotNil(t, msg)
	assert.Equal(t, 1, len(msg.events))
	msg = msg.events[0]
	assert.Equal(t, "me", msg.doc["name"])
	assert.Equal(t, 10.0, msg.doc["line"])
}

func testStructuredEvent(t *testing.T, factory clientFactory) {
	enableLogging([]string{"*"})
	server := newMockServerTCP(t, 1*time.Second, "")

	sock, transp, err := server.connectPair(1 * time.Second)
	if err != nil {
		t.Fatalf("Failed to connect server and client: %v", err)
	}
	client := factory(transp)
	conn := &mockConn{sock, streambuf.New(nil)}
	defer transp.Close()
	defer sock.Close()

	event := common.MapStr{
		"name": "test",
		"struct": common.MapStr{
			"field1": 1,
			"field2": true,
			"field3": []int{1, 2, 3},
			"field4": []interface{}{
				1,
				"test",
				common.MapStr{
					"sub": "field",
				},
			},
			"field5": common.MapStr{
				"sub1": 2,
			},
		},
	}
	client.Publish([]common.MapStr{event})

	win, err := conn.recvMessage()
	assert.Nil(t, err)

	msg, err := conn.recvMessage()
	assert.Nil(t, err)

	conn.sendACK(1)
	defer client.Stop()

	// validate
	assert.NotNil(t, win)
	assert.NotNil(t, msg)
	assert.Equal(t, 1, len(msg.events))
	msg = msg.events[0]
	assert.Equal(t, "test", msg.doc["name"])
	assert.Equal(t, 1.0, msg.doc.get("struct.field1"))
	assert.Equal(t, true, msg.doc.get("struct.field2"))
	assert.Equal(t, 2.0, msg.doc.get("struct.field5.sub1"))
}

func testCloseAfterWindowSize(t *testing.T, factory clientFactory) {
	enableLogging([]string{"*"})
	server := newMockServerTCP(t, 100*time.Millisecond, "")

	sock, transp, err := server.connectPair(100 * time.Millisecond)
	if err != nil {
		t.Fatalf("Failed to connect server and client: %v", err)
	}
	client := factory(transp)
	conn := &mockConn{sock, streambuf.New(nil)}
	defer transp.Close()
	defer sock.Close()

	client.Publish([]common.MapStr{common.MapStr{
		"message": "hello world",
	}})

	_, err = conn.recvMessage()
	if err != nil {
		t.Fatalf("failed to read window size message: %v", err)
	}

	defer client.Stop()
}
