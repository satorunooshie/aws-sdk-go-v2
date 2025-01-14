// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// The values of a given attribute, such as Throughput Optimized HDD or Provisioned
// IOPS for the Amazon EC2volumeType attribute.
type AttributeValue struct {

	// The specific value of an attributeName.
	Value *string

	noSmithyDocumentSerde
}

// The constraints that you want all returned products to match.
type Filter struct {

	// The product metadata field that you want to filter on. You can filter by just
	// the service code to see all products for a specific service, filter by just the
	// attribute name to see a specific attribute for multiple services, or use both a
	// service code and an attribute name to retrieve only products that match both
	// fields. Valid values include: ServiceCode, and all attribute names For example,
	// you can filter by the AmazonEC2 service code and the volumeType attribute name
	// to get the prices for only Amazon EC2 volumes.
	//
	// This member is required.
	Field *string

	// The type of filter that you want to use. Valid values are: TERM_MATCH.
	// TERM_MATCH returns only products that match both the given filter field and the
	// given value.
	//
	// This member is required.
	Type FilterType

	// The service code or attribute value that you want to filter by. If you're
	// filtering by service code this is the actual service code, such as AmazonEC2. If
	// you're filtering by attribute name, this is the attribute value that you want
	// the returned products to match, such as a Provisioned IOPS volume.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// This feature is in preview release and is subject to change. Your use of Amazon
// Web Services Price List API is subject to the Beta Service Participation terms
// of the Amazon Web Services Service Terms (https://aws.amazon.com/service-terms/)
// (Section 1.10). This is the type of price list references that match your
// request.
type PriceList struct {

	// The three alphabetical character ISO-4217 currency code the Price List files are
	// denominated in.
	CurrencyCode *string

	// The format you want to retrieve your Price List files. The FileFormat can be
	// obtained from the ListPriceList
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_ListPriceLists.html)
	// response.
	FileFormats []string

	// The unique identifier that maps to where your Price List files are located.
	// PriceListArn can be obtained from the ListPriceList
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_ListPriceLists.html)
	// response.
	PriceListArn *string

	// This is used to filter the Price List by Amazon Web Services Region. For
	// example, to get the price list only for the US East (N. Virginia) Region, use
	// us-east-1. If nothing is specified, you retrieve price lists for all applicable
	// Regions. The available RegionCode list can be retrieved from GetAttributeValues
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_GetAttributeValues.html)
	// API.
	RegionCode *string

	noSmithyDocumentSerde
}

// The metadata for a service, such as the service code and available attribute
// names.
type Service struct {

	// The code for the Amazon Web Services service.
	//
	// This member is required.
	ServiceCode *string

	// The attributes that are available for this service.
	AttributeNames []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
