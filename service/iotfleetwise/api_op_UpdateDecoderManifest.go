// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a decoder manifest. A decoder manifest can only be updated when the
// status is DRAFT. Only ACTIVE decoder manifests can be associated with vehicles.
func (c *Client) UpdateDecoderManifest(ctx context.Context, params *UpdateDecoderManifestInput, optFns ...func(*Options)) (*UpdateDecoderManifestOutput, error) {
	if params == nil {
		params = &UpdateDecoderManifestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDecoderManifest", params, optFns, c.addOperationUpdateDecoderManifestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDecoderManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDecoderManifestInput struct {

	// The name of the decoder manifest to update.
	//
	// This member is required.
	Name *string

	// A brief description of the decoder manifest to update.
	Description *string

	// A list of information about the network interfaces to add to the decoder
	// manifest.
	NetworkInterfacesToAdd []types.NetworkInterface

	// A list of network interfaces to remove from the decoder manifest.
	NetworkInterfacesToRemove []string

	// A list of information about the network interfaces to update in the decoder
	// manifest.
	NetworkInterfacesToUpdate []types.NetworkInterface

	// A list of information about decoding additional signals to add to the decoder
	// manifest.
	SignalDecodersToAdd []types.SignalDecoder

	// A list of signal decoders to remove from the decoder manifest.
	SignalDecodersToRemove []string

	// A list of updated information about decoding signals to update in the decoder
	// manifest.
	SignalDecodersToUpdate []types.SignalDecoder

	// The state of the decoder manifest. If the status is ACTIVE, the decoder manifest
	// can't be edited. If the status is DRAFT, you can edit the decoder manifest.
	Status types.ManifestStatus

	noSmithyDocumentSerde
}

type UpdateDecoderManifestOutput struct {

	// The Amazon Resource Name (ARN) of the updated decoder manifest.
	//
	// This member is required.
	Arn *string

	// The name of the updated decoder manifest.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDecoderManifestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateDecoderManifest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateDecoderManifest{}, middleware.After)
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
	if err = addOpUpdateDecoderManifestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDecoderManifest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDecoderManifest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "UpdateDecoderManifest",
	}
}
