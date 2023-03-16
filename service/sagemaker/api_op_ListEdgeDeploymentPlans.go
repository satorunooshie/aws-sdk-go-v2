// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all edge deployment plans.
func (c *Client) ListEdgeDeploymentPlans(ctx context.Context, params *ListEdgeDeploymentPlansInput, optFns ...func(*Options)) (*ListEdgeDeploymentPlansOutput, error) {
	if params == nil {
		params = &ListEdgeDeploymentPlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEdgeDeploymentPlans", params, optFns, c.addOperationListEdgeDeploymentPlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEdgeDeploymentPlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEdgeDeploymentPlansInput struct {

	// Selects edge deployment plans created after this time.
	CreationTimeAfter *time.Time

	// Selects edge deployment plans created before this time.
	CreationTimeBefore *time.Time

	// Selects edge deployment plans with a device fleet name containing this name.
	DeviceFleetNameContains *string

	// Selects edge deployment plans that were last updated after this time.
	LastModifiedTimeAfter *time.Time

	// Selects edge deployment plans that were last updated before this time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of results to select (50 by default).
	MaxResults *int32

	// Selects edge deployment plans with names containing this name.
	NameContains *string

	// The response from the last list when returning a list large enough to need
	// tokening.
	NextToken *string

	// The column by which to sort the edge deployment plans. Can be one of NAME,
	// DEVICEFLEETNAME, CREATIONTIME, LASTMODIFIEDTIME.
	SortBy types.ListEdgeDeploymentPlansSortBy

	// The direction of the sorting (ascending or descending).
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListEdgeDeploymentPlansOutput struct {

	// List of summaries of edge deployment plans.
	//
	// This member is required.
	EdgeDeploymentPlanSummaries []types.EdgeDeploymentPlanSummary

	// The token to use when calling the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEdgeDeploymentPlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEdgeDeploymentPlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEdgeDeploymentPlans{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEdgeDeploymentPlans(options.Region), middleware.Before); err != nil {
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

// ListEdgeDeploymentPlansAPIClient is a client that implements the
// ListEdgeDeploymentPlans operation.
type ListEdgeDeploymentPlansAPIClient interface {
	ListEdgeDeploymentPlans(context.Context, *ListEdgeDeploymentPlansInput, ...func(*Options)) (*ListEdgeDeploymentPlansOutput, error)
}

var _ ListEdgeDeploymentPlansAPIClient = (*Client)(nil)

// ListEdgeDeploymentPlansPaginatorOptions is the paginator options for
// ListEdgeDeploymentPlans
type ListEdgeDeploymentPlansPaginatorOptions struct {
	// The maximum number of results to select (50 by default).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEdgeDeploymentPlansPaginator is a paginator for ListEdgeDeploymentPlans
type ListEdgeDeploymentPlansPaginator struct {
	options   ListEdgeDeploymentPlansPaginatorOptions
	client    ListEdgeDeploymentPlansAPIClient
	params    *ListEdgeDeploymentPlansInput
	nextToken *string
	firstPage bool
}

// NewListEdgeDeploymentPlansPaginator returns a new
// ListEdgeDeploymentPlansPaginator
func NewListEdgeDeploymentPlansPaginator(client ListEdgeDeploymentPlansAPIClient, params *ListEdgeDeploymentPlansInput, optFns ...func(*ListEdgeDeploymentPlansPaginatorOptions)) *ListEdgeDeploymentPlansPaginator {
	if params == nil {
		params = &ListEdgeDeploymentPlansInput{}
	}

	options := ListEdgeDeploymentPlansPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEdgeDeploymentPlansPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEdgeDeploymentPlansPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEdgeDeploymentPlans page.
func (p *ListEdgeDeploymentPlansPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEdgeDeploymentPlansOutput, error) {
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

	result, err := p.client.ListEdgeDeploymentPlans(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEdgeDeploymentPlans(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListEdgeDeploymentPlans",
	}
}
