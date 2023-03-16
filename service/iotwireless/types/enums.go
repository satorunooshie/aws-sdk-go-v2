// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplicationConfigType string

// Enum values for ApplicationConfigType
const (
	ApplicationConfigTypeSemtechGeoLocation ApplicationConfigType = "SemtechGeolocation"
)

// Values returns all known values for ApplicationConfigType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationConfigType) Values() []ApplicationConfigType {
	return []ApplicationConfigType{
		"SemtechGeolocation",
	}
}

type BatteryLevel string

// Enum values for BatteryLevel
const (
	BatteryLevelNormal   BatteryLevel = "normal"
	BatteryLevelLow      BatteryLevel = "low"
	BatteryLevelCritical BatteryLevel = "critical"
)

// Values returns all known values for BatteryLevel. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BatteryLevel) Values() []BatteryLevel {
	return []BatteryLevel{
		"normal",
		"low",
		"critical",
	}
}

type ConnectionStatus string

// Enum values for ConnectionStatus
const (
	ConnectionStatusConnected    ConnectionStatus = "Connected"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
)

// Values returns all known values for ConnectionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionStatus) Values() []ConnectionStatus {
	return []ConnectionStatus{
		"Connected",
		"Disconnected",
	}
}

type DeviceState string

// Enum values for DeviceState
const (
	DeviceStateProvisioned           DeviceState = "Provisioned"
	DeviceStateRegisterednotseen     DeviceState = "RegisteredNotSeen"
	DeviceStateRegisteredreachable   DeviceState = "RegisteredReachable"
	DeviceStateRegisteredunreachable DeviceState = "RegisteredUnreachable"
)

// Values returns all known values for DeviceState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DeviceState) Values() []DeviceState {
	return []DeviceState{
		"Provisioned",
		"RegisteredNotSeen",
		"RegisteredReachable",
		"RegisteredUnreachable",
	}
}

type DlClass string

// Enum values for DlClass
const (
	DlClassClassB DlClass = "ClassB"
	DlClassClassC DlClass = "ClassC"
)

// Values returns all known values for DlClass. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (DlClass) Values() []DlClass {
	return []DlClass{
		"ClassB",
		"ClassC",
	}
}

type DownlinkMode string

// Enum values for DownlinkMode
const (
	DownlinkModeSequential         DownlinkMode = "SEQUENTIAL"
	DownlinkModeConcurrent         DownlinkMode = "CONCURRENT"
	DownlinkModeUsingUplinkGateway DownlinkMode = "USING_UPLINK_GATEWAY"
)

// Values returns all known values for DownlinkMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DownlinkMode) Values() []DownlinkMode {
	return []DownlinkMode{
		"SEQUENTIAL",
		"CONCURRENT",
		"USING_UPLINK_GATEWAY",
	}
}

type Event string

// Enum values for Event
const (
	EventDiscovered  Event = "discovered"
	EventLost        Event = "lost"
	EventAck         Event = "ack"
	EventNack        Event = "nack"
	EventPassthrough Event = "passthrough"
)

// Values returns all known values for Event. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Event) Values() []Event {
	return []Event{
		"discovered",
		"lost",
		"ack",
		"nack",
		"passthrough",
	}
}

type EventNotificationPartnerType string

// Enum values for EventNotificationPartnerType
const (
	EventNotificationPartnerTypeSidewalk EventNotificationPartnerType = "Sidewalk"
)

// Values returns all known values for EventNotificationPartnerType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventNotificationPartnerType) Values() []EventNotificationPartnerType {
	return []EventNotificationPartnerType{
		"Sidewalk",
	}
}

type EventNotificationResourceType string

// Enum values for EventNotificationResourceType
const (
	EventNotificationResourceTypeSidewalkAccount EventNotificationResourceType = "SidewalkAccount"
	EventNotificationResourceTypeWirelessDevice  EventNotificationResourceType = "WirelessDevice"
	EventNotificationResourceTypeWirelessGateway EventNotificationResourceType = "WirelessGateway"
)

