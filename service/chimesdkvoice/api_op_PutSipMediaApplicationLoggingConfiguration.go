// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) PutSipMediaApplicationLoggingConfiguration(ctx context.Context, params *PutSipMediaApplicationLoggingConfigurationInput, optFns ...func(*Options)) (*PutSipMediaApplicationLoggingConfigurationOutput, error) {
	if params == nil {
		params = &PutSipMediaApplicationLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSipMediaApplicationLoggingConfiguration", params, optFns, c.addOperationPutSipMediaApplicationLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSipMediaApplicationLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSipMediaApplicationLoggingConfigurationInput struct {

	// This member is required.
	SipMediaApplicationId *string

	SipMediaApplicationLoggingConfiguration *types.SipMediaApplicationLoggingConfiguration

	noSmithyDocumentSerde
}

type PutSipMediaApplicationLoggingConfigurationOutput struct {
	SipMediaApplicationLoggingConfiguration *types.SipMediaApplicationLoggingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSipMediaApplicationLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSipMediaApplicationLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSipMediaApplicationLoggingConfiguration{}, middleware.After)
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
	if err = addOpPutSipMediaApplicationLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSipMediaApplicationLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSipMediaApplicationLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutSipMediaApplicationLoggingConfiguration",
	}
}
