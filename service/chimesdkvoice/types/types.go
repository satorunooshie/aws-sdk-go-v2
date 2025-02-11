// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

type Address struct {
	City *string

	Country *string

	PostDirectional *string

	PostalCode *string

	PostalCodePlus4 *string

	PreDirectional *string

	State *string

	StreetName *string

	StreetNumber *string

	StreetSuffix *string

	noSmithyDocumentSerde
}

type CandidateAddress struct {
	City *string

	Country *string

	PostalCode *string

	PostalCodePlus4 *string

	State *string

	StreetInfo *string

	StreetNumber *string

	noSmithyDocumentSerde
}

type Credential struct {
	Password *string

	Username *string

	noSmithyDocumentSerde
}

type DNISEmergencyCallingConfiguration struct {

	// This member is required.
	CallingCountry *string

	// This member is required.
	EmergencyPhoneNumber *string

	TestPhoneNumber *string

	noSmithyDocumentSerde
}

type EmergencyCallingConfiguration struct {
	DNIS []DNISEmergencyCallingConfiguration

	noSmithyDocumentSerde
}

type GeoMatchParams struct {

	// This member is required.
	AreaCode *string

	// This member is required.
	Country *string

	noSmithyDocumentSerde
}

type LoggingConfiguration struct {
	EnableMediaMetricLogs *bool

	EnableSIPLogs *bool

	noSmithyDocumentSerde
}

type OrderedPhoneNumber struct {
	E164PhoneNumber *string

	Status OrderedPhoneNumberStatus

	noSmithyDocumentSerde
}

type Origination struct {
	Disabled *bool

	Routes []OriginationRoute

	noSmithyDocumentSerde
}

type OriginationRoute struct {
	Host *string

	Port *int32

	Priority *int32

	Protocol OriginationRouteProtocol

	Weight *int32

	noSmithyDocumentSerde
}

type Participant struct {
	PhoneNumber *string

	ProxyPhoneNumber *string

	noSmithyDocumentSerde
}

type PhoneNumber struct {
	Associations []PhoneNumberAssociation

	CallingName *string

	CallingNameStatus CallingNameStatus

	Capabilities *PhoneNumberCapabilities

	Country *string

	CreatedTimestamp *time.Time

	DeletionTimestamp *time.Time

	E164PhoneNumber *string

	OrderId *string

	PhoneNumberId *string

	ProductType PhoneNumberProductType

	Status PhoneNumberStatus

	Type PhoneNumberType

	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

type PhoneNumberAssociation struct {
	AssociatedTimestamp *time.Time

	Name PhoneNumberAssociationName

	Value *string

	noSmithyDocumentSerde
}

type PhoneNumberCapabilities struct {
	InboundCall *bool

	InboundMMS *bool

	InboundSMS *bool

	OutboundCall *bool

	OutboundMMS *bool

	OutboundSMS *bool

	noSmithyDocumentSerde
}

type PhoneNumberCountry struct {
	CountryCode *string

	SupportedPhoneNumberTypes []PhoneNumberType

	noSmithyDocumentSerde
}

type PhoneNumberError struct {
	ErrorCode ErrorCode

	ErrorMessage *string

	PhoneNumberId *string

	noSmithyDocumentSerde
}

type PhoneNumberOrder struct {
	CreatedTimestamp *time.Time

	OrderType PhoneNumberOrderType

	OrderedPhoneNumbers []OrderedPhoneNumber

	PhoneNumberOrderId *string

	ProductType PhoneNumberProductType

	Status PhoneNumberOrderStatus

	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

type Proxy struct {
	DefaultSessionExpiryMinutes *int32

	Disabled *bool

	FallBackPhoneNumber *string

	PhoneNumberCountries []string

	noSmithyDocumentSerde
}

type ProxySession struct {
	Capabilities []Capability

	CreatedTimestamp *time.Time

	EndedTimestamp *time.Time

	ExpiryMinutes *int32

	GeoMatchLevel GeoMatchLevel

	GeoMatchParams *GeoMatchParams

	Name *string

	NumberSelectionBehavior NumberSelectionBehavior

	Participants []Participant

	ProxySessionId *string

	Status ProxySessionStatus

	UpdatedTimestamp *time.Time

	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type SipMediaApplication struct {
	AwsRegion *string

	CreatedTimestamp *time.Time

	Endpoints []SipMediaApplicationEndpoint

	Name *string

	SipMediaApplicationId *string

	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

type SipMediaApplicationAlexaSkillConfiguration struct {

	// This member is required.
	AlexaSkillIds []string

	// This member is required.
	AlexaSkillStatus AlexaSkillStatus

	noSmithyDocumentSerde
}

type SipMediaApplicationCall struct {
	TransactionId *string

	noSmithyDocumentSerde
}

type SipMediaApplicationEndpoint struct {
	LambdaArn *string

	noSmithyDocumentSerde
}

type SipMediaApplicationLoggingConfiguration struct {
	EnableSipMediaApplicationMessageLogs *bool

	noSmithyDocumentSerde
}

type SipRule struct {
	CreatedTimestamp *time.Time

	Disabled *bool

	Name *string

	SipRuleId *string

	TargetApplications []SipRuleTargetApplication

	TriggerType SipRuleTriggerType

	TriggerValue *string

	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

type SipRuleTargetApplication struct {
	AwsRegion *string

	Priority *int32

	SipMediaApplicationId *string

	noSmithyDocumentSerde
}

type StreamingConfiguration struct {

	// This member is required.
	DataRetentionInHours *int32

	// This member is required.
	Disabled *bool

	StreamingNotificationTargets []StreamingNotificationTarget

	noSmithyDocumentSerde
}

type StreamingNotificationTarget struct {
	NotificationTarget NotificationTarget

	noSmithyDocumentSerde
}

type Termination struct {
	CallingRegions []string

	CidrAllowedList []string

	CpsLimit *int32

	DefaultPhoneNumber *string

	Disabled *bool

	noSmithyDocumentSerde
}

type TerminationHealth struct {
	Source *string

	Timestamp *time.Time

	noSmithyDocumentSerde
}

type UpdatePhoneNumberRequestItem struct {

	// This member is required.
	PhoneNumberId *string

	CallingName *string

	ProductType PhoneNumberProductType

	noSmithyDocumentSerde
}

type VoiceConnector struct {
	AwsRegion VoiceConnectorAwsRegion

	CreatedTimestamp *time.Time

	Name *string

	OutboundHostName *string

	RequireEncryption *bool

	UpdatedTimestamp *time.Time

	VoiceConnectorArn *string

	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type VoiceConnectorGroup struct {
	CreatedTimestamp *time.Time

	Name *string

	UpdatedTimestamp *time.Time

	VoiceConnectorGroupArn *string

	VoiceConnectorGroupId *string

	VoiceConnectorItems []VoiceConnectorItem

	noSmithyDocumentSerde
}

type VoiceConnectorItem struct {

	// This member is required.
	Priority *int32

	// This member is required.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type VoiceConnectorSettings struct {
	CdrBucket *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
