// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of branches in a specified source repository.
func (c *Client) ListSourceRepositoryBranches(ctx context.Context, params *ListSourceRepositoryBranchesInput, optFns ...func(*Options)) (*ListSourceRepositoryBranchesOutput, error) {
	if params == nil {
		params = &ListSourceRepositoryBranchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSourceRepositoryBranches", params, optFns, c.addOperationListSourceRepositoryBranchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSourceRepositoryBranchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSourceRepositoryBranchesInput struct {

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the source repository.
	//
	// This member is required.
	SourceRepositoryName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	MaxResults *int32

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSourceRepositoryBranchesOutput struct {

	// Information about the source branches.
	//
	// This member is required.
	Items []types.ListSourceRepositoryBranchesItem

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSourceRepositoryBranchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSourceRepositoryBranches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSourceRepositoryBranches{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addBearerAuthSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListSourceRepositoryBranchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSourceRepositoryBranches(options.Region), middleware.Before); err != nil {
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

// ListSourceRepositoryBranchesAPIClient is a client that implements the
// ListSourceRepositoryBranches operation.
type ListSourceRepositoryBranchesAPIClient interface {
	ListSourceRepositoryBranches(context.Context, *ListSourceRepositoryBranchesInput, ...func(*Options)) (*ListSourceRepositoryBranchesOutput, error)
}

var _ ListSourceRepositoryBranchesAPIClient = (*Client)(nil)

// ListSourceRepositoryBranchesPaginatorOptions is the paginator options for
// ListSourceRepositoryBranches
type ListSourceRepositoryBranchesPaginatorOptions struct {
	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSourceRepositoryBranchesPaginator is a paginator for
// ListSourceRepositoryBranches
type ListSourceRepositoryBranchesPaginator struct {
	options   ListSourceRepositoryBranchesPaginatorOptions
	client    ListSourceRepositoryBranchesAPIClient
	params    *ListSourceRepositoryBranchesInput
	nextToken *string
	firstPage bool
}

// NewListSourceRepositoryBranchesPaginator returns a new
// ListSourceRepositoryBranchesPaginator
func NewListSourceRepositoryBranchesPaginator(client ListSourceRepositoryBranchesAPIClient, params *ListSourceRepositoryBranchesInput, optFns ...func(*ListSourceRepositoryBranchesPaginatorOptions)) *ListSourceRepositoryBranchesPaginator {
	if params == nil {
		params = &ListSourceRepositoryBranchesInput{}
	}

	options := ListSourceRepositoryBranchesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSourceRepositoryBranchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSourceRepositoryBranchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSourceRepositoryBranches page.
func (p *ListSourceRepositoryBranchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSourceRepositoryBranchesOutput, error) {
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

	result, err := p.client.ListSourceRepositoryBranches(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSourceRepositoryBranches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSourceRepositoryBranches",
	}
}
