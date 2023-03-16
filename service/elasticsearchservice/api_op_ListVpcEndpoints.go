// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all Amazon OpenSearch Service-managed VPC endpoints in the current
// account and Region.
func (c *Client) ListVpcEndpoints(ctx context.Context, params *ListVpcEndpointsInput, optFns ...func(*Options)) (*ListVpcEndpointsOutput, error) {
	if params == nil {
		params = &ListVpcEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVpcEndpoints", params, optFns, c.addOperationListVpcEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVpcEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to the ListVpcEndpoints operation.
type ListVpcEndpointsInput struct {

	// Identifier to allow retrieval of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

// Container for response parameters to the ListVpcEndpoints operation. Returns a
// list containing summarized details of the VPC endpoints.
type ListVpcEndpointsOutput struct {

	// Provides an identifier to allow retrieval of paginated results.
	//
	// This member is required.
	NextToken *string

	// Information about each endpoint.
	//
	// This member is required.
	VpcEndpointSummaryList []types.VpcEndpointSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVpcEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVpcEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVpcEndpoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVpcEndpoints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListVpcEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListVpcEndpoints",
	}
}
