// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypeReplace ChangeType = "REPLACE"
	ChangeTypeAppend  ChangeType = "APPEND"
	ChangeTypeModify  ChangeType = "MODIFY"
)

// Values returns all known values for ChangeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeType) Values() []ChangeType {
	return []ChangeType{
		"REPLACE",
		"APPEND",
		"MODIFY",
	}
}

type ColumnDataType string

// Enum values for ColumnDataType
const (
	ColumnDataTypeString   ColumnDataType = "STRING"
	ColumnDataTypeChar     ColumnDataType = "CHAR"
	ColumnDataTypeInteger  ColumnDataType = "INTEGER"
	ColumnDataTypeTinyint  ColumnDataType = "TINYINT"
	ColumnDataTypeSmallint ColumnDataType = "SMALLINT"
	ColumnDataTypeBigint   ColumnDataType = "BIGINT"
	ColumnDataTypeFloat    ColumnDataType = "FLOAT"
	ColumnDataTypeDouble   ColumnDataType = "DOUBLE"
	ColumnDataTypeDate     ColumnDataType = "DATE"
	ColumnDataTypeDatetime ColumnDataType = "DATETIME"
	ColumnDataTypeBoolean  ColumnDataType = "BOOLEAN"
	ColumnDataTypeBinary   ColumnDataType = "BINARY"
)

// Values returns all known values for ColumnDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ColumnDataType) Values() []ColumnDataType {
	return []ColumnDataType{
		"STRING",
		"CHAR",
		"INTEGER",
		"TINYINT",
		"SMALLINT",
		"BIGINT",
		"FLOAT",
		"DOUBLE",
		"DATE",
		"DATETIME",
		"BOOLEAN",
		"BINARY",
	}
}

type DatasetKind string

// Enum values for DatasetKind
const (
	DatasetKindTabular    DatasetKind = "TABULAR"
	DatasetKindNonTabular DatasetKind = "NON_TABULAR"
)

// Values returns all known values for DatasetKind. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DatasetKind) Values() []DatasetKind {
	return []DatasetKind{
		"TABULAR",
		"NON_TABULAR",
	}
}

type DatasetStatus string

// Enum values for DatasetStatus
const (
	DatasetStatusPending DatasetStatus = "PENDING"
	DatasetStatusFailed  DatasetStatus = "FAILED"
	DatasetStatusSuccess DatasetStatus = "SUCCESS"
	DatasetStatusRunning DatasetStatus = "RUNNING"
)

// Values returns all known values for DatasetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DatasetStatus) Values() []DatasetStatus {
	return []DatasetStatus{
		"PENDING",
		"FAILED",
		"SUCCESS",
		"RUNNING",
	}
}

type DataViewStatus string

// Enum values for DataViewStatus
const (
	DataViewStatusRunning             DataViewStatus = "RUNNING"
	DataViewStatusStarting            DataViewStatus = "STARTING"
	DataViewStatusFailed              DataViewStatus = "FAILED"
	DataViewStatusCancelled           DataViewStatus = "CANCELLED"
	DataViewStatusTimeout             DataViewStatus = "TIMEOUT"
	DataViewStatusSuccess             DataViewStatus = "SUCCESS"
	DataViewStatusPending             DataViewStatus = "PENDING"
	DataViewStatusFailedCleanupFailed DataViewStatus = "FAILED_CLEANUP_FAILED"
)

// Values returns all known values for DataViewStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataViewStatus) Values() []DataViewStatus {
	return []DataViewStatus{
		"RUNNING",
		"STARTING",
		"FAILED",
		"CANCELLED",
		"TIMEOUT",
		"SUCCESS",
		"PENDING",
		"FAILED_CLEANUP_FAILED",
	}
}

type ErrorCategory string

// Enum values for ErrorCategory
const (
	ErrorCategoryValidation               ErrorCategory = "VALIDATION"
	ErrorCategoryServiceQuotaExceeded     ErrorCategory = "SERVICE_QUOTA_EXCEEDED"
	ErrorCategoryAccessDenied             ErrorCategory = "ACCESS_DENIED"
	ErrorCategoryResourceNotFound         ErrorCategory = "RESOURCE_NOT_FOUND"
	ErrorCategoryThrottling               ErrorCategory = "THROTTLING"
	ErrorCategoryInternalServiceException ErrorCategory = "INTERNAL_SERVICE_EXCEPTION"
	ErrorCategoryCancelled                ErrorCategory = "CANCELLED"
	ErrorCategoryUserRecoverable          ErrorCategory = "USER_RECOVERABLE"
)

// Values returns all known values for ErrorCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ErrorCategory) Values() []ErrorCategory {
	return []ErrorCategory{
		"VALIDATION",
		"SERVICE_QUOTA_EXCEEDED",
		"ACCESS_DENIED",
		"RESOURCE_NOT_FOUND",
		"THROTTLING",
		"INTERNAL_SERVICE_EXCEPTION",
		"CANCELLED",
		"USER_RECOVERABLE",
	}
}

type IngestionStatus string

// Enum values for IngestionStatus
const (
	IngestionStatusPending       IngestionStatus = "PENDING"
	IngestionStatusFailed        IngestionStatus = "FAILED"
	IngestionStatusSuccess       IngestionStatus = "SUCCESS"
	IngestionStatusRunning       IngestionStatus = "RUNNING"
	IngestionStatusStopRequested IngestionStatus = "STOP_REQUESTED"
)

// Values returns all known values for IngestionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IngestionStatus) Values() []IngestionStatus {
	return []IngestionStatus{
		"PENDING",
		"FAILED",
		"SUCCESS",
		"RUNNING",
		"STOP_REQUESTED",
	}
}

type LocationType string

// Enum values for LocationType
const (
	LocationTypeIngestion LocationType = "INGESTION"
	LocationTypeSagemaker LocationType = "SAGEMAKER"
)

// Values returns all known values for LocationType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LocationType) Values() []LocationType {
	return []LocationType{
		"INGESTION",
		"SAGEMAKER",
	}
}
