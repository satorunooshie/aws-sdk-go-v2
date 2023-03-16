// Code generated by smithy-go-codegen DO NOT EDIT.

package supportapp

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/supportapp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a Slack workspace for your Amazon Web Services account. To call this
// API, your account must be part of an organization in Organizations. If you're
// the management account and you want to register Slack workspaces for your
// organization, you must complete the following tasks:
//
// * Sign in to the Amazon
// Web Services Support Center (https://console.aws.amazon.com/support/app) and
// authorize the Slack workspaces where you want your organization to have access
// to. See Authorize a Slack workspace
// (https://docs.aws.amazon.com/awssupport/latest/user/authorize-slack-workspace.html)
// in the Amazon Web Services Support User Guide.
//
// * Call the
// RegisterSlackWorkspaceForOrganization API to authorize each Slack workspace for
// the organization.
//
// After the management account authorizes the Slack workspace,
// member accounts can call this API to authorize the same Slack workspace for
// their individual accounts. Member accounts don't need to authorize the Slack
// workspace manually through the Amazon Web Services Support Center
// (https://console.aws.amazon.com/support/app). To use the Amazon Web Services
// Support App, each account must then complete the following tasks:
//
// * Create an
// Identity and Access Management (IAM) role with the required permission. For more
// information, see Managing access to the Amazon Web Services Support App
// (https://docs.aws.amazon.com/awssupport/latest/user/support-app-permissions.html).
//
// *
// Configure a Slack channel to use the Amazon Web Services Support App for support
// cases for that account. For more information, see Configuring a Slack channel
// (https://docs.aws.amazon.com/awssupport/latest/user/add-your-slack-channel.html).
func (c *Client) RegisterSlackWorkspaceForOrganization(ctx context.Context, params *RegisterSlackWorkspaceForOrganizationInput, optFns ...func(*Options)) (*RegisterSlackWorkspaceForOrganizationOutput, error) {
	if params == nil {
		params = &RegisterSlackWorkspaceForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterSlackWorkspaceForOrganization", params, optFns, c.addOperationRegisterSlackWorkspaceForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterSlackWorkspaceForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterSlackWorkspaceForOrganizationInput struct {

	// The team ID in Slack. This ID uniquely identifies a Slack workspace, such as
	// T012ABCDEFG. Specify the Slack workspace that you want to use for your
	// organization.
	//
	// This member is required.
	TeamId *string

	noSmithyDocumentSerde
}

type RegisterSlackWorkspaceForOrganizationOutput struct {

	// Whether the Amazon Web Services account is a management or member account that's
	// part of an organization in Organizations.
	AccountType types.AccountType

	// The team ID in Slack. This ID uniquely identifies a Slack workspace, such as
	// T012ABCDEFG.
	TeamId *string

	// The name of the Slack workspace.
	TeamName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterSlackWorkspaceForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterSlackWorkspaceForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterSlackWorkspaceForOrganization{}, middleware.After)
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
	if err = addOpRegisterSlackWorkspaceForOrganizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterSlackWorkspaceForOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterSlackWorkspaceForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "supportapp",
		OperationName: "RegisterSlackWorkspaceForOrganization",
	}
}
