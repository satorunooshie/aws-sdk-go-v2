// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AnalysisType string

// Enum values for AnalysisType
const (
	AnalysisTypeSourceCodeAnalysis AnalysisType = "SOURCE_CODE_ANALYSIS"
	AnalysisTypeDatabaseAnalysis   AnalysisType = "DATABASE_ANALYSIS"
	AnalysisTypeRuntimeAnalysis    AnalysisType = "RUNTIME_ANALYSIS"
	AnalysisTypeBinaryAnalysis     AnalysisType = "BINARY_ANALYSIS"
)

// Values returns all known values for AnalysisType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AnalysisType) Values() []AnalysisType {
	return []AnalysisType{
		"SOURCE_CODE_ANALYSIS",
		"DATABASE_ANALYSIS",
		"RUNTIME_ANALYSIS",
		"BINARY_ANALYSIS",
	}
}

type AntipatternReportStatus string

// Enum values for AntipatternReportStatus
const (
	AntipatternReportStatusFailed     AntipatternReportStatus = "FAILED"
	AntipatternReportStatusInProgress AntipatternReportStatus = "IN_PROGRESS"
	AntipatternReportStatusSuccess    AntipatternReportStatus = "SUCCESS"
)

// Values returns all known values for AntipatternReportStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AntipatternReportStatus) Values() []AntipatternReportStatus {
	return []AntipatternReportStatus{
		"FAILED",
		"IN_PROGRESS",
		"SUCCESS",
	}
}

type ApplicationComponentCriteria string

// Enum values for ApplicationComponentCriteria
const (
	ApplicationComponentCriteriaNotDefined     ApplicationComponentCriteria = "NOT_DEFINED"
	ApplicationComponentCriteriaAppName        ApplicationComponentCriteria = "APP_NAME"
	ApplicationComponentCriteriaServerId       ApplicationComponentCriteria = "SERVER_ID"
	ApplicationComponentCriteriaAppType        ApplicationComponentCriteria = "APP_TYPE"
	ApplicationComponentCriteriaStrategy       ApplicationComponentCriteria = "STRATEGY"
	ApplicationComponentCriteriaDestination    ApplicationComponentCriteria = "DESTINATION"
	ApplicationComponentCriteriaAnalysisStatus ApplicationComponentCriteria = "ANALYSIS_STATUS"
	ApplicationComponentCriteriaErrorCategory  ApplicationComponentCriteria = "ERROR_CATEGORY"
)

// Values returns all known values for ApplicationComponentCriteria. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationComponentCriteria) Values() []ApplicationComponentCriteria {
	return []ApplicationComponentCriteria{
		"NOT_DEFINED",
		"APP_NAME",
		"SERVER_ID",
		"APP_TYPE",
		"STRATEGY",
		"DESTINATION",
		"ANALYSIS_STATUS",
		"ERROR_CATEGORY",
	}
}

type ApplicationMode string

// Enum values for ApplicationMode
const (
	ApplicationModeAll     ApplicationMode = "ALL"
	ApplicationModeKnown   ApplicationMode = "KNOWN"
	ApplicationModeUnknown ApplicationMode = "UNKNOWN"
)

// Values returns all known values for ApplicationMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationMode) Values() []ApplicationMode {
	return []ApplicationMode{
		"ALL",
		"KNOWN",
		"UNKNOWN",
	}
}

type AppType string

// Enum values for AppType
const (
	AppTypeDotNetFramework  AppType = "DotNetFramework"
	AppTypeJava             AppType = "Java"
	AppTypeSqlServer        AppType = "SQLServer"
	AppTypeIis              AppType = "IIS"
	AppTypeOracle           AppType = "Oracle"
	AppTypeOther            AppType = "Other"
	AppTypeTomcat           AppType = "Tomcat"
	AppTypeJboss            AppType = "JBoss"
	AppTypeSpring           AppType = "Spring"
	AppTypeMongodb          AppType = "Mongo DB"
	AppTypeDb2              AppType = "DB2"
	AppTypeMariadb          AppType = "Maria DB"
	AppTypeMysql            AppType = "MySQL"
	AppTypeSybase           AppType = "Sybase"
	AppTypePostgresqlserver AppType = "PostgreSQLServer"
	AppTypeCassandra        AppType = "Cassandra"
	AppTypeWebsphere        AppType = "IBM WebSphere"
	AppTypeWeblogic         AppType = "Oracle WebLogic"
	AppTypeVisualbasic      AppType = "Visual Basic"
	AppTypeUnknown          AppType = "Unknown"
	AppTypeDotnetcore       AppType = "DotnetCore"
	AppTypeDotnet           AppType = "Dotnet"
)

