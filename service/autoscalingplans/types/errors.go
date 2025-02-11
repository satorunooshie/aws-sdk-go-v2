// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Concurrent updates caused an exception, for example, if you request an update to
// a scaling plan that already has a pending update.
type ConcurrentUpdateException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConcurrentUpdateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentUpdateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentUpdateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConcurrentUpdateException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConcurrentUpdateException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The service encountered an internal error.
type InternalServiceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServiceException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The token provided is not valid.
type InvalidNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidNextTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your account exceeded a limit. This exception is thrown when a per-account
// resource limit is exceeded.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified object could not be found.
type ObjectNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ObjectNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ObjectNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ObjectNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception was thrown for a validation issue. Review the parameters provided.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
