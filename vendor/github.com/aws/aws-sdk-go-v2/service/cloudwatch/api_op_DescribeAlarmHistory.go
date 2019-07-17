// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DescribeAlarmHistoryInput
type DescribeAlarmHistoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the alarm.
	AlarmName *string `min:"1" type:"string"`

	// The ending date to retrieve alarm history.
	EndDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The type of alarm histories to retrieve.
	HistoryItemType HistoryItemType `type:"string" enum:"true"`

	// The maximum number of alarm history records to retrieve.
	MaxRecords *int64 `min:"1" type:"integer"`

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string `type:"string"`

	// The starting date to retrieve alarm history.
	StartDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s DescribeAlarmHistoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAlarmHistoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAlarmHistoryInput"}
	if s.AlarmName != nil && len(*s.AlarmName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AlarmName", 1))
	}
	if s.MaxRecords != nil && *s.MaxRecords < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxRecords", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DescribeAlarmHistoryOutput
type DescribeAlarmHistoryOutput struct {
	_ struct{} `type:"structure"`

	// The alarm histories, in JSON format.
	AlarmHistoryItems []AlarmHistoryItem `type:"list"`

	// The token that marks the start of the next batch of returned results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAlarmHistoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAlarmHistory = "DescribeAlarmHistory"

// DescribeAlarmHistoryRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Retrieves the history for the specified alarm. You can filter the results
// by date range or item type. If an alarm name is not specified, the histories
// for all alarms are returned.
//
// CloudWatch retains the history of an alarm even if you delete the alarm.
//
//    // Example sending a request using DescribeAlarmHistoryRequest.
//    req := client.DescribeAlarmHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DescribeAlarmHistory
func (c *Client) DescribeAlarmHistoryRequest(input *DescribeAlarmHistoryInput) DescribeAlarmHistoryRequest {
	op := &aws.Operation{
		Name:       opDescribeAlarmHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeAlarmHistoryInput{}
	}

	req := c.newRequest(op, input, &DescribeAlarmHistoryOutput{})
	return DescribeAlarmHistoryRequest{Request: req, Input: input, Copy: c.DescribeAlarmHistoryRequest}
}

// DescribeAlarmHistoryRequest is the request type for the
// DescribeAlarmHistory API operation.
type DescribeAlarmHistoryRequest struct {
	*aws.Request
	Input *DescribeAlarmHistoryInput
	Copy  func(*DescribeAlarmHistoryInput) DescribeAlarmHistoryRequest
}

// Send marshals and sends the DescribeAlarmHistory API request.
func (r DescribeAlarmHistoryRequest) Send(ctx context.Context) (*DescribeAlarmHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAlarmHistoryResponse{
		DescribeAlarmHistoryOutput: r.Request.Data.(*DescribeAlarmHistoryOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeAlarmHistoryRequestPaginator returns a paginator for DescribeAlarmHistory.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeAlarmHistoryRequest(input)
//   p := cloudwatch.NewDescribeAlarmHistoryRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeAlarmHistoryPaginator(req DescribeAlarmHistoryRequest) DescribeAlarmHistoryPaginator {
	return DescribeAlarmHistoryPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeAlarmHistoryInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeAlarmHistoryPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeAlarmHistoryPaginator struct {
	aws.Pager
}

func (p *DescribeAlarmHistoryPaginator) CurrentPage() *DescribeAlarmHistoryOutput {
	return p.Pager.CurrentPage().(*DescribeAlarmHistoryOutput)
}

// DescribeAlarmHistoryResponse is the response type for the
// DescribeAlarmHistory API operation.
type DescribeAlarmHistoryResponse struct {
	*DescribeAlarmHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAlarmHistory request.
func (r *DescribeAlarmHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
