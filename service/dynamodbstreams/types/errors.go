// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The shard iterator has expired and can no longer be used to retrieve stream
// records. A shard iterator expires 15 minutes after it is retrieved using the
// GetShardIterator action.
type ExpiredIteratorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ExpiredIteratorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExpiredIteratorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExpiredIteratorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ExpiredIteratorException"
	}
	return *e.ErrorCodeOverride
}
func (e *ExpiredIteratorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred on the server side.
type InternalServerError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerError"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// There is no limit to the number of daily on-demand backups that can be taken.
// For most purposes, up to 500 simultaneous table operations are allowed per
// account. These operations include CreateTable, UpdateTable,
// DeleteTable,UpdateTimeToLive, RestoreTableFromBackup, and
// RestoreTableToPointInTime. When you are creating a table with one or more
// secondary indexes, you can have up to 250 such requests running at a time.
// However, if the table or index specifications are complex, then DynamoDB might
// temporarily reduce the number of concurrent operations. When importing into
// DynamoDB, up to 50 simultaneous import table operations are allowed per account.
// There is a soft account quota of 2,500 tables.
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

// The operation tried to access a nonexistent table or index. The resource might
// not be specified correctly, or its status might not be ACTIVE.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

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

// The operation attempted to read past the oldest stream record in a shard. In
// DynamoDB Streams, there is a 24 hour limit on data retention. Stream records
// whose age exceeds this limit are subject to removal (trimming) from the stream.
// You might receive a TrimmedDataAccessException if:
//
// * You request a shard
// iterator with a sequence number older than the trim point (24 hours).
//
// * You
// obtain a shard iterator, but before you use the iterator in a GetRecords
// request, a stream record in the shard exceeds the 24 hour period and is trimmed.
// This causes the iterator to access a record that no longer exists.
type TrimmedDataAccessException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TrimmedDataAccessException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrimmedDataAccessException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrimmedDataAccessException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TrimmedDataAccessException"
	}
	return *e.ErrorCodeOverride
}
func (e *TrimmedDataAccessException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