// Values returns all known values for AppType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AppType) Values() []AppType {
	return []AppType{
		"DotNetFramework",
		"Java",
		"SQLServer",
		"IIS",
		"Oracle",
		"Other",
		"Tomcat",
		"JBoss",
		"Spring",
		"Mongo DB",
		"DB2",
		"Maria DB",
		"MySQL",
		"Sybase",
		"PostgreSQLServer",
		"Cassandra",
		"IBM WebSphere",
		"Oracle WebLogic",
		"Visual Basic",
		"Unknown",
		"DotnetCore",
		"Dotnet",
	}
}

type AppUnitErrorCategory string

// Enum values for AppUnitErrorCategory
const (
	AppUnitErrorCategoryCredentialError   AppUnitErrorCategory = "CREDENTIAL_ERROR"
	AppUnitErrorCategoryConnectivityError AppUnitErrorCategory = "CONNECTIVITY_ERROR"
	AppUnitErrorCategoryPermissionError   AppUnitErrorCategory = "PERMISSION_ERROR"
	AppUnitErrorCategoryUnsupportedError  AppUnitErrorCategory = "UNSUPPORTED_ERROR"
	AppUnitErrorCategoryOtherError        AppUnitErrorCategory = "OTHER_ERROR"
)

// Values returns all known values for AppUnitErrorCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AppUnitErrorCategory) Values() []AppUnitErrorCategory {
	return []AppUnitErrorCategory{
		"CREDENTIAL_ERROR",
		"CONNECTIVITY_ERROR",
		"PERMISSION_ERROR",
		"UNSUPPORTED_ERROR",
		"OTHER_ERROR",
	}
}

type AssessmentStatus string

// Enum values for AssessmentStatus
const (
	AssessmentStatusInProgress AssessmentStatus = "IN_PROGRESS"
	AssessmentStatusComplete   AssessmentStatus = "COMPLETE"
	AssessmentStatusFailed     AssessmentStatus = "FAILED"
	AssessmentStatusStopped    AssessmentStatus = "STOPPED"
)

// Values returns all known values for AssessmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssessmentStatus) Values() []AssessmentStatus {
	return []AssessmentStatus{
		"IN_PROGRESS",
		"COMPLETE",
		"FAILED",
		"STOPPED",
	}
}

type AuthType string

// Enum values for AuthType
const (
	AuthTypeNtlm AuthType = "NTLM"
	AuthTypeSsh  AuthType = "SSH"
	AuthTypeCert AuthType = "CERT"
)

// Values returns all known values for AuthType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AuthType) Values() []AuthType {
	return []AuthType{
		"NTLM",
		"SSH",
		"CERT",
	}
}

type AwsManagedTargetDestination string

// Enum values for AwsManagedTargetDestination
const (
	AwsManagedTargetDestinationNoneSpecified       AwsManagedTargetDestination = "None specified"
	AwsManagedTargetDestinationAwsElasticBeanstalk AwsManagedTargetDestination = "AWS Elastic BeanStalk"
	AwsManagedTargetDestinationAwsFargate          AwsManagedTargetDestination = "AWS Fargate"
)

// Values returns all known values for AwsManagedTargetDestination. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AwsManagedTargetDestination) Values() []AwsManagedTargetDestination {
	return []AwsManagedTargetDestination{
		"None specified",
		"AWS Elastic BeanStalk",
		"AWS Fargate",
	}
}

type BinaryAnalyzerName string

// Enum values for BinaryAnalyzerName
const (
	BinaryAnalyzerNameDllAnalyzer      BinaryAnalyzerName = "DLL_ANALYZER"
	BinaryAnalyzerNameBytecodeAnalyzer BinaryAnalyzerName = "BYTECODE_ANALYZER"
)

// Values returns all known values for BinaryAnalyzerName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BinaryAnalyzerName) Values() []BinaryAnalyzerName {
	return []BinaryAnalyzerName{
		"DLL_ANALYZER",
		"BYTECODE_ANALYZER",
	}
}

type CollectorHealth string

// Enum values for CollectorHealth
const (
	CollectorHealthCollectorHealthy   CollectorHealth = "COLLECTOR_HEALTHY"
	CollectorHealthCollectorUnhealthy CollectorHealth = "COLLECTOR_UNHEALTHY"
)

// Values returns all known values for CollectorHealth. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CollectorHealth) Values() []CollectorHealth {
	return []CollectorHealth{
		"COLLECTOR_HEALTHY",
		"COLLECTOR_UNHEALTHY",
	}
}

type Condition string

