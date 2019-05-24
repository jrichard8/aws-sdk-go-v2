// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetDiscoveredResourceCountsRequest
type GetDiscoveredResourceCountsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of ResourceCount objects returned on each page. The default
	// is 100. You cannot specify a number greater than 100. If you specify 0, AWS
	// Config uses the default.
	Limit *int64 `locationName:"limit" type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The comma-separated list that specifies the resource types that you want
	// AWS Config to return (for example, "AWS::EC2::Instance", "AWS::IAM::User").
	//
	// If a value for resourceTypes is not specified, AWS Config returns all resource
	// types that AWS Config is recording in the region for your account.
	//
	// If the configuration recorder is turned off, AWS Config returns an empty
	// list of ResourceCount objects. If the configuration recorder is not recording
	// a specific resource type (for example, S3 buckets), that resource type is
	// not returned in the list of ResourceCount objects.
	ResourceTypes []string `locationName:"resourceTypes" type:"list"`
}

// String returns the string representation
func (s GetDiscoveredResourceCountsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetDiscoveredResourceCountsResponse
type GetDiscoveredResourceCountsOutput struct {
	_ struct{} `type:"structure"`

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of ResourceCount objects. Each object is listed in descending order
	// by the number of resources.
	ResourceCounts []ResourceCount `locationName:"resourceCounts" type:"list"`

	// The total number of resources that AWS Config is recording in the region
	// for your account. If you specify resource types in the request, AWS Config
	// returns only the total number of resources for those resource types.
	//
	// Example
	//
	// AWS Config is recording three resource types in the US East (Ohio) Region
	// for your account: 25 EC2 instances, 20 IAM users, and 15 S3 buckets, for
	// a total of 60 resources.
	//
	// You make a call to the GetDiscoveredResourceCounts action and specify the
	// resource type, "AWS::EC2::Instances", in the request.
	//
	// AWS Config returns 25 for totalDiscoveredResources.
	TotalDiscoveredResources *int64 `locationName:"totalDiscoveredResources" type:"long"`
}

// String returns the string representation
func (s GetDiscoveredResourceCountsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDiscoveredResourceCounts = "GetDiscoveredResourceCounts"

// GetDiscoveredResourceCountsRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the resource types, the number of each resource type, and the total
// number of resources that AWS Config is recording in this region for your
// AWS account.
//
// Example
//
// AWS Config is recording three resource types in the US East (Ohio) Region
// for your account: 25 EC2 instances, 20 IAM users, and 15 S3 buckets.
//
// You make a call to the GetDiscoveredResourceCounts action and specify that
// you want all resource types.
//
// AWS Config returns the following:
//
// The resource types (EC2 instances, IAM users, and S3 buckets).
//
// The number of each resource type (25, 20, and 15).
//
// The total number of all resources (60).
//
// The response is paginated. By default, AWS Config lists 100 ResourceCount
// objects on each page. You can customize this number with the limit parameter.
// The response includes a nextToken string. To get the next page of results,
// run the request again and specify the string for the nextToken parameter.
//
// If you make a call to the GetDiscoveredResourceCounts action, you might not
// immediately receive resource counts in the following situations:
//
// You are a new AWS Config customer.
//
// You just enabled resource recording.
//
// It might take a few minutes for AWS Config to record and count your resources.
// Wait a few minutes and then retry the GetDiscoveredResourceCounts action.
//
//    // Example sending a request using GetDiscoveredResourceCountsRequest.
//    req := client.GetDiscoveredResourceCountsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetDiscoveredResourceCounts
func (c *Client) GetDiscoveredResourceCountsRequest(input *GetDiscoveredResourceCountsInput) GetDiscoveredResourceCountsRequest {
	op := &aws.Operation{
		Name:       opGetDiscoveredResourceCounts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDiscoveredResourceCountsInput{}
	}

	req := c.newRequest(op, input, &GetDiscoveredResourceCountsOutput{})
	return GetDiscoveredResourceCountsRequest{Request: req, Input: input, Copy: c.GetDiscoveredResourceCountsRequest}
}

// GetDiscoveredResourceCountsRequest is the request type for the
// GetDiscoveredResourceCounts API operation.
type GetDiscoveredResourceCountsRequest struct {
	*aws.Request
	Input *GetDiscoveredResourceCountsInput
	Copy  func(*GetDiscoveredResourceCountsInput) GetDiscoveredResourceCountsRequest
}

// Send marshals and sends the GetDiscoveredResourceCounts API request.
func (r GetDiscoveredResourceCountsRequest) Send(ctx context.Context) (*GetDiscoveredResourceCountsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDiscoveredResourceCountsResponse{
		GetDiscoveredResourceCountsOutput: r.Request.Data.(*GetDiscoveredResourceCountsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDiscoveredResourceCountsResponse is the response type for the
// GetDiscoveredResourceCounts API operation.
type GetDiscoveredResourceCountsResponse struct {
	*GetDiscoveredResourceCountsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDiscoveredResourceCounts request.
func (r *GetDiscoveredResourceCountsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}