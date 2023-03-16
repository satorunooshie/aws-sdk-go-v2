// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakergeospatial

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops the Vector Enrichment job for a given job ARN.
func (c *Client) StopVectorEnrichmentJob(ctx context.Context, params *StopVectorEnrichmentJobInput, optFns ...func(*Options)) (*StopVectorEnrichmentJobOutput, error) {
	if params == nil {
		params = &StopVectorEnrichmentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopVectorEnrichmentJob", params, optFns, c.addOperationStopVectorEnrichmentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopVectorEnrichmentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopVectorEnrichmentJobInput struct {

	// The Amazon Resource Name (ARN) of the Vector Enrichment job.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type StopVectorEnrichmentJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopVectorEnrichmentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopVectorEnrichmentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopVectorEnrichmentJob{}, middleware.After)
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
	if err = addOpStopVectorEnrichmentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopVectorEnrichmentJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopVectorEnrichmentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker-geospatial",
		OperationName: "StopVectorEnrichmentJob",
	}
}
