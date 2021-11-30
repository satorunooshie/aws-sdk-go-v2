// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdeviceadvisor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a Device Advisor test suite. Requires permission to
// access the GetSuiteDefinition
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) GetSuiteDefinition(ctx context.Context, params *GetSuiteDefinitionInput, optFns ...func(*Options)) (*GetSuiteDefinitionOutput, error) {
	if params == nil {
		params = &GetSuiteDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSuiteDefinition", params, optFns, c.addOperationGetSuiteDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSuiteDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSuiteDefinitionInput struct {

	// Suite definition ID of the test suite to get.
	//
	// This member is required.
	SuiteDefinitionId *string

	// Suite definition version of the test suite to get.
	SuiteDefinitionVersion *string

	noSmithyDocumentSerde
}

type GetSuiteDefinitionOutput struct {

	// Date (in Unix epoch time) when the suite definition was created.
	CreatedAt *time.Time

	// Date (in Unix epoch time) when the suite definition was last modified.
	LastModifiedAt *time.Time

	// Latest suite definition version of the suite definition.
	LatestVersion *string

	// The ARN of the suite definition.
	SuiteDefinitionArn *string

	// Suite configuration of the suite definition.
	SuiteDefinitionConfiguration *types.SuiteDefinitionConfiguration

	// Suite definition ID of the suite definition.
	SuiteDefinitionId *string

	// Suite definition version of the suite definition.
	SuiteDefinitionVersion *string

	// Tags attached to the suite definition.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSuiteDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSuiteDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSuiteDefinition{}, middleware.After)
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
	if err = addOpGetSuiteDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSuiteDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSuiteDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdeviceadvisor",
		OperationName: "GetSuiteDefinition",
	}
}
