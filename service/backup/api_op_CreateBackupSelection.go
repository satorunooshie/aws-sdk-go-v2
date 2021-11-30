// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a JSON document that specifies a set of resources to assign to a backup
// plan. For examples, see Assigning resources programmatically
// (https://docs.aws.amazon.com/assigning-resources.html#assigning-resources-json).
func (c *Client) CreateBackupSelection(ctx context.Context, params *CreateBackupSelectionInput, optFns ...func(*Options)) (*CreateBackupSelectionOutput, error) {
	if params == nil {
		params = &CreateBackupSelectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBackupSelection", params, optFns, c.addOperationCreateBackupSelectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBackupSelectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBackupSelectionInput struct {

	// Uniquely identifies the backup plan to be associated with the selection of
	// resources.
	//
	// This member is required.
	BackupPlanId *string

	// Specifies the body of a request to assign a set of resources to a backup plan.
	//
	// This member is required.
	BackupSelection *types.BackupSelection

	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of running the operation twice. This parameter is
	// optional. If used, this parameter must contain 1 to 50 alphanumeric or '-_.'
	// characters.
	CreatorRequestId *string

	noSmithyDocumentSerde
}

type CreateBackupSelectionOutput struct {

	// Uniquely identifies a backup plan.
	BackupPlanId *string

	// The date and time a backup selection is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Uniquely identifies the body of a request to assign a set of resources to a
	// backup plan.
	SelectionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBackupSelectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBackupSelection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBackupSelection{}, middleware.After)
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
	if err = addOpCreateBackupSelectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackupSelection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBackupSelection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "CreateBackupSelection",
	}
}
