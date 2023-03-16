// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ArrayJobDependency string

// Enum values for ArrayJobDependency
const (
	ArrayJobDependencyNToN       ArrayJobDependency = "N_TO_N"
	ArrayJobDependencySequential ArrayJobDependency = "SEQUENTIAL"
)

// Values returns all known values for ArrayJobDependency. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArrayJobDependency) Values() []ArrayJobDependency {
	return []ArrayJobDependency{
		"N_TO_N",
		"SEQUENTIAL",
	}
}

type AssignPublicIp string

// Enum values for AssignPublicIp
const (
	AssignPublicIpEnabled  AssignPublicIp = "ENABLED"
	AssignPublicIpDisabled AssignPublicIp = "DISABLED"
)

// Values returns all known values for AssignPublicIp. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssignPublicIp) Values() []AssignPublicIp {
	return []AssignPublicIp{
		"ENABLED",
		"DISABLED",
	}
}

type CEState string

// Enum values for CEState
const (
	CEStateEnabled  CEState = "ENABLED"
	CEStateDisabled CEState = "DISABLED"
)

// Values returns all known values for CEState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (CEState) Values() []CEState {
	return []CEState{
		"ENABLED",
		"DISABLED",
	}
}

type CEStatus string

// Enum values for CEStatus
const (
	CEStatusCreating CEStatus = "CREATING"
	CEStatusUpdating CEStatus = "UPDATING"
	CEStatusDeleting CEStatus = "DELETING"
	CEStatusDeleted  CEStatus = "DELETED"
	CEStatusValid    CEStatus = "VALID"
	CEStatusInvalid  CEStatus = "INVALID"
)

// Values returns all known values for CEStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (CEStatus) Values() []CEStatus {
	return []CEStatus{
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"VALID",
		"INVALID",
	}
}

type CEType string

// Enum values for CEType
const (
	CETypeManaged   CEType = "MANAGED"
	CETypeUnmanaged CEType = "UNMANAGED"
)

// Values returns all known values for CEType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (CEType) Values() []CEType {
	return []CEType{
		"MANAGED",
		"UNMANAGED",
	}
}

type CRAllocationStrategy string

// Enum values for CRAllocationStrategy
const (
	CRAllocationStrategyBestFit               CRAllocationStrategy = "BEST_FIT"
	CRAllocationStrategyBestFitProgressive    CRAllocationStrategy = "BEST_FIT_PROGRESSIVE"
	CRAllocationStrategySpotCapacityOptimized CRAllocationStrategy = "SPOT_CAPACITY_OPTIMIZED"
)

// Values returns all known values for CRAllocationStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CRAllocationStrategy) Values() []CRAllocationStrategy {
	return []CRAllocationStrategy{
		"BEST_FIT",
		"BEST_FIT_PROGRESSIVE",
		"SPOT_CAPACITY_OPTIMIZED",
	}
}

type CRType string

// Enum values for CRType
const (
	CRTypeEc2         CRType = "EC2"
	CRTypeSpot        CRType = "SPOT"
	CRTypeFargate     CRType = "FARGATE"
	CRTypeFargateSpot CRType = "FARGATE_SPOT"
)

// Values returns all known values for CRType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (CRType) Values() []CRType {
	return []CRType{
		"EC2",
		"SPOT",
		"FARGATE",
		"FARGATE_SPOT",
	}
}

type CRUpdateAllocationStrategy string

// Enum values for CRUpdateAllocationStrategy
const (
	CRUpdateAllocationStrategyBestFitProgressive    CRUpdateAllocationStrategy = "BEST_FIT_PROGRESSIVE"
	CRUpdateAllocationStrategySpotCapacityOptimized CRUpdateAllocationStrategy = "SPOT_CAPACITY_OPTIMIZED"
)

// Values returns all known values for CRUpdateAllocationStrategy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (CRUpdateAllocationStrategy) Values() []CRUpdateAllocationStrategy {
	return []CRUpdateAllocationStrategy{
		"BEST_FIT_PROGRESSIVE",
		"SPOT_CAPACITY_OPTIMIZED",
	}
}

type DeviceCgroupPermission string

// Enum values for DeviceCgroupPermission
const (
	DeviceCgroupPermissionRead  DeviceCgroupPermission = "READ"
	DeviceCgroupPermissionWrite DeviceCgroupPermission = "WRITE"
	DeviceCgroupPermissionMknod DeviceCgroupPermission = "MKNOD"
)

// Values returns all known values for DeviceCgroupPermission. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeviceCgroupPermission) Values() []DeviceCgroupPermission {
	return []DeviceCgroupPermission{
		"READ",
		"WRITE",
		"MKNOD",
	}
}

type EFSAuthorizationConfigIAM string