// Values returns all known values for EventNotificationResourceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (EventNotificationResourceType) Values() []EventNotificationResourceType {
	return []EventNotificationResourceType{
		"SidewalkAccount",
		"WirelessDevice",
		"WirelessGateway",
	}
}

type EventNotificationTopicStatus string

// Enum values for EventNotificationTopicStatus
const (
	EventNotificationTopicStatusEnabled  EventNotificationTopicStatus = "Enabled"
	EventNotificationTopicStatusDisabled EventNotificationTopicStatus = "Disabled"
)

// Values returns all known values for EventNotificationTopicStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventNotificationTopicStatus) Values() []EventNotificationTopicStatus {
	return []EventNotificationTopicStatus{
		"Enabled",
		"Disabled",
	}
}

type ExpressionType string

// Enum values for ExpressionType
const (
	ExpressionTypeRuleName  ExpressionType = "RuleName"
	ExpressionTypeMqttTopic ExpressionType = "MqttTopic"
)

// Values returns all known values for ExpressionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExpressionType) Values() []ExpressionType {
	return []ExpressionType{
		"RuleName",
		"MqttTopic",
	}
}

type FuotaDeviceStatus string

// Enum values for FuotaDeviceStatus
const (
	FuotaDeviceStatusInitial              FuotaDeviceStatus = "Initial"
	FuotaDeviceStatusPackageNotSupported  FuotaDeviceStatus = "Package_Not_Supported"
	FuotaDeviceStatusFragAlgoUnsupported  FuotaDeviceStatus = "FragAlgo_unsupported"
	FuotaDeviceStatusNotEnoughMemory      FuotaDeviceStatus = "Not_enough_memory"
	FuotaDeviceStatusFragIndexUnsupported FuotaDeviceStatus = "FragIndex_unsupported"
	FuotaDeviceStatusWrongDescriptor      FuotaDeviceStatus = "Wrong_descriptor"
	FuotaDeviceStatusSessionCntReplay     FuotaDeviceStatus = "SessionCnt_replay"
	FuotaDeviceStatusMissingFrag          FuotaDeviceStatus = "MissingFrag"
	FuotaDeviceStatusMemoryError          FuotaDeviceStatus = "MemoryError"
	FuotaDeviceStatusMICError             FuotaDeviceStatus = "MICError"
	FuotaDeviceStatusSuccessful           FuotaDeviceStatus = "Successful"
)

// Values returns all known values for FuotaDeviceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FuotaDeviceStatus) Values() []FuotaDeviceStatus {
	return []FuotaDeviceStatus{
		"Initial",
		"Package_Not_Supported",
		"FragAlgo_unsupported",
		"Not_enough_memory",
		"FragIndex_unsupported",
		"Wrong_descriptor",
		"SessionCnt_replay",
		"MissingFrag",
		"MemoryError",
		"MICError",
		"Successful",
	}
}

type FuotaTaskStatus string

// Enum values for FuotaTaskStatus
const (
	FuotaTaskStatusPending             FuotaTaskStatus = "Pending"
	FuotaTaskStatusFuotaSessionWaiting FuotaTaskStatus = "FuotaSession_Waiting"
	FuotaTaskStatusInFuotaSession      FuotaTaskStatus = "In_FuotaSession"
	FuotaTaskStatusFuotaDone           FuotaTaskStatus = "FuotaDone"
	FuotaTaskStatusDeleteWaiting       FuotaTaskStatus = "Delete_Waiting"
)

// Values returns all known values for FuotaTaskStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FuotaTaskStatus) Values() []FuotaTaskStatus {
	return []FuotaTaskStatus{
		"Pending",
		"FuotaSession_Waiting",
		"In_FuotaSession",
		"FuotaDone",
		"Delete_Waiting",
	}
}

type IdentifierType string

