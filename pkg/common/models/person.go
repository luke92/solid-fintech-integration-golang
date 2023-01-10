package models

import "time"

type PersonIDType string

const (
	PersonIDTypeSSN      PersonIDType = "ssn"
	PersonIDTypePassport PersonIDType = "passport"
	PersonIDTypeOther    PersonIDType = "otherId"
)

type PersonStatus string

const (
	PersonStatusActive      PersonStatus = "active"
	PersonStatusInactive    PersonStatus = "inactive"    // (use to temporarily freeze)
	PersonStatusDeactivated PersonStatus = "deactivated" // (use to permanently block)
)

// Note: If KYC is not needed, the only fields that are required to create a person are firstName, lastName, phone, and email.
type PersonDataPartial struct {
	// 	first name of the person (max 100 chars)
	FirstName string `json:"firstName"`
	// middle name of the person (max 100 chars)
	MiddleName string `json:"middleName"`
	// last name of the person (max 100 chars)
	LastName string `json:"lastName"`
	// mobile phone of the person (E.164, max 16 chars, starts with +)
	Phone string `json:"phone"`
	// email of the person (max 100 chars)
	Email string `json:"email"`
	// 	date of birth of the person (YYYY-MM-DD)
	DateOfBirth string       `json:"dateOfBirth"`
	IDType      PersonIDType `json:"idType"`
	/*
		- if idType is ssn, idNumber can be either Last 4 or full SSN
		- if idType is passport, idNumber must be passport number
		- if idType is otherId, idNumber must be unique number of the id
		- idNumber must be unique, as in, you cannot use the same idNumber for two different persons
		- 9 to 50 chars
	*/
	IDNumber string      `json:"idNumber"`
	Address  Address     `json:"address"`
	Metadata interface{} `json:"metadata"`
}

type PersonDataFull struct {
	ID string `json:"id"`
	PersonDataPartial
	// 	person's phone is verified (Solid internal use)
	PhoneVerified bool `json:"phoneVerified"`
	// person's email is verified (Solid internal use)
	EmailVerified bool         `json:"emailVerified"`
	KYC           KYC          `json:"kyc"`
	ProgramID     string       `json:"programId"`
	CreatedAt     time.Time    `json:"createdAt"`
	ModifiedAt    time.Time    `json:"modifiedAt"`
	Language      string       `json:"language"`
	Status        PersonStatus `json:"status"`
}

type KYCResults struct {
	Idv         string `json:"idv"`
	Address     string `json:"address"`
	DateOfBirth string `json:"dateOfBirth"`
	Fraud       string `json:"fraud"`
}

type KYCStatus string

const (
	KYCStatusNotStarted KYCStatus = "notStarted"
	KYCStatusSubmitted  KYCStatus = "submitted"
	KYCStatusApproved   KYCStatus = "approved"
	KYCStatusDeclined   KYCStatus = "declined"
	KYCStatusInReview   KYCStatus = "inReview"
)

type KYC struct {
	ID       string    `json:"id"`
	PersonID string    `json:"personId"`
	Status   KYCStatus `json:"status"`
	// https://www.solidfi.com/docs/kyc-in-review-codes
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