// Enum values for Condition
const (
	ConditionEquals      Condition = "EQUALS"
	ConditionNotEquals   Condition = "NOT_EQUALS"
	ConditionContains    Condition = "CONTAINS"
	ConditionNotContains Condition = "NOT_CONTAINS"
)

// Values returns all known values for Condition. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Condition) Values() []Condition {
	return []Condition{
		"EQUALS",
		"NOT_EQUALS",
		"CONTAINS",
		"NOT_CONTAINS",
	}
}

type DatabaseManagementPreference string

// Enum values for DatabaseManagementPreference
const (
	DatabaseManagementPreferenceAwsManaged   DatabaseManagementPreference = "AWS-managed"
	DatabaseManagementPreferenceSelfManage   DatabaseManagementPreference = "Self-manage"
	DatabaseManagementPreferenceNoPreference DatabaseManagementPreference = "No preference"
)

// Values returns all known values for DatabaseManagementPreference. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DatabaseManagementPreference) Values() []DatabaseManagementPreference {
	return []DatabaseManagementPreference{
		"AWS-managed",
		"Self-manage",
		"No preference",
	}
}

type DataSourceType string

// Enum values for DataSourceType
const (
	DataSourceTypeAds    DataSourceType = "ApplicationDiscoveryService"
	DataSourceTypeMpa    DataSourceType = "MPA"
	DataSourceTypeImport DataSourceType = "Import"
)

// Values returns all known values for DataSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceType) Values() []DataSourceType {
	return []DataSourceType{
		"ApplicationDiscoveryService",
		"MPA",
		"Import",
	}
}

type GroupName string

// Enum values for GroupName
const (
	GroupNameExternalId         GroupName = "ExternalId"
	GroupNameExternalSourceType GroupName = "ExternalSourceType"
)

// Values returns all known values for GroupName. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (GroupName) Values() []GroupName {
	return []GroupName{
		"ExternalId",
		"ExternalSourceType",
	}
}

type HeterogeneousTargetDatabaseEngine string

// Enum values for HeterogeneousTargetDatabaseEngine
const (
	HeterogeneousTargetDatabaseEngineNoneSpecified      HeterogeneousTargetDatabaseEngine = "None specified"
	HeterogeneousTargetDatabaseEngineAmazonAurora       HeterogeneousTargetDatabaseEngine = "Amazon Aurora"
	HeterogeneousTargetDatabaseEngineAwsPostgresql      HeterogeneousTargetDatabaseEngine = "AWS PostgreSQL"
	HeterogeneousTargetDatabaseEngineMysql              HeterogeneousTargetDatabaseEngine = "MySQL"
	HeterogeneousTargetDatabaseEngineMicrosoftSqlServer HeterogeneousTargetDatabaseEngine = "Microsoft SQL Server"
	HeterogeneousTargetDatabaseEngineOracleDatabase     HeterogeneousTargetDatabaseEngine = "Oracle Database"
	HeterogeneousTargetDatabaseEngineMariaDb            HeterogeneousTargetDatabaseEngine = "MariaDB"
	HeterogeneousTargetDatabaseEngineSap                HeterogeneousTargetDatabaseEngine = "SAP"
	HeterogeneousTargetDatabaseEngineDb2Luw             HeterogeneousTargetDatabaseEngine = "Db2 LUW"
	HeterogeneousTargetDatabaseEngineMongoDb            HeterogeneousTargetDatabaseEngine = "MongoDB"
)

// Values returns all known values for HeterogeneousTargetDatabaseEngine. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (HeterogeneousTargetDatabaseEngine) Values() []HeterogeneousTargetDatabaseEngine {
	return []HeterogeneousTargetDatabaseEngine{
		"None specified",
		"Amazon Aurora",
		"AWS PostgreSQL",
		"MySQL",
		"Microsoft SQL Server",
		"Oracle Database",
		"MariaDB",
		"SAP",
		"Db2 LUW",
		"MongoDB",
	}
}

type HomogeneousTargetDatabaseEngine string

// Enum values for HomogeneousTargetDatabaseEngine
const (
	HomogeneousTargetDatabaseEngineNoneSpecified HomogeneousTargetDatabaseEngine = "None specified"
)

// Values returns all known values for HomogeneousTargetDatabaseEngine. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (HomogeneousTargetDatabaseEngine) Values() []HomogeneousTargetDatabaseEngine {
	return []HomogeneousTargetDatabaseEngine{
		"None specified",
	}
}

type ImportFileTaskStatus string

