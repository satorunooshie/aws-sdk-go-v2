// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCancelQuantumTask struct {
}

func (*validateOpCancelQuantumTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelQuantumTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelQuantumTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelQuantumTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateQuantumTask struct {
}

func (*validateOpCreateQuantumTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateQuantumTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateQuantumTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateQuantumTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDevice struct {
}

func (*validateOpGetDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQuantumTask struct {
}

func (*validateOpGetQuantumTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQuantumTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQuantumTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQuantumTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchDevices struct {
}

func (*validateOpSearchDevices) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchDevices) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchDevicesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchDevicesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchQuantumTasks struct {
}

func (*validateOpSearchQuantumTasks) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchQuantumTasks) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchQuantumTasksInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchQuantumTasksInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelQuantumTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelQuantumTask{}, middleware.After)
}

func addOpCreateQuantumTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateQuantumTask{}, middleware.After)
}

func addOpGetDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDevice{}, middleware.After)
}

func addOpGetQuantumTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQuantumTask{}, middleware.After)
}

func addOpSearchDevicesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchDevices{}, middleware.After)
}

func addOpSearchQuantumTasksValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchQuantumTasks{}, middleware.After)
}

func validateSearchDevicesFilter(v *types.SearchDevicesFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchDevicesFilter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchDevicesFilterList(v []*types.SearchDevicesFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchDevicesFilterList"}
	for i := range v {
		if err := validateSearchDevicesFilter(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchQuantumTasksFilter(v *types.SearchQuantumTasksFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchQuantumTasksFilter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if len(v.Operator) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Operator"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchQuantumTasksFilterList(v []*types.SearchQuantumTasksFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchQuantumTasksFilterList"}
	for i := range v {
		if err := validateSearchQuantumTasksFilter(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelQuantumTaskInput(v *CancelQuantumTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelQuantumTaskInput"}
	if v.QuantumTaskArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QuantumTaskArn"))
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateQuantumTaskInput(v *CreateQuantumTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateQuantumTaskInput"}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.DeviceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceArn"))
	}
	if v.Shots == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Shots"))
	}
	if v.OutputS3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputS3Bucket"))
	}
	if v.OutputS3KeyPrefix == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputS3KeyPrefix"))
	}
	if v.Action == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDeviceInput(v *GetDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDeviceInput"}
	if v.DeviceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQuantumTaskInput(v *GetQuantumTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQuantumTaskInput"}
	if v.QuantumTaskArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QuantumTaskArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchDevicesInput(v *SearchDevicesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchDevicesInput"}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	} else if v.Filters != nil {
		if err := validateSearchDevicesFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchQuantumTasksInput(v *SearchQuantumTasksInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchQuantumTasksInput"}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	} else if v.Filters != nil {
		if err := validateSearchQuantumTasksFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}