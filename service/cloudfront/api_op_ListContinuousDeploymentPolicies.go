// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the continuous deployment policies in your Amazon Web Services
// account. You can optionally specify the maximum number of items to receive in
// the response. If the total number of items in the list exceeds the maximum that
// you specify, or the default maximum, the response is paginated. To get the next
// page of items, send a subsequent request that specifies the NextMarker value
// from the current response as the Marker value in the subsequent request.
func (c *Client) ListContinuousDeploymentPolicies(ctx context.Context, params *ListContinuousDeploymentPoliciesInput, optFns ...func(*Options)) (*ListContinuousDeploymentPoliciesOutput, error) {
	if params == nil {
		params = &ListContinuousDeploymentPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContinuousDeploymentPolicies", params, optFns, c.addOperationListContinuousDeploymentPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContinuousDeploymentPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContinuousDeploymentPoliciesInput struct {

	// Use this field when paginating results to indicate where to begin in your list
	// of continuous deployment policies. The response includes policies in the list
	// that occur after the marker. To get the next page of the list, set this field's
	// value to the value of NextMarker from the current page's response.
	Marker *string

	// The maximum number of continuous deployment policies that you want returned in
	// the response.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListContinuousDeploymentPoliciesOutput struct {

	// A list of continuous deployment policies.
	ContinuousDeploymentPolicyList *types.ContinuousDeploymentPolicyList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContinuousDeploymentPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListContinuousDeploymentPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListContinuousDeploymentPolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContinuousDeploymentPolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListContinuousDeploymentPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListContinuousDeploymentPolicies",
	}
}
