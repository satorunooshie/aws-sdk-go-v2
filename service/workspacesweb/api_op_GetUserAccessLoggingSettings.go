// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets user access logging settings.
func (c *Client) GetUserAccessLoggingSettings(ctx context.Context, params *GetUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*GetUserAccessLoggingSettingsOutput, error) {
	if params == nil {
		params = &GetUserAccessLoggingSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUserAccessLoggingSettings", params, optFns, c.addOperationGetUserAccessLoggingSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserAccessLoggingSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserAccessLoggingSettingsInput struct {

	// The ARN of the user access logging settings.
	//
	// This member is required.
	UserAccessLoggingSettingsArn *string

	noSmithyDocumentSerde
}

type GetUserAccessLoggingSettingsOutput struct {

	// The user access logging settings.
	UserAccessLoggingSettings *types.UserAccessLoggingSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUserAccessLoggingSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUserAccessLoggingSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUserAccessLoggingSettings{}, middleware.After)
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
	if err = addOpGetUserAccessLoggingSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserAccessLoggingSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUserAccessLoggingSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces-web",
		OperationName: "GetUserAccessLoggingSettings",
	}
}
