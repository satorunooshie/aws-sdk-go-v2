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

// Updates a campaign.
func (c *Client) UpdateCampaign(ctx context.Context, params *UpdateCampaignInput, optFns ...func(*Options)) (*UpdateCampaignOutput, error) {
	if params == nil {
		params = &UpdateCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCampaign", params, optFns, c.addOperationUpdateCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCampaignInput struct {

	// Specifies how to update a campaign. The action can be one of the following:
	//
	// *
	// APPROVE - To approve delivering a data collection scheme to vehicles.
	//
	// * SUSPEND
	// - To suspend collecting signal data.
	//
	// * RESUME - To resume collecting signal
	// data.
	//
	// * UPDATE - To update a campaign.
	//
	// This member is required.
	Action types.UpdateCampaignAction

	// The name of the campaign to update.
	//
	// This member is required.
	Name *string

	// A list of vehicle attributes to associate with a signal. Default: An empty array
	DataExtraDimensions []string

	// The description of the campaign.
	Description *string

	noSmithyDocumentSerde
}

type UpdateCampaignOutput struct {

	// The Amazon Resource Name (ARN) of the campaign.
	Arn *string

	// The name of the updated campaign.
	Name *string

	// The state of a campaign. The status can be one of:
	//
	// * CREATING - Amazon Web
	// Services IoT FleetWise is processing your request to create the campaign.
	//
	// *
	// WAITING_FOR_APPROVAL - After a campaign is created, it enters the
	// WAITING_FOR_APPROVAL state. To allow Amazon Web Services IoT FleetWise to deploy
	// the campaign to the target vehicle or fleet, use the API operation to approve
	// the campaign.
	//
	// * RUNNING - The campaign is active.
	//
	// * SUSPENDED - The campaign
	// is suspended. To resume the campaign, use the API operation.
	Status types.CampaignStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateCampaign{}, middleware.After)
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
	if err = addOpUpdateCampaignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCampaign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "UpdateCampaign",
	}
}
