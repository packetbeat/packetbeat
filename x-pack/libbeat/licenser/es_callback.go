// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package licenser

import (
	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/common/atomic"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs/elasticsearch"
)

// Enforce setups the corresponding callbacks in libbeat to verify the license on the
// remote elasticsearch cluster.
func Enforce(log *logp.Logger, checks ...CheckFunc) {
	validLicense := atomic.MakeBool(false)
	cb := func(client *elasticsearch.Client) error {
		if validLicense.Load() {
			return nil
		}

		fetcher := NewElasticFetcher(client)
		license, err := fetcher.Fetch()
		if err != nil {
			return errors.Wrapf(err, "cannot retrieve the elasticsearch license or no license endpoint")
		}

		if !Validate(log, *license, checks...) {
			return errors.New("could not find a valid license")
		}

		validLicense.Store(true)

		return nil
	}

	elasticsearch.RegisterLicenseCallback(cb)
}
