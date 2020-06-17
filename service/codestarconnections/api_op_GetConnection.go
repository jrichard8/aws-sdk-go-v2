// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarconnections

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetConnectionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of a connection.
	//
	// ConnectionArn is a required field
	ConnectionArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConnectionInput"}

	if s.ConnectionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The connection details, such as status, owner, and provider type.
	Connection *Connection `type:"structure"`
}

// String returns the string representation
func (s GetConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetConnection = "GetConnection"

// GetConnectionRequest returns a request value for making API operation for
// AWS CodeStar connections.
//
// Returns the connection ARN and details such as status, owner, and provider
// type.
//
//    // Example sending a request using GetConnectionRequest.
//    req := client.GetConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/GetConnection
func (c *Client) GetConnectionRequest(input *GetConnectionInput) GetConnectionRequest {
	op := &aws.Operation{
		Name:       opGetConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConnectionInput{}
	}

	req := c.newRequest(op, input, &GetConnectionOutput{})

	return GetConnectionRequest{Request: req, Input: input, Copy: c.GetConnectionRequest}
}

// GetConnectionRequest is the request type for the
// GetConnection API operation.
type GetConnectionRequest struct {
	*aws.Request
	Input *GetConnectionInput
	Copy  func(*GetConnectionInput) GetConnectionRequest
}

// Send marshals and sends the GetConnection API request.
func (r GetConnectionRequest) Send(ctx context.Context) (*GetConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectionResponse{
		GetConnectionOutput: r.Request.Data.(*GetConnectionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConnectionResponse is the response type for the
// GetConnection API operation.
type GetConnectionResponse struct {
	*GetConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnection request.
func (r *GetConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}