package stdin

import (
	"fmt"

	"github.com/elastic/beats/filebeat/channel"
	"github.com/elastic/beats/filebeat/harvester"
	"github.com/elastic/beats/filebeat/input/file"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

// Stdin is a prospector for stdin
type Stdin struct {
	harvester *harvester.Harvester
	started   bool
	cfg       *common.Config
	outlet    channel.Outleter
}

// NewStdin creates a new stdin prospector
// This prospector contains one harvester which is reading from stdin
func NewProspector(cfg *common.Config, outlet channel.Outleter) (*Stdin, error) {

	prospectorer := &Stdin{
		started: false,
		cfg:     cfg,
		outlet:  outlet,
	}

	var err error

	prospectorer.harvester, err = prospectorer.createHarvester(file.State{Source: "-"})
	if err != nil {
		return nil, fmt.Errorf("Error initializing stdin harvester: %v", err)
	}

	return prospectorer, nil
}

// Run runs the prospector
func (s *Stdin) Run() {

	// Make sure stdin harvester is only started once
	if !s.started {
		reader, err := s.harvester.Setup()
		if err != nil {
			logp.Err("Error starting stdin harvester: %s", err)
			return
		}
		go s.harvester.Harvest(reader)
		s.started = true
	}
}

// createHarvester creates a new harvester instance from the given state
func (s *Stdin) createHarvester(state file.State) (*harvester.Harvester, error) {

	// Each harvester gets its own copy of the outlet
	outlet := s.outlet.Copy()
	h, err := harvester.NewHarvester(
		s.cfg,
		state,
		nil,
		outlet,
	)

	return h, err
}
