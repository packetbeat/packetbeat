// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common/reload"
	"github.com/elastic/elastic-agent-client/v7/pkg/client"
	"github.com/elastic/elastic-agent-client/v7/pkg/proto"
	conf "github.com/elastic/elastic-agent-libs/config"
)

func packetbeatCfg(rawIn *proto.UnitExpectedConfig, agentInfo *client.AgentInfo) ([]*reload.ConfigWithMeta, error) {
	configList := make([]*reload.ConfigWithMeta, 1)
	uconfig, err := conf.NewConfigFrom(rawIn.Source.AsMap())
	if err != nil {
		return nil, fmt.Errorf("error in conversion to conf.C: %w", err)
	}

	configList[0] = &reload.ConfigWithMeta{Config: uconfig}

	return configList, nil
}
