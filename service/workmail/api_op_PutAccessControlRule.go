// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a new access control rule for the specified organization. The rule allows
// or denies access to the organization for the specified IPv4 addresses, access
// protocol actions, user IDs and impersonation IDs. Adding a new rule with the
// same name as an existing rule replaces the older rule.
func (c *Client) PutAccessControlRule(ctx context.Context, params *PutAccessControlRuleInput, optFns ...func(*Options)) (*PutAccessControlRuleOutput, error) {
	if params == nil {
		params = &PutAccessControlRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccessControlRule", params, optFns, c.addOperationPutAccessControlRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccessControlRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccessControlRuleInput struct {

	// The rule description.
	//
	// This member is required.
	Description *string

	// The rule effect.
	//
	// This member is required.
	Effect types.AccessControlRuleEffect

	// The rule name.
	//
	// This member is required.
	Name *string

	// The identifier of the organization.
	//
	// This member is required.
	OrganizationId *string

	// Access protocol actions to include in the rule. Valid values include ActiveSync,
	// AutoDiscover, EWS, IMAP, SMTP, WindowsOutlook, and WebMail.
	Actions []string

	// Impersonation role IDs to include in the rule.
	ImpersonationRoleIds []string

	// IPv4 CIDR ranges to include in the rule.
	IpRanges []string

	// Access protocol actions to exclude from the rule. Valid values include
	// ActiveSync, AutoDiscover, EWS, IMAP, SMTP, WindowsOutlook, and WebMail.
	NotActions []string

	// Impersonation role IDs to exclude from the rule.
	NotImpersonationRoleIds []string

	// IPv4 CIDR ranges to exclude from the rule.
	NotIpRanges []string

	// User IDs to exclude from the rule.
	NotUserIds []string

	// User IDs to include in the rule.
	UserIds []string

	noSmithyDocumentSerde
}

type PutAccessControlRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccessControlRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccessControlRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccessControlRule{}, middleware.After)
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
	if err = addOpPutAccessControlRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccessControlRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccessControlRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "PutAccessControlRule",
	}
}