// Enum values for ImportFileTaskStatus
const (
	ImportFileTaskStatusImportInProgress     ImportFileTaskStatus = "ImportInProgress"
	ImportFileTaskStatusImportFailed         ImportFileTaskStatus = "ImportFailed"
	ImportFileTaskStatusImportPartialSuccess ImportFileTaskStatus = "ImportPartialSuccess"
	ImportFileTaskStatusImportSuccess        ImportFileTaskStatus = "ImportSuccess"
	ImportFileTaskStatusDeleteInProgress     ImportFileTaskStatus = "DeleteInProgress"
	ImportFileTaskStatusDeleteFailed         ImportFileTaskStatus = "DeleteFailed"
	ImportFileTaskStatusDeletePartialSuccess ImportFileTaskStatus = "DeletePartialSuccess"
	ImportFileTaskStatusDeleteSuccess        ImportFileTaskStatus = "DeleteSuccess"
)

// Values returns all known values for ImportFileTaskStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportFileTaskStatus) Values() []ImportFileTaskStatus {
	return []ImportFileTaskStatus{
		"ImportInProgress",
		"ImportFailed",
		"ImportPartialSuccess",
		"ImportSuccess",
		"DeleteInProgress",
		"DeleteFailed",
		"DeletePartialSuccess",
		"DeleteSuccess",
	}
}

type InclusionStatus string

// Enum values for InclusionStatus
const (
	InclusionStatusExcludeFromRecommendation InclusionStatus = "excludeFromAssessment"
	InclusionStatusIncludeInRecommendation   InclusionStatus = "includeInAssessment"
)

// Values returns all known values for InclusionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InclusionStatus) Values() []InclusionStatus {
	return []InclusionStatus{
		"excludeFromAssessment",
		"includeInAssessment",
	}
}

type NoPreferenceTargetDestination string

// Enum values for NoPreferenceTargetDestination
const (
	NoPreferenceTargetDestinationNoneSpecified                  NoPreferenceTargetDestination = "None specified"
	NoPreferenceTargetDestinationAwsElasticBeanstalk            NoPreferenceTargetDestination = "AWS Elastic BeanStalk"
	NoPreferenceTargetDestinationAwsFargate                     NoPreferenceTargetDestination = "AWS Fargate"
	NoPreferenceTargetDestinationAmazonElasticCloudCompute      NoPreferenceTargetDestination = "Amazon Elastic Cloud Compute (EC2)"
	NoPreferenceTargetDestinationAmazonElasticContainerService  NoPreferenceTargetDestination = "Amazon Elastic Container Service (ECS)"
	NoPreferenceTargetDestinationAmazonElasticKubernetesService NoPreferenceTargetDestination = "Amazon Elastic Kubernetes Service (EKS)"
)

// Values returns all known values for NoPreferenceTargetDestination. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (NoPreferenceTargetDestination) Values() []NoPreferenceTargetDestination {
	return []NoPreferenceTargetDestination{
		"None specified",
		"AWS Elastic BeanStalk",
		"AWS Fargate",
		"Amazon Elastic Cloud Compute (EC2)",
		"Amazon Elastic Container Service (ECS)",
		"Amazon Elastic Kubernetes Service (EKS)",
	}
}

type OSType string

// Enum values for OSType
const (
	OSTypeLinux   OSType = "LINUX"
	OSTypeWindows OSType = "WINDOWS"
)

// Values returns all known values for OSType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OSType) Values() []OSType {
	return []OSType{
		"LINUX",
		"WINDOWS",
	}
}

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatExcel OutputFormat = "Excel"
	OutputFormatJson  OutputFormat = "Json"
)

// Values returns all known values for OutputFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		"Excel",
		"Json",
	}
}

type PipelineType string

// Enum values for PipelineType
const (
	PipelineTypeAzureDevops PipelineType = "AZURE_DEVOPS"
)

// Values returns all known values for PipelineType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PipelineType) Values() []PipelineType {
	return []PipelineType{
		"AZURE_DEVOPS",
	}
}

type RecommendationReportStatus string

// Enum values for RecommendationReportStatus
const (
	RecommendationReportStatusFailed     RecommendationReportStatus = "FAILED"
	RecommendationReportStatusInProgress RecommendationReportStatus = "IN_PROGRESS"
	RecommendationReportStatusSuccess    RecommendationReportStatus = "SUCCESS"
)

// Values returns all known values for RecommendationReportStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationReportStatus) Values() []RecommendationReportStatus {
	return []RecommendationReportStatus{
		"FAILED",
		"IN_PROGRESS",
		"SUCCESS",
	}
}

type ResourceSubType string

// Enum values for ResourceSubType
const (
	ResourceSubTypeDatabase        ResourceSubType = "Database"
	ResourceSubTypeProcess         ResourceSubType = "Process"
	ResourceSubTypeDatabaseProcess ResourceSubType = "DatabaseProcess"
)

