// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the FinSpace Changesets for a Dataset.
func (c *Client) ListChangesets(ctx context.Context, params *ListChangesetsInput, optFns ...func(*Options)) (*ListChangesetsOutput, error) {
	if params == nil {
		params = &ListChangesetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChangesets", params, optFns, c.addOperationListChangesetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChangesetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to ListChangesetsRequest. It exposes minimal query filters.
type ListChangesetsInput struct {

	// The unique identifier for the FinSpace Dataset to which the Changeset belongs.
	//
	// This member is required.
	DatasetId *string

	// The maximum number of results per page.
	MaxResults int32

	// A token indicating where a results page should begin.
	NextToken *string

	noSmithyDocumentSerde
}

// Response to ListChangesetsResponse. This returns a list of dataset changesets
// that match the query criteria.
type ListChangesetsOutput struct {

	// List of Changesets found.
	Changesets []types.ChangesetSummary

	// A token indicating where a results page should begin.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChangesetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChangesets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChangesets{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListChangesetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChangesets(options.Region), middleware.Before); err != nil {
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

// ListChangesetsAPIClient is a client that implements the ListChangesets
// operation.
type ListChangesetsAPIClient interface {
	ListChangesets(context.Context, *ListChangesetsInput, ...func(*Options)) (*ListChangesetsOutput, error)
}

var _ ListChangesetsAPIClient = (*Client)(nil)

// ListChangesetsPaginatorOptions is the paginator options for ListChangesets
type ListChangesetsPaginatorOptions struct {
	// The maximum number of results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChangesetsPaginator is a paginator for ListChangesets
type ListChangesetsPaginator struct {
	options   ListChangesetsPaginatorOptions
	client    ListChangesetsAPIClient
	params    *ListChangesetsInput
	nextToken *string
	firstPage bool
}

// NewListChangesetsPaginator returns a new ListChangesetsPaginator
func NewListChangesetsPaginator(client ListChangesetsAPIClient, params *ListChangesetsInput, optFns ...func(*ListChangesetsPaginatorOptions)) *ListChangesetsPaginator {
	if params == nil {
		params = &ListChangesetsInput{}
	}

	options := ListChangesetsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChangesetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChangesetsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListChangesets page.
func (p *ListChangesetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChangesetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListChangesets(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListChangesets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "ListChangesets",
	}
}
