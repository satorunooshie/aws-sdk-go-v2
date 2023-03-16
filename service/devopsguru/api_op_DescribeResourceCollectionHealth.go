// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the number of open proactive insights, open reactive insights, and the
// Mean Time to Recover (MTTR) for all closed insights in resource collections in
// your account. You specify the type of Amazon Web Services resources collection.
// The two types of Amazon Web Services resource collections supported are Amazon
// Web Services CloudFormation stacks and Amazon Web Services resources that
// contain the same Amazon Web Services tag. DevOps Guru can be configured to
// analyze the Amazon Web Services resources that are defined in the stacks or that
// are tagged using the same tag key. You can specify up to 500 Amazon Web Services
// CloudFormation stacks.
func (c *Client) DescribeResourceCollectionHealth(ctx context.Context, params *DescribeResourceCollectionHealthInput, optFns ...func(*Options)) (*DescribeResourceCollectionHealthOutput, error) {
	if params == nil {
		params = &DescribeResourceCollectionHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeResourceCollectionHealth", params, optFns, c.addOperationDescribeResourceCollectionHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeResourceCollectionHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeResourceCollectionHealthInput struct {

	// An Amazon Web Services resource collection type. This type specifies how
	// analyzed Amazon Web Services resources are defined. The two types of Amazon Web
	// Services resource collections supported are Amazon Web Services CloudFormation
	// stacks and Amazon Web Services resources that contain the same Amazon Web
	// Services tag. DevOps Guru can be configured to analyze the Amazon Web Services
	// resources that are defined in the stacks or that are tagged using the same tag
	// key. You can specify up to 500 Amazon Web Services CloudFormation stacks.
	//
	// This member is required.
	ResourceCollectionType types.ResourceCollectionType

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeResourceCollectionHealthOutput struct {

	// The returned CloudFormationHealthOverview object that contains an
	// InsightHealthOverview object with the requested system health information.
	CloudFormation []types.CloudFormationHealth

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// An array of ServiceHealth objects that describes the health of the Amazon Web
	// Services services associated with the resources in the collection.
	Service []types.ServiceHealth

	// The Amazon Web Services tags that are used by resources in the resource
	// collection. Tags help you identify and organize your Amazon Web Services
	// resources. Many Amazon Web Services services support tagging, so you can assign
	// the same tag to resources from different services to indicate that the resources
	// are related. For example, you can assign the same tag to an Amazon DynamoDB
	// table resource that you assign to an Lambda function. For more information about
	// using tags, see the Tagging best practices
	// (https://docs.aws.amazon.com/whitepapers/latest/tagging-best-practices/tagging-best-practices.html)
	// whitepaper. Each Amazon Web Services tag has two parts.
	//
	// * A tag key (for
	// example, CostCenter, Environment, Project, or Secret). Tag keys are
	// case-sensitive.
	//
	// * An optional field known as a tag value (for example,
	// 111122223333, Production, or a team name). Omitting the tag value is the same as
	// using an empty string. Like tag keys, tag values are case-sensitive.
	//
	// Together
	// these are known as key-value pairs. The string used for a key in a tag that you
	// use to define your resource coverage must begin with the prefix Devops-guru-.
	// The tag key might be DevOps-Guru-deployment-application or
	// devops-guru-rds-application. When you create a key, the case of characters in
	// the key can be whatever you choose. After you create a key, it is
	// case-sensitive. For example, DevOps Guru works with a key named devops-guru-rds
	// and a key named DevOps-Guru-RDS, and these act as two different keys. Possible
	// key/value pairs in your application might be
	// Devops-Guru-production-application/RDS or
	// Devops-Guru-production-application/containers.
	Tags []types.TagHealth

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeResourceCollectionHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeResourceCollectionHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeResourceCollectionHealth{}, middleware.After)
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
	if err = addOpDescribeResourceCollectionHealthValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeResourceCollectionHealth(options.Region), middleware.Before); err != nil {
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

// DescribeResourceCollectionHealthAPIClient is a client that implements the
// DescribeResourceCollectionHealth operation.
type DescribeResourceCollectionHealthAPIClient interface {
	DescribeResourceCollectionHealth(context.Context, *DescribeResourceCollectionHealthInput, ...func(*Options)) (*DescribeResourceCollectionHealthOutput, error)
}

var _ DescribeResourceCollectionHealthAPIClient = (*Client)(nil)

// DescribeResourceCollectionHealthPaginatorOptions is the paginator options for
// DescribeResourceCollectionHealth
type DescribeResourceCollectionHealthPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeResourceCollectionHealthPaginator is a paginator for
// DescribeResourceCollectionHealth
type DescribeResourceCollectionHealthPaginator struct {
	options   DescribeResourceCollectionHealthPaginatorOptions
	client    DescribeResourceCollectionHealthAPIClient
	params    *DescribeResourceCollectionHealthInput
	nextToken *string
	firstPage bool
}

// NewDescribeResourceCollectionHealthPaginator returns a new
// DescribeResourceCollectionHealthPaginator
func NewDescribeResourceCollectionHealthPaginator(client DescribeResourceCollectionHealthAPIClient, params *DescribeResourceCollectionHealthInput, optFns ...func(*DescribeResourceCollectionHealthPaginatorOptions)) *DescribeResourceCollectionHealthPaginator {
	if params == nil {
		params = &DescribeResourceCollectionHealthInput{}
	}

	options := DescribeResourceCollectionHealthPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeResourceCollectionHealthPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeResourceCollectionHealthPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeResourceCollectionHealth page.
func (p *DescribeResourceCollectionHealthPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeResourceCollectionHealthOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeResourceCollectionHealth(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeResourceCollectionHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "DescribeResourceCollectionHealth",
	}
}
