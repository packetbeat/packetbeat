// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package awss3

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/elastic/beats/v7/libbeat/common/cfgtype"
	"github.com/elastic/beats/v7/libbeat/common/match"
	"github.com/elastic/beats/v7/libbeat/reader/parser"
	"github.com/elastic/beats/v7/libbeat/reader/readfile"
	awscommon "github.com/elastic/beats/v7/x-pack/libbeat/common/aws"
)

type config struct {
	APITimeout          time.Duration        `config:"api_timeout"`
	VisibilityTimeout   time.Duration        `config:"visibility_timeout"`
	FIPSEnabled         bool                 `config:"fips_enabled"`
	MaxNumberOfMessages int                  `config:"max_number_of_messages"`
<<<<<<< HEAD
	QueueURL            string               `config:"queue_url" validate:"required"`
=======
	QueueURL            string               `config:"queue_url"`
	BucketARN           string               `config:"bucket_arn"`
	BucketListInterval  time.Duration        `config:"bucket_list_interval"`
	NumberOfWorkers     int                  `config:"number_of_workers"`
>>>>>>> b4ecc29bb (Add vars in modules.d/aws.yml.disabled (#27454))
	AWSConfig           awscommon.ConfigAWS  `config:",inline"`
	FileSelectors       []fileSelectorConfig `config:"file_selectors"`
	ReaderConfig        readerConfig         `config:",inline"` // Reader options to apply when no file_selectors are used.
}

func defaultConfig() config {
	c := config{
		APITimeout:          120 * time.Second,
		VisibilityTimeout:   300 * time.Second,
		FIPSEnabled:         false,
		MaxNumberOfMessages: 5,
	}
	c.ReaderConfig.InitDefaults()
	return c
}

func (c *config) Validate() error {
<<<<<<< HEAD
	if c.VisibilityTimeout <= 0 || c.VisibilityTimeout.Hours() > 12 {
=======
	if c.QueueURL == "" && c.BucketARN == "" {
		return fmt.Errorf("queue_url or bucket_arn must provided")
	}

	if c.QueueURL != "" && c.BucketARN != "" {
		return fmt.Errorf("queue_url <%v> and bucket_arn <%v> "+
			"cannot be set at the same time", c.QueueURL, c.BucketARN)
	}

	if c.BucketARN != "" && c.BucketListInterval <= 0 {
		return fmt.Errorf("bucket_list_interval <%v> must be greater than 0", c.BucketListInterval)
	}

	if c.BucketARN != "" && c.NumberOfWorkers <= 0 {
		return fmt.Errorf("number_of_workers <%v> must be greater than 0", c.NumberOfWorkers)
	}

	if c.QueueURL != "" && (c.VisibilityTimeout <= 0 || c.VisibilityTimeout.Hours() > 12) {
>>>>>>> b4ecc29bb (Add vars in modules.d/aws.yml.disabled (#27454))
		return fmt.Errorf("visibility_timeout <%v> must be greater than 0 and "+
			"less than or equal to 12h", c.VisibilityTimeout)
	}

	if c.APITimeout <= 0 || c.APITimeout > c.VisibilityTimeout/2 {
		return fmt.Errorf("api_timeout <%v> must be greater than 0 and less "+
			"than 1/2 of the visibility_timeout (%v)", c.APITimeout, c.VisibilityTimeout/2)
	}

	if c.MaxNumberOfMessages <= 0 || c.MaxNumberOfMessages > 10 {
		return fmt.Errorf("max_number_of_messages <%v> must be greater than "+
			"0 and less than or equal to 10", c.MaxNumberOfMessages)
	}
	return nil
}

// fileSelectorConfig defines reader configuration that applies to a subset
// of S3 objects whose URL matches the given regex.
type fileSelectorConfig struct {
	Regex        *match.Matcher `config:"regex" validate:"required"`
	ReaderConfig readerConfig   `config:",inline"`
}

// readerConfig defines the options for reading the content of an S3 object.
type readerConfig struct {
	BufferSize               cfgtype.ByteSize        `config:"buffer_size"`
	ContentType              string                  `config:"content_type"`
	Encoding                 string                  `config:"encoding"`
	ExpandEventListFromField string                  `config:"expand_event_list_from_field"`
	IncludeS3Metadata        []string                `config:"include_s3_metadata"`
	LineTerminator           readfile.LineTerminator `config:"line_terminator"`
	MaxBytes                 cfgtype.ByteSize        `config:"max_bytes"`
	Parsers                  parser.Config           `config:",inline"`
}

func (f *readerConfig) Validate() error {
	if f.BufferSize <= 0 {
		return fmt.Errorf("buffer_size <%v> must be greater than 0", f.BufferSize)
	}

	if f.MaxBytes <= 0 {
		return fmt.Errorf("max_bytes <%v> must be greater than 0", f.MaxBytes)
	}
	if f.ExpandEventListFromField != "" && f.ContentType != "" && f.ContentType != "application/json" {
		return fmt.Errorf("content_type must be `application/json` when expand_event_list_from_field is used")
	}

	return nil
}

func (f *readerConfig) InitDefaults() {
	f.BufferSize = 16 * humanize.KiByte
	f.MaxBytes = 10 * humanize.MiByte
	f.LineTerminator = readfile.AutoLineTerminator
}
