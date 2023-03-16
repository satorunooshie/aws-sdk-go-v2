// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttributeType string

// Enum values for AttributeType
const (
	AttributeTypeString      AttributeType = "string"
	AttributeTypeInteger     AttributeType = "integer"
	AttributeTypeFloat       AttributeType = "float"
	AttributeTypeTimestamp   AttributeType = "timestamp"
	AttributeTypeGeolocation AttributeType = "geolocation"
)

// Values returns all known values for AttributeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AttributeType) Values() []AttributeType {
	return []AttributeType{
		"string",
		"integer",
		"float",
		"timestamp",
		"geolocation",
	}
}

type AutoMLOverrideStrategy string

// Enum values for AutoMLOverrideStrategy
const (
	AutoMLOverrideStrategyLatencyOptimized  AutoMLOverrideStrategy = "LatencyOptimized"
	AutoMLOverrideStrategyAccuracyOptimized AutoMLOverrideStrategy = "AccuracyOptimized"
)

// Values returns all known values for AutoMLOverrideStrategy. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoMLOverrideStrategy) Values() []AutoMLOverrideStrategy {
	return []AutoMLOverrideStrategy{
		"LatencyOptimized",
		"AccuracyOptimized",
	}
}

type Condition string

// Enum values for Condition
const (
	ConditionEquals      Condition = "EQUALS"
	ConditionNotEquals   Condition = "NOT_EQUALS"
	ConditionLessThan    Condition = "LESS_THAN"
	ConditionGreaterThan Condition = "GREATER_THAN"
)

// Values returns all known values for Condition. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Condition) Values() []Condition {
	return []Condition{
		"EQUALS",
		"NOT_EQUALS",
		"LESS_THAN",
		"GREATER_THAN",
	}
}

type DatasetType string

// Enum values for DatasetType
const (
	DatasetTypeTargetTimeSeries  DatasetType = "TARGET_TIME_SERIES"
	DatasetTypeRelatedTimeSeries DatasetType = "RELATED_TIME_SERIES"
	DatasetTypeItemMetadata      DatasetType = "ITEM_METADATA"
)

// Values returns all known values for DatasetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DatasetType) Values() []DatasetType {
	return []DatasetType{
		"TARGET_TIME_SERIES",
		"RELATED_TIME_SERIES",
		"ITEM_METADATA",
	}
}

type DayOfWeek string

// Enum values for DayOfWeek
const (
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
)

// Values returns all known values for DayOfWeek. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
		"SUNDAY",
	}
}

type Domain string

// Enum values for Domain
const (
	DomainRetail            Domain = "RETAIL"
	DomainCustom            Domain = "CUSTOM"
	DomainInventoryPlanning Domain = "INVENTORY_PLANNING"
	DomainEc2Capacity       Domain = "EC2_CAPACITY"
	DomainWorkForce         Domain = "WORK_FORCE"
	DomainWebTraffic        Domain = "WEB_TRAFFIC"
	DomainMetrics           Domain = "METRICS"
)

// Values returns all known values for Domain. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Domain) Values() []Domain {
	return []Domain{
		"RETAIL",
		"CUSTOM",
		"INVENTORY_PLANNING",
		"EC2_CAPACITY",
		"WORK_FORCE",
		"WEB_TRAFFIC",
		"METRICS",
	}
}

type EvaluationType string

// Enum values for EvaluationType
const (
	EvaluationTypeSummary  EvaluationType = "SUMMARY"
	EvaluationTypeComputed EvaluationType = "COMPUTED"
)

// Values returns all known values for EvaluationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EvaluationType) Values() []EvaluationType {
	return []EvaluationType{
		"SUMMARY",
		"COMPUTED",
	}
}

type FeaturizationMethodName string

// Enum values for FeaturizationMethodName
const (
	FeaturizationMethodNameFilling FeaturizationMethodName = "filling"
)

// Values returns all known values for FeaturizationMethodName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FeaturizationMethodName) Values() []FeaturizationMethodName {
	return []FeaturizationMethodName{
		"filling",
	}
}

type FilterConditionString string

// Enum values for FilterConditionString
const (
	FilterConditionStringIs    FilterConditionString = "IS"
	FilterConditionStringIsNot FilterConditionString = "IS_NOT"
)

// Values returns all known values for FilterConditionString. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterConditionString) Values() []FilterConditionString {
	return []FilterConditionString{
		"IS",
		"IS_NOT",
	}
}

