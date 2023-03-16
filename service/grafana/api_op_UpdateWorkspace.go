// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/grafana/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an existing Amazon Managed Grafana workspace. If you use this operation
// and omit any optional parameters, the existing values of those parameters are
// not changed. To modify the user authentication methods that the workspace uses,
// such as SAML or IAM Identity Center, use UpdateWorkspaceAuthentication
// (https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdateWorkspaceAuthentication.html).
// To modify which users in the workspace have the Admin and Editor Grafana roles,
// use UpdatePermissions
// (https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdatePermissions.html).
func (c *Client) UpdateWorkspace(ctx context.Context, params *UpdateWorkspaceInput, optFns ...func(*Options)) (*UpdateWorkspaceOutput, error) {
	if params == nil {
		params = &UpdateWorkspaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkspace", params, optFns, c.addOperationUpdateWorkspaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkspaceInput struct {

	// The ID of the workspace to update.
	//
	// This member is required.
	WorkspaceId *string

	// Specifies whether the workspace can access Amazon Web Services resources in this
	// Amazon Web Services account only, or whether it can also access Amazon Web
	// Services resources in other accounts in the same organization. If you specify
	// ORGANIZATION, you must specify which organizational units the workspace can
	// access in the workspaceOrganizationalUnits parameter.
	AccountAccessType types.AccountAccessType

	// The configuration settings for network access to your workspace. When this is
	// configured, only listed IP addresses and VPC endpoints will be able to access
	// your workspace. Standard Grafana authentication and authorization will still be
	// required. If this is not configured, or is removed, then all IP addresses and
	// VPC endpoints will be allowed. Standard Grafana authentication and authorization
	// will still be required.
	NetworkAccessControl *types.NetworkAccessConfiguration

	// The name of an IAM role that already exists to use to access resources through
	// Organizations. This can only be used with a workspace that has the
	// permissionType set to CUSTOMER_MANAGED.
	OrganizationRoleName *string

	// Use this parameter if you want to change a workspace from SERVICE_MANAGED to
	// CUSTOMER_MANAGED. This allows you to manage the permissions that the workspace
	// uses to access datasources and notification channels. If the workspace is in a
	// member Amazon Web Services account of an organization, and that account is not a
	// delegated administrator account, and you want the workspace to access data
	// sources in other Amazon Web Services accounts in the organization, you must
	// choose CUSTOMER_MANAGED. If you specify this as CUSTOMER_MANAGED, you must also
	// specify a workspaceRoleArn that the workspace will use for accessing Amazon Web
	// Services resources. For more information on the role and permissions needed, see
	// Amazon Managed Grafana permissions and policies for Amazon Web Services data
	// sources and notification channels
	// (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-permissions.html)
	// Do not use this to convert a CUSTOMER_MANAGED workspace to SERVICE_MANAGED. Do
	// not include this parameter if you want to leave the workspace as
	// SERVICE_MANAGED. You can convert a CUSTOMER_MANAGED workspace to SERVICE_MANAGED
	// using the Amazon Managed Grafana console. For more information, see Managing
	// permissions for data sources and notification channels
	// (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-datasource-and-notification.html).
	PermissionType types.PermissionType

	// Whether to remove the network access configuration from the workspace. Setting
	// this to true and providing a networkAccessControl to set will return an error.
	// If you remove this configuration by setting this to true, then all IP addresses
	// and VPC endpoints will be allowed. Standard Grafana authentication and
	// authorization will still be required.
	RemoveNetworkAccessConfiguration *bool

	// Whether to remove the VPC configuration from the workspace. Setting this to true
	// and providing a vpcConfiguration to set will return an error.
	RemoveVpcConfiguration *bool

	// The name of the CloudFormation stack set to use to generate IAM roles to be used
	// for this workspace.
	StackSetName *string

	// The configuration settings for an Amazon VPC that contains data sources for your
	// Grafana workspace to connect to.
	VpcConfiguration *types.VpcConfiguration

	// This parameter is for internal use only, and should not be used.
	WorkspaceDataSources []types.DataSourceType

	// A description for the workspace. This is used only to help you identify this
	// workspace.
	WorkspaceDescription *string

	// A new name for the workspace to update.
	WorkspaceName *string

	// Specify the Amazon Web Services notification channels that you plan to use in
	// this workspace. Specifying these data sources here enables Amazon Managed
	// Grafana to create IAM roles and permissions that allow Amazon Managed Grafana to
	// use these channels.
	WorkspaceNotificationDestinations []types.NotificationDestinationType

	// Specifies the organizational units that this workspace is allowed to use data
	// sources from, if this workspace is in an account that is part of an
	// organization.
	WorkspaceOrganizationalUnits []string

	// Specifies an IAM role that grants permissions to Amazon Web Services resources
	// that the workspace accesses, such as data sources and notification channels. If
	// this workspace has permissionTypeCUSTOMER_MANAGED, then this role is required.
	WorkspaceRoleArn *string

	noSmithyDocumentSerde
}

type UpdateWorkspaceOutput struct {

	// A structure containing data about the workspace that was created.
	//
	// This member is required.
	Workspace *types.WorkspaceDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkspaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWorkspace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWorkspace{}, middleware.After)
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
	if err = addOpUpdateWorkspaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkspace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkspace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "grafana",
		OperationName: "UpdateWorkspace",
	}
}