// Enum values for IdentifierType
const (
	IdentifierTypePartnerAccountId  IdentifierType = "PartnerAccountId"
	IdentifierTypeDevEui            IdentifierType = "DevEui"
	IdentifierTypeGatewayEui        IdentifierType = "GatewayEui"
	IdentifierTypeWirelessDeviceId  IdentifierType = "WirelessDeviceId"
	IdentifierTypeWirelessGatewayId IdentifierType = "WirelessGatewayId"
)

// Values returns all known values for IdentifierType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IdentifierType) Values() []IdentifierType {
	return []IdentifierType{
		"PartnerAccountId",
		"DevEui",
		"GatewayEui",
		"WirelessDeviceId",
		"WirelessGatewayId",
	}
}

type LogLevel string

// Enum values for LogLevel
const (
	LogLevelInfo     LogLevel = "INFO"
	LogLevelError    LogLevel = "ERROR"
	LogLevelDisabled LogLevel = "DISABLED"
)

// Values returns all known values for LogLevel. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogLevel) Values() []LogLevel {
	return []LogLevel{
		"INFO",
		"ERROR",
		"DISABLED",
	}
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeCustomCommandIdNotify MessageType = "CUSTOM_COMMAND_ID_NOTIFY"
	MessageTypeCustomCommandIdGet    MessageType = "CUSTOM_COMMAND_ID_GET"
	MessageTypeCustomCommandIdSet    MessageType = "CUSTOM_COMMAND_ID_SET"
	MessageTypeCustomCommandIdResp   MessageType = "CUSTOM_COMMAND_ID_RESP"
)

// Values returns all known values for MessageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MessageType) Values() []MessageType {
	return []MessageType{
		"CUSTOM_COMMAND_ID_NOTIFY",
		"CUSTOM_COMMAND_ID_GET",
		"CUSTOM_COMMAND_ID_SET",
		"CUSTOM_COMMAND_ID_RESP",
	}
}

type PartnerType string

// Enum values for PartnerType
const (
	PartnerTypeSidewalk PartnerType = "Sidewalk"
)

// Values returns all known values for PartnerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PartnerType) Values() []PartnerType {
	return []PartnerType{
		"Sidewalk",
	}
}

type PositionConfigurationFec string

// Enum values for PositionConfigurationFec
const (
	PositionConfigurationFecRose PositionConfigurationFec = "ROSE"
	PositionConfigurationFecNone PositionConfigurationFec = "NONE"
)

// Values returns all known values for PositionConfigurationFec. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PositionConfigurationFec) Values() []PositionConfigurationFec {
	return []PositionConfigurationFec{
		"ROSE",
		"NONE",
	}
}

type PositionConfigurationStatus string

// Enum values for PositionConfigurationStatus
const (
	PositionConfigurationStatusEnabled  PositionConfigurationStatus = "Enabled"
	PositionConfigurationStatusDisabled PositionConfigurationStatus = "Disabled"
)

// Values returns all known values for PositionConfigurationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PositionConfigurationStatus) Values() []PositionConfigurationStatus {
	return []PositionConfigurationStatus{
		"Enabled",
		"Disabled",
	}
}

type PositioningConfigStatus string

// Enum values for PositioningConfigStatus
const (
	PositioningConfigStatusEnabled  PositioningConfigStatus = "Enabled"
	PositioningConfigStatusDisabled PositioningConfigStatus = "Disabled"
)

// Values returns all known values for PositioningConfigStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PositioningConfigStatus) Values() []PositioningConfigStatus {
	return []PositioningConfigStatus{
		"Enabled",
		"Disabled",
	}
}

type PositionResourceType string

// Enum values for PositionResourceType
const (
	PositionResourceTypeWirelessDevice  PositionResourceType = "WirelessDevice"
	PositionResourceTypeWirelessGateway PositionResourceType = "WirelessGateway"
)

// Values returns all known values for PositionResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PositionResourceType) Values() []PositionResourceType {
	return []PositionResourceType{
		"WirelessDevice",
		"WirelessGateway",
	}
}

type PositionSolverProvider string

