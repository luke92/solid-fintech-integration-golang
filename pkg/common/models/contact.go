package models

import "time"

type IntrabankContactDataPartial struct {
	AccountNumber string `json:"accountNumber"`
}

type IntrabankContactDataFull struct {
	IntrabankContactDataPartial
	AccountID string `json:"accountId"`
}

type AchContactDataPartial struct {
	IntrabankContactDataPartial
	RoutingNumber string `json:"routingNumber"`
	AccountType   string `json:"accountType"`
	BankName      string `json:"bankName"`
}

type AchContactDataFull struct {
	AchContactDataPartial
	Verified bool `json:"verified"`
}

type contactData struct {
	AccountID string `json:"accountId"`
	// 	person full name or business legal name (100 character max, alphanumeric + special characters allowed)
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ContactDataPartial struct {
	contactData
	Intrabank IntrabankContactDataPartial `json:"intrabank"`
	Ach       AchContactDataPartial       `json:"ach"`
}

type DebitCard struct {
	CardNumber  string `json:"cardNumber"`
	Last4       string `json:"last4"`
	ExpiryMonth string `json:"expiryMonth"`
	// 	CVV code (required for debit pull verification)
	CVV        string  `json:"cvv"`
	ExpiryYear string  `json:"expiryYear"`
	Address    Address `json:"address"`
	//	if debit pull is enabled or disabled by the debit card issuer
	Pull interface{} `json:"pull"`
	// if debit push is enabled or disabled by the debit card issuer
	Push      interface{} `json:"push"`
	UsageType string      `json:"usageType"`
}

type ContactStatus string

const (
	ContactStatusActive  ContactStatus = "active"
	ContactStatusDeleted ContactStatus = "deleted"
)

type ContactType string

const (
	// If users try to include a self-owned intrabank account when creating a standard "others" contact, they receive an error such as "EC_CONTACT_INVALID_SELF_INTRABANK_REQUEST" and "Invalid self intrabank request".
	// They should remove the intrabank sub-object and create a separate contact for the selfIntrabank link.
	ContactTypeOthers  ContactType = "others"
	ContactTypeSelfAch ContactType = "selfAch"
	// If the user is linking intrabank accounts owned by the same entity, this is referred to as a selfIntrabank contact.
	// In this situation, they should create a separate contact for the selfIntrabank link using only the intrabank sub-object.
	ContactTypeSelfIntrabank ContactType = "selfIntrabank"
)

type ContactDataFull struct {
	ID string `json:"id"`
	contactData
	Label       string                   `json:"label"`
	Status      ContactStatus            `json:"status"`
	Type        ContactType              `json:"type"`
	ProgramID   string                   `json:"programId"`
	CreatedAt   time.Time                `json:"createdAt"`
	ModifiedAt  time.Time                `json:"modifiedAt"`
	Intrabank   IntrabankContactDataFull `json:"intrabank"`
	Ach         AchContactDataFull       `json:"ach"`
	Wire        interface{}              `json:"wire"`
	Check       interface{}              `json:"check"`
	Card        interface{}              `json:"card"`
	DebitCard   DebitCard                `json:"debitCard"`
	CrossBorder interface{}              `json:"crossBorder"`
	Metadata    interface{}              `json:"metadata"`
}

type AccountBankType string

const (
	AccountBankTypeABA   AccountBankType = "aba"
	AccountBankTypeSwift AccountBankType = "swift"
)

// You can Retrieve a Bank information in two ways:
// - Pass type=aba and the routingNumber=Bank's Routing Number (for Domestic US Banks)
// - Pass type=swift and the routingNumber=SWIFT code (for International Banks)
type BankInfo struct {
	RoutingNumber  string          `json:"routingNumber"`
	BankName       string          `json:"bankName"`
	Type           AccountBankType `json:"type"`
	PaymentMethods []string        `json:"paymentMethods"`
	Address        Address         `json:"address"`
}

type DebitCardTokenResponse struct {
	ID             string `json:"id"`
	DebitCardToken string `json:"debitCardToken"`
}
