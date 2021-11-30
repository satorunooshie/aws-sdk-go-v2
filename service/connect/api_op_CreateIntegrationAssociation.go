// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an AWS resource association with an Amazon Connect instance.
func (c *Client) CreateIntegrationAssociation(ctx context.Context, params *CreateIntegrationAssociationInput, optFns ...func(*Options)) (*CreateIntegrationAssociationOutput, error) {
	if params == nil {
		params = &CreateIntegrationAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIntegrationAssociation", params, optFns, c.addOperationCreateIntegrationAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIntegrationAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIntegrationAssociationInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The Amazon Resource Name (ARN) of the integration.
	//
	// This member is required.
	IntegrationArn *string

	// The type of information to be ingested.
	//
	// This member is required.
	IntegrationType types.IntegrationType

	// The name of the external application. This field is only required for the EVENT
	// integration type.
	SourceApplicationName *string

	// The URL for the external application. This field is only required for the EVENT
	// integration type.
	SourceApplicationUrl *string

	// The type of the data source. This field is only required for the EVENT
	// integration type.
	SourceType types.SourceType

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateIntegrationAssociationOutput struct {

	// The Amazon Resource Name (ARN) for the association.
	IntegrationAssociationArn *string

	// The identifier for the integration association.
	IntegrationAssociationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIntegrationAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIntegrationAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIntegrationAssociation{}, middleware.After)
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
	if err = addOpCreateIntegrationAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIntegrationAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIntegrationAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "CreateIntegrationAssociation",
	}
}
