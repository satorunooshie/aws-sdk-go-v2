// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a response headers policy. When you update a response headers policy,
// the entire policy is replaced. You cannot update some policy fields independent
// of others. To update a response headers policy configuration:
//
// * Use
// GetResponseHeadersPolicyConfig to get the current policy's configuration.
//
// *
// Modify the fields in the response headers policy configuration that you want to
// update.
//
// * Call UpdateResponseHeadersPolicy, providing the entire response
// headers policy configuration, including the fields that you modified and those
// that you didn't.
func (c *Client) UpdateResponseHeadersPolicy(ctx context.Context, params *UpdateResponseHeadersPolicyInput, optFns ...func(*Options)) (*UpdateResponseHeadersPolicyOutput, error) {
	if params == nil {
		params = &UpdateResponseHeadersPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResponseHeadersPolicy", params, optFns, c.addOperationUpdateResponseHeadersPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResponseHeadersPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResponseHeadersPolicyInput struct {

	// The identifier for the response headers policy that you are updating.
	//
	// This member is required.
	Id *string

	// A response headers policy configuration.
	//
	// This member is required.
	ResponseHeadersPolicyConfig *types.ResponseHeadersPolicyConfig

	// The version of the response headers policy that you are updating. The version is
	// returned in the cache policy's ETag field in the response to
	// GetResponseHeadersPolicyConfig.
	IfMatch *string

	noSmithyDocumentSerde
}

type UpdateResponseHeadersPolicyOutput struct {

	// The current version of the response headers policy.
	ETag *string

	// A response headers policy.
	ResponseHeadersPolicy *types.ResponseHeadersPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResponseHeadersPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateResponseHeadersPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateResponseHeadersPolicy{}, middleware.After)
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
	if err = addOpUpdateResponseHeadersPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResponseHeadersPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResponseHeadersPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateResponseHeadersPolicy",
	}
}
