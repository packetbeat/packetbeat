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

package beater

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/packetbeat/npcap"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

const installTimeout = 120 * time.Second

func installNpcap(b *beat.Beat, cfg *conf.C) error {
	if !b.Info.ElasticLicensed {
		return nil
	}

	defer func() {
		log := logp.NewLogger("npcap")
		npcapVersion := npcap.Version()
		if npcapVersion == "" {
			log.Warn("no version available for npcap")
		} else {
			log.Infof("npcap version: %s", npcapVersion)
		}
	}()

	log := logp.NewLogger("npcap_install")

	// Get a complete diagnostic config from the *beat.Beat.
	// Move this before calling canInstallNpcap and don't
	// do more than log the error. We are not here to disrupt,
	// but to observe.
	rawConfig := make(mapstr.M)
	err := b.BeatConfig.Unpack(&rawConfig)
	if err != nil {
		log.Errorf("failed to unpack complete beat config from *config.C: %v", err)
	}
	log.Infow("complete config", "beat config", rawConfig)

	rawConfig = make(mapstr.M)
	err = cfg.Unpack(&rawConfig)
	if err != nil {
		log.Errorf("failed to unpack complete raw config from *config.C: %v", err)
	}
	log.Infow("complete config", "raw config", rawConfig)

	canInstall, err := canInstallNpcap(b, cfg)
	if err != nil {
		return err
	}
	if !canInstall {
		log.Warn("npcap installation/upgrade disabled by user")
		return nil
	}

	if runtime.GOOS != "windows" {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), installTimeout)
	defer cancel()

	if npcap.Installer == nil {
		return nil
	}
	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		return fmt.Errorf("could not create installation temporary directory: %w", err)
	}
	defer func() {
		// The init sequence duplicates the embedded binary.
		// Get rid of the part we can. The remainder is in
		// the packetbeat text section as a string.
		npcap.Installer = nil
		// Remove the installer from the file system.
		os.RemoveAll(tmp)
	}()
	installerPath := filepath.Join(tmp, "npcap.exe")
	err = os.WriteFile(installerPath, npcap.Installer, 0o700)
	if err != nil {
		return fmt.Errorf("could not create installation temporary file: %w", err)
	}
	return npcap.Install(ctx, log, installerPath, "", false)
}

// canInstallNpcap returns whether the Npcap DLL installation can proceed or has been
// blocked by the user. This needs special consideration because we have not yet had
// configurations from agent normalised to the internal packetbeat format by this point.
// In the case that the beat is managed, any data stream that has npcap.never_install
// set to true will result in a block on the installation.
func canInstallNpcap(b *beat.Beat, rawcfg *conf.C) (bool, error) {
	type npcapInstallCfg struct {
		NeverInstall bool `config:"npcap.never_install"`
	}

	// Agent managed case.
	if b.Manager.Enabled() {
		var cfg struct {
			Streams []npcapInstallCfg `config:"streams"`
		}
		err := rawcfg.Unpack(&cfg)
		if err != nil {
			return false, fmt.Errorf("failed to unpack npcap config from agent configuration: %w", err)
		}
		for _, c := range cfg.Streams {
			if c.NeverInstall {
				return false, nil
			}
		}
		return true, nil
	}

	// Packetbeat case.
	var cfg npcapInstallCfg
	err := rawcfg.Unpack(&cfg)
	if err != nil {
		return false, fmt.Errorf("failed to unpack npcap config from packetbeat configuration: %w", err)
	}
	return !cfg.NeverInstall, err
}
