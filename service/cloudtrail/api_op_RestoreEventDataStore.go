// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores a deleted event data store specified by EventDataStore, which accepts
// an event data store ARN. You can only restore a deleted event data store within
// the seven-day wait period after deletion. Restoring an event data store can take
// several minutes, depending on the size of the event data store.
func (c *Client) RestoreEventDataStore(ctx context.Context, params *RestoreEventDataStoreInput, optFns ...func(*Options)) (*RestoreEventDataStoreOutput, error) {
	if params == nil {
		params = &RestoreEventDataStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreEventDataStore", params, optFns, c.addOperationRestoreEventDataStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreEventDataStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreEventDataStoreInput struct {

	// The ARN (or the ID suffix of the ARN) of the event data store that you want to
	// restore.
	//
	// This member is required.
	EventDataStore *string

	noSmithyDocumentSerde
}

type RestoreEventDataStoreOutput struct {

	// The advanced event selectors that were used to select events.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// The timestamp of an event data store's creation.
	CreatedTimestamp *time.Time

	// The event data store ARN.
	EventDataStoreArn *string

	// Specifies the KMS key ID that encrypts the events delivered by CloudTrail. The
	// value is a fully specified ARN to a KMS key in the following format.
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Indicates whether the event data store is collecting events from all regions, or
	// only from the region in which the event data store was created.
	MultiRegionEnabled *bool

	// The name of the event data store.
	Name *string

	// Indicates whether an event data store is collecting logged events for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period, in days.
	RetentionPeriod *int32

	// The status of the event data store.
	Status types.EventDataStoreStatus

	// Indicates that termination protection is enabled and the event data store cannot
	// be automatically deleted.
	TerminationProtectionEnabled *bool

	// The timestamp that shows when an event data store was updated, if applicable.
	// UpdatedTimestamp is always either the same or newer than the time shown in
	// CreatedTimestamp.
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreEventDataStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreEventDataStore{}, middleware.After)
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
	if err = addOpRestoreEventDataStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreEventDataStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreEventDataStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "RestoreEventDataStore",
	}
}
