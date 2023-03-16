// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Exception that indicates the specified AttackId does not exist, or the requester
// does not have the appropriate permissions to access the AttackId.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// In order to grant the necessary access to the Shield Response Team (SRT) the
// user submitting the request must have the iam:PassRole permission. This error
// indicates the user did not have the appropriate permissions. For more
// information, see Granting a User Permissions to Pass a Role to an Amazon Web
// Services Service
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html).
type AccessDeniedForDependencyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedForDependencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedForDependencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedForDependencyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedForDependencyException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedForDependencyException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Exception that indicates that a problem occurred with the service
// infrastructure. You can retry the request.
type InternalErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Exception that indicates that the operation would not cause any change to occur.
type InvalidOperationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOperationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidOperationException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception that indicates that the NextToken specified in the request is invalid.
// Submit the request using the NextToken value that was returned in the prior
// response.
type InvalidPaginationTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidPaginationTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPaginationTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPaginationTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidPaginationTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidPaginationTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception that indicates that the parameters passed to the API are invalid. If
// available, this exception includes details in additional properties.
type InvalidParameterException struct {
	Message *string

	ErrorCodeOverride *string

	Reason ValidationExceptionReason
	Fields []ValidationExceptionField

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception that indicates that the resource is invalid. You might not have access
// to the resource, or the resource might not exist.
type InvalidResourceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidResourceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidResourceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidResourceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidResourceException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidResourceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception that indicates that the operation would exceed a limit.
type LimitsExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Type  *string
	Limit int64

	noSmithyDocumentSerde
}

func (e *LimitsExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitsExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitsExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitsExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitsExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You are trying to update a subscription that has not yet completed the 1-year
// commitment. You can change the AutoRenew parameter during the last 30 days of
// your subscription. This exception indicates that you are attempting to change
// AutoRenew prior to that period.
type LockedSubscriptionException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LockedSubscriptionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LockedSubscriptionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LockedSubscriptionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LockedSubscriptionException"
	}
	return *e.ErrorCodeOverride
}
func (e *LockedSubscriptionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The ARN of the role that you specified does not exist.
type NoAssociatedRoleException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NoAssociatedRoleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoAssociatedRoleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoAssociatedRoleException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NoAssociatedRoleException"
	}
	return *e.ErrorCodeOverride
}
func (e *NoAssociatedRoleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception that indicates that the resource state has been modified by another
// client. Retrieve the resource and then retry your request.
type OptimisticLockException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *OptimisticLockException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OptimisticLockException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OptimisticLockException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OptimisticLockException"
	}
	return *e.ErrorCodeOverride
}
func (e *OptimisticLockException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception indicating the specified resource already exists. If available, this
// exception includes details in additional properties.
type ResourceAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceType *string

	noSmithyDocumentSerde
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception indicating the specified resource does not exist. If available, this
// exception includes details in additional properties.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceType *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
