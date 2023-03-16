// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a replication configuration that replicates an existing EFS file system
// to a new, read-only file system. For more information, see Amazon EFS
// replication (https://docs.aws.amazon.com/efs/latest/ug/efs-replication.html) in
// the Amazon EFS User Guide. The replication configuration specifies the
// following:
//
// * Source file system - An existing EFS file system that you want
// replicated. The source file system cannot be a destination file system in an
// existing replication configuration.
//
// * Destination file system configuration -
// The configuration of the destination file system to which the source file system
// will be replicated. There can only be one destination file system in a
// replication configuration. The destination file system configuration consists of
// the following properties:
//
// * Amazon Web Services Region - The Amazon Web
// Services Region in which the destination file system is created. Amazon EFS
// replication is available in all Amazon Web Services Regions that Amazon EFS is
// available in, except Africa (Cape Town), Asia Pacific (Hong Kong), Asia Pacific
// (Jakarta), Europe (Milan), and Middle East (Bahrain).
//
// * Availability Zone - If
// you want the destination file system to use EFS One Zone availability and
// durability, you must specify the Availability Zone to create the file system in.
// For more information about EFS storage classes, see  Amazon EFS storage classes
// (https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the Amazon
// EFS User Guide.
//
// * Encryption - All destination file systems are created with
// encryption at rest enabled. You can specify the Key Management Service (KMS) key
// that is used to encrypt the destination file system. If you don't specify a KMS
// key, your service-managed KMS key for Amazon EFS is used. After the file system
// is created, you cannot change the KMS key.
//
// The following properties are set by
// default:
//
// * Performance mode - The destination file system's performance mode
// matches that of the source file system, unless the destination file system uses
// EFS One Zone storage. In that case, the General Purpose performance mode is
// used. The performance mode cannot be changed.
//
// * Throughput mode - The
// destination file system's throughput mode matches that of the source file
// system. After the file system is created, you can modify the throughput
// mode.
//
// The following properties are turned off by default:
//
// * Lifecycle
// management - EFS lifecycle management and EFS Intelligent-Tiering are not
// enabled on the destination file system. After the destination file system is
// created, you can enable EFS lifecycle management and EFS Intelligent-Tiering.
//
// *
// Automatic backups - Automatic daily backups not enabled on the destination file
// system. After the file system is created, you can change this setting.
//
// For more
// information, see Amazon EFS replication
// (https://docs.aws.amazon.com/efs/latest/ug/efs-replication.html) in the Amazon
// EFS User Guide.
func (c *Client) CreateReplicationConfiguration(ctx context.Context, params *CreateReplicationConfigurationInput, optFns ...func(*Options)) (*CreateReplicationConfigurationOutput, error) {
	if params == nil {
		params = &CreateReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationConfiguration", params, optFns, c.addOperationCreateReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplicationConfigurationInput struct {

	// An array of destination configuration objects. Only one destination
	// configuration object is supported.
	//
	// This member is required.
	Destinations []types.DestinationToCreate

	// Specifies the Amazon EFS file system that you want to replicate. This file
	// system cannot already be a source or destination file system in another
	// replication configuration.
	//
	// This member is required.
	SourceFileSystemId *string

	noSmithyDocumentSerde
}

type CreateReplicationConfigurationOutput struct {

	// Describes when the replication configuration was created.
	//
	// This member is required.
	CreationTime *time.Time

	// An array of destination objects. Only one destination object is supported.
	//
	// This member is required.
	Destinations []types.Destination

	// The Amazon Resource Name (ARN) of the original source Amazon EFS file system in
	// the replication configuration.
	//
	// This member is required.
	OriginalSourceFileSystemArn *string

	// The Amazon Resource Name (ARN) of the current source file system in the
	// replication configuration.
	//
	// This member is required.
	SourceFileSystemArn *string

	// The ID of the source Amazon EFS file system that is being replicated.
	//
	// This member is required.
	SourceFileSystemId *string

	// The Amazon Web Services Region in which the source Amazon EFS file system is
	// located.
	//
	// This member is required.
	SourceFileSystemRegion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateReplicationConfiguration{}, middleware.After)
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
	if err = addOpCreateReplicationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "CreateReplicationConfiguration",
	}
}
