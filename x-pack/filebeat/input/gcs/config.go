// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package gcs

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"

	"github.com/elastic/beats/v7/libbeat/common/match"
	"github.com/elastic/beats/v7/libbeat/reader/parser"
)

// MaxWorkers, Poll, PollInterval, BucketTimeOut, ParseJSON, FileSelectors, TimeStampEpoch & ExpandEventListFromField
// can be configured at a global level, which applies to all buckets, as well as at the bucket level.
// Bucket level configurations will always override global level values.
type config struct {
	// ProjectId - Defines the project id of the concerned gcs bucket in Google Cloud.
	ProjectId string `config:"project_id" validate:"required"`
	// Auth - Defines the authentication mechanism to be used for accessing the gcs bucket.
	Auth authConfig `config:"auth"`
	// MaxWorkers - Defines the maximum number of go routines that will be spawned.
	MaxWorkers int `config:"max_workers,omitempty" validate:"max=5000"`
	// Poll - Defines if polling should be performed on the input bucket source.
	Poll bool `config:"poll,omitempty"`
	// PollInterval - Defines the maximum amount of time to wait before polling for the next batch of objects from the bucket.
	PollInterval time.Duration `config:"poll_interval,omitempty"`
	// ParseJSON - Informs the publisher whether to parse & objectify json data or not. By default this is set to
	// false, since it can get expensive dealing with highly nested json data.
	ParseJSON bool `config:"parse_json,omitempty"`
	// BucketTimeOut - Defines the maximum time that the sdk will wait for a bucket api response before timing out.
	BucketTimeOut time.Duration `config:"bucket_timeout,omitempty"`
	// Buckets - Defines a list of buckets that will be polled for objects.
	Buckets []bucket `config:"buckets" validate:"required"`
	// FileSelectors - Defines a list of regex patterns that can be used to filter out objects from the bucket.
	FileSelectors []fileSelectorConfig `config:"file_selectors"`
	// ReaderConfig is the default parser and decoder configuration.
	ReaderConfig readerConfig `config:",inline"`
	// TimeStampEpoch - Defines the epoch time in seconds, which is used to filter out objects that are older than the specified timestamp.
	TimeStampEpoch *int64 `config:"timestamp_epoch"`
	// ExpandEventListFromField - Defines the field name that will be used to expand the event into separate events.
	ExpandEventListFromField string `config:"expand_event_list_from_field"`
	// This field is only used for system test purposes, to override the HTTP endpoint.
	AlternativeHost string `config:"alternative_host,omitempty"`
}

// bucket contains the config for each specific object storage bucket in the root account
type bucket struct {
	Name                     string               `config:"name" validate:"required"`
	MaxWorkers               *int                 `config:"max_workers,omitempty" validate:"max=5000"`
	BucketTimeOut            *time.Duration       `config:"bucket_timeout,omitempty"`
	Poll                     *bool                `config:"poll,omitempty"`
	PollInterval             *time.Duration       `config:"poll_interval,omitempty"`
	ParseJSON                *bool                `config:"parse_json,omitempty"`
	FileSelectors            []fileSelectorConfig `config:"file_selectors"`
	ReaderConfig             readerConfig         `config:",inline"`
	TimeStampEpoch           *int64               `config:"timestamp_epoch"`
	ExpandEventListFromField string               `config:"expand_event_list_from_field"`
}

// fileSelectorConfig helps filter out gcs objects based on a regex pattern
type fileSelectorConfig struct {
	Regex *match.Matcher `config:"regex" validate:"required"`
	// TODO: Add support for reader config in future
}

// readerConfig defines the options for reading the content of an GCS object.
type readerConfig struct {
	Parsers  parser.Config `config:",inline"`
	Decoding decoderConfig `config:"decoding"`
}

type authConfig struct {
	CredentialsJSON *jsonCredentialsConfig `config:"credentials_json,omitempty"`
	CredentialsFile *fileCredentialsConfig `config:"credentials_file,omitempty"`
}

type fileCredentialsConfig struct {
	Path string `config:"path,omitempty"`
}
type jsonCredentialsConfig struct {
	AccountKey string `config:"account_key"`
}

func (c authConfig) Validate() error {
	// credentials_file
	if c.CredentialsFile != nil {
		_, err := os.Stat(c.CredentialsFile.Path)
		if errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("credentials_file is configured, but the file %q cannot be found", c.CredentialsFile.Path)
		} else {
			return nil
		}
	}

	// credentials_json
	if c.CredentialsJSON != nil && len(c.CredentialsJSON.AccountKey) > 0 {
		return nil
	}

	// Application Default Credentials (ADC)
	_, err := google.FindDefaultCredentials(context.Background(), storage.ScopeReadOnly)
	if err == nil {
		return nil
	}

	return fmt.Errorf("no authentication credentials were configured or detected " +
		"(credentials_file, credentials_json, and application default credentials (ADC))")
}
