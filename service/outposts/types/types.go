// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about an address.
type Address struct {

	// The first line of the address.
	//
	// This member is required.
	AddressLine1 *string

	// The city for the address.
	//
	// This member is required.
	City *string

	// The ISO-3166 two-letter country code for the address.
	//
	// This member is required.
	CountryCode *string

	// The postal code for the address.
	//
	// This member is required.
	PostalCode *string

	// The state for the address.
	//
	// This member is required.
	StateOrRegion *string

	// The second line of the address.
	AddressLine2 *string

	// The third line of the address.
	AddressLine3 *string

	// The name of the contact.
	ContactName *string

	// The phone number of the contact.
	ContactPhoneNumber *string

	// The district or county for the address.
	DistrictOrCounty *string

	// The municipality for the address.
	Municipality *string

	noSmithyDocumentSerde
}

// Information about a catalog item.
type CatalogItem struct {

	// The ID of the catalog item.
	CatalogItemId *string

	// Information about the EC2 capacity of an item.
	EC2Capacities []EC2Capacity

	// The status of a catalog item.
	ItemStatus CatalogItemStatus

	// Information about the power draw of an item.
	PowerKva *float32

	// The supported storage options for the catalog item.
	SupportedStorage []SupportedStorageEnum

	// The uplink speed this catalog item requires for the connection to the Region.
	SupportedUplinkGbps []int32

	// The weight of the item in pounds.
	WeightLbs *int32

	noSmithyDocumentSerde
}

// Information about EC2 capacity.
type EC2Capacity struct {

	// The family of the EC2 capacity.
	Family *string

	// The maximum size of the EC2 capacity.
	MaxSize *string

	// The quantity of the EC2 capacity.
	Quantity *string

	noSmithyDocumentSerde
}

// Information about an instance type.
type InstanceTypeItem struct {

	// The instance type.
	InstanceType *string

	noSmithyDocumentSerde
}

// Information about a line item.
type LineItem struct {

	// The ID of the catalog item.
	CatalogItemId *string

	// The ID of the line item.
	LineItemId *string

	// The quantity of the line item.
	Quantity int32

	// The status of the line item.
	Status LineItemStatus

	noSmithyDocumentSerde
}

// Information about a line item request.
type LineItemRequest struct {

	// The ID of the catalog item.
	CatalogItemId *string

	// The quantity of a line item request.
	Quantity int32

	noSmithyDocumentSerde
}

// Information about an order.
type Order struct {

	// The line items for the order
	LineItems []LineItem

	// The fulfillment date of the order.
	OrderFulfilledDate *time.Time

	// The ID of the order.
	OrderId *string

	// The submission date for the order.
	OrderSubmissionDate *time.Time

	// The ID of the Outpost in the order.
	OutpostId *string

	// The payment option for the order.
	PaymentOption PaymentOption

	// The status of the order.
	//
	// * PREPARING - Order is received and being prepared.
	//
	// *
	// IN_PROGRESS - Order is either being built, shipped, or installed. To get more
	// details, see the LineItem status.
	//
	// * COMPLETED - Order is complete.
	//
	// * CANCELLED
	// - Order is cancelled.
	//
	// * ERROR - Customer should contact support.
	//
	// The following
	// status are deprecated: RECEIVED, PENDING, PROCESSING, INSTALLING, and FULFILLED.
	Status OrderStatus

	noSmithyDocumentSerde
}

// A summary of line items in your order.
type OrderSummary struct {

	// The status of all line items in the order.
	LineItemCountsByStatus map[string]int32

	// Fulfilment date for the order.
	OrderFulfilledDate *time.Time

	// The ID of the order.
	OrderId *string

	// Submission date for the order.
	OrderSubmissionDate *time.Time

	// The type of order.
	OrderType OrderType

	// The ID of the Outpost.
	OutpostId *string

	// The status of the order.
	//
	// * PREPARING - Order is received and is being
	// prepared.
	//
	// * IN_PROGRESS - Order is either being built, shipped, or installed.
	// For more information, see the LineItem status.
	//
	// * COMPLETED - Order is
	// complete.
	//
	// * CANCELLED - Order is cancelled.
	//
	// * ERROR - Customer should contact
	// support.
	//
	// The following statuses are deprecated: RECEIVED, PENDING, PROCESSING,
	// INSTALLING, and FULFILLED.
	Status OrderStatus

	noSmithyDocumentSerde
}

// Information about an Outpost.
type Outpost struct {

	// The Availability Zone.
	AvailabilityZone *string

	// The ID of the Availability Zone.
	AvailabilityZoneId *string

	// The description of the Outpost.
	Description *string

	// The life cycle status.
	LifeCycleStatus *string

	// The name of the Outpost.
	Name *string

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The ID of the Outpost.
	OutpostId *string

	// The Amazon Web Services account ID of the Outpost owner.
	OwnerId *string

	// The Amazon Resource Name (ARN) of the site.
	SiteArn *string

	// The ID of the site.
	SiteId *string

	// The Outpost tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Information about the physical and logistical details for racks at sites. For
// more information about hardware requirements for racks, see Network readiness
// checklist
// (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#checklist)
// in the Amazon Web Services Outposts User Guide.
type RackPhysicalProperties struct {

	// The type of fiber used to attach the Outpost to the network.
	FiberOpticCableType FiberOpticCableType

	// The maximum rack weight that this site can support. NO_LIMIT is over 2000 lbs
	// (907 kg).
	MaximumSupportedWeightLbs MaximumSupportedWeightLbs

	// The type of optical standard used to attach the Outpost to the network. This
	// field is dependent on uplink speed, fiber type, and distance to the upstream
	// device. For more information about networking requirements for racks, see
	// Network
	// (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#facility-networking)
	// in the Amazon Web Services Outposts User Guide.
	OpticalStandard OpticalStandard

	// The power connector for the hardware.
	PowerConnector PowerConnector

	// The power draw available at the hardware placement position for the rack.
	PowerDrawKva PowerDrawKva

	// The position of the power feed.
	PowerFeedDrop PowerFeedDrop

	// The power option that you can provide for hardware.
	PowerPhase PowerPhase

	// The number of uplinks each Outpost network device.
	UplinkCount UplinkCount

	// The uplink speed the rack supports for the connection to the Region.
	UplinkGbps UplinkGbps

	noSmithyDocumentSerde
}

// Information about a site.
type Site struct {

	// The ID of the Amazon Web Services account.
	AccountId *string

	// The description of the site.
	Description *string

	// The name of the site.
	Name *string

	// Notes about a site.
	Notes *string

	// City where the hardware is installed and powered on.
	OperatingAddressCity *string

	// The ISO-3166 two-letter country code where the hardware is installed and powered
	// on.
	OperatingAddressCountryCode *string

	// State or region where the hardware is installed and powered on.
	OperatingAddressStateOrRegion *string

	// Information about the physical and logistical details for a rack at the site.
	RackPhysicalProperties *RackPhysicalProperties

	// The Amazon Resource Name (ARN) of the site.
	SiteArn *string

	// The ID of the site.
	SiteId *string

	// The site tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
