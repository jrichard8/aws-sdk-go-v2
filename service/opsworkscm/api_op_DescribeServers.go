// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DescribeServersRequest
type DescribeServersInput struct {
	_ struct{} `type:"structure"`

	// This is not currently implemented for DescribeServers requests.
	MaxResults *int64 `min:"1" type:"integer"`

	// This is not currently implemented for DescribeServers requests.
	NextToken *string `type:"string"`

	// Describes the server with the specified ServerName.
	ServerName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeServersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeServersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.ServerName != nil && len(*s.ServerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DescribeServersResponse
type DescribeServersOutput struct {
	_ struct{} `type:"structure"`

	// This is not currently implemented for DescribeServers requests.
	NextToken *string `type:"string"`

	// Contains the response to a DescribeServers request.
	//
	// For Puppet Server:DescribeServersResponse$Servers$EngineAttributes contains
	// PUPPET_API_CA_CERT. This is the PEM-encoded CA certificate that is used by
	// the Puppet API over TCP port number 8140. The CA certificate is also used
	// to sign node certificates.
	Servers []Server `type:"list"`
}

// String returns the string representation
func (s DescribeServersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeServers = "DescribeServers"

// DescribeServersRequest returns a request value for making API operation for
// AWS OpsWorks for Chef Automate.
//
// Lists all configuration management servers that are identified with your
// account. Only the stored results from Amazon DynamoDB are returned. AWS OpsWorks
// CM does not query other services.
//
// This operation is synchronous.
//
// A ResourceNotFoundException is thrown when the server does not exist. A ValidationException
// is raised when parameters of the request are not valid.
//
//    // Example sending a request using DescribeServersRequest.
//    req := client.DescribeServersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DescribeServers
func (c *Client) DescribeServersRequest(input *DescribeServersInput) DescribeServersRequest {
	op := &aws.Operation{
		Name:       opDescribeServers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeServersInput{}
	}

	req := c.newRequest(op, input, &DescribeServersOutput{})
	return DescribeServersRequest{Request: req, Input: input, Copy: c.DescribeServersRequest}
}

// DescribeServersRequest is the request type for the
// DescribeServers API operation.
type DescribeServersRequest struct {
	*aws.Request
	Input *DescribeServersInput
	Copy  func(*DescribeServersInput) DescribeServersRequest
}

// Send marshals and sends the DescribeServers API request.
func (r DescribeServersRequest) Send(ctx context.Context) (*DescribeServersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeServersResponse{
		DescribeServersOutput: r.Request.Data.(*DescribeServersOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeServersResponse is the response type for the
// DescribeServers API operation.
type DescribeServersResponse struct {
	*DescribeServersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeServers request.
func (r *DescribeServersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}