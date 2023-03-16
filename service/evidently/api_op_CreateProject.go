// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a project, which is the logical object in Evidently that can contain
// features, launches, and experiments. Use projects to group similar features
// together. To update an existing project, use UpdateProject
// (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_UpdateProject.html).
func (c *Client) CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) {
	if params == nil {
		params = &CreateProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProject", params, optFns, c.addOperationCreateProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProjectInput struct {

	// The name for the project.
	//
	// This member is required.
	Name *string

	// Use this parameter if the project will use client-side evaluation powered by
	// AppConfig. Client-side evaluation allows your application to assign variations
	// to user sessions locally instead of by calling the EvaluateFeature
	// (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_EvaluateFeature.html)
	// operation. This mitigates the latency and availability risks that come with an
	// API call. For more information, see  Client-side evaluation - powered by
	// AppConfig.
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-client-side-evaluation.html)
	// This parameter is a structure that contains information about the AppConfig
	// application and environment that will be used as for client-side evaluation. To
	// create a project that uses client-side evaluation, you must have the
	// evidently:ExportProjectAsConfiguration permission.
	AppConfigResource *types.ProjectAppConfigResourceConfig

	// A structure that contains information about where Evidently is to store
	// evaluation events for longer term storage, if you choose to do so. If you choose
	// not to store these events, Evidently deletes them after using them to produce
	// metrics and other experiment results that you can view.
	DataDelivery *types.ProjectDataDeliveryConfig

	// An optional description of the project.
	Description *string

	// Assigns one or more tags (key-value pairs) to the project. Tags can help you
	// organize and categorize your resources. You can also use them to scope user
	// permissions by granting a user permission to access or change only resources
	// with certain tag values. Tags don't have any semantic meaning to Amazon Web
	// Services and are interpreted strictly as strings of characters. You can
	// associate as many as 50 tags with a project. For more information, see Tagging
	// Amazon Web Services resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateProjectOutput struct {

	// A structure that contains information about the created project.
	//
	// This member is required.
	Project *types.Project

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProject{}, middleware.After)
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
	if err = addOpCreateProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "CreateProject",
	}
}
