package models

import "time"

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

type TimeWrapper struct {
	time.Time
}

func (t *TimeWrapper) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*t = TimeWrapper{tt}
	return err
}
