// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Turns on Amazon Web Services Resource Explorer in the Amazon Web Services Region
// in which you called this operation by creating an index. Resource Explorer
// begins discovering the resources in this Region and stores the details about the
// resources in the index so that they can be queried by using the Search
// operation. You can create only one index in a Region. This operation creates
// only a local index. To promote the local index in one Amazon Web Services Region
// into the aggregator index for the Amazon Web Services account, use the
// UpdateIndexType operation. For more information, see Turning on cross-Region
// search by creating an aggregator index
// (https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html)
// in the Amazon Web Services Resource Explorer User Guide. For more details about
// what happens when you turn on Resource Explorer in an Amazon Web Services
// Region, see Turn on Resource Explorer to index your resources in an Amazon Web
// Services Region
// (https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-service-activate.html)
// in the Amazon Web Services Resource Explorer User Guide. If this is the first
// Amazon Web Services Region in which you've created an index for Resource
// Explorer, then this operation also creates a service-linked role
// (https://docs.aws.amazon.com/resource-explorer/latest/userguide/security_iam_service-linked-roles.html)
// in your Amazon Web Services account that allows Resource Explorer to enumerate
// your resources to populate the index.
//
// * Action: resource-explorer-2:CreateIndex
// Resource: The ARN of the index (as it will exist after the operation completes)
// in the Amazon Web Services Region and account in which you're trying to create
// the index. Use the wildcard character (*) at the end of the string to match the
// eventual UUID. For example, the following Resource element restricts the role or
// user to creating an index in only the us-east-2 Region of the specified account.
// "Resource": "arn:aws:resource-explorer-2:us-west-2:<account-id>:index/*"
// Alternatively, you can use "Resource": "*" to allow the role or user to create
// an index in any Region.
//
// * Action: iam:CreateServiceLinkedRole Resource: No
// specific resource (*). This permission is required only the first time you
// create an index to turn on Resource Explorer in the account. Resource Explorer
// uses this to create the service-linked role needed to index the resources in
// your account
// (https://docs.aws.amazon.com/resource-explorer/latest/userguide/security_iam_service-linked-roles.html).
// Resource Explorer uses the same service-linked role for all additional indexes
// you create afterwards.
func (c *Client) CreateIndex(ctx context.Context, params *CreateIndexInput, optFns ...func(*Options)) (*CreateIndexOutput, error) {
	if params == nil {
		params = &CreateIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIndex", params, optFns, c.addOperationCreateIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIndexInput struct {

	// This value helps ensure idempotency. Resource Explorer uses this value to
	// prevent the accidental creation of duplicate versions. We recommend that you
	// generate a UUID-type value
	// (https://wikipedia.org/wiki/Universally_unique_identifier) to ensure the
	// uniqueness of your views.
	ClientToken *string

	// The specified tags are attached only to the index created in this Amazon Web
	// Services Region. The tags aren't attached to any of the resources listed in the
	// index.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateIndexOutput struct {

	// The ARN of the new local index for the Region. You can reference this ARN in IAM
	// permission policies to authorize the following operations: DeleteIndex |
	// GetIndex | UpdateIndexType | CreateView
	Arn *string

	// The date and timestamp when the index was created.
	CreatedAt *time.Time

	// Indicates the current state of the index. You can check for changes to the state
	// for asynchronous operations by calling the GetIndex operation. The state can
	// remain in the CREATING or UPDATING state for several hours as Resource Explorer
	// discovers the information about your resources and populates the index.
	State types.IndexState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIndex{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateIndexMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIndex(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateIndex struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIndex) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIndex) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIndexInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIndexInput ")
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
func addIdempotencyToken_opCreateIndexMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateIndex{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-explorer-2",
		OperationName: "CreateIndex",
	}
}
