// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26/GetResourcesInput
type GetResourcesInput struct {
	_ struct{} `type:"structure"`

	// A string that indicates that additional data is available. Leave this value
	// empty for your initial request. If the response includes a PaginationToken,
	// use that string for this value to request an additional page of data.
	PaginationToken *string `type:"string"`

	// The constraints on the resources that you want returned. The format of each
	// resource type is service[:resourceType]. For example, specifying a resource
	// type of ec2 returns all tagged Amazon EC2 resources (which includes tagged
	// EC2 instances). Specifying a resource type of ec2:instance returns only EC2
	// instances.
	//
	// The string for each service name and resource type is the same as that embedded
	// in a resource's Amazon Resource Name (ARN). Consult the AWS General Reference
	// for the following:
	//
	//    * For a list of service name strings, see AWS Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces).
	//
	//    * For resource type strings, see Example ARNs (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arns-syntax).
	//
	//    * For more information about ARNs, see Amazon Resource Names (ARNs) and
	//    AWS Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	ResourceTypeFilters []string `type:"list"`

	// A limit that restricts the number of resources returned by GetResources in
	// paginated output. You can set ResourcesPerPage to a minimum of 1 item and
	// the maximum of 50 items.
	ResourcesPerPage *int64 `type:"integer"`

	// A list of tags (keys and values). A request can include up to 50 keys, and
	// each key can include up to 20 values.
	//
	// If you specify multiple filters connected by an AND operator in a single
	// request, the response returns only those resources that are associated with
	// every specified filter.
	//
	// If you specify multiple filters connected by an OR operator in a single request,
	// the response returns all resources that are associated with at least one
	// or possibly more of the specified filters.
	TagFilters []TagFilter `type:"list"`

	// A limit that restricts the number of tags (key and value pairs) returned
	// by GetResources in paginated output. A resource with no tags is counted as
	// having one tag (one key and value pair).
	//
	// GetResources does not split a resource and its associated tags across pages.
	// If the specified TagsPerPage would cause such a break, a PaginationToken
	// is returned in place of the affected resource and its tags. Use that token
	// in another request to get the remaining data. For example, if you specify
	// a TagsPerPage of 100 and the account has 22 resources with 10 tags each (meaning
	// that each resource has 10 key and value pairs), the output will consist of
	// 3 pages, with the first page displaying the first 10 resources, each with
	// its 10 tags, the second page displaying the next 10 resources each with its
	// 10 tags, and the third page displaying the remaining 2 resources, each with
	// its 10 tags.
	//
	// You can set TagsPerPage to a minimum of 100 items and the maximum of 500
	// items.
	TagsPerPage *int64 `type:"integer"`
}

// String returns the string representation
func (s GetResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourcesInput"}
	if s.TagFilters != nil {
		for i, v := range s.TagFilters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagFilters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26/GetResourcesOutput
type GetResourcesOutput struct {
	_ struct{} `type:"structure"`

	// A string that indicates that the response contains more data than can be
	// returned in a single response. To receive additional data, specify this string
	// for the PaginationToken value in a subsequent request.
	PaginationToken *string `type:"string"`

	// A list of resource ARNs and the tags (keys and values) associated with each.
	ResourceTagMappingList []ResourceTagMapping `type:"list"`
}

// String returns the string representation
func (s GetResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetResources = "GetResources"

// GetResourcesRequest returns a request value for making API operation for
// AWS Resource Groups Tagging API.
//
// Returns all the tagged resources that are associated with the specified tags
// (keys and values) located in the specified region for the AWS account. The
// tags and the resource types that you specify in the request are known as
// filters. The response includes all tags that are associated with the requested
// resources. If no filter is provided, this action returns a paginated resource
// list with the associated tags.
//
//    // Example sending a request using GetResourcesRequest.
//    req := client.GetResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26/GetResources
func (c *Client) GetResourcesRequest(input *GetResourcesInput) GetResourcesRequest {
	op := &aws.Operation{
		Name:       opGetResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PaginationToken"},
			OutputTokens:    []string{"PaginationToken"},
			LimitToken:      "ResourcesPerPage",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetResourcesInput{}
	}

	req := c.newRequest(op, input, &GetResourcesOutput{})
	return GetResourcesRequest{Request: req, Input: input, Copy: c.GetResourcesRequest}
}

// GetResourcesRequest is the request type for the
// GetResources API operation.
type GetResourcesRequest struct {
	*aws.Request
	Input *GetResourcesInput
	Copy  func(*GetResourcesInput) GetResourcesRequest
}

// Send marshals and sends the GetResources API request.
func (r GetResourcesRequest) Send(ctx context.Context) (*GetResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourcesResponse{
		GetResourcesOutput: r.Request.Data.(*GetResourcesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetResourcesRequestPaginator returns a paginator for GetResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetResourcesRequest(input)
//   p := resourcegroupstaggingapi.NewGetResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetResourcesPaginator(req GetResourcesRequest) GetResourcesPaginator {
	return GetResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetResourcesInput
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

// GetResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetResourcesPaginator struct {
	aws.Pager
}

func (p *GetResourcesPaginator) CurrentPage() *GetResourcesOutput {
	return p.Pager.CurrentPage().(*GetResourcesOutput)
}

// GetResourcesResponse is the response type for the
// GetResources API operation.
type GetResourcesResponse struct {
	*GetResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResources request.
func (r *GetResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
