// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Kinesis data stream and all its shards and data. You must shut down
// any applications that are operating on the stream before you delete the stream.
// If an application attempts to operate on a deleted stream, it receives the
// exception ResourceNotFoundException. When invoking this API, it is recommended
// you use the StreamARN input parameter rather than the StreamName input
// parameter. If the stream is in the ACTIVE state, you can delete it. After a
// DeleteStream request, the specified stream is in the DELETING state until
// Kinesis Data Streams completes the deletion. Note: Kinesis Data Streams might
// continue to accept data read and write operations, such as PutRecord,
// PutRecords, and GetRecords, on a stream in the DELETING state until the stream
// deletion is complete. When you delete a stream, any shards in that stream are
// also deleted, and any tags are dissociated from the stream. You can use the
// DescribeStreamSummary operation to check the state of the stream, which is
// returned in StreamStatus. DeleteStream has a limit of five transactions per
// second per account.
func (c *Client) DeleteStream(ctx context.Context, params *DeleteStreamInput, optFns ...func(*Options)) (*DeleteStreamOutput, error) {
	if params == nil {
		params = &DeleteStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStream", params, optFns, c.addOperationDeleteStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for DeleteStream.
type DeleteStreamInput struct {

	// If this parameter is unset (null) or if you set it to false, and the stream has
	// registered consumers, the call to DeleteStream fails with a
	// ResourceInUseException.
	EnforceConsumerDeletion *bool

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream to delete.
	StreamName *string

	noSmithyDocumentSerde
}

type DeleteStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteStream{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DeleteStream",
	}
}