// Enum values for PositionSolverProvider
const (
	PositionSolverProviderSemtech PositionSolverProvider = "Semtech"
)

// Values returns all known values for PositionSolverProvider. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PositionSolverProvider) Values() []PositionSolverProvider {
	return []PositionSolverProvider{
		"Semtech",
	}
}

type PositionSolverType string

// Enum values for PositionSolverType
const (
	PositionSolverTypeGnss PositionSolverType = "GNSS"
)

// Values returns all known values for PositionSolverType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PositionSolverType) Values() []PositionSolverType {
	return []PositionSolverType{
		"GNSS",
	}
}

type SigningAlg string

// Enum values for SigningAlg
const (
	SigningAlgEd25519 SigningAlg = "Ed25519"
	SigningAlgP256r1  SigningAlg = "P256r1"
)

// Values returns all known values for SigningAlg. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SigningAlg) Values() []SigningAlg {
	return []SigningAlg{
		"Ed25519",
		"P256r1",
	}
}

type SupportedRfRegion string

// Enum values for SupportedRfRegion
const (
	SupportedRfRegionEu868  SupportedRfRegion = "EU868"
	SupportedRfRegionUs915  SupportedRfRegion = "US915"
	SupportedRfRegionAu915  SupportedRfRegion = "AU915"
	SupportedRfRegionAs9231 SupportedRfRegion = "AS923-1"
)

// Values returns all known values for SupportedRfRegion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SupportedRfRegion) Values() []SupportedRfRegion {
	return []SupportedRfRegion{
		"EU868",
		"US915",
		"AU915",
		"AS923-1",
	}
}

type WirelessDeviceEvent string

// Enum values for WirelessDeviceEvent
const (
	WirelessDeviceEventJoin         WirelessDeviceEvent = "Join"
	WirelessDeviceEventRejoin       WirelessDeviceEvent = "Rejoin"
	WirelessDeviceEventUplinkData   WirelessDeviceEvent = "Uplink_Data"
	WirelessDeviceEventDownlinkData WirelessDeviceEvent = "Downlink_Data"
	WirelessDeviceEventRegistration WirelessDeviceEvent = "Registration"
)

// Values returns all known values for WirelessDeviceEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceEvent) Values() []WirelessDeviceEvent {
	return []WirelessDeviceEvent{
		"Join",
		"Rejoin",
		"Uplink_Data",
		"Downlink_Data",
		"Registration",
	}
}

type WirelessDeviceFrameInfo string

// Enum values for WirelessDeviceFrameInfo
const (
	WirelessDeviceFrameInfoEnabled  WirelessDeviceFrameInfo = "ENABLED"
	WirelessDeviceFrameInfoDisabled WirelessDeviceFrameInfo = "DISABLED"
)

// Values returns all known values for WirelessDeviceFrameInfo. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceFrameInfo) Values() []WirelessDeviceFrameInfo {
	return []WirelessDeviceFrameInfo{
		"ENABLED",
		"DISABLED",
	}
}

type WirelessDeviceIdType string

// Enum values for WirelessDeviceIdType
const (
	WirelessDeviceIdTypeWirelessDeviceId        WirelessDeviceIdType = "WirelessDeviceId"
	WirelessDeviceIdTypeDevEui                  WirelessDeviceIdType = "DevEui"
	WirelessDeviceIdTypeThingName               WirelessDeviceIdType = "ThingName"
	WirelessDeviceIdTypeSidewalkManufacturingSn WirelessDeviceIdType = "SidewalkManufacturingSn"
)

// Values returns all known values for WirelessDeviceIdType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceIdType) Values() []WirelessDeviceIdType {
	return []WirelessDeviceIdType{
		"WirelessDeviceId",
		"DevEui",
		"ThingName",
		"SidewalkManufacturingSn",
	}
}

type WirelessDeviceType string

// Enum values for WirelessDeviceType
const (
	WirelessDeviceTypeSidewalk WirelessDeviceType = "Sidewalk"
	WirelessDeviceTypeLoRaWAN  WirelessDeviceType = "LoRaWAN"
)

