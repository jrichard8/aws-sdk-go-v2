// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a core network Connect peer for a specified core network connect
// attachment between a core network and an appliance. The peer address and transit
// gateway address must be the same IP address family (IPv4 or IPv6).
func (c *Client) CreateConnectPeer(ctx context.Context, params *CreateConnectPeerInput, optFns ...func(*Options)) (*CreateConnectPeerOutput, error) {
	if params == nil {
		params = &CreateConnectPeerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnectPeer", params, optFns, c.addOperationCreateConnectPeerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectPeerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectPeerInput struct {

	// The ID of the connection attachment.
	//
	// This member is required.
	ConnectAttachmentId *string

	// The inside IP addresses used for BGP peering.
	//
	// This member is required.
	InsideCidrBlocks []string

	// The Connect peer address.
	//
	// This member is required.
	PeerAddress *string

	// The Connect peer BGP options.
	BgpOptions *types.BgpOptions

	// The client token associated with the request.
	ClientToken *string

	// A Connect peer core network address.
	CoreNetworkAddress *string

	// The tags associated with the peer request.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateConnectPeerOutput struct {

	// The response to the request.
	ConnectPeer *types.ConnectPeer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConnectPeerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConnectPeer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConnectPeer{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateConnectPeerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateConnectPeerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnectPeer(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateConnectPeer struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateConnectPeer) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateConnectPeer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateConnectPeerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateConnectPeerInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateConnectPeerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateConnectPeer{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateConnectPeer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "CreateConnectPeer",
	}
}
