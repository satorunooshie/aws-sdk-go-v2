// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Automatically enables Amazon Security Lake for new member accounts in your
// organization. Security Lake is not automatically enabled for any existing member
// accounts in your organization.
func (c *Client) CreateDatalakeAutoEnable(ctx context.Context, params *CreateDatalakeAutoEnableInput, optFns ...func(*Options)) (*CreateDatalakeAutoEnableOutput, error) {
	if params == nil {
		params = &CreateDatalakeAutoEnableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatalakeAutoEnable", params, optFns, c.addOperationCreateDatalakeAutoEnableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatalakeAutoEnableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatalakeAutoEnableInput struct {

	// Enable Security Lake with the specified configuration settings to begin
	// collecting security data for new accounts in your organization.
	//
	// This member is required.
	ConfigurationForNewAccounts []types.AutoEnableNewRegionConfiguration

	noSmithyDocumentSerde
}

type CreateDatalakeAutoEnableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatalakeAutoEnableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDatalakeAutoEnable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDatalakeAutoEnable{}, middleware.After)
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
	if err = addOpCreateDatalakeAutoEnableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatalakeAutoEnable(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatalakeAutoEnable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateDatalakeAutoEnable",
	}
}
