// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Casts a vote for a specified ProposalId on behalf of a member. The member to
// vote as, specified by VoterMemberId, must be in the same Amazon Web Services
// account as the principal that calls the action. Applies only to Hyperledger
// Fabric.
func (c *Client) VoteOnProposal(ctx context.Context, params *VoteOnProposalInput, optFns ...func(*Options)) (*VoteOnProposalOutput, error) {
	if params == nil {
		params = &VoteOnProposalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VoteOnProposal", params, optFns, c.addOperationVoteOnProposalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VoteOnProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VoteOnProposalInput struct {

	// The unique identifier of the network.
	//
	// This member is required.
	NetworkId *string

	// The unique identifier of the proposal.
	//
	// This member is required.
	ProposalId *string

	// The value of the vote.
	//
	// This member is required.
	Vote types.VoteValue

	// The unique identifier of the member casting the vote.
	//
	// This member is required.
	VoterMemberId *string

	noSmithyDocumentSerde
}

type VoteOnProposalOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVoteOnProposalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpVoteOnProposal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpVoteOnProposal{}, middleware.After)
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
	if err = addOpVoteOnProposalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVoteOnProposal(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVoteOnProposal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "VoteOnProposal",
	}
}
