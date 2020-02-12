// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Parameter input for DescribeDBInstanceAutomatedBackups.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBInstanceAutomatedBackupsMessage
type DescribeDBInstanceAutomatedBackupsInput struct {
	_ struct{} `type:"structure"`

	// (Optional) The user-supplied instance identifier. If this parameter is specified,
	// it must match the identifier of an existing DB instance. It returns information
	// from the specific DB instance' automated backup. This parameter isn't case-sensitive.
	DBInstanceIdentifier *string `type:"string"`

	// The resource ID of the DB instance that is the source of the automated backup.
	// This parameter isn't case-sensitive.
	DbiResourceId *string `type:"string"`

	// A filter that specifies which resources to return based on status.
	//
	// Supported filters are the following:
	//
	//    * status active - automated backups for current instances retained - automated
	//    backups for deleted instances creating - automated backups that are waiting
	//    for the first automated snapshot to be available
	//
	//    * db-instance-id - Accepts DB instance identifiers and Amazon Resource
	//    Names (ARNs) for DB instances. The results list includes only information
	//    about the DB instance automated backupss identified by these ARNs.
	//
	//    * dbi-resource-id - Accepts DB instance resource identifiers and DB Amazon
	//    Resource Names (ARNs) for DB instances. The results list includes only
	//    information about the DB instance resources identified by these ARNs.
	//
	// Returns all resources by default. The status for each resource is specified
	// in the response.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// The pagination token provided in the previous request. If this parameter
	// is specified the response includes only records beyond the marker, up to
	// MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDBInstanceAutomatedBackupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstanceAutomatedBackupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBInstanceAutomatedBackupsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the result of a successful invocation of the DescribeDBInstanceAutomatedBackups
// action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBInstanceAutomatedBackupMessage
type DescribeDBInstanceAutomatedBackupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DBInstanceAutomatedBackup instances.
	DBInstanceAutomatedBackups []DBInstanceAutomatedBackup `locationNameList:"DBInstanceAutomatedBackup" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords .
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBInstanceAutomatedBackupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBInstanceAutomatedBackups = "DescribeDBInstanceAutomatedBackups"

// DescribeDBInstanceAutomatedBackupsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Displays backups for both current and deleted instances. For example, use
// this operation to find details about automated backups for previously deleted
// instances. Current instances with retention periods greater than zero (0)
// are returned for both the DescribeDBInstanceAutomatedBackups and DescribeDBInstances
// operations.
//
// All parameters are optional.
//
//    // Example sending a request using DescribeDBInstanceAutomatedBackupsRequest.
//    req := client.DescribeDBInstanceAutomatedBackupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBInstanceAutomatedBackups
func (c *Client) DescribeDBInstanceAutomatedBackupsRequest(input *DescribeDBInstanceAutomatedBackupsInput) DescribeDBInstanceAutomatedBackupsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBInstanceAutomatedBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBInstanceAutomatedBackupsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBInstanceAutomatedBackupsOutput{})
	return DescribeDBInstanceAutomatedBackupsRequest{Request: req, Input: input, Copy: c.DescribeDBInstanceAutomatedBackupsRequest}
}

// DescribeDBInstanceAutomatedBackupsRequest is the request type for the
// DescribeDBInstanceAutomatedBackups API operation.
type DescribeDBInstanceAutomatedBackupsRequest struct {
	*aws.Request
	Input *DescribeDBInstanceAutomatedBackupsInput
	Copy  func(*DescribeDBInstanceAutomatedBackupsInput) DescribeDBInstanceAutomatedBackupsRequest
}

// Send marshals and sends the DescribeDBInstanceAutomatedBackups API request.
func (r DescribeDBInstanceAutomatedBackupsRequest) Send(ctx context.Context) (*DescribeDBInstanceAutomatedBackupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBInstanceAutomatedBackupsResponse{
		DescribeDBInstanceAutomatedBackupsOutput: r.Request.Data.(*DescribeDBInstanceAutomatedBackupsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBInstanceAutomatedBackupsRequestPaginator returns a paginator for DescribeDBInstanceAutomatedBackups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBInstanceAutomatedBackupsRequest(input)
//   p := rds.NewDescribeDBInstanceAutomatedBackupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBInstanceAutomatedBackupsPaginator(req DescribeDBInstanceAutomatedBackupsRequest) DescribeDBInstanceAutomatedBackupsPaginator {
	return DescribeDBInstanceAutomatedBackupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBInstanceAutomatedBackupsInput
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

// DescribeDBInstanceAutomatedBackupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBInstanceAutomatedBackupsPaginator struct {
	aws.Pager
}

func (p *DescribeDBInstanceAutomatedBackupsPaginator) CurrentPage() *DescribeDBInstanceAutomatedBackupsOutput {
	return p.Pager.CurrentPage().(*DescribeDBInstanceAutomatedBackupsOutput)
}

// DescribeDBInstanceAutomatedBackupsResponse is the response type for the
// DescribeDBInstanceAutomatedBackups API operation.
type DescribeDBInstanceAutomatedBackupsResponse struct {
	*DescribeDBInstanceAutomatedBackupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBInstanceAutomatedBackups request.
func (r *DescribeDBInstanceAutomatedBackupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