// Enum values for EFSAuthorizationConfigIAM
const (
	EFSAuthorizationConfigIAMEnabled  EFSAuthorizationConfigIAM = "ENABLED"
	EFSAuthorizationConfigIAMDisabled EFSAuthorizationConfigIAM = "DISABLED"
)

// Values returns all known values for EFSAuthorizationConfigIAM. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EFSAuthorizationConfigIAM) Values() []EFSAuthorizationConfigIAM {
	return []EFSAuthorizationConfigIAM{
		"ENABLED",
		"DISABLED",
	}
}

type EFSTransitEncryption string

// Enum values for EFSTransitEncryption
const (
	EFSTransitEncryptionEnabled  EFSTransitEncryption = "ENABLED"
	EFSTransitEncryptionDisabled EFSTransitEncryption = "DISABLED"
)

// Values returns all known values for EFSTransitEncryption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EFSTransitEncryption) Values() []EFSTransitEncryption {
	return []EFSTransitEncryption{
		"ENABLED",
		"DISABLED",
	}
}

type JobDefinitionType string

// Enum values for JobDefinitionType
const (
	JobDefinitionTypeContainer JobDefinitionType = "container"
	JobDefinitionTypeMultinode JobDefinitionType = "multinode"
)

// Values returns all known values for JobDefinitionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JobDefinitionType) Values() []JobDefinitionType {
	return []JobDefinitionType{
		"container",
		"multinode",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusSubmitted JobStatus = "SUBMITTED"
	JobStatusPending   JobStatus = "PENDING"
	JobStatusRunnable  JobStatus = "RUNNABLE"
	JobStatusStarting  JobStatus = "STARTING"
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusSucceeded JobStatus = "SUCCEEDED"
	JobStatusFailed    JobStatus = "FAILED"
)

// Values returns all known values for JobStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"SUBMITTED",
		"PENDING",
		"RUNNABLE",
		"STARTING",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
	}
}

type JQState string

// Enum values for JQState
const (
	JQStateEnabled  JQState = "ENABLED"
	JQStateDisabled JQState = "DISABLED"
)

// Values returns all known values for JQState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JQState) Values() []JQState {
	return []JQState{
		"ENABLED",
		"DISABLED",
	}
}

type JQStatus string

// Enum values for JQStatus
const (
	JQStatusCreating JQStatus = "CREATING"
	JQStatusUpdating JQStatus = "UPDATING"
	JQStatusDeleting JQStatus = "DELETING"
	JQStatusDeleted  JQStatus = "DELETED"
	JQStatusValid    JQStatus = "VALID"
	JQStatusInvalid  JQStatus = "INVALID"
)

// Values returns all known values for JQStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JQStatus) Values() []JQStatus {
	return []JQStatus{
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"VALID",
		"INVALID",
	}
}

type LogDriver string

// Enum values for LogDriver
const (
	LogDriverJsonFile LogDriver = "json-file"
	LogDriverSyslog   LogDriver = "syslog"
	LogDriverJournald LogDriver = "journald"
	LogDriverGelf     LogDriver = "gelf"
	LogDriverFluentd  LogDriver = "fluentd"
	LogDriverAwslogs  LogDriver = "awslogs"
	LogDriverSplunk   LogDriver = "splunk"
)

// Values returns all known values for LogDriver. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogDriver) Values() []LogDriver {
	return []LogDriver{
		"json-file",
		"syslog",
		"journald",
		"gelf",
		"fluentd",
		"awslogs",
		"splunk",
	}
}

type OrchestrationType string

// Enum values for OrchestrationType
const (
	OrchestrationTypeEcs OrchestrationType = "ECS"
	OrchestrationTypeEks OrchestrationType = "EKS"
)

// Values returns all known values for OrchestrationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrchestrationType) Values() []OrchestrationType {
	return []OrchestrationType{
		"ECS",
		"EKS",
	}
}

type PlatformCapability string

// Enum values for PlatformCapability
const (
	PlatformCapabilityEc2     PlatformCapability = "EC2"
	PlatformCapabilityFargate PlatformCapability = "FARGATE"
)

// Values returns all known values for PlatformCapability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlatformCapability) Values() []PlatformCapability {
	return []PlatformCapability{
		"EC2",
		"FARGATE",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeGpu    ResourceType = "GPU"
	ResourceTypeVcpu   ResourceType = "VCPU"
	ResourceTypeMemory ResourceType = "MEMORY"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"GPU",
		"VCPU",
		"MEMORY",
	}
}

type RetryAction string

// Enum values for RetryAction
const (
	RetryActionRetry RetryAction = "RETRY"
	RetryActionExit  RetryAction = "EXIT"
)

// Values returns all known values for RetryAction. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RetryAction) Values() []RetryAction {
	return []RetryAction{
		"RETRY",
		"EXIT",
	}
}