// Values returns all known values for ResourceSubType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceSubType) Values() []ResourceSubType {
	return []ResourceSubType{
		"Database",
		"Process",
		"DatabaseProcess",
	}
}

type RuntimeAnalysisStatus string

// Enum values for RuntimeAnalysisStatus
const (
	RuntimeAnalysisStatusAnalysisToBeScheduled RuntimeAnalysisStatus = "ANALYSIS_TO_BE_SCHEDULED"
	RuntimeAnalysisStatusAnalysisStarted       RuntimeAnalysisStatus = "ANALYSIS_STARTED"
	RuntimeAnalysisStatusAnalysisSuccess       RuntimeAnalysisStatus = "ANALYSIS_SUCCESS"
	RuntimeAnalysisStatusAnalysisFailed        RuntimeAnalysisStatus = "ANALYSIS_FAILED"
)

// Values returns all known values for RuntimeAnalysisStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RuntimeAnalysisStatus) Values() []RuntimeAnalysisStatus {
	return []RuntimeAnalysisStatus{
		"ANALYSIS_TO_BE_SCHEDULED",
		"ANALYSIS_STARTED",
		"ANALYSIS_SUCCESS",
		"ANALYSIS_FAILED",
	}
}

type RunTimeAnalyzerName string

// Enum values for RunTimeAnalyzerName
const (
	RunTimeAnalyzerNameA2cAnalyzer      RunTimeAnalyzerName = "A2C_ANALYZER"
	RunTimeAnalyzerNameRehostAnalyzer   RunTimeAnalyzerName = "REHOST_ANALYZER"
	RunTimeAnalyzerNameEmpPaAnalyzer    RunTimeAnalyzerName = "EMP_PA_ANALYZER"
	RunTimeAnalyzerNameDatabaseAnalyzer RunTimeAnalyzerName = "DATABASE_ANALYZER"
	RunTimeAnalyzerNameSctAnalyzer      RunTimeAnalyzerName = "SCT_ANALYZER"
)

// Values returns all known values for RunTimeAnalyzerName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RunTimeAnalyzerName) Values() []RunTimeAnalyzerName {
	return []RunTimeAnalyzerName{
		"A2C_ANALYZER",
		"REHOST_ANALYZER",
		"EMP_PA_ANALYZER",
		"DATABASE_ANALYZER",
		"SCT_ANALYZER",
	}
}

type RunTimeAssessmentStatus string

// Enum values for RunTimeAssessmentStatus
const (
	RunTimeAssessmentStatusDcToBeScheduled  RunTimeAssessmentStatus = "dataCollectionTaskToBeScheduled"
	RunTimeAssessmentStatusDcReqSent        RunTimeAssessmentStatus = "dataCollectionTaskScheduled"
	RunTimeAssessmentStatusDcStarted        RunTimeAssessmentStatus = "dataCollectionTaskStarted"
	RunTimeAssessmentStatusDcStopped        RunTimeAssessmentStatus = "dataCollectionTaskStopped"
	RunTimeAssessmentStatusDcSuccess        RunTimeAssessmentStatus = "dataCollectionTaskSuccess"
	RunTimeAssessmentStatusDcFailed         RunTimeAssessmentStatus = "dataCollectionTaskFailed"
	RunTimeAssessmentStatusDcPartialSuccess RunTimeAssessmentStatus = "dataCollectionTaskPartialSuccess"
)

// Values returns all known values for RunTimeAssessmentStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RunTimeAssessmentStatus) Values() []RunTimeAssessmentStatus {
	return []RunTimeAssessmentStatus{
		"dataCollectionTaskToBeScheduled",
		"dataCollectionTaskScheduled",
		"dataCollectionTaskStarted",
		"dataCollectionTaskStopped",
		"dataCollectionTaskSuccess",
		"dataCollectionTaskFailed",
		"dataCollectionTaskPartialSuccess",
	}
}

type SelfManageTargetDestination string

// Enum values for SelfManageTargetDestination
const (
	SelfManageTargetDestinationNoneSpecified                  SelfManageTargetDestination = "None specified"
	SelfManageTargetDestinationAmazonElasticCloudCompute      SelfManageTargetDestination = "Amazon Elastic Cloud Compute (EC2)"
	SelfManageTargetDestinationAmazonElasticContainerService  SelfManageTargetDestination = "Amazon Elastic Container Service (ECS)"
	SelfManageTargetDestinationAmazonElasticKubernetesService SelfManageTargetDestination = "Amazon Elastic Kubernetes Service (EKS)"
)

