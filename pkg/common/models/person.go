package models

import "time"

type PersonDataPartial struct {
	FirstName   string        `json:"firstName"`
	MiddleName  string        `json:"middleName"`
	LastName    string        `json:"lastName"`
	Phone       string        `json:"phone"`
	Email       string        `json:"email"`
	DateOfBirth string        `json:"dateOfBirth"`
	IDType      string        `json:"idType"`
	IDNumber    string        `json:"idNumber"`
	Address     PersonAddress `json:"address"`
	Metadata    interface{}   `json:"metadata"`
}

type PersonDataFull struct {
	ID string `json:"id"`
	PersonDataPartial
	PhoneVerified bool      `json:"phoneVerified"`
	EmailVerified bool      `json:"emailVerified"`
	KYC           KYC       `json:"kyc"`
	ProgramID     string    `json:"programId"`
	CreatedAt     time.Time `json:"createdAt"`
	ModifiedAt    time.Time `json:"modifiedAt"`
	Language      string    `json:"language"`
	Status        string    `json:"status"`
}

type PersonAddress struct {
	AddressType string `json:"addressType"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	PostalCode  string `json:"postalCode"`
}

type KYCResults struct {
	Idv         string `json:"idv"`
	Address     string `json:"address"`
	DateOfBirth string `json:"dateOfBirth"`
	Fraud       string `json:"fraud"`
}

type KYC struct {
	ID            string        `json:"id"`
	PersonID      string        `json:"personId"`
	Status        string        `json:"status"`
	ReviewCode    string        `json:"reviewCode"`
	ReviewMessage string        `json:"reviewMessage"`
	Results       KYCResults    `json:"results"`
	CreatedAt     time.Time     `json:"createdAt"`
	ModifiedAt    time.Time     `json:"modifiedAt"`
	ReviewReasons ReviewReasons `json:"reviewReasons"`
}

type ReviewReasons struct {
	Idv         []interface{} `json:"idv"`
	Address     []interface{} `json:"address"`
	DateOfBirth []interface{} `json:"dateOfBirth"`
	Fraud       []interface{} `json:"fraud"`
}

type IDV struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}
