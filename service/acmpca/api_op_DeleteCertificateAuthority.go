// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a private certificate authority (CA). You must provide the Amazon
// Resource Name (ARN) of the private CA that you want to delete. You can find the
// ARN by calling the ListCertificateAuthorities
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html)
// action. Deleting a CA will invalidate other CAs and certificates below it in
// your CA hierarchy. Before you can delete a CA that you have created and
// activated, you must disable it. To do this, call the UpdateCertificateAuthority
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html)
// action and set the CertificateAuthorityStatus parameter to DISABLED.
// Additionally, you can delete a CA if you are waiting for it to be created (that
// is, the status of the CA is CREATING). You can also delete it if the CA has been
// created but you haven't yet imported the signed certificate into Amazon Web
// Services Private CA (that is, the status of the CA is PENDING_CERTIFICATE). When
// you successfully call DeleteCertificateAuthority
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeleteCertificateAuthority.html),
// the CA's status changes to DELETED. However, the CA won't be permanently deleted
// until the restoration period has passed. By default, if you do not set the
// PermanentDeletionTimeInDays parameter, the CA remains restorable for 30 days.
// You can set the parameter from 7 to 30 days. The DescribeCertificateAuthority
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_DescribeCertificateAuthority.html)
// action returns the time remaining in the restoration window of a private CA in
// the DELETED state. To restore an eligible CA, call the
// RestoreCertificateAuthority
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_RestoreCertificateAuthority.html)
// action.
func (c *Client) DeleteCertificateAuthority(ctx context.Context, params *DeleteCertificateAuthorityInput, optFns ...func(*Options)) (*DeleteCertificateAuthorityOutput, error) {
	if params == nil {
		params = &DeleteCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCertificateAuthority", params, optFns, c.addOperationDeleteCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCertificateAuthorityInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority
	// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).
	// This must have the following form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string

	// The number of days to make a CA restorable after it has been deleted. This can
	// be anywhere from 7 to 30 days, with 30 being the default.
	PermanentDeletionTimeInDays *int32

	noSmithyDocumentSerde
}

type DeleteCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCertificateAuthority{}, middleware.After)
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
	if err = addOpDeleteCertificateAuthorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCertificateAuthority(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCertificateAuthority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "DeleteCertificateAuthority",
	}
}
