// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified data quality ruleset.
func (c *Client) UpdateDataQualityRuleset(ctx context.Context, params *UpdateDataQualityRulesetInput, optFns ...func(*Options)) (*UpdateDataQualityRulesetOutput, error) {
	if params == nil {
		params = &UpdateDataQualityRulesetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataQualityRuleset", params, optFns, c.addOperationUpdateDataQualityRulesetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataQualityRulesetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataQualityRulesetInput struct {

	// The name of the data quality ruleset.
	//
	// This member is required.
	Name *string

	// A description of the ruleset.
	Description *string

	// A Data Quality Definition Language (DQDL) ruleset. For more information, see the
	// Glue developer guide.
	Ruleset *string

	// The new name of the ruleset, if you are renaming it.
	UpdatedName *string

	noSmithyDocumentSerde
}

type UpdateDataQualityRulesetOutput struct {

	// A description of the ruleset.
	Description *string

	// The name of the data quality ruleset.
	Name *string

	// A Data Quality Definition Language (DQDL) ruleset. For more information, see the
	// Glue developer guide.
	Ruleset *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataQualityRulesetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDataQualityRuleset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDataQualityRuleset{}, middleware.After)
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
	if err = addOpUpdateDataQualityRulesetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataQualityRuleset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataQualityRuleset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateDataQualityRuleset",
	}
}
