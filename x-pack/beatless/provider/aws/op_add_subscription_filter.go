// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package aws

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

	"github.com/elastic/beats/libbeat/logp"
)

type subscriptionFilter struct {
	LogGroupName  string
	FilterName    string
	FilterPattern string
}
type opAddSubscriptionFilter struct {
	log          *logp.Logger
	subscription subscriptionFilter
	svc          *cloudwatchlogs.CloudWatchLogs
}

func (o *opAddSubscriptionFilter) Execute(ctx *executorContext) error {
	o.log.Debugf(
		"adding subscription filter for LogGroupName: %, FilterName: %s, FilterPattern: %s",
		o.subscription.LogGroupName,
		o.subscription.FilterName,
		o.subscription.FilterPattern,
	)

	if ctx.AliasArn == "" {
		return errors.New("AliasArn not found, need to create or update the alias")
	}

	req := &cloudwatchlogs.PutSubscriptionFilterInput{
		DestinationArn: aws.String(ctx.AliasArn),
		LogGroupName:   aws.String(o.subscription.LogGroupName),
		FilterName:     aws.String(o.subscription.FilterName),
		FilterPattern:  aws.String(o.subscription.FilterPattern),
	}

	api := o.svc.PutSubscriptionFilterRequest(req)
	resp, err := api.Send()
	if err != nil {
		o.log.Debugf("could not add subscription filter, error: %s, response: %s", err, resp)
		return err
	}

	o.log.Debugf("adding subscription filter successfully")
	return nil
}

func (o *opAddSubscriptionFilter) Rollback(ctx *executorContext) error {
	o.log.Debugf(
		"remove subscription filter for LogGroupName: %, FilterName: %s, FilterPattern: %s",
		o.subscription.LogGroupName,
		o.subscription.FilterName,
		o.subscription.FilterPattern,
	)
	req := &cloudwatchlogs.DeleteSubscriptionFilterInput{
		FilterName:   aws.String(o.subscription.FilterName),
		LogGroupName: aws.String(o.subscription.LogGroupName),
	}

	api := o.svc.DeleteSubscriptionFilterRequest(req)
	resp, err := api.Send()
	if err != nil {
		o.log.Debugf("could not remove subscription filter, error: %s, response: %s", err, resp)
		return err
	}

	o.log.Debugf("subscription filter removed successfully")
	return nil
}

func newOpAddSubscriptionFilter(
	log *logp.Logger,
	awsCfg aws.Config,
	subscription subscriptionFilter,
) *opAddSubscriptionFilter {
	if log == nil {
		log = logp.NewLogger("opAddSubscriptionFilter")
	}
	return &opAddSubscriptionFilter{
		log:          log,
		subscription: subscription,
		svc:          cloudwatchlogs.New(awsCfg),
	}
}
