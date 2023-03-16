// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous document classification job. Use the
// DescribeDocumentClassificationJob operation to track the progress of the job.
func (c *Client) StartDocumentClassificationJob(ctx context.Context, params *StartDocumentClassificationJobInput, optFns ...func(*Options)) (*StartDocumentClassificationJobOutput, error) {
	if params == nil {
		params = &StartDocumentClassificationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDocumentClassificationJob", params, optFns, c.addOperationStartDocumentClassificationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDocumentClassificationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDocumentClassificationJobInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Amazon Comprehend read access to your input data.
	//
	// This member is required.
	DataAccessRoleArn *string

	// Specifies the format and location of the input data for the job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// Specifies where to send the output files.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique identifier for the request. If you do not set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string

	// The Amazon Resource Name (ARN) of the document classifier to use to process the
	// job.
	DocumentClassifierArn *string

	// The Amazon Resource Number (ARN) of the flywheel associated with the model to
	// use.
	FlywheelArn *string

	// The identifier of the job.
	JobName *string

	// Tags to associate with the document classification job. A tag is a key-value
	// pair that adds metadata to a resource used by Amazon Comprehend. For example, a
	// tag with "Sales" as the key might be added to a resource to indicate its use by
	// the sales department.
	Tags []types.Tag

	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to
	// encrypt data on the storage volume attached to the ML compute instance(s) that
	// process the analysis job. The VolumeKmsKeyId can be either of the following
	// formats:
	//
	// * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * Amazon
	// Resource Name (ARN) of a KMS Key:
	// "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your document classification job. For
	// more information, see Amazon VPC
	// (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type StartDocumentClassificationJobOutput struct {

	// The ARN of the custom classification model.
	DocumentClassifierArn *string

	// The Amazon Resource Name (ARN) of the document classification job. It is a
	// unique, fully qualified identifier for the job. It includes the AWS account,
	// Region, and the job ID. The format of the ARN is as follows:
	// arn::comprehend:::document-classification-job/ The following is an example job
	// ARN:
	// arn:aws:comprehend:us-west-2:111122223333:document-classification-job/1234abcd12ab34cd56ef1234567890ab
	JobArn *string

	// The identifier generated for the job. To get the status of the job, use this
	// identifier with the DescribeDocumentClassificationJob operation.
	JobId *string

	// The status of the job:
	//
	// * SUBMITTED - The job has been received and queued for
	// processing.
	//
	// * IN_PROGRESS - Amazon Comprehend is processing the job.
	//
	// *
	// COMPLETED - The job was successfully completed and the output is available.
	//
	// *
	// FAILED - The job did not complete. For details, use the
	// DescribeDocumentClassificationJob operation.
	//
	// * STOP_REQUESTED - Amazon
	// Comprehend has received a stop request for the job and is processing the
	// request.
	//
	// * STOPPED - The job was successfully stopped without completing.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDocumentClassificationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDocumentClassificationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDocumentClassificationJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartDocumentClassificationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartDocumentClassificationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDocumentClassificationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartDocumentClassificationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartDocumentClassificationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartDocumentClassificationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartDocumentClassificationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartDocumentClassificationJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartDocumentClassificationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartDocumentClassificationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartDocumentClassificationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "StartDocumentClassificationJob",
	}
}
