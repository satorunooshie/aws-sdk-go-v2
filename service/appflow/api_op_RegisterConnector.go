// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a new custom connector with your Amazon Web Services account. Before
// you can register the connector, you must deploy the associated AWS lambda
// function in your account.
func (c *Client) RegisterConnector(ctx context.Context, params *RegisterConnectorInput, optFns ...func(*Options)) (*RegisterConnectorOutput, error) {
	if params == nil {
		params = &RegisterConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterConnector", params, optFns, c.addOperationRegisterConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterConnectorInput struct {

	// The name of the connector. The name is unique for each ConnectorRegistration in
	// your Amazon Web Services account.
	ConnectorLabel *string

	// The provisioning type of the connector. Currently the only supported value is
	// LAMBDA.
	ConnectorProvisioningConfig *types.ConnectorProvisioningConfig

	// The provisioning type of the connector. Currently the only supported value is
	// LAMBDA.
	ConnectorProvisioningType types.ConnectorProvisioningType

	// A description about the connector that's being registered.
	Description *string

	noSmithyDocumentSerde
}

type RegisterConnectorOutput struct {

	// The ARN of the connector being registered.
	ConnectorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterConnector{}, middleware.After)
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
	if err = addOpRegisterConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "RegisterConnector",
	}
}
