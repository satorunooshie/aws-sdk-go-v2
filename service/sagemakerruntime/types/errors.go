// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Your request caused an exception with an internal dependency. Contact customer
// support.
type InternalDependencyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalDependencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalDependencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalDependencyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalDependencyException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalDependencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// An internal failure occurred.
type InternalFailure struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalFailure) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailure) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailure) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalFailure"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalFailure) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Model (owned by the customer in the container) returned 4xx or 5xx error code.
type ModelError struct {
	Message *string

	ErrorCodeOverride *string

	OriginalStatusCode *int32
	OriginalMessage    *string
	LogStreamArn       *string

	noSmithyDocumentSerde
}

func (e *ModelError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ModelError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ModelError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ModelError"
	}
	return *e.ErrorCodeOverride
}
func (e *ModelError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Either a serverless endpoint variant's resources are still being provisioned, or
// a multi-model endpoint is still downloading or loading the target model. Wait
// and try your request again.
type ModelNotReadyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ModelNotReadyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ModelNotReadyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ModelNotReadyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ModelNotReadyException"
	}
	return *e.ErrorCodeOverride
}
func (e *ModelNotReadyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is unavailable. Try your call again.
type ServiceUnavailable struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailable) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailable) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailable) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceUnavailable"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailable) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Inspect your request and try again.
type ValidationError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationError"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
