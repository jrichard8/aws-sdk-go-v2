// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the dataflow endpoint group.
func (c *Client) GetDataflowEndpointGroup(ctx context.Context, params *GetDataflowEndpointGroupInput, optFns ...func(*Options)) (*GetDataflowEndpointGroupOutput, error) {
	if params == nil {
		params = &GetDataflowEndpointGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataflowEndpointGroup", params, optFns, c.addOperationGetDataflowEndpointGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataflowEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataflowEndpointGroupInput struct {

	// UUID of a dataflow endpoint group.
	//
	// This member is required.
	DataflowEndpointGroupId *string

	noSmithyDocumentSerde
}

type GetDataflowEndpointGroupOutput struct {

	// Amount of time, in seconds, after a contact ends for the contact to remain in a
	// POSTPASS state. A CloudWatch event is emitted when the contact enters and exits
	// the POSTPASS state.
	ContactPostPassDurationSeconds *int32

	// Amount of time, in seconds, prior to contact start for the contact to remain in
	// a PREPASS state. A CloudWatch event is emitted when the contact enters and exits
	// the PREPASS state.
	ContactPrePassDurationSeconds *int32

	// ARN of a dataflow endpoint group.
	DataflowEndpointGroupArn *string

	// UUID of a dataflow endpoint group.
	DataflowEndpointGroupId *string

	// Details of a dataflow endpoint.
	EndpointsDetails []types.EndpointDetails

	// Tags assigned to a dataflow endpoint group.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataflowEndpointGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataflowEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataflowEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetDataflowEndpointGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataflowEndpointGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetDataflowEndpointGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "groundstation",
		OperationName: "GetDataflowEndpointGroup",
	}
}