// Values returns all known values for SelfManageTargetDestination. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SelfManageTargetDestination) Values() []SelfManageTargetDestination {
	return []SelfManageTargetDestination{
		"None specified",
		"Amazon Elastic Cloud Compute (EC2)",
		"Amazon Elastic Container Service (ECS)",
		"Amazon Elastic Kubernetes Service (EKS)",
	}
}

type ServerCriteria string

// Enum values for ServerCriteria
const (
	ServerCriteriaNotDefined     ServerCriteria = "NOT_DEFINED"
	ServerCriteriaOsName         ServerCriteria = "OS_NAME"
	ServerCriteriaStrategy       ServerCriteria = "STRATEGY"
	ServerCriteriaDestination    ServerCriteria = "DESTINATION"
	ServerCriteriaServerId       ServerCriteria = "SERVER_ID"
	ServerCriteriaAnalysisStatus ServerCriteria = "ANALYSIS_STATUS"
	ServerCriteriaErrorCategory  ServerCriteria = "ERROR_CATEGORY"
)

// Values returns all known values for ServerCriteria. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServerCriteria) Values() []ServerCriteria {
	return []ServerCriteria{
		"NOT_DEFINED",
		"OS_NAME",
		"STRATEGY",
		"DESTINATION",
		"SERVER_ID",
		"ANALYSIS_STATUS",
		"ERROR_CATEGORY",
	}
}

type ServerErrorCategory string

// Enum values for ServerErrorCategory
const (
	ServerErrorCategoryConnectivityError ServerErrorCategory = "CONNECTIVITY_ERROR"
	ServerErrorCategoryCredentialError   ServerErrorCategory = "CREDENTIAL_ERROR"
	ServerErrorCategoryPermissionError   ServerErrorCategory = "PERMISSION_ERROR"
	ServerErrorCategoryArchitectureError ServerErrorCategory = "ARCHITECTURE_ERROR"
	ServerErrorCategoryOtherError        ServerErrorCategory = "OTHER_ERROR"
)

// Values returns all known values for ServerErrorCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServerErrorCategory) Values() []ServerErrorCategory {
	return []ServerErrorCategory{
		"CONNECTIVITY_ERROR",
		"CREDENTIAL_ERROR",
		"PERMISSION_ERROR",
		"ARCHITECTURE_ERROR",
		"OTHER_ERROR",
	}
}

type ServerOsType string

// Enum values for ServerOsType
const (
	ServerOsTypeWindowsServer             ServerOsType = "WindowsServer"
	ServerOsTypeAmazonLinux               ServerOsType = "AmazonLinux"
	ServerOsTypeEndOfSupportWindowsServer ServerOsType = "EndOfSupportWindowsServer"
	ServerOsTypeRedhat                    ServerOsType = "Redhat"
	ServerOsTypeOther                     ServerOsType = "Other"
)

// Values returns all known values for ServerOsType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ServerOsType) Values() []ServerOsType {
	return []ServerOsType{
		"WindowsServer",
		"AmazonLinux",
		"EndOfSupportWindowsServer",
		"Redhat",
		"Other",
	}
}

type Severity string

// Enum values for Severity
const (
	SeverityHigh   Severity = "HIGH"
	SeverityMedium Severity = "MEDIUM"
	SeverityLow    Severity = "LOW"
)

// Values returns all known values for Severity. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Severity) Values() []Severity {
	return []Severity{
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASC",
		"DESC",
	}
}

type SourceCodeAnalyzerName string

// Enum values for SourceCodeAnalyzerName
const (
	SourceCodeAnalyzerNameCsharpAnalyzer   SourceCodeAnalyzerName = "CSHARP_ANALYZER"
	SourceCodeAnalyzerNameJavaAnalyzer     SourceCodeAnalyzerName = "JAVA_ANALYZER"
	SourceCodeAnalyzerNameBytecodeAnalyzer SourceCodeAnalyzerName = "BYTECODE_ANALYZER"
	SourceCodeAnalyzerNamePortingAssistant SourceCodeAnalyzerName = "PORTING_ASSISTANT"
)

// Values returns all known values for SourceCodeAnalyzerName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SourceCodeAnalyzerName) Values() []SourceCodeAnalyzerName {
	return []SourceCodeAnalyzerName{
		"CSHARP_ANALYZER",
		"JAVA_ANALYZER",
		"BYTECODE_ANALYZER",
		"PORTING_ASSISTANT",
	}
}

type SrcCodeOrDbAnalysisStatus string

