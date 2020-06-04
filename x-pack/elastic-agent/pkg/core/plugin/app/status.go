// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package app

import (
	"context"
	"fmt"

	"github.com/elastic/elastic-agent-client/v7/pkg/proto"
	"github.com/sanathkr/go-yaml"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/errors"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/server"
)

type ApplicationStatusHandler struct{}

func (*ApplicationStatusHandler) OnStatusChange(state *server.ApplicationState, status proto.StateObserved_Status, _ string) {
	app, ok := state.App().(*Application)
	if !ok {
		panic(errors.New("only *Application can be registered when using the ApplicationStatusHandler", errors.TypeUnexpected))
	}

	app.appLock.Lock()
	if status == proto.StateObserved_FAILED {
		// ignore when expected state is stopping
		if state.Expected() == proto.StateExpected_STOPPING {
			app.appLock.Unlock()
			return
		}

		// it was a crash, report it async not to block
		// process management with networking issues
		go app.reportCrash(context.Background())

		// kill the process
		if app.proc != nil {
			_ = app.proc.Process.Kill()
			app.proc = nil
		}
		ctx := app.startContext
		tag := app.tag
		app.appLock.Unlock()

		// it was marshalled to pass into the state, so unmarshall will always succeed
		var cfg map[string]interface{}
		_ = yaml.Unmarshal([]byte(state.Config()), &cfg)

		err := app.Start(ctx, tag, cfg)
		if err != nil {
			app.logger.Error(errors.New(
				fmt.Sprintf("application '%s' failed to restart", app.id),
				errors.TypeApplicationCrash,
				errors.M(errors.MetaKeyAppName, app.name),
				errors.M(errors.MetaKeyAppName, app.id)))
		}
	}
}

