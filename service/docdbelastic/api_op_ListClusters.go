// Code generated by smithy-go-codegen DO NOT EDIT.

package docdbelastic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdbelastic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about provisioned Elastic DocumentDB clusters.
func (c *Client) ListClusters(ctx context.Context, params *ListClustersInput, optFns ...func(*Options)) (*ListClustersOutput, error) {
	if params == nil {
		params = &ListClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusters", params, optFns, c.addOperationListClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClustersInput struct {

	// The maximum number of entries to recieve in the response.
	MaxResults *int32

	// The nextToken which is used the get the next page of data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListClustersOutput struct {

	// A list of Elastic DocumentDB cluster.
	Clusters []types.ClusterInList

	// The response will provide a nextToken if there is more data beyond the
	// maxResults. If there is no more data in the responce, the nextToken will not be
	// returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListClusters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusters(options.Region), middleware.Before); err != nil {
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

// ListClustersAPIClient is a client that implements the ListClusters operation.
type ListClustersAPIClient interface {
	ListClusters(context.Context, *ListClustersInput, ...func(*Options)) (*ListClustersOutput, error)
}

var _ ListClustersAPIClient = (*Client)(nil)

// ListClustersPaginatorOptions is the paginator options for ListClusters
type ListClustersPaginatorOptions struct {
	// The maximum number of entries to recieve in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClustersPaginator is a paginator for ListClusters
type ListClustersPaginator struct {
	options   ListClustersPaginatorOptions
	client    ListClustersAPIClient
	params    *ListClustersInput
	nextToken *string
	firstPage bool
}

// NewListClustersPaginator returns a new ListClustersPaginator
func NewListClustersPaginator(client ListClustersAPIClient, params *ListClustersInput, optFns ...func(*ListClustersPaginatorOptions)) *ListClustersPaginator {
	if params == nil {
		params = &ListClustersInput{}
	}

	options := ListClustersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClustersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListClusters page.
func (p *ListClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClustersOutput, error) {
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

	result, err := p.client.ListClusters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "docdb-elastic",
		OperationName: "ListClusters",
	}
}
