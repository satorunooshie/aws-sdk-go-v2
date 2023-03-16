// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates function package content. This can be used as a dry run before
// uploading function package content with PutSolFunctionPackageContent
// (https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolFunctionPackageContent.html).
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
func (c *Client) ValidateSolFunctionPackageContent(ctx context.Context, params *ValidateSolFunctionPackageContentInput, optFns ...func(*Options)) (*ValidateSolFunctionPackageContentOutput, error) {
	if params == nil {
		params = &ValidateSolFunctionPackageContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateSolFunctionPackageContent", params, optFns, c.addOperationValidateSolFunctionPackageContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateSolFunctionPackageContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateSolFunctionPackageContentInput struct {

	// Function package file.
	//
	// This member is required.
	File []byte

	// Function package ID.
	//
	// This member is required.
	VnfPkgId *string

	// Function package content type.
	ContentType types.PackageContentType

	noSmithyDocumentSerde
}

type ValidateSolFunctionPackageContentOutput struct {

	// Function package ID.
	//
	// This member is required.
	Id *string

	// Function package metadata.
	//
	// This member is required.
	Metadata *types.ValidateSolFunctionPackageContentMetadata

	// Network function product name.
	//
	// This member is required.
	VnfProductName *string

	// Network function provider.
	//
	// This member is required.
	VnfProvider *string

	// Function package descriptor ID.
	//
	// This member is required.
	VnfdId *string

	// Function package descriptor version.
	//
	// This member is required.
	VnfdVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidateSolFunctionPackageContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidateSolFunctionPackageContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidateSolFunctionPackageContent{}, middleware.After)
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
	if err = addOpValidateSolFunctionPackageContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateSolFunctionPackageContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateSolFunctionPackageContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "ValidateSolFunctionPackageContent",
	}
}
