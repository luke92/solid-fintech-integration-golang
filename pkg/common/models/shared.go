package models

type AddressType string

const (
	AddressTypeMailing  AddressType = "mailing"
	AddressTypeBilling  AddressType = "billing"
	AddressTypeShipping AddressType = "shipping"
	AddressTypeCard     AddressType = "card"
	AddressTypeWire     AddressType = "wire"
	AddressTypeCheck    AddressType = "check"
)

type Address struct {
	AddressType AddressType `json:"addressType"`
	// 	line 1 of the address (60 characters max)
	Line1 string `json:"line1"`
	// line 2 of the address (optional, 60 characters max)
	Line2 string `json:"line2"`
	// 	city of the address (60 characters max)
	City string `json:"city"`
	// 2-Letter US state abbreviation or 60 characters max (ex: CA)
	State string `json:"state"`
	// 	2-letter abbreviated country code(ex: US)
	Country string `json:"country"`
	// postal code (60 characters max)
	PostalCode string `json:"postalCode"`
}
