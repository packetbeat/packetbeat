// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// RDS provides the API operation methods for making requests to
// Amazon Relational Database Service. See this package's package overview docs
// for details on the service.
//
// RDS methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type RDS struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*RDS)

// Used for custom request initialization logic
var initRequest func(*RDS, *aws.Request)

// Service information constants
const (
	ServiceName = "rds"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the RDS client with a config.
//
// Example:
//     // Create a RDS client from just a config.
//     svc := rds.New(myConfig)
func New(config aws.Config) *RDS {
	var signingName string
	signingRegion := config.Region

	svc := &RDS{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2014-10-31",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a RDS operation and runs any
// custom request initialization.
func (c *RDS) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
