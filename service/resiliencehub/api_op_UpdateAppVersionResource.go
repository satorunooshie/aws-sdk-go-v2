// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the resource details in the AWS Resilience Hub application.
//
// * This
// action has no effect outside AWS Resilience Hub.
//
// * This API updates the AWS
// Resilience Hub application draft version. To use this resource for running
// resiliency assessments, you must publish the AWS Resilience Hub application
// using the PublishAppVersion API.
//
// * To update application version with new
// physicalResourceID, you must call ResolveAppVersionResources API.
func (c *Client) UpdateAppVersionResource(ctx context.Context, params *UpdateAppVersionResourceInput, optFns ...func(*Options)) (*UpdateAppVersionResourceOutput, error) {
	if params == nil {
		params = &UpdateAppVersionResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAppVersionResource", params, optFns, c.addOperationUpdateAppVersionResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAppVersionResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAppVersionResourceInput struct {

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// Currently, there is no supported additional information for resources.
	AdditionalInfo map[string][]string

	// The list of Application Components that this resource belongs to. If an
	// Application Component is not part of the AWS Resilience Hub application, it will
	// be added.
	AppComponents []string

	// The Amazon Web Services account that owns the physical resource.
	AwsAccountId *string

	// The Amazon Web Services region that owns the physical resource.
	AwsRegion *string

	// Indicates if a resource is excluded from an AWS Resilience Hub application. You
	// can exclude only imported resources from an AWS Resilience Hub application.
	Excluded *bool

	// The logical identifier of the resource.
	LogicalResourceId *types.LogicalResourceId

	// The physical identifier of the resource.
	PhysicalResourceId *string

	// The name of the resource.
	ResourceName *string

	// The type of resource.
	ResourceType *string

	noSmithyDocumentSerde
}

type UpdateAppVersionResourceOutput struct {

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The AWS Resilience Hub application version.
	//
	// This member is required.
	AppVersion *string

	// Defines a physical resource. A physical resource is a resource that exists in
	// your account. It can be identified using an Amazon Resource Name (ARN) or a
	// Resilience Hub-native identifier.
	PhysicalResource *types.PhysicalResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAppVersionResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAppVersionResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAppVersionResource{}, middleware.After)
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
	if err = addOpUpdateAppVersionResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAppVersionResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAppVersionResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "UpdateAppVersionResource",
	}
}
