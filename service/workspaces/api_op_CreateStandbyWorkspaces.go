// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a standby WorkSpace in a secondary Region.
func (c *Client) CreateStandbyWorkspaces(ctx context.Context, params *CreateStandbyWorkspacesInput, optFns ...func(*Options)) (*CreateStandbyWorkspacesOutput, error) {
	if params == nil {
		params = &CreateStandbyWorkspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStandbyWorkspaces", params, optFns, c.addOperationCreateStandbyWorkspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStandbyWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStandbyWorkspacesInput struct {

	// The Region of the primary WorkSpace.
	//
	// This member is required.
	PrimaryRegion *string

	// Information about the standby WorkSpace to be created.
	//
	// This member is required.
	StandbyWorkspaces []types.StandbyWorkspace

	noSmithyDocumentSerde
}

type CreateStandbyWorkspacesOutput struct {

	// Information about the standby WorkSpace that could not be created.
	FailedStandbyRequests []types.FailedCreateStandbyWorkspacesRequest

	// Information about the standby WorkSpace that was created.
	PendingStandbyRequests []types.PendingCreateStandbyWorkspacesRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStandbyWorkspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStandbyWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStandbyWorkspaces{}, middleware.After)
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
	if err = addOpCreateStandbyWorkspacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStandbyWorkspaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStandbyWorkspaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "CreateStandbyWorkspaces",
	}
}