// Enum values for SrcCodeOrDbAnalysisStatus
const (
	SrcCodeOrDbAnalysisStatusAnalysisToBeScheduled  SrcCodeOrDbAnalysisStatus = "ANALYSIS_TO_BE_SCHEDULED"
	SrcCodeOrDbAnalysisStatusAnalysisStarted        SrcCodeOrDbAnalysisStatus = "ANALYSIS_STARTED"
	SrcCodeOrDbAnalysisStatusAnalysisSuccess        SrcCodeOrDbAnalysisStatus = "ANALYSIS_SUCCESS"
	SrcCodeOrDbAnalysisStatusAnalysisFailed         SrcCodeOrDbAnalysisStatus = "ANALYSIS_FAILED"
	SrcCodeOrDbAnalysisStatusAnalysisPartialSuccess SrcCodeOrDbAnalysisStatus = "ANALYSIS_PARTIAL_SUCCESS"
	SrcCodeOrDbAnalysisStatusUnconfigured           SrcCodeOrDbAnalysisStatus = "UNCONFIGURED"
	SrcCodeOrDbAnalysisStatusConfigured             SrcCodeOrDbAnalysisStatus = "CONFIGURED"
)

// Values returns all known values for SrcCodeOrDbAnalysisStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SrcCodeOrDbAnalysisStatus) Values() []SrcCodeOrDbAnalysisStatus {
	return []SrcCodeOrDbAnalysisStatus{
		"ANALYSIS_TO_BE_SCHEDULED",
		"ANALYSIS_STARTED",
		"ANALYSIS_SUCCESS",
		"ANALYSIS_FAILED",
		"ANALYSIS_PARTIAL_SUCCESS",
		"UNCONFIGURED",
		"CONFIGURED",
	}
}

type Strategy string

// Enum values for Strategy
const (
	StrategyRehost     Strategy = "Rehost"
	StrategyRetirement Strategy = "Retirement"
	StrategyRefactor   Strategy = "Refactor"
	StrategyReplatform Strategy = "Replatform"
	StrategyRetain     Strategy = "Retain"
	StrategyRelocate   Strategy = "Relocate"
	StrategyRepurchase Strategy = "Repurchase"
)

// Values returns all known values for Strategy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Strategy) Values() []Strategy {
	return []Strategy{
		"Rehost",
		"Retirement",
		"Refactor",
		"Replatform",
		"Retain",
		"Relocate",
		"Repurchase",
	}
}

type StrategyRecommendation string

// Enum values for StrategyRecommendation
const (
	StrategyRecommendationRecommended    StrategyRecommendation = "recommended"
	StrategyRecommendationViableOption   StrategyRecommendation = "viableOption"
	StrategyRecommendationNotRecommended StrategyRecommendation = "notRecommended"
	StrategyRecommendationPotential      StrategyRecommendation = "potential"
)

// Values returns all known values for StrategyRecommendation. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StrategyRecommendation) Values() []StrategyRecommendation {
	return []StrategyRecommendation{
		"recommended",
		"viableOption",
		"notRecommended",
		"potential",
	}
}

type TargetDatabaseEngine string

// Enum values for TargetDatabaseEngine
const (
	TargetDatabaseEngineNoneSpecified      TargetDatabaseEngine = "None specified"
	TargetDatabaseEngineAmazonAurora       TargetDatabaseEngine = "Amazon Aurora"
	TargetDatabaseEngineAwsPostgresql      TargetDatabaseEngine = "AWS PostgreSQL"
	TargetDatabaseEngineMysql              TargetDatabaseEngine = "MySQL"
	TargetDatabaseEngineMicrosoftSqlServer TargetDatabaseEngine = "Microsoft SQL Server"
	TargetDatabaseEngineOracleDatabase     TargetDatabaseEngine = "Oracle Database"
	TargetDatabaseEngineMariaDb            TargetDatabaseEngine = "MariaDB"
	TargetDatabaseEngineSap                TargetDatabaseEngine = "SAP"
	TargetDatabaseEngineDb2Luw             TargetDatabaseEngine = "Db2 LUW"
	TargetDatabaseEngineMongoDb            TargetDatabaseEngine = "MongoDB"
)

// Values returns all known values for TargetDatabaseEngine. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TargetDatabaseEngine) Values() []TargetDatabaseEngine {
	return []TargetDatabaseEngine{
		"None specified",
		"Amazon Aurora",
		"AWS PostgreSQL",
		"MySQL",
		"Microsoft SQL Server",
		"Oracle Database",
		"MariaDB",
		"SAP",
		"Db2 LUW",
		"MongoDB",
	}
}

type TargetDestination string

