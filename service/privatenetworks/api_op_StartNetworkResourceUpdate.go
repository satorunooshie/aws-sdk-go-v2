// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/privatenetworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an update of the specified network resource. After you submit a request
// to replace or return a network resource, the status of the network resource is
// CREATING_SHIPPING_LABEL. The shipping label is available when the status of the
// network resource is PENDING_RETURN. After the network resource is successfully
// returned, its status is DELETED. For more information, see Return a radio unit
// (https://docs.aws.amazon.com/private-networks/latest/userguide/radio-units.html#return-radio-unit).
func (c *Client) StartNetworkResourceUpdate(ctx context.Context, params *StartNetworkResourceUpdateInput, optFns ...func(*Options)) (*StartNetworkResourceUpdateOutput, error) {
	if params == nil {
		params = &StartNetworkResourceUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartNetworkResourceUpdate", params, optFns, c.addOperationStartNetworkResourceUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartNetworkResourceUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartNetworkResourceUpdateInput struct {

	// The Amazon Resource Name (ARN) of the network resource.
	//
	// This member is required.
	NetworkResourceArn *string

	// The update type.
	//
	// * REPLACE - Submits a request to replace a defective radio
	// unit. We provide a shipping label that you can use for the return process and we
	// ship a replacement radio unit to you.
	//
	// * RETURN - Submits a request to replace a
	// radio unit that you no longer need. We provide a shipping label that you can use
	// for the return process.
	//
	// This member is required.
	UpdateType types.UpdateType

	// The reason for the return. Providing a reason for a return is optional.
	ReturnReason *string

	// The shipping address. If you don't provide a shipping address when replacing or
	// returning a network resource, we use the address from the original order for the
	// network resource.
	ShippingAddress *types.Address

	noSmithyDocumentSerde
}

type StartNetworkResourceUpdateOutput struct {

	// The network resource.
	NetworkResource *types.NetworkResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartNetworkResourceUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartNetworkResourceUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartNetworkResourceUpdate{}, middleware.After)
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
	if err = addOpStartNetworkResourceUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartNetworkResourceUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartNetworkResourceUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "private-networks",
		OperationName: "StartNetworkResourceUpdate",
	}
}
