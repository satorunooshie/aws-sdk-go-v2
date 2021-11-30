// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the given recommender, including its status. A recommender can be in
// one of the following states:
//
// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE
// -or- CREATE FAILED
//
// * DELETE PENDING > DELETE IN_PROGRESS
//
// When the status is
// CREATE FAILED, the response includes the failureReason key, which describes why.
// For more information on recommenders, see CreateRecommender
// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateRecommender.html).
func (c *Client) DescribeRecommender(ctx context.Context, params *DescribeRecommenderInput, optFns ...func(*Options)) (*DescribeRecommenderOutput, error) {
	if params == nil {
		params = &DescribeRecommenderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecommender", params, optFns, c.addOperationDescribeRecommenderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecommenderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecommenderInput struct {

	// The Amazon Resource Name (ARN) of the recommender to describe.
	//
	// This member is required.
	RecommenderArn *string

	noSmithyDocumentSerde
}

type DescribeRecommenderOutput struct {

	// The properties of the recommender.
	Recommender *types.Recommender

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecommenderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRecommender{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRecommender{}, middleware.After)
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
	if err = addOpDescribeRecommenderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecommender(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRecommender(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeRecommender",
	}
}
