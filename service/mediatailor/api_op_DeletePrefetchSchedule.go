// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a prefetch schedule for a specific playback configuration. If you call
// DeletePrefetchSchedule on an expired prefetch schedule, MediaTailor returns an
// HTTP 404 status code. For more information about ad prefetching, see Using ad
// prefetching
// (https://docs.aws.amazon.com/mediatailor/latest/ug/prefetching-ads.html) in the
// MediaTailor User Guide.
func (c *Client) DeletePrefetchSchedule(ctx context.Context, params *DeletePrefetchScheduleInput, optFns ...func(*Options)) (*DeletePrefetchScheduleOutput, error) {
	if params == nil {
		params = &DeletePrefetchScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePrefetchSchedule", params, optFns, c.addOperationDeletePrefetchScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePrefetchScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePrefetchScheduleInput struct {

	// The name of the prefetch schedule. If the action is successful, the service
	// sends back an HTTP 204 response with an empty HTTP body.
	//
	// This member is required.
	Name *string

	// The name of the playback configuration for this prefetch schedule.
	//
	// This member is required.
	PlaybackConfigurationName *string

	noSmithyDocumentSerde
}

type DeletePrefetchScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePrefetchScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeletePrefetchSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeletePrefetchSchedule{}, middleware.After)
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
	if err = addOpDeletePrefetchScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePrefetchSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePrefetchSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "DeletePrefetchSchedule",
	}
}