// Values returns all known values for WirelessDeviceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WirelessDeviceType) Values() []WirelessDeviceType {
	return []WirelessDeviceType{
		"Sidewalk",
		"LoRaWAN",
	}
}

type WirelessGatewayEvent string

// Enum values for WirelessGatewayEvent
const (
	WirelessGatewayEventCupsRequest WirelessGatewayEvent = "CUPS_Request"
	WirelessGatewayEventCertificate WirelessGatewayEvent = "Certificate"
)

// Values returns all known values for WirelessGatewayEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayEvent) Values() []WirelessGatewayEvent {
	return []WirelessGatewayEvent{
		"CUPS_Request",
		"Certificate",
	}
}

type WirelessGatewayIdType string

// Enum values for WirelessGatewayIdType
const (
	WirelessGatewayIdTypeGatewayEui        WirelessGatewayIdType = "GatewayEui"
	WirelessGatewayIdTypeWirelessGatewayId WirelessGatewayIdType = "WirelessGatewayId"
	WirelessGatewayIdTypeThingName         WirelessGatewayIdType = "ThingName"
)

// Values returns all known values for WirelessGatewayIdType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayIdType) Values() []WirelessGatewayIdType {
	return []WirelessGatewayIdType{
		"GatewayEui",
		"WirelessGatewayId",
		"ThingName",
	}
}

type WirelessGatewayServiceType string

// Enum values for WirelessGatewayServiceType
const (
	WirelessGatewayServiceTypeCups WirelessGatewayServiceType = "CUPS"
	WirelessGatewayServiceTypeLns  WirelessGatewayServiceType = "LNS"
)

// Values returns all known values for WirelessGatewayServiceType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayServiceType) Values() []WirelessGatewayServiceType {
	return []WirelessGatewayServiceType{
		"CUPS",
		"LNS",
	}
}

type WirelessGatewayTaskDefinitionType string

// Enum values for WirelessGatewayTaskDefinitionType
const (
	WirelessGatewayTaskDefinitionTypeUpdate WirelessGatewayTaskDefinitionType = "UPDATE"
)

// Values returns all known values for WirelessGatewayTaskDefinitionType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (WirelessGatewayTaskDefinitionType) Values() []WirelessGatewayTaskDefinitionType {
	return []WirelessGatewayTaskDefinitionType{
		"UPDATE",
	}
}

type WirelessGatewayTaskStatus string

// Enum values for WirelessGatewayTaskStatus
const (
	WirelessGatewayTaskStatusPending     WirelessGatewayTaskStatus = "PENDING"
	WirelessGatewayTaskStatusInProgress  WirelessGatewayTaskStatus = "IN_PROGRESS"
	WirelessGatewayTaskStatusFirstRetry  WirelessGatewayTaskStatus = "FIRST_RETRY"
	WirelessGatewayTaskStatusSecondRetry WirelessGatewayTaskStatus = "SECOND_RETRY"
	WirelessGatewayTaskStatusCompleted   WirelessGatewayTaskStatus = "COMPLETED"
	WirelessGatewayTaskStatusFailed      WirelessGatewayTaskStatus = "FAILED"
)

// Values returns all known values for WirelessGatewayTaskStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayTaskStatus) Values() []WirelessGatewayTaskStatus {
	return []WirelessGatewayTaskStatus{
		"PENDING",
		"IN_PROGRESS",
		"FIRST_RETRY",
		"SECOND_RETRY",
		"COMPLETED",
		"FAILED",
	}
}

type WirelessGatewayType string

// Enum values for WirelessGatewayType
const (
	WirelessGatewayTypeLoRaWAN WirelessGatewayType = "LoRaWAN"
)

// Values returns all known values for WirelessGatewayType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WirelessGatewayType) Values() []WirelessGatewayType {
	return []WirelessGatewayType{
		"LoRaWAN",
	}
}
