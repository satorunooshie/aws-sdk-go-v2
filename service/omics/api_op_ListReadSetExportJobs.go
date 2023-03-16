// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of read set export jobs.
func (c *Client) ListReadSetExportJobs(ctx context.Context, params *ListReadSetExportJobsInput, optFns ...func(*Options)) (*ListReadSetExportJobsOutput, error) {
	if params == nil {
		params = &ListReadSetExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReadSetExportJobs", params, optFns, c.addOperationListReadSetExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReadSetExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReadSetExportJobsInput struct {

	// The jobs' sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// A filter to apply to the list.
	Filter *types.ExportReadSetFilter

	// The maximum number of jobs to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReadSetExportJobsOutput struct {

	// A list of jobs.
	ExportJobs []types.ExportReadSetJobDetail

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReadSetExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReadSetExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReadSetExportJobs{}, middleware.After)
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
	if err = addEndpointPrefix_opListReadSetExportJobsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListReadSetExportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReadSetExportJobs(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListReadSetExportJobsMiddleware struct {
}

func (*endpointPrefix_opListReadSetExportJobsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListReadSetExportJobsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListReadSetExportJobsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListReadSetExportJobsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListReadSetExportJobsAPIClient is a client that implements the
// ListReadSetExportJobs operation.
type ListReadSetExportJobsAPIClient interface {
	ListReadSetExportJobs(context.Context, *ListReadSetExportJobsInput, ...func(*Options)) (*ListReadSetExportJobsOutput, error)
}

var _ ListReadSetExportJobsAPIClient = (*Client)(nil)

// ListReadSetExportJobsPaginatorOptions is the paginator options for
// ListReadSetExportJobs
type ListReadSetExportJobsPaginatorOptions struct {
	// The maximum number of jobs to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReadSetExportJobsPaginator is a paginator for ListReadSetExportJobs
type ListReadSetExportJobsPaginator struct {
	options   ListReadSetExportJobsPaginatorOptions
	client    ListReadSetExportJobsAPIClient
	params    *ListReadSetExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListReadSetExportJobsPaginator returns a new ListReadSetExportJobsPaginator
func NewListReadSetExportJobsPaginator(client ListReadSetExportJobsAPIClient, params *ListReadSetExportJobsInput, optFns ...func(*ListReadSetExportJobsPaginatorOptions)) *ListReadSetExportJobsPaginator {
	if params == nil {
		params = &ListReadSetExportJobsInput{}
	}

	options := ListReadSetExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReadSetExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReadSetExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReadSetExportJobs page.
func (p *ListReadSetExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReadSetExportJobsOutput, error) {
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

	result, err := p.client.ListReadSetExportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReadSetExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "ListReadSetExportJobs",
	}
}
