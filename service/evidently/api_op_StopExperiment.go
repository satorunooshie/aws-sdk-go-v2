// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Stops an experiment that is currently running. If you stop an experiment, you
// can't resume it or restart it.
func (c *Client) StopExperiment(ctx context.Context, params *StopExperimentInput, optFns ...func(*Options)) (*StopExperimentOutput, error) {
	if params == nil {
		params = &StopExperimentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopExperiment", params, optFns, c.addOperationStopExperimentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopExperimentInput struct {

	// The name of the experiment to stop.
	//
	// This member is required.
	Experiment *string

	// The name or ARN of the project that contains the experiment to stop.
	//
	// This member is required.
	Project *string

	// Specify whether the experiment is to be considered COMPLETED or CANCELLED after
	// it stops.
	DesiredState types.ExperimentStopDesiredState

	// A string that describes why you are stopping the experiment.
	Reason *string

	noSmithyDocumentSerde
}

type StopExperimentOutput struct {

	// The date and time that the experiment stopped.
	EndedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopExperimentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopExperiment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopExperiment{}, middleware.After)
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
	if err = addOpStopExperimentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopExperiment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopExperiment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "StopExperiment",
	}
}
