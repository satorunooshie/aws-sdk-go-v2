// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists a game server groups.
func (c *Client) ListGameServerGroups(ctx context.Context, params *ListGameServerGroupsInput, optFns ...func(*Options)) (*ListGameServerGroupsOutput, error) {
	if params == nil {
		params = &ListGameServerGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGameServerGroups", params, optFns, c.addOperationListGameServerGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGameServerGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGameServerGroupsInput struct {

	// The game server groups' limit.
	Limit *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGameServerGroupsOutput struct {

	// The game server groups' game server groups.
	GameServerGroups []types.GameServerGroup

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGameServerGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGameServerGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGameServerGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGameServerGroups(options.Region), middleware.Before); err != nil {
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

// ListGameServerGroupsAPIClient is a client that implements the
// ListGameServerGroups operation.
type ListGameServerGroupsAPIClient interface {
	ListGameServerGroups(context.Context, *ListGameServerGroupsInput, ...func(*Options)) (*ListGameServerGroupsOutput, error)
}

var _ ListGameServerGroupsAPIClient = (*Client)(nil)

// ListGameServerGroupsPaginatorOptions is the paginator options for
// ListGameServerGroups
type ListGameServerGroupsPaginatorOptions struct {
	// The game server groups' limit.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGameServerGroupsPaginator is a paginator for ListGameServerGroups
type ListGameServerGroupsPaginator struct {
	options   ListGameServerGroupsPaginatorOptions
	client    ListGameServerGroupsAPIClient
	params    *ListGameServerGroupsInput
	nextToken *string
	firstPage bool
}

// NewListGameServerGroupsPaginator returns a new ListGameServerGroupsPaginator
func NewListGameServerGroupsPaginator(client ListGameServerGroupsAPIClient, params *ListGameServerGroupsInput, optFns ...func(*ListGameServerGroupsPaginatorOptions)) *ListGameServerGroupsPaginator {
	if params == nil {
		params = &ListGameServerGroupsInput{}
	}

	options := ListGameServerGroupsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGameServerGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGameServerGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGameServerGroups page.
func (p *ListGameServerGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGameServerGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListGameServerGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGameServerGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListGameServerGroups",
	}
}
