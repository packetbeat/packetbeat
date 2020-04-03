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

package unix

import (
	"fmt"
	"net"
	"sync"

	"golang.org/x/net/netutil"

	"github.com/elastic/beats/filebeat/inputsource/tcp"
	"github.com/elastic/beats/v7/filebeat/inputsource/common"
	"github.com/elastic/beats/v7/libbeat/common/atomic"
	"github.com/elastic/beats/v7/libbeat/logp"
)

// Server represent a Unix socket server
type Server struct {
	config       *Config
	Listener     net.Listener
	wg           sync.WaitGroup
	done         chan struct{}
	factory      tcp.HandlerFactory
	log          *logp.Logger
	closer       *tcp.Closer
	clientsCount atomic.Int
}

// New creates a new unix server
func New(
	config *Config,
	factory tcp.HandlerFactory,
) (*Server, error) {
	if factory == nil {
		return nil, fmt.Errorf("HandlerFactory can't be empty")
	}

	return &Server{
		config:  config,
		done:    make(chan struct{}),
		factory: factory,
		log:     logp.NewLogger("unix").With("path", config.Path),
		closer:  NewCloser(nil),
	}, nil
}

// Start listens on the Unix socket.
func (s *Server) Start() error {
	var err error
	s.Listener, err = s.createServer()
	if err != nil {
		return err
	}

	s.closer.callback = func() { s.Listener.Close() }
	s.log.Info("Started listening for Unix socket connection")

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.run()
	}()
	return nil
}

// Run start and run a new Unix socket listener to receive new data. When a new connection is accepted, the factory
// is used  to create a ConnectionHandler. The ConnectionHandler takes the connection as input and handles the data
// that is being received via tha io.Reader. Most clients use the splitHandler which can take a bufio.SplitFunc and
// parse out each message into an appropriate event. The Close() of the ConnectionHandler can be used to clean up
// the connection either by client or server based on need.
func (s *Server) run() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			select {
			case <-s.closer.Done():
				return
			default:
				s.log.Debugw("Can not accept the connection", "error", err)
				continue
			}
		}

		handler := s.factory(*s.config)
		closer := common.WithCloser(s.closer, func() { conn.Close() })

		s.wg.Add(1)
		go func() {
			defer logp.Recover("recovering from a unix client crash")
			defer s.wg.Done()
			defer closer.Close()

			s.registerHandler()
			defer s.unregisterHandler()
			s.log.Debugw("New client", "total", s.clientsCount.Load())

			err := handler.Handle(closer, conn)
			if err != nil {
				s.log.Debugw("client error", "error", err)
			}

			defer s.log.Debugw("client disconnected", "total", s.clientsCount.Load())
		}()
	}
}

// Stop stops accepting new incoming Unix socket connection and Close any active clients
func (s *Server) Stop() {
	s.log.Info("Stopping Unix socket server")
	s.closer.Close()
	s.wg.Wait()
	s.log.Info("Unix socket server stopped")
}

func (s *Server) registerHandler() {
	s.clientsCount.Inc()
}

func (s *Server) unregisterHandler() {
	s.clientsCount.Dec()
}

func (s *Server) createServer() (net.Listener, error) {
	l, err := net.Listen("unix", s.config.Path)
	if err != nil {
		return nil, err
	}

	if s.config.MaxConnections > 0 {
		return netutil.LimitListener(l, s.config.MaxConnections), nil
	}
	return l, nil
}