// Enum values for TargetDestination
const (
	TargetDestinationNoneSpecified                  TargetDestination = "None specified"
	TargetDestinationAwsElasticBeanstalk            TargetDestination = "AWS Elastic BeanStalk"
	TargetDestinationAwsFargate                     TargetDestination = "AWS Fargate"
	TargetDestinationAmazonElasticCloudCompute      TargetDestination = "Amazon Elastic Cloud Compute (EC2)"
	TargetDestinationAmazonElasticContainerService  TargetDestination = "Amazon Elastic Container Service (ECS)"
	TargetDestinationAmazonElasticKubernetesService TargetDestination = "Amazon Elastic Kubernetes Service (EKS)"
	TargetDestinationAuroraMysql                    TargetDestination = "Aurora MySQL"
	TargetDestinationAuroraPostgresql               TargetDestination = "Aurora PostgreSQL"
	TargetDestinationAmazonRdsMysql                 TargetDestination = "Amazon Relational Database Service on MySQL"
	TargetDestinationAmazonRdsPostgresql            TargetDestination = "Amazon Relational Database Service on PostgreSQL"
	TargetDestinationAmazonDocumentdb               TargetDestination = "Amazon DocumentDB"
	TargetDestinationAmazonDynamodb                 TargetDestination = "Amazon DynamoDB"
	TargetDestinationAmazonRds                      TargetDestination = "Amazon Relational Database Service"
	TargetDestinationBabelfishAuroraPostgresql      TargetDestination = "Babelfish for Aurora PostgreSQL"
)

// Values returns all known values for TargetDestination. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TargetDestination) Values() []TargetDestination {
	return []TargetDestination{
		"None specified",
		"AWS Elastic BeanStalk",
		"AWS Fargate",
		"Amazon Elastic Cloud Compute (EC2)",
		"Amazon Elastic Container Service (ECS)",
		"Amazon Elastic Kubernetes Service (EKS)",
		"Aurora MySQL",
		"Aurora PostgreSQL",
		"Amazon Relational Database Service on MySQL",
		"Amazon Relational Database Service on PostgreSQL",
		"Amazon DocumentDB",
		"Amazon DynamoDB",
		"Amazon Relational Database Service",
		"Babelfish for Aurora PostgreSQL",
	}
}

type TransformationToolName string

// Enum values for TransformationToolName
const (
	TransformationToolNameApp2container                 TransformationToolName = "App2Container"
	TransformationToolNamePortingAssistant              TransformationToolName = "Porting Assistant For .NET"
	TransformationToolNameEmp                           TransformationToolName = "End of Support Migration"
	TransformationToolNameWwama                         TransformationToolName = "Windows Web Application Migration Assistant"
	TransformationToolNameMgn                           TransformationToolName = "Application Migration Service"
	TransformationToolNameStrategyRecommendationSupport TransformationToolName = "Strategy Recommendation Support"
	TransformationToolNameInPlaceOsUpgrade              TransformationToolName = "In Place Operating System Upgrade"
	TransformationToolNameSct                           TransformationToolName = "Schema Conversion Tool"
	TransformationToolNameDms                           TransformationToolName = "Database Migration Service"
	TransformationToolNameNativeSql                     TransformationToolName = "Native SQL Server Backup/Restore"
)

// Values returns all known values for TransformationToolName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransformationToolName) Values() []TransformationToolName {
	return []TransformationToolName{
		"App2Container",
		"Porting Assistant For .NET",
		"End of Support Migration",
		"Windows Web Application Migration Assistant",
		"Application Migration Service",
		"Strategy Recommendation Support",
		"In Place Operating System Upgrade",
		"Schema Conversion Tool",
		"Database Migration Service",
		"Native SQL Server Backup/Restore",
	}
}

type VersionControl string

// Enum values for VersionControl
const (
	VersionControlGithub           VersionControl = "GITHUB"
	VersionControlGithubEnterprise VersionControl = "GITHUB_ENTERPRISE"
	VersionControlAzureDevopsGit   VersionControl = "AZURE_DEVOPS_GIT"
)

// Values returns all known values for VersionControl. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VersionControl) Values() []VersionControl {
	return []VersionControl{
		"GITHUB",
		"GITHUB_ENTERPRISE",
		"AZURE_DEVOPS_GIT",
	}
}

type VersionControlType string

// Enum values for VersionControlType
const (
	VersionControlTypeGithub           VersionControlType = "GITHUB"
	VersionControlTypeGithubEnterprise VersionControlType = "GITHUB_ENTERPRISE"
	VersionControlTypeAzureDevopsGit   VersionControlType = "AZURE_DEVOPS_GIT"
)

// Values returns all known values for VersionControlType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VersionControlType) Values() []VersionControlType {
	return []VersionControlType{
		"GITHUB",
		"GITHUB_ENTERPRISE",
		"AZURE_DEVOPS_GIT",
	}
}