type ImportMode string

// Enum values for ImportMode
const (
	ImportModeFull        ImportMode = "FULL"
	ImportModeIncremental ImportMode = "INCREMENTAL"
)

// Values returns all known values for ImportMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ImportMode) Values() []ImportMode {
	return []ImportMode{
		"FULL",
		"INCREMENTAL",
	}
}

type Month string

// Enum values for Month
const (
	MonthJanuary   Month = "JANUARY"
	MonthFebruary  Month = "FEBRUARY"
	MonthMarch     Month = "MARCH"
	MonthApril     Month = "APRIL"
	MonthMay       Month = "MAY"
	MonthJune      Month = "JUNE"
	MonthJuly      Month = "JULY"
	MonthAugust    Month = "AUGUST"
	MonthSeptember Month = "SEPTEMBER"
	MonthOctober   Month = "OCTOBER"
	MonthNovember  Month = "NOVEMBER"
	MonthDecember  Month = "DECEMBER"
)

// Values returns all known values for Month. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Month) Values() []Month {
	return []Month{
		"JANUARY",
		"FEBRUARY",
		"MARCH",
		"APRIL",
		"MAY",
		"JUNE",
		"JULY",
		"AUGUST",
		"SEPTEMBER",
		"OCTOBER",
		"NOVEMBER",
		"DECEMBER",
	}
}

type Operation string

// Enum values for Operation
const (
	OperationAdd      Operation = "ADD"
	OperationSubtract Operation = "SUBTRACT"
	OperationMultiply Operation = "MULTIPLY"
	OperationDivide   Operation = "DIVIDE"
)

// Values returns all known values for Operation. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Operation) Values() []Operation {
	return []Operation{
		"ADD",
		"SUBTRACT",
		"MULTIPLY",
		"DIVIDE",
	}
}

type OptimizationMetric string

// Enum values for OptimizationMetric
const (
	OptimizationMetricWape                        OptimizationMetric = "WAPE"
	OptimizationMetricRmse                        OptimizationMetric = "RMSE"
	OptimizationMetricAverageWeightedQuantileLoss OptimizationMetric = "AverageWeightedQuantileLoss"
	OptimizationMetricMase                        OptimizationMetric = "MASE"
	OptimizationMetricMape                        OptimizationMetric = "MAPE"
)

// Values returns all known values for OptimizationMetric. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OptimizationMetric) Values() []OptimizationMetric {
	return []OptimizationMetric{
		"WAPE",
		"RMSE",
		"AverageWeightedQuantileLoss",
		"MASE",
		"MAPE",
	}
}

type ScalingType string

// Enum values for ScalingType
const (
	ScalingTypeAuto               ScalingType = "Auto"
	ScalingTypeLinear             ScalingType = "Linear"
	ScalingTypeLogarithmic        ScalingType = "Logarithmic"
	ScalingTypeReverseLogarithmic ScalingType = "ReverseLogarithmic"
)

// Values returns all known values for ScalingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScalingType) Values() []ScalingType {
	return []ScalingType{
		"Auto",
		"Linear",
		"Logarithmic",
		"ReverseLogarithmic",
	}
}

type State string

// Enum values for State
const (
	StateActive  State = "Active"
	StateDeleted State = "Deleted"
)

// Values returns all known values for State. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (State) Values() []State {
	return []State{
		"Active",
		"Deleted",
	}
}

type TimePointGranularity string

// Enum values for TimePointGranularity
const (
	TimePointGranularityAll      TimePointGranularity = "ALL"
	TimePointGranularitySpecific TimePointGranularity = "SPECIFIC"
)

// Values returns all known values for TimePointGranularity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TimePointGranularity) Values() []TimePointGranularity {
	return []TimePointGranularity{
		"ALL",
		"SPECIFIC",
	}
}

type TimeSeriesGranularity string

// Enum values for TimeSeriesGranularity
const (
	TimeSeriesGranularityAll      TimeSeriesGranularity = "ALL"
	TimeSeriesGranularitySpecific TimeSeriesGranularity = "SPECIFIC"
)

// Values returns all known values for TimeSeriesGranularity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TimeSeriesGranularity) Values() []TimeSeriesGranularity {
	return []TimeSeriesGranularity{
		"ALL",
		"SPECIFIC",
	}
}
