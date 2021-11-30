// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a time series (data stream) with an asset property.
func (c *Client) AssociateTimeSeriesToAssetProperty(ctx context.Context, params *AssociateTimeSeriesToAssetPropertyInput, optFns ...func(*Options)) (*AssociateTimeSeriesToAssetPropertyOutput, error) {
	if params == nil {
		params = &AssociateTimeSeriesToAssetPropertyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateTimeSeriesToAssetProperty", params, optFns, c.addOperationAssociateTimeSeriesToAssetPropertyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateTimeSeriesToAssetPropertyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTimeSeriesToAssetPropertyInput struct {

	// The alias that identifies the time series.
	//
	// This member is required.
	Alias *string

	// The ID of the asset in which the asset property was created.
	//
	// This member is required.
	AssetId *string

	// The ID of the asset property.
	//
	// This member is required.
	PropertyId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	noSmithyDocumentSerde
}

type AssociateTimeSeriesToAssetPropertyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateTimeSeriesToAssetPropertyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateTimeSeriesToAssetProperty{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateTimeSeriesToAssetProperty{}, middleware.After)
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
	if err = addEndpointPrefix_opAssociateTimeSeriesToAssetPropertyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opAssociateTimeSeriesToAssetPropertyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAssociateTimeSeriesToAssetPropertyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTimeSeriesToAssetProperty(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opAssociateTimeSeriesToAssetPropertyMiddleware struct {
}

func (*endpointPrefix_opAssociateTimeSeriesToAssetPropertyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opAssociateTimeSeriesToAssetPropertyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opAssociateTimeSeriesToAssetPropertyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opAssociateTimeSeriesToAssetPropertyMiddleware{}, `OperationSerializer`, middleware.After)
}

type idempotencyToken_initializeOpAssociateTimeSeriesToAssetProperty struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAssociateTimeSeriesToAssetProperty) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAssociateTimeSeriesToAssetProperty) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AssociateTimeSeriesToAssetPropertyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AssociateTimeSeriesToAssetPropertyInput ")
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
func addIdempotencyToken_opAssociateTimeSeriesToAssetPropertyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpAssociateTimeSeriesToAssetProperty{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAssociateTimeSeriesToAssetProperty(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "AssociateTimeSeriesToAssetProperty",
	}
}
