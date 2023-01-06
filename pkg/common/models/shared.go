package models

type Address struct {
	AddressType string `json:"addressType"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	PostalCode  string `json:"postalCode"`
}
