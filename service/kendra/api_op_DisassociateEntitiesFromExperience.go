// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Prevents users or groups in your IAM Identity Center identity source from
// accessing your Amazon Kendra experience. You can create an Amazon Kendra
// experience such as a search application. For more information on creating a
// search application experience, see Building a search experience with no code
// (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html).
func (c *Client) DisassociateEntitiesFromExperience(ctx context.Context, params *DisassociateEntitiesFromExperienceInput, optFns ...func(*Options)) (*DisassociateEntitiesFromExperienceOutput, error) {
	if params == nil {
		params = &DisassociateEntitiesFromExperienceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateEntitiesFromExperience", params, optFns, c.addOperationDisassociateEntitiesFromExperienceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateEntitiesFromExperienceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateEntitiesFromExperienceInput struct {

	// Lists users or groups in your IAM Identity Center identity source.
	//
	// This member is required.
	EntityList []types.EntityConfiguration

	// The identifier of your Amazon Kendra experience.
	//
	// This member is required.
	Id *string

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DisassociateEntitiesFromExperienceOutput struct {

	// Lists the users or groups in your IAM Identity Center identity source that
	// failed to properly remove access to your Amazon Kendra experience.
	FailedEntityList []types.FailedEntity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateEntitiesFromExperienceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateEntitiesFromExperience{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateEntitiesFromExperience{}, middleware.After)
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
	if err = addOpDisassociateEntitiesFromExperienceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateEntitiesFromExperience(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateEntitiesFromExperience(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DisassociateEntitiesFromExperience",
	}
}
