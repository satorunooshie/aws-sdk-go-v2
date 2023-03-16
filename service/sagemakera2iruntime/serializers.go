// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakera2iruntime

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpDeleteHumanLoop struct {
}

func (*awsRestjson1_serializeOpDeleteHumanLoop) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteHumanLoop) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteHumanLoopInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/human-loops/{HumanLoopName}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteHumanLoopInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteHumanLoopInput(v *DeleteHumanLoopInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.HumanLoopName == nil || len(*v.HumanLoopName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member HumanLoopName must not be empty")}
	}
	if v.HumanLoopName != nil {
		if err := encoder.SetURI("HumanLoopName").String(*v.HumanLoopName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDescribeHumanLoop struct {
}

func (*awsRestjson1_serializeOpDescribeHumanLoop) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeHumanLoop) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeHumanLoopInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/human-loops/{HumanLoopName}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDescribeHumanLoopInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeHumanLoopInput(v *DescribeHumanLoopInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.HumanLoopName == nil || len(*v.HumanLoopName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member HumanLoopName must not be empty")}
	}
	if v.HumanLoopName != nil {
		if err := encoder.SetURI("HumanLoopName").String(*v.HumanLoopName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListHumanLoops struct {
}

func (*awsRestjson1_serializeOpListHumanLoops) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListHumanLoops) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListHumanLoopsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/human-loops")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListHumanLoopsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListHumanLoopsInput(v *ListHumanLoopsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.CreationTimeAfter != nil {
		encoder.SetQuery("CreationTimeAfter").String(smithytime.FormatDateTime(*v.CreationTimeAfter))
	}

	if v.CreationTimeBefore != nil {
		encoder.SetQuery("CreationTimeBefore").String(smithytime.FormatDateTime(*v.CreationTimeBefore))
	}

	if v.FlowDefinitionArn != nil {
		encoder.SetQuery("FlowDefinitionArn").String(*v.FlowDefinitionArn)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("MaxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("NextToken").String(*v.NextToken)
	}

	if len(v.SortOrder) > 0 {
		encoder.SetQuery("SortOrder").String(string(v.SortOrder))
	}

	return nil
}

type awsRestjson1_serializeOpStartHumanLoop struct {
}

func (*awsRestjson1_serializeOpStartHumanLoop) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartHumanLoop) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartHumanLoopInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/human-loops")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentStartHumanLoopInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStartHumanLoopInput(v *StartHumanLoopInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStartHumanLoopInput(v *StartHumanLoopInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DataAttributes != nil {
		ok := object.Key("DataAttributes")
		if err := awsRestjson1_serializeDocumentHumanLoopDataAttributes(v.DataAttributes, ok); err != nil {
			return err
		}
	}

	if v.FlowDefinitionArn != nil {
		ok := object.Key("FlowDefinitionArn")
		ok.String(*v.FlowDefinitionArn)
	}

	if v.HumanLoopInput != nil {
		ok := object.Key("HumanLoopInput")
		if err := awsRestjson1_serializeDocumentHumanLoopInput(v.HumanLoopInput, ok); err != nil {
			return err
		}
	}

	if v.HumanLoopName != nil {
		ok := object.Key("HumanLoopName")
		ok.String(*v.HumanLoopName)
	}

	return nil
}

type awsRestjson1_serializeOpStopHumanLoop struct {
}

func (*awsRestjson1_serializeOpStopHumanLoop) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStopHumanLoop) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StopHumanLoopInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/human-loops/stop")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentStopHumanLoopInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStopHumanLoopInput(v *StopHumanLoopInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStopHumanLoopInput(v *StopHumanLoopInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.HumanLoopName != nil {
		ok := object.Key("HumanLoopName")
		ok.String(*v.HumanLoopName)
	}

	return nil
}

func awsRestjson1_serializeDocumentContentClassifiers(v []types.ContentClassifier, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentHumanLoopDataAttributes(v *types.HumanLoopDataAttributes, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContentClassifiers != nil {
		ok := object.Key("ContentClassifiers")
		if err := awsRestjson1_serializeDocumentContentClassifiers(v.ContentClassifiers, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentHumanLoopInput(v *types.HumanLoopInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.InputContent != nil {
		ok := object.Key("InputContent")
		ok.String(*v.InputContent)
	}

	return nil
}
