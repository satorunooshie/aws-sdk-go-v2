// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of storage configs for the identified instance and
// resource type.
func (c *Client) ListInstanceStorageConfigs(ctx context.Context, params *ListInstanceStorageConfigsInput, optFns ...func(*Options)) (*ListInstanceStorageConfigsOutput, error) {
	if params == nil {
		params = &ListInstanceStorageConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstanceStorageConfigs", params, optFns, c.addOperationListInstanceStorageConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstanceStorageConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstanceStorageConfigsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// A valid resource type.
	//
	// This member is required.
	ResourceType types.InstanceStorageResourceType

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListInstanceStorageConfigsOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// A valid storage type.
	StorageConfigs []types.InstanceStorageConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInstanceStorageConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInstanceStorageConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInstanceStorageConfigs{}, middleware.After)
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
	if err = addOpListInstanceStorageConfigsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceStorageConfigs(options.Region), middleware.Before); err != nil {
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

// ListInstanceStorageConfigsAPIClient is a client that implements the
// ListInstanceStorageConfigs operation.
type ListInstanceStorageConfigsAPIClient interface {
	ListInstanceStorageConfigs(context.Context, *ListInstanceStorageConfigsInput, ...func(*Options)) (*ListInstanceStorageConfigsOutput, error)
}

var _ ListInstanceStorageConfigsAPIClient = (*Client)(nil)

// ListInstanceStorageConfigsPaginatorOptions is the paginator options for
// ListInstanceStorageConfigs
type ListInstanceStorageConfigsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstanceStorageConfigsPaginator is a paginator for
// ListInstanceStorageConfigs
type ListInstanceStorageConfigsPaginator struct {
	options   ListInstanceStorageConfigsPaginatorOptions
	client    ListInstanceStorageConfigsAPIClient
	params    *ListInstanceStorageConfigsInput
	nextToken *string
	firstPage bool
}

// NewListInstanceStorageConfigsPaginator returns a new
// ListInstanceStorageConfigsPaginator
func NewListInstanceStorageConfigsPaginator(client ListInstanceStorageConfigsAPIClient, params *ListInstanceStorageConfigsInput, optFns ...func(*ListInstanceStorageConfigsPaginatorOptions)) *ListInstanceStorageConfigsPaginator {
	if params == nil {
		params = &ListInstanceStorageConfigsInput{}
	}

	options := ListInstanceStorageConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInstanceStorageConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstanceStorageConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInstanceStorageConfigs page.
func (p *ListInstanceStorageConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstanceStorageConfigsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListInstanceStorageConfigs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListInstanceStorageConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListInstanceStorageConfigs",
	}
}
