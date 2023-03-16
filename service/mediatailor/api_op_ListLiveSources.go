// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the live sources contained in a source location. A source represents a
// piece of content.
func (c *Client) ListLiveSources(ctx context.Context, params *ListLiveSourcesInput, optFns ...func(*Options)) (*ListLiveSourcesOutput, error) {
	if params == nil {
		params = &ListLiveSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLiveSources", params, optFns, c.addOperationListLiveSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLiveSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLiveSourcesInput struct {

	// The name of the source location associated with this Live Sources list.
	//
	// This member is required.
	SourceLocationName *string

	// The maximum number of live sources that you want MediaTailor to return in
	// response to the current request. If there are more than MaxResults live sources,
	// use the value of NextToken in the response to get the next page of results.
	MaxResults int32

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLiveSourcesOutput struct {

	// Lists the live sources.
	Items []types.LiveSource

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLiveSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLiveSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLiveSources{}, middleware.After)
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
	if err = addOpListLiveSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLiveSources(options.Region), middleware.Before); err != nil {
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

// ListLiveSourcesAPIClient is a client that implements the ListLiveSources
// operation.
type ListLiveSourcesAPIClient interface {
	ListLiveSources(context.Context, *ListLiveSourcesInput, ...func(*Options)) (*ListLiveSourcesOutput, error)
}

var _ ListLiveSourcesAPIClient = (*Client)(nil)

// ListLiveSourcesPaginatorOptions is the paginator options for ListLiveSources
type ListLiveSourcesPaginatorOptions struct {
	// The maximum number of live sources that you want MediaTailor to return in
	// response to the current request. If there are more than MaxResults live sources,
	// use the value of NextToken in the response to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLiveSourcesPaginator is a paginator for ListLiveSources
type ListLiveSourcesPaginator struct {
	options   ListLiveSourcesPaginatorOptions
	client    ListLiveSourcesAPIClient
	params    *ListLiveSourcesInput
	nextToken *string
	firstPage bool
}

// NewListLiveSourcesPaginator returns a new ListLiveSourcesPaginator
func NewListLiveSourcesPaginator(client ListLiveSourcesAPIClient, params *ListLiveSourcesInput, optFns ...func(*ListLiveSourcesPaginatorOptions)) *ListLiveSourcesPaginator {
	if params == nil {
		params = &ListLiveSourcesInput{}
	}

	options := ListLiveSourcesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLiveSourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLiveSourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLiveSources page.
func (p *ListLiveSourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLiveSourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListLiveSources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLiveSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "ListLiveSources",
	}
}
