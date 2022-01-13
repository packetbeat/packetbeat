package cmd

import (
	"github.com/elastic/beats/v7/kubebeat/beater"

	cmd "github.com/elastic/beats/v7/libbeat/cmd"
	"github.com/elastic/beats/v7/libbeat/cmd/instance"

	_ "github.com/elastic/beats/v7/x-pack/libbeat/include"
)

// Name of this beat
var Name = "kubebeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: Name})
