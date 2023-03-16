// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A list of the pricing plans that are associated with a pricing rule.
func (c *Client) ListPricingPlansAssociatedWithPricingRule(ctx context.Context, params *ListPricingPlansAssociatedWithPricingRuleInput, optFns ...func(*Options)) (*ListPricingPlansAssociatedWithPricingRuleOutput, error) {
	if params == nil {
		params = &ListPricingPlansAssociatedWithPricingRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPricingPlansAssociatedWithPricingRule", params, optFns, c.addOperationListPricingPlansAssociatedWithPricingRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPricingPlansAssociatedWithPricingRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPricingPlansAssociatedWithPricingRuleInput struct {

	// The pricing rule Amazon Resource Name (ARN) for which associations will be
	// listed.
	//
	// This member is required.
	PricingRuleArn *string

	// The pricing plan billing period for which associations will be listed.
	BillingPeriod *string

	// The optional maximum number of pricing rule associations to retrieve.
	MaxResults *int32

	// The optional pagination token returned by a previous call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPricingPlansAssociatedWithPricingRuleOutput struct {

	// The pricing plan billing period for which associations will be listed.
	BillingPeriod *string

	// The pagination token to be used on subsequent calls.
	NextToken *string

	// The list containing pricing plans that are associated with the requested pricing
	// rule.
	PricingPlanArns []string

	// The pricing rule Amazon Resource Name (ARN) for which associations will be
	// listed.
	PricingRuleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPricingPlansAssociatedWithPricingRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPricingPlansAssociatedWithPricingRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPricingPlansAssociatedWithPricingRule{}, middleware.After)
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
	if err = addOpListPricingPlansAssociatedWithPricingRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPricingPlansAssociatedWithPricingRule(options.Region), middleware.Before); err != nil {
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

// ListPricingPlansAssociatedWithPricingRuleAPIClient is a client that implements
// the ListPricingPlansAssociatedWithPricingRule operation.
type ListPricingPlansAssociatedWithPricingRuleAPIClient interface {
	ListPricingPlansAssociatedWithPricingRule(context.Context, *ListPricingPlansAssociatedWithPricingRuleInput, ...func(*Options)) (*ListPricingPlansAssociatedWithPricingRuleOutput, error)
}

var _ ListPricingPlansAssociatedWithPricingRuleAPIClient = (*Client)(nil)

// ListPricingPlansAssociatedWithPricingRulePaginatorOptions is the paginator
// options for ListPricingPlansAssociatedWithPricingRule
type ListPricingPlansAssociatedWithPricingRulePaginatorOptions struct {
	// The optional maximum number of pricing rule associations to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPricingPlansAssociatedWithPricingRulePaginator is a paginator for
// ListPricingPlansAssociatedWithPricingRule
type ListPricingPlansAssociatedWithPricingRulePaginator struct {
	options   ListPricingPlansAssociatedWithPricingRulePaginatorOptions
	client    ListPricingPlansAssociatedWithPricingRuleAPIClient
	params    *ListPricingPlansAssociatedWithPricingRuleInput
	nextToken *string
	firstPage bool
}

// NewListPricingPlansAssociatedWithPricingRulePaginator returns a new
// ListPricingPlansAssociatedWithPricingRulePaginator
func NewListPricingPlansAssociatedWithPricingRulePaginator(client ListPricingPlansAssociatedWithPricingRuleAPIClient, params *ListPricingPlansAssociatedWithPricingRuleInput, optFns ...func(*ListPricingPlansAssociatedWithPricingRulePaginatorOptions)) *ListPricingPlansAssociatedWithPricingRulePaginator {
	if params == nil {
		params = &ListPricingPlansAssociatedWithPricingRuleInput{}
	}

	options := ListPricingPlansAssociatedWithPricingRulePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPricingPlansAssociatedWithPricingRulePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPricingPlansAssociatedWithPricingRulePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPricingPlansAssociatedWithPricingRule page.
func (p *ListPricingPlansAssociatedWithPricingRulePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPricingPlansAssociatedWithPricingRuleOutput, error) {
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

	result, err := p.client.ListPricingPlansAssociatedWithPricingRule(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPricingPlansAssociatedWithPricingRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "ListPricingPlansAssociatedWithPricingRule",
	}
}
