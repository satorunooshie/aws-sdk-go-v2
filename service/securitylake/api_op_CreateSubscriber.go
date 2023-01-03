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

// Creates a subscription permission for accounts that are already enabled in
// Amazon Security Lake. You can create a subscriber with access to data in the
// current Amazon Web Services Region.
func (c *Client) CreateSubscriber(ctx context.Context, params *CreateSubscriberInput, optFns ...func(*Options)) (*CreateSubscriberOutput, error) {
	if params == nil {
		params = &CreateSubscriberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubscriber", params, optFns, c.addOperationCreateSubscriberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubscriberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubscriberInput struct {

	// The Amazon Web Services account ID used to access your data.
	//
	// This member is required.
	AccountId *string

	// The external ID of the subscriber. This lets the user that is assuming the role
	// assert the circumstances in which they are operating. It also provides a way for
	// the account owner to permit the role to be assumed only under specific
	// circumstances.
	//
	// This member is required.
	ExternalId *string

	// The supported Amazon Web Services from which logs and events are collected.
	// Security Lake supports log and event collection for natively supported Amazon
	// Web Services.
	//
	// This member is required.
	SourceTypes []types.SourceType

	// The name of your Security Lake subscriber account.
	//
	// This member is required.
	SubscriberName *string

	// The Amazon S3 or Lake Formation access type.
	AccessTypes []types.AccessType

	// The description for your subscriber account in Security Lake.
	SubscriberDescription *string

	noSmithyDocumentSerde
}

type CreateSubscriberOutput struct {

	// The subscriptionId created by the CreateSubscriber API call.
	//
	// This member is required.
	SubscriptionId *string

	// The Amazon Resource Name (ARN) created by you to provide to the subscriber. For
	// more information about ARNs and how to use them in policies, see IAM identifiers
	// in the Identity and Access Management (IAM) User Guide
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html). .
	RoleArn *string

	// The ARN for the Amazon S3 bucket.
	S3BucketArn *string

	// The ARN for the Amazon Simple Notification Service.
	SnsArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubscriberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSubscriber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSubscriber{}, middleware.After)
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
	if err = addOpCreateSubscriberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubscriber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSubscriber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateSubscriber",
	}
}
