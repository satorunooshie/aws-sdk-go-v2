// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the assets in package versions and sets the package versions' status to
// Disposed. A disposed package version cannot be restored in your repository
// because its assets are deleted. To view all disposed package versions in a
// repository, use ListPackageVersions
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html)
// and set the status
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html#API_ListPackageVersions_RequestSyntax)
// parameter to Disposed. To view information about a disposed package version, use
// DescribePackageVersion
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DescribePackageVersion.html).
func (c *Client) DisposePackageVersions(ctx context.Context, params *DisposePackageVersionsInput, optFns ...func(*Options)) (*DisposePackageVersionsOutput, error) {
	if params == nil {
		params = &DisposePackageVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisposePackageVersions", params, optFns, c.addOperationDisposePackageVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisposePackageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisposePackageVersionsInput struct {

	// The name of the domain that contains the repository you want to dispose.
	//
	// This member is required.
	Domain *string

	// A format that specifies the type of package versions you want to dispose.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package with the versions you want to dispose.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the package versions you want to
	// dispose.
	//
	// This member is required.
	Repository *string

	// The versions of the package you want to dispose.
	//
	// This member is required.
	Versions []string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The expected status of the package version to dispose.
	ExpectedStatus types.PackageVersionStatus

	// The namespace of the package versions to be disposed. The package version
	// component that specifies its namespace depends on its type. For example:
	//
	// * The
	// namespace of a Maven package version is its groupId.
	//
	// * The namespace of an npm
	// package version is its scope.
	//
	// * Python and NuGet package versions do not
	// contain a corresponding component, package versions of those formats do not have
	// a namespace.
	//
	// * The namespace of a generic package is it’s namespace.
	Namespace *string

	// The revisions of the package versions you want to dispose.
	VersionRevisions map[string]string

	noSmithyDocumentSerde
}

type DisposePackageVersionsOutput struct {

	// A PackageVersionError object that contains a map of errors codes for the
	// disposed package versions that failed. The possible error codes are:
	//
	// *
	// ALREADY_EXISTS
	//
	// * MISMATCHED_REVISION
	//
	// * MISMATCHED_STATUS
	//
	// * NOT_ALLOWED
	//
	// *
	// NOT_FOUND
	//
	// * SKIPPED
	FailedVersions map[string]types.PackageVersionError

	// A list of the package versions that were successfully disposed.
	SuccessfulVersions map[string]types.SuccessfulPackageVersionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisposePackageVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisposePackageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisposePackageVersions{}, middleware.After)
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
	if err = addOpDisposePackageVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisposePackageVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisposePackageVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DisposePackageVersions",
	}
}
