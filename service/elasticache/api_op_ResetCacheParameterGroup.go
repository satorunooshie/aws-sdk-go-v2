// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ResetCacheParameterGroup operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ResetCacheParameterGroupMessage
type ResetCacheParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache parameter group to reset.
	//
	// CacheParameterGroupName is a required field
	CacheParameterGroupName *string `type:"string" required:"true"`

	// An array of parameter names to reset to their default values. If ResetAllParameters
	// is true, do not use ParameterNameValues. If ResetAllParameters is false,
	// you must specify the name of at least one parameter to reset.
	ParameterNameValues []ParameterNameValue `locationNameList:"ParameterNameValue" type:"list"`

	// If true, all parameters in the cache parameter group are reset to their default
	// values. If false, only the parameters listed by ParameterNameValues are reset
	// to their default values.
	//
	// Valid values: true | false
	ResetAllParameters *bool `type:"boolean"`
}

// String returns the string representation
func (s ResetCacheParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetCacheParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetCacheParameterGroupInput"}

	if s.CacheParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of one of the following operations:
//
//    * ModifyCacheParameterGroup
//
//    * ResetCacheParameterGroup
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CacheParameterGroupNameMessage
type ResetCacheParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of the cache parameter group.
	CacheParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ResetCacheParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opResetCacheParameterGroup = "ResetCacheParameterGroup"

// ResetCacheParameterGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Modifies the parameters of a cache parameter group to the engine or system
// default value. You can reset specific parameters by submitting a list of
// parameter names. To reset the entire cache parameter group, specify the ResetAllParameters
// and CacheParameterGroupName parameters.
//
//    // Example sending a request using ResetCacheParameterGroupRequest.
//    req := client.ResetCacheParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ResetCacheParameterGroup
func (c *Client) ResetCacheParameterGroupRequest(input *ResetCacheParameterGroupInput) ResetCacheParameterGroupRequest {
	op := &aws.Operation{
		Name:       opResetCacheParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetCacheParameterGroupInput{}
	}

	req := c.newRequest(op, input, &ResetCacheParameterGroupOutput{})
	return ResetCacheParameterGroupRequest{Request: req, Input: input, Copy: c.ResetCacheParameterGroupRequest}
}

// ResetCacheParameterGroupRequest is the request type for the
// ResetCacheParameterGroup API operation.
type ResetCacheParameterGroupRequest struct {
	*aws.Request
	Input *ResetCacheParameterGroupInput
	Copy  func(*ResetCacheParameterGroupInput) ResetCacheParameterGroupRequest
}

// Send marshals and sends the ResetCacheParameterGroup API request.
func (r ResetCacheParameterGroupRequest) Send(ctx context.Context) (*ResetCacheParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetCacheParameterGroupResponse{
		ResetCacheParameterGroupOutput: r.Request.Data.(*ResetCacheParameterGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetCacheParameterGroupResponse is the response type for the
// ResetCacheParameterGroup API operation.
type ResetCacheParameterGroupResponse struct {
	*ResetCacheParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetCacheParameterGroup request.
func (r *ResetCacheParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}