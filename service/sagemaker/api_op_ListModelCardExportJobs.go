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

// List the export jobs for the Amazon SageMaker Model Card.
func (c *Client) ListModelCardExportJobs(ctx context.Context, params *ListModelCardExportJobsInput, optFns ...func(*Options)) (*ListModelCardExportJobsOutput, error) {
	if params == nil {
		params = &ListModelCardExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelCardExportJobs", params, optFns, c.addOperationListModelCardExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelCardExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelCardExportJobsInput struct {

	// List export jobs for the model card with the specified name.
	//
	// This member is required.
	ModelCardName *string

	// Only list model card export jobs that were created after the time specified.
	CreationTimeAfter *time.Time

	// Only list model card export jobs that were created before the time specified.
	CreationTimeBefore *time.Time

	// The maximum number of model card export jobs to list.
	MaxResults *int32

	// Only list model card export jobs with names that contain the specified string.
	ModelCardExportJobNameContains *string

	// List export jobs for the model card with the specified version.
	ModelCardVersion int32

	// If the response to a previous ListModelCardExportJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of model card export
	// jobs, use the token in the next request.
	NextToken *string

	// Sort model card export jobs by either name or creation time. Sorts by creation
	// time by default.
	SortBy types.ModelCardExportJobSortBy

	// Sort model card export jobs by ascending or descending order.
	SortOrder types.ModelCardExportJobSortOrder

	// Only list model card export jobs with the specified status.
	StatusEquals types.ModelCardExportJobStatus

	noSmithyDocumentSerde
}

type ListModelCardExportJobsOutput struct {

	// The summaries of the listed model card export jobs.
	//
	// This member is required.
	ModelCardExportJobSummaries []types.ModelCardExportJobSummary

	// If the response is truncated, SageMaker returns this token. To retrieve the next
	// set of model card export jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelCardExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModelCardExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModelCardExportJobs{}, middleware.After)
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
	if err = addOpListModelCardExportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelCardExportJobs(options.Region), middleware.Before); err != nil {
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

// ListModelCardExportJobsAPIClient is a client that implements the
// ListModelCardExportJobs operation.
type ListModelCardExportJobsAPIClient interface {
	ListModelCardExportJobs(context.Context, *ListModelCardExportJobsInput, ...func(*Options)) (*ListModelCardExportJobsOutput, error)
}

var _ ListModelCardExportJobsAPIClient = (*Client)(nil)

// ListModelCardExportJobsPaginatorOptions is the paginator options for
// ListModelCardExportJobs
type ListModelCardExportJobsPaginatorOptions struct {
	// The maximum number of model card export jobs to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelCardExportJobsPaginator is a paginator for ListModelCardExportJobs
type ListModelCardExportJobsPaginator struct {
	options   ListModelCardExportJobsPaginatorOptions
	client    ListModelCardExportJobsAPIClient
	params    *ListModelCardExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListModelCardExportJobsPaginator returns a new
// ListModelCardExportJobsPaginator
func NewListModelCardExportJobsPaginator(client ListModelCardExportJobsAPIClient, params *ListModelCardExportJobsInput, optFns ...func(*ListModelCardExportJobsPaginatorOptions)) *ListModelCardExportJobsPaginator {
	if params == nil {
		params = &ListModelCardExportJobsInput{}
	}

	options := ListModelCardExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelCardExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelCardExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelCardExportJobs page.
func (p *ListModelCardExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelCardExportJobsOutput, error) {
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

	result, err := p.client.ListModelCardExportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModelCardExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModelCardExportJobs",
	}
}
