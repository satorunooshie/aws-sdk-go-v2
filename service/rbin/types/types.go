// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Information about a retention rule lock configuration.
type LockConfiguration struct {

	// Information about the retention rule unlock delay.
	//
	// This member is required.
	UnlockDelay *UnlockDelay

	noSmithyDocumentSerde
}

// Information about the resource tags used to identify resources that are retained
// by the retention rule.
type ResourceTag struct {

	// The tag key.
	//
	// This member is required.
	ResourceTagKey *string

	// The tag value.
	ResourceTagValue *string

	noSmithyDocumentSerde
}

// Information about the retention period for which the retention rule is to retain
// resources.
type RetentionPeriod struct {

	// The unit of time in which the retention period is measured. Currently, only DAYS
	// is supported.
	//
	// This member is required.
	RetentionPeriodUnit RetentionPeriodUnit

	// The period value for which the retention rule is to retain resources. The period
	// is measured using the unit specified for RetentionPeriodUnit.
	//
	// This member is required.
	RetentionPeriodValue *int32

	noSmithyDocumentSerde
}

// Information about a Recycle Bin retention rule.
type RuleSummary struct {

	// The retention rule description.
	Description *string

	// The unique ID of the retention rule.
	Identifier *string

	// The lock state for the retention rule.
	//
	// * locked - The retention rule is locked
	// and can't be modified or deleted.
	//
	// * pending_unlock - The retention rule has
	// been unlocked but it is still within the unlock delay period. The retention rule
	// can be modified or deleted only after the unlock delay period has expired.
	//
	// *
	// unlocked - The retention rule is unlocked and it can be modified or deleted by
	// any user with the required permissions.
	//
	// * null - The retention rule has never
	// been locked. Once a retention rule has been locked, it can transition between
	// the locked and unlocked states only; it can never transition back to null.
	LockState LockState

	// Information about the retention period for which the retention rule is to retain
	// resources.
	RetentionPeriod *RetentionPeriod

	noSmithyDocumentSerde
}

// Information about the tags to assign to the retention rule.
type Tag struct {

	// The tag key.
	//
	// This member is required.
	Key *string

	// The tag value.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Information about the retention rule unlock delay. The unlock delay is the
// period after which a retention rule can be modified or edited after it has been
// unlocked by a user with the required permissions. The retention rule can't be
// modified or deleted during the unlock delay.
type UnlockDelay struct {

	// The unit of time in which to measure the unlock delay. Currently, the unlock
	// delay can be measure only in days.
	//
	// This member is required.
	UnlockDelayUnit UnlockDelayUnit

	// The unlock delay period, measured in the unit specified for UnlockDelayUnit.
	//
	// This member is required.
	UnlockDelayValue *int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
