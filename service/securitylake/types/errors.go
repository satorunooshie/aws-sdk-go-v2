// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You do not have sufficient access to perform this action. Access denied errors
// appear when Amazon Security Lake explicitly or implicitly denies an
// authorization request. An explicit denial occurs when a policy contains a Deny
// statement for the specific Amazon Web Services action. An implicit denial occurs
// when there is no applicable Deny statement and also no applicable Allow
// statement.
type AccessDeniedException struct {
	Message *string

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
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Security Lake cannot find an Amazon Web Services account with the
// accountID that you specified, or the account whose credentials you used to make
// this request isn't a member of an organization.
type AccountNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AccountNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccountNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccountNotFoundException) ErrorCode() string             { return "AccountNotFoundException" }
func (e *AccountNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Security Lake generally returns 404 errors if the requested object is
// missing from the bucket.
type BucketNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BucketNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BucketNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BucketNotFoundException) ErrorCode() string             { return "BucketNotFoundException" }
func (e *BucketNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// More than one process tried to modify a resource at the same time.
type ConcurrentModificationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Occurs when a conflict with a previous successful write is detected. This
// generally occurs when the previous write did not have time to propagate to the
// host serving the current request. A retry (with appropriate backoff logic) is
// the recommended response to this exception.
type ConflictException struct {
	Message *string

	ResourceId   *string
	ResourceType *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was a conflict when you attempted to modify a Security Lake source name.
type ConflictSourceNamesException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConflictSourceNamesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictSourceNamesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictSourceNamesException) ErrorCode() string             { return "ConflictSourceNamesException" }
func (e *ConflictSourceNamesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A conflicting subscription exception operation is in progress.
type ConflictSubscriptionException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConflictSubscriptionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictSubscriptionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictSubscriptionException) ErrorCode() string             { return "ConflictSubscriptionException" }
func (e *ConflictSubscriptionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Represents an error interacting with the Amazon EventBridge service.
type EventBridgeException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *EventBridgeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EventBridgeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EventBridgeException) ErrorCode() string             { return "EventBridgeException" }
func (e *EventBridgeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Internal service exceptions are sometimes caused by transient issues. Before you
// start troubleshooting, perform the operation again.
type InternalServerException struct {
	Message *string

	RetryAfterSeconds *int32

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string             { return "InternalServerException" }
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request was rejected because a value that's not valid or is out of range was
// supplied for an input parameter.
type InvalidInputException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource could not be found.
type ResourceNotFoundException struct {
	Message *string

	ResourceId   *string
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
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Provides an extension of the AmazonServiceException for errors reported by
// Amazon S3 while processing a request. In particular, this class provides access
// to the Amazon S3 extended request ID. If Amazon S3 is incorrectly handling a
// request and you need to contact Amazon, this extended request ID may provide
// useful debugging information.
type S3Exception struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *S3Exception) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *S3Exception) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *S3Exception) ErrorCode() string             { return "S3Exception" }
func (e *S3Exception) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded your service quota. To perform the requested action, remove
// some of the relevant resources, or use Service Quotas to request a service quota
// increase.
type ServiceQuotaExceededException struct {
	Message *string

	ResourceId   *string
	ResourceType *string
	ServiceCode  *string
	QuotaCode    *string

	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string             { return "ServiceQuotaExceededException" }
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The limit on the number of requests per second was exceeded.
type ThrottlingException struct {
	Message *string

	ServiceCode       *string
	QuotaCode         *string
	RetryAfterSeconds *int32

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your signing certificate could not be validated.
type ValidationException struct {
	Message *string

	Reason    ValidationExceptionReason
	FieldList []ValidationExceptionField

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
func (e *ValidationException) ErrorCode() string             { return "ValidationException" }
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
