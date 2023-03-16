// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a security policy to be used by one or more OpenSearch Serverless
// collections. Security policies provide access to a collection and its OpenSearch
// Dashboards endpoint from public networks or specific VPC endpoints. They also
// allow you to secure a collection with a KMS encryption key. For more
// information, see Network access for Amazon OpenSearch Serverless
// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-network.html)
// and Encryption at rest for Amazon OpenSearch Serverless
// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-encryption.html).
func (c *Client) CreateSecurityPolicy(ctx context.Context, params *CreateSecurityPolicyInput, optFns ...func(*Options)) (*CreateSecurityPolicyOutput, error) {
	if params == nil {
		params = &CreateSecurityPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityPolicy", params, optFns, c.addOperationCreateSecurityPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityPolicyInput struct {

	// The name of the policy.
	//
	// This member is required.
	Name *string

	// The JSON policy document to use as the content for the new policy.
	//
	// This member is required.
	Policy *string

	// The type of security policy.
	//
	// This member is required.
	Type types.SecurityPolicyType

	// Unique, case-sensitive identifier to ensure idempotency of the request.
	ClientToken *string

	// A description of the policy. Typically used to store information about the
	// permissions defined in the policy.
	Description *string

	noSmithyDocumentSerde
}

type CreateSecurityPolicyOutput struct {

	// Details about the created security policy.
	SecurityPolicyDetail *types.SecurityPolicyDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSecurityPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateSecurityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateSecurityPolicy{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateSecurityPolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSecurityPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityPolicy(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateSecurityPolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSecurityPolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSecurityPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSecurityPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSecurityPolicyInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateSecurityPolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSecurityPolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSecurityPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aoss",
		OperationName: "CreateSecurityPolicy",
	}
}
