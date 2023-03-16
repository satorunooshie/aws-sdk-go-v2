// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List studios in your Amazon Web Services accounts in the requested Amazon Web
// Services Region.
func (c *Client) ListStudios(ctx context.Context, params *ListStudiosInput, optFns ...func(*Options)) (*ListStudiosOutput, error) {
	if params == nil {
		params = &ListStudiosInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStudios", params, optFns, c.addOperationListStudiosMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStudiosOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStudiosInput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStudiosOutput struct {

	// A collection of studios.
	//
	// This member is required.
	Studios []types.Studio

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStudiosMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStudios{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStudios{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStudios(options.Region), middleware.Before); err != nil {
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

// ListStudiosAPIClient is a client that implements the ListStudios operation.
type ListStudiosAPIClient interface {
	ListStudios(context.Context, *ListStudiosInput, ...func(*Options)) (*ListStudiosOutput, error)
}

var _ ListStudiosAPIClient = (*Client)(nil)

// ListStudiosPaginatorOptions is the paginator options for ListStudios
type ListStudiosPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStudiosPaginator is a paginator for ListStudios
type ListStudiosPaginator struct {
	options   ListStudiosPaginatorOptions
	client    ListStudiosAPIClient
	params    *ListStudiosInput
	nextToken *string
	firstPage bool
}

// NewListStudiosPaginator returns a new ListStudiosPaginator
func NewListStudiosPaginator(client ListStudiosAPIClient, params *ListStudiosInput, optFns ...func(*ListStudiosPaginatorOptions)) *ListStudiosPaginator {
	if params == nil {
		params = &ListStudiosInput{}
	}

	options := ListStudiosPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStudiosPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStudiosPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStudios page.
func (p *ListStudiosPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStudiosOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListStudios(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStudios(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "ListStudios",
	}
}
