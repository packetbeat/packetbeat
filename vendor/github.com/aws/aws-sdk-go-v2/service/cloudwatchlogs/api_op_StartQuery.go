// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/StartQueryRequest
type StartQueryInput struct {
	_ struct{} `type:"structure"`

	// The end of the time range to query. The range is inclusive, so the specified
	// end time is included in the query. Specified as epoch time, the number of
	// seconds since January 1, 1970, 00:00:00 UTC.
	//
	// EndTime is a required field
	EndTime *int64 `locationName:"endTime" type:"long" required:"true"`

	// The maximum number of log events to return in the query. If the query string
	// uses the fields command, only the specified fields and their values are returned.
	Limit *int64 `locationName:"limit" min:"1" type:"integer"`

	// The log group on which to perform the query.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The query string to use. For more information, see CloudWatch Logs Insights
	// Query Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	//
	// QueryString is a required field
	QueryString *string `locationName:"queryString" type:"string" required:"true"`

	// The beginning of the time range to query. The range is inclusive, so the
	// specified start time is included in the query. Specified as epoch time, the
	// number of seconds since January 1, 1970, 00:00:00 UTC.
	//
	// StartTime is a required field
	StartTime *int64 `locationName:"startTime" type:"long" required:"true"`
}

// String returns the string representation
func (s StartQueryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartQueryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartQueryInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if s.QueryString == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryString"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/StartQueryResponse
type StartQueryOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the query.
	QueryId *string `locationName:"queryId" type:"string"`
}

// String returns the string representation
func (s StartQueryOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartQuery = "StartQuery"

// StartQueryRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Schedules a query of a log group using CloudWatch Logs Insights. You specify
// the log group and time range to query, and the query string to use.
//
// For more information, see CloudWatch Logs Insights Query Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
//
//    // Example sending a request using StartQueryRequest.
//    req := client.StartQueryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/StartQuery
func (c *Client) StartQueryRequest(input *StartQueryInput) StartQueryRequest {
	op := &aws.Operation{
		Name:       opStartQuery,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartQueryInput{}
	}

	req := c.newRequest(op, input, &StartQueryOutput{})
	return StartQueryRequest{Request: req, Input: input, Copy: c.StartQueryRequest}
}

// StartQueryRequest is the request type for the
// StartQuery API operation.
type StartQueryRequest struct {
	*aws.Request
	Input *StartQueryInput
	Copy  func(*StartQueryInput) StartQueryRequest
}

// Send marshals and sends the StartQuery API request.
func (r StartQueryRequest) Send(ctx context.Context) (*StartQueryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartQueryResponse{
		StartQueryOutput: r.Request.Data.(*StartQueryOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartQueryResponse is the response type for the
// StartQuery API operation.
type StartQueryResponse struct {
	*StartQueryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartQuery request.
func (r *StartQueryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
