// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a fleet. Before you delete a fleet, all vehicles must be dissociated
// from the fleet. For more information, see Delete a fleet (AWS CLI)
// (https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/delete-fleet-cli.html)
// in the Amazon Web Services IoT FleetWise Developer Guide. If the fleet is
// successfully deleted, Amazon Web Services IoT FleetWise sends back an HTTP 200
// response with an empty body.
func (c *Client) DeleteFleet(ctx context.Context, params *DeleteFleetInput, optFns ...func(*Options)) (*DeleteFleetOutput, error) {
	if params == nil {
		params = &DeleteFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFleet", params, optFns, c.addOperationDeleteFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFleetInput struct {

	// The ID of the fleet to delete.
	//
	// This member is required.
	FleetId *string

	noSmithyDocumentSerde
}

type DeleteFleetOutput struct {

	// The Amazon Resource Name (ARN) of the deleted fleet.
	Arn *string

	// The ID of the deleted fleet.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteFleet{}, middleware.After)
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
	if err = addOpDeleteFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "DeleteFleet",
	}
}
