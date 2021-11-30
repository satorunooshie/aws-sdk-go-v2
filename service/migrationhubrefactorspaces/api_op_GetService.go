// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubrefactorspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets an Amazon Web Services Migration Hub Refactor Spaces service.
func (c *Client) GetService(ctx context.Context, params *GetServiceInput, optFns ...func(*Options)) (*GetServiceOutput, error) {
	if params == nil {
		params = &GetServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetService", params, optFns, c.addOperationGetServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceInput struct {

	// The ID of the application.
	//
	// This member is required.
	ApplicationIdentifier *string

	// The ID of the environment.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The ID of the service.
	//
	// This member is required.
	ServiceIdentifier *string

	noSmithyDocumentSerde
}

type GetServiceOutput struct {

	// The ID of the application.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the service.
	Arn *string

	// The Amazon Web Services account ID of the service creator.
	CreatedByAccountId *string

	// The timestamp of when the service is created.
	CreatedTime *time.Time

	// The description of the service.
	Description *string

	// The endpoint type of the service.
	EndpointType types.ServiceEndpointType

	// The unique identifier of the environment.
	EnvironmentId *string

	// Any error associated with the service resource.
	Error *types.ErrorResponse

	// The configuration for the Lambda endpoint type. The Arn is the Amazon Resource
	// Name (ARN) of the Lambda function associated with this service.
	LambdaEndpoint *types.LambdaEndpointConfig

	// A timestamp that indicates when the service was last updated.
	LastUpdatedTime *time.Time

	// The name of the service.
	Name *string

	// The Amazon Web Services account ID of the service owner.
	OwnerAccountId *string

	// The unique identifier of the service.
	ServiceId *string

	// The current state of the service.
	State types.ServiceState

	// The tags assigned to the service. A tag is a label that you assign to an Amazon
	// Web Services resource. Each tag consists of a key-value pair.
	Tags map[string]string

	// The configuration for the URL endpoint type. The Url isthe URL of the endpoint
	// type. The HealthUrl is the health check URL of the endpoint type.
	UrlEndpoint *types.UrlEndpointConfig

	// The ID of the virtual private cloud (VPC).
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetService{}, middleware.After)
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
	if err = addOpGetServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "refactor-spaces",
		OperationName: "GetService",
	}
}
