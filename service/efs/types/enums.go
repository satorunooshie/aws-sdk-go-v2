// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type LifeCycleState string

// Enum values for LifeCycleState
const (
	LifeCycleStateCreating  LifeCycleState = "creating"
	LifeCycleStateAvailable LifeCycleState = "available"
	LifeCycleStateUpdating  LifeCycleState = "updating"
	LifeCycleStateDeleting  LifeCycleState = "deleting"
	LifeCycleStateDeleted   LifeCycleState = "deleted"
	LifeCycleStateError     LifeCycleState = "error"
)

// Values returns all known values for LifeCycleState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LifeCycleState) Values() []LifeCycleState {
	return []LifeCycleState{
		"creating",
		"available",
		"updating",
		"deleting",
		"deleted",
		"error",
	}
}

type PerformanceMode string

// Enum values for PerformanceMode
const (
	PerformanceModeGeneralPurpose PerformanceMode = "generalPurpose"
	PerformanceModeMaxIo          PerformanceMode = "maxIO"
)

// Values returns all known values for PerformanceMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PerformanceMode) Values() []PerformanceMode {
	return []PerformanceMode{
		"generalPurpose",
		"maxIO",
	}
}

type ReplicationStatus string

// Enum values for ReplicationStatus
const (
	ReplicationStatusEnabled  ReplicationStatus = "ENABLED"
	ReplicationStatusEnabling ReplicationStatus = "ENABLING"
	ReplicationStatusDeleting ReplicationStatus = "DELETING"
	ReplicationStatusError    ReplicationStatus = "ERROR"
)

// Values returns all known values for ReplicationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationStatus) Values() []ReplicationStatus {
	return []ReplicationStatus{
		"ENABLED",
		"ENABLING",
		"DELETING",
		"ERROR",
	}
}

type Resource string

// Enum values for Resource
const (
	ResourceFileSystem  Resource = "FILE_SYSTEM"
	ResourceMountTarget Resource = "MOUNT_TARGET"
)

// Values returns all known values for Resource. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Resource) Values() []Resource {
	return []Resource{
		"FILE_SYSTEM",
		"MOUNT_TARGET",
	}
}

type ResourceIdType string

// Enum values for ResourceIdType
const (
	ResourceIdTypeLongId  ResourceIdType = "LONG_ID"
	ResourceIdTypeShortId ResourceIdType = "SHORT_ID"
)

// Values returns all known values for ResourceIdType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceIdType) Values() []ResourceIdType {
	return []ResourceIdType{
		"LONG_ID",
		"SHORT_ID",
	}
}

type Status string

// Enum values for Status
const (
	StatusEnabled   Status = "ENABLED"
	StatusEnabling  Status = "ENABLING"
	StatusDisabled  Status = "DISABLED"
	StatusDisabling Status = "DISABLING"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"ENABLED",
		"ENABLING",
		"DISABLED",
		"DISABLING",
	}
}

type ThroughputMode string

// Enum values for ThroughputMode
const (
	ThroughputModeBursting    ThroughputMode = "bursting"
	ThroughputModeProvisioned ThroughputMode = "provisioned"
	ThroughputModeElastic     ThroughputMode = "elastic"
)

// Values returns all known values for ThroughputMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThroughputMode) Values() []ThroughputMode {
	return []ThroughputMode{
		"bursting",
		"provisioned",
		"elastic",
	}
}

type TransitionToIARules string

// Enum values for TransitionToIARules
const (
	TransitionToIARulesAfter7Days  TransitionToIARules = "AFTER_7_DAYS"
	TransitionToIARulesAfter14Days TransitionToIARules = "AFTER_14_DAYS"
	TransitionToIARulesAfter30Days TransitionToIARules = "AFTER_30_DAYS"
	TransitionToIARulesAfter60Days TransitionToIARules = "AFTER_60_DAYS"
	TransitionToIARulesAfter90Days TransitionToIARules = "AFTER_90_DAYS"
	TransitionToIARulesAfter1Day   TransitionToIARules = "AFTER_1_DAY"
)

// Values returns all known values for TransitionToIARules. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransitionToIARules) Values() []TransitionToIARules {
	return []TransitionToIARules{
		"AFTER_7_DAYS",
		"AFTER_14_DAYS",
		"AFTER_30_DAYS",
		"AFTER_60_DAYS",
		"AFTER_90_DAYS",
		"AFTER_1_DAY",
	}
}

type TransitionToPrimaryStorageClassRules string

// Enum values for TransitionToPrimaryStorageClassRules
const (
	TransitionToPrimaryStorageClassRulesAfter1Access TransitionToPrimaryStorageClassRules = "AFTER_1_ACCESS"
)

// Values returns all known values for TransitionToPrimaryStorageClassRules. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TransitionToPrimaryStorageClassRules) Values() []TransitionToPrimaryStorageClassRules {
	return []TransitionToPrimaryStorageClassRules{
		"AFTER_1_ACCESS",
	}
}
