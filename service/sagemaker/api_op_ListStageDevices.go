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
)

// Lists devices allocated to the stage, containing detailed device information and
// deployment status.
func (c *Client) ListStageDevices(ctx context.Context, params *ListStageDevicesInput, optFns ...func(*Options)) (*ListStageDevicesOutput, error) {
	if params == nil {
		params = &ListStageDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStageDevices", params, optFns, c.addOperationListStageDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStageDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStageDevicesInput struct {

	// The name of the edge deployment plan.
	//
	// This member is required.
	EdgeDeploymentPlanName *string

	// The name of the stage in the deployment.
	//
	// This member is required.
	StageName *string

	// Toggle for excluding devices deployed in other stages.
	ExcludeDevicesDeployedInOtherStage bool

	// The maximum number of requests to select.
	MaxResults *int32

	// The response from the last list when returning a list large enough to neeed
	// tokening.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStageDevicesOutput struct {

	// List of summaries of devices allocated to the stage.
	//
	// This member is required.
	DeviceDeploymentSummaries []types.DeviceDeploymentSummary

	// The token to use when calling the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStageDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStageDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStageDevices{}, middleware.After)
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
	if err = addOpListStageDevicesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStageDevices(options.Region), middleware.Before); err != nil {
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

// ListStageDevicesAPIClient is a client that implements the ListStageDevices
// operation.
type ListStageDevicesAPIClient interface {
	ListStageDevices(context.Context, *ListStageDevicesInput, ...func(*Options)) (*ListStageDevicesOutput, error)
}

var _ ListStageDevicesAPIClient = (*Client)(nil)

// ListStageDevicesPaginatorOptions is the paginator options for ListStageDevices
type ListStageDevicesPaginatorOptions struct {
	// The maximum number of requests to select.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStageDevicesPaginator is a paginator for ListStageDevices
type ListStageDevicesPaginator struct {
	options   ListStageDevicesPaginatorOptions
	client    ListStageDevicesAPIClient
	params    *ListStageDevicesInput
	nextToken *string
	firstPage bool
}

// NewListStageDevicesPaginator returns a new ListStageDevicesPaginator
func NewListStageDevicesPaginator(client ListStageDevicesAPIClient, params *ListStageDevicesInput, optFns ...func(*ListStageDevicesPaginatorOptions)) *ListStageDevicesPaginator {
	if params == nil {
		params = &ListStageDevicesInput{}
	}

	options := ListStageDevicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStageDevicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStageDevicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStageDevices page.
func (p *ListStageDevicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStageDevicesOutput, error) {
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

	result, err := p.client.ListStageDevices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStageDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListStageDevices",
	}
}
